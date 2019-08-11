package internal

import (
	"context"
	"github.com/ega-forever/otus_go/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EventService struct {
}

var events = make(map[int64]event.Event, 0)

func (*EventService) CreateEvent(ctx context.Context, req *event.CreateEventReq) (*event.CreateEventRes, error) {

	req.Event.Id = int64(len(events) + 1)

	events[req.Event.Id] = *req.Event
	res := event.CreateEventRes{Event: req.Event}
	return &res, nil
}
func (*EventService) UpdateEvent(ctx context.Context, req *event.UpdateEventReq) (*event.UpdateEventRes, error) {

	ev := events[req.Event.Id]
	if ev.Id == 0 {
		return nil, status.Error(codes.NotFound, "record not found")
	}

	events[req.Event.Id] = *req.Event

	res := event.UpdateEventRes{Event: req.Event}
	return &res, nil
}
func (*EventService) DeleteEvent(ctx context.Context, req *event.DeleteEventReq) (*event.DeleteEventRes, error) {

	ev := events[req.Id]
	if ev.Id == 0 {
		return nil, status.Error(codes.NotFound, "record not found")
	}

	delete(events, req.Id)

	res := event.DeleteEventRes{Id: req.Id}
	return &res, nil
}
func (*EventService) GetEvent(ctx context.Context, req *event.GetEventReq) (*event.GetEventRes, error) {

	ev := events[req.Id]

	if ev.Id == 0 {
		return nil, status.Error(codes.NotFound, "record not found")
	}

	res := event.GetEventRes{Event: &ev}
	return &res, nil
}
func (*EventService) ListEvents(ctx context.Context, req *event.ListEventReq) (*event.ListEventRes, error) {

	eventsArr := make([]*event.Event, 0)

	for _, ev := range events {
		eventsArr = append(eventsArr, &ev)
	}

	res := event.ListEventRes{Events: eventsArr}

	return &res, nil
}
