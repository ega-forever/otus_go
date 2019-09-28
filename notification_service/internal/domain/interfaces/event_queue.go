package interfaces

import "github.com/ega-forever/otus_go/notification_service/internal/domain/models"

type EventQueue interface {
	Subscribe() <-chan models.Event
}
