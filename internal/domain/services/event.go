package services

import (
	"context"
	"github.com/ega-forever/otus_go/internal/domain/interfaces"
	"github.com/ega-forever/otus_go/internal/domain/models"
)

type EventService struct {
	eventStorage interfaces.EventStorage
}

func NewEventService(eventStorage interfaces.EventStorage) *EventService {

	return &EventService{eventStorage: eventStorage}
}

func (es *EventService) CreateEvent(ctx context.Context, text string, timestamp int64) (*models.Event, error) {

	savedEvent, err := es.eventStorage.SaveEvent(ctx, &models.Event{Text: text, Timestamp: timestamp})

	if err != nil {
		return nil, err
	}

	return savedEvent, nil
}
func (es *EventService) UpdateEvent(ctx context.Context, id int64, text string, timestamp int64) (*models.Event, error) {

	ev, err := es.eventStorage.UpdateEventById(ctx, id, text, timestamp)

	if err != nil {
		return nil, err
	}

	return ev, nil
}
func (es *EventService) DeleteEvent(ctx context.Context, id int64) error {
	return es.eventStorage.DeleteEventById(ctx, id)
}

func (es *EventService) GetEvent(ctx context.Context, id int64) (*models.Event, error) {

	ev, err := es.eventStorage.GetEventById(ctx, id)

	if err != nil {
		return nil, err
	}

	return ev, nil
}

func (es *EventService) ListEvents(ctx context.Context) ([]*models.Event, error) {

	evs, err := es.eventStorage.ListEvents(ctx)

	if err != nil {
		return nil, err
	}

	return evs, nil
}
