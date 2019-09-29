package queue

import (
	"encoding/json"
	"github.com/ega-forever/otus_go/scan_service/internal/domain/models"
	"github.com/streadway/amqp"
)

type Queue struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    *amqp.Queue
}

func New(uri string) (*Queue, error) {

	conn, err := amqp.Dial(uri)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"platform", // name
		"topic",    // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)

	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"events_publisher",
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,
		"events",
		"platform",
		false,
		nil)

	return &Queue{conn: conn, ch: ch, q: &q}, nil
}

func (q *Queue) PushEvent(event *models.Event) error {

	marshaled, err := json.Marshal(event)

	if err != nil {
		return err
	}

	err = q.ch.Publish(
		"platform",
		"events",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			ReplyTo:     q.q.Name,
			Body:        marshaled,
		})

	if err != nil {
		return err
	}

	return nil
}
