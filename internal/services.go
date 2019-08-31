package internal

import (
	"context"
	"github.com/ega-forever/otus_go/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EventService struct {
	events map[int64]event.Event
}

func NewEventService() *EventService {
	es := EventService{}
	es.events = make(map[int64]event.Event, 0)
	return &es
}

func (es *EventService) CreateEvent(ctx context.Context, req *event.CreateEventReq) (*event.CreateEventRes, error) {

	req.Event.Id = int64(len(es.events) + 1)

	es.events[req.Event.Id] = *req.Event
	res := event.CreateEventRes{Event: req.Event}
	return &res, nil
}
func (es *EventService) UpdateEvent(ctx context.Context, req *event.UpdateEventReq) (*event.UpdateEventRes, error) {

	_, ok := es.events[req.Event.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "record not found")
	}

	es.events[req.Event.Id] = *req.Event

	res := event.UpdateEventRes{Event: req.Event}
	return &res, nil
}
func (es *EventService) DeleteEvent(ctx context.Context, req *event.DeleteEventReq) (*event.DeleteEventRes, error) {

	_, ok := es.events[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "record not found")
	}

	delete(es.events, req.Id)

	res := event.DeleteEventRes{Id: req.Id}
	return &res, nil
}
func (es *EventService) GetEvent(ctx context.Context, req *event.GetEventReq) (*event.GetEventRes, error) {

	ev, ok := es.events[req.Id]

	if !ok {
		return nil, status.Error(codes.NotFound, "record not found")
	}

	res := event.GetEventRes{Event: &ev}
	return &res, nil
}
func (es *EventService) ListEvents(ctx context.Context, req *event.ListEventReq) (*event.ListEventRes, error) {

	eventsArr := make([]*event.Event, 0)

	for _, ev := range es.events {
		eventsArr = append(eventsArr, &ev)
	}

	res := event.ListEventRes{Events: eventsArr}

	return &res, nil
}
