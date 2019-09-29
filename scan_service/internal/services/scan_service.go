package services

import (
	"context"
	"github.com/ega-forever/otus_go/scan_service/internal/domain/interfaces"
	log "github.com/sirupsen/logrus"
	"time"
)

type ScanService struct {
	ctxCancel context.CancelFunc
	Ctx       context.Context
	storage   interfaces.EventStorage
	queue     interfaces.EventQueue
}

func New(storage interfaces.EventStorage, queue interfaces.EventQueue) *ScanService {

	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)

	return &ScanService{Ctx: ctx, ctxCancel: cancelFunc, storage: storage, queue: queue}
}

func (ss *ScanService) scan(timestamp int64) error {

	events, err := ss.storage.FindEventsAfterTimestamp(timestamp)

	if err != nil {
		ss.ctxCancel()
	}

	for _, ev := range events {
		err := ss.queue.PushEvent(ev)

		if err != nil {
			return err
		}

	}

	return nil
}

func (ss *ScanService) Job(seconds time.Duration, timestamp int64) {

	go func() {

		timeout := time.NewTicker(time.Second * seconds)

		for {

			select {
			case <-timeout.C:
				err := ss.scan(timestamp)

				if err != nil {
					ss.ctxCancel()
					return
				}

				log.Println("pushed events")

			case <-ss.Ctx.Done():
				return
			}

		}

	}()

}
