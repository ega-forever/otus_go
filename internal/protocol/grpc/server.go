package grpc

import (
	"context"
	"github.com/ega-forever/otus_go/internal/domain/services"
	"github.com/ega-forever/otus_go/internal/protocol/grpc/api"
)

type EventServer struct {
	EventService *services.EventService
}

func New(es *services.EventService) *EventServer {
	return &EventServer{es}
}

// implementation of CreateEvent
func (server *EventServer) CreateEvent(ctx context.Context, req *event.CreateEventReq) (*event.CreateEventRes, error) {

	ev, err := server.EventService.CreateEvent(ctx, req.GetEvent().Text, req.GetEvent().Timestamp)
	if err != nil {
		resp := &event.CreateEventRes{
			Result: &event.CreateEventRes_Error{Error: err.Error()},
		}
		return resp, nil
	}

	createdEv := event.Event{Id: ev.Id, Text: ev.Text, Timestamp: ev.Timestamp}

	resp := &event.CreateEventRes{
		Result: &event.CreateEventRes_Event{
			Event: &createdEv,
		},
	}

	return resp, nil
}

func (server *EventServer) UpdateEvent(ctx context.Context, req *event.UpdateEventReq) (*event.UpdateEventRes, error) {

	ev, err := server.EventService.UpdateEvent(ctx, req.GetEvent().Id, req.GetEvent().Text, req.GetEvent().Timestamp)
	if err != nil {
		resp := &event.UpdateEventRes{
			Result: &event.UpdateEventRes_Error{Error: err.Error()},
		}
		return resp, nil
	}

	resp := &event.UpdateEventRes{
		Result: &event.UpdateEventRes_Event{
			Event: &event.Event{Id: ev.Id, Text: ev.Text},
		},
	}

	return resp, nil
}

func (server *EventServer) GetEvent(ctx context.Context, req *event.GetEventReq) (*event.GetEventRes, error) {

	ev, err := server.EventService.GetEvent(ctx, req.Id)
	if err != nil {
		resp := &event.GetEventRes{
			Result: &event.GetEventRes_Error{Error: err.Error()},
		}
		return resp, nil
	}

	resp := &event.GetEventRes{
		Result: &event.GetEventRes_Event{
			Event: &event.Event{Id: ev.Id, Text: ev.Text},
		},
	}

	return resp, nil
}

func (server *EventServer) DeleteEvent(ctx context.Context, req *event.DeleteEventReq) (*event.DeleteEventRes, error) {

	err := server.EventService.DeleteEvent(ctx, req.Id)
	if err != nil {
		resp := &event.DeleteEventRes{
			Result: &event.DeleteEventRes_Error{Error: err.Error()},
		}
		return resp, nil
	}

	resp := &event.DeleteEventRes{
		Result: &event.DeleteEventRes_Event{
			Event: req.Id,
		},
	}

	return resp, nil
}

func (server *EventServer) ListEvents(ctx context.Context, req *event.ListEventReq) (*event.ListEventRes, error) {

	events, err := server.EventService.ListEvents(ctx)

	gEvents := make([]*event.Event, 0)

	for _, eventElem := range events {
		gEvents = append(gEvents, &event.Event{Id: eventElem.Id, Text: eventElem.Text})
	}

	evsList := event.EventList{Events: gEvents}

	if err != nil {
		resp := &event.ListEventRes{
			Result: &event.ListEventRes_Error{Error: err.Error()},
		}
		return resp, nil
	}

	resp := &event.ListEventRes{
		Result: &event.ListEventRes_Events{
			Events: &evsList,
		},
	}

	return resp, nil
}
