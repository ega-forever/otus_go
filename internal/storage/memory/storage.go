package memory

import (
	"context"
	"errors"
	"github.com/ega-forever/otus_go/internal/domain/models"
)

type Storage struct {
	events map[int64]*models.Event
}

func New() *Storage {
	return &Storage{make(map[int64]*models.Event)}
}

func (storage *Storage) SaveEvent(ctx context.Context, event *models.Event) (*models.Event, error) {

	id := int64(len(storage.events)) + 1
	ev := models.Event{Id: id, Text: event.Text, Timestamp: event.Timestamp}

	storage.events[id] = &ev
	return &ev, nil
}

func (storage *Storage) UpdateEventById(ctx context.Context, id int64, text string, timestamp int64) (*models.Event, error) {
	elem, ok := storage.events[id]

	if ok == false {
		return nil, errors.New("record not found")
	}

	elem.Text = text
	elem.Timestamp = timestamp
	return elem, nil
}

func (storage *Storage) GetEventById(ctx context.Context, id int64) (*models.Event, error) {
	elem, ok := storage.events[id]

	if ok == false {
		return nil, errors.New("record not found")
	}

	return elem, nil
}

func (storage *Storage) DeleteEventById(ctx context.Context, id int64) error {
	_, ok := storage.events[id]

	if ok == false {
		return errors.New("record not found")
	}

	delete(storage.events, id)

	return nil
}

func (storage *Storage) ListEvents(ctx context.Context) ([]*models.Event, error) {

	envs := make([]*models.Event, 0)

	for _, env := range storage.events {
		envs = append(envs, env)
	}

	return envs, nil
}
