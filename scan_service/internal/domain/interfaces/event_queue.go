package interfaces

import (
	"github.com/ega-forever/otus_go/scan_service/internal/domain/models"
)

type EventQueue interface {
	PushEvent(event *models.Event) error
}
