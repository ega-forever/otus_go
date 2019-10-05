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

func NewScanService(storage interfaces.EventStorage, queue interfaces.EventQueue) *ScanService {

	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)

	return &ScanService{Ctx: ctx, ctxCancel: cancelFunc, storage: storage, queue: queue}
}

func (ss *ScanService) scan(eventCreatedEarliestSeconds int64) (int, error) {

	startTimestamp := (time.Now().Unix() - eventCreatedEarliestSeconds) * 1000

	events, err := ss.storage.FindEventsAfterTimestamp(startTimestamp)

	if err != nil {
		ss.ctxCancel()
	}

	for _, ev := range events {
		err := ss.queue.PushEvent(ev)

		if err != nil {
			return 0, err
		}

	}

	return len(events), nil
}

func (ss *ScanService) Job(scanSeconds time.Duration, eventCreatedEarliestSeconds int64) {

	go func() {

		timeout := time.NewTicker(time.Second * scanSeconds)

		for {

			select {
			case <-timeout.C:
				count, err := ss.scan(eventCreatedEarliestSeconds)

				if err != nil {
					ss.ctxCancel()
					return
				}

				if count > 0 {
					log.Printf("pushed %d events", count)
				}

			case <-ss.Ctx.Done():
				return
			}

		}

	}()

}
