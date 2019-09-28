package interfaces

import (
	"github.com/ega-forever/otus_go/scan_service/internal/domain/models"
)

// todo add more actions
type EventStorage interface {
	SaveEvent(event *models.Event) (*models.Event, error)
	FindEventsAfterTimestamp(timestamp int64) ([]*models.Event, error)
}
