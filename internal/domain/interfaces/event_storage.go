package interfaces

import (
	"context"
	"github.com/ega-forever/otus_go/internal/domain/models"
)

// todo add more actions
type EventStorage interface {
	SaveEvent(ctx context.Context, event *models.Event) (*models.Event, error)
	UpdateEventById(ctx context.Context, id int64, text string, timestamp int64) (*models.Event, error)
	GetEventById(ctx context.Context, id int64) (*models.Event, error)
	DeleteEventById(ctx context.Context, id int64) error
	ListEvents(ctx context.Context) ([]*models.Event, error)
}
