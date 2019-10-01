package services

import (
	"context"
	"github.com/ega-forever/otus_go/notification_service/internal/domain/interfaces"
	log "github.com/sirupsen/logrus"
)

type ScanService struct {
	ctxCancel context.CancelFunc
	Ctx       context.Context
	queue     interfaces.EventQueue
}

func NewScanService(queue interfaces.EventQueue) *ScanService {

	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)

	return &ScanService{Ctx: ctx, ctxCancel: cancelFunc, queue: queue}
}

func (ss *ScanService) Job() {

	go func() {

		subscription := ss.queue.Subscribe()

		for {

			select {
			case ev := <-subscription:
				log.Printf("remind about event %s, started at %d", ev.Text, ev.Timestamp)

			case <-ss.Ctx.Done():
				return
			}

		}

	}()

}
