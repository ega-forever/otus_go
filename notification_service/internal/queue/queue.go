package queue

import (
	"encoding/json"
	"github.com/ega-forever/otus_go/notification_service/internal/domain/models"
	"github.com/streadway/amqp"
	"log"
)

type Queue struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    *amqp.Queue
	Msg  <-chan amqp.Delivery
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

	q, err := ch.QueueDeclare(
		"events_subscriber", // name
		false,               // durable
		false,               // delete when unused
		true,                // exclusive
		false,               // noWait
		nil,                 // arguments
	)

	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,     // queue name
		"events",   // routing key
		"platform", // exchange
		false,
		nil)

	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	return &Queue{conn: conn, ch: ch, q: &q, Msg: msgs}, nil
}

func (q *Queue) Subscribe() <-chan models.Event {

	ch := make(chan models.Event)

	go func() {

		for ev := range q.Msg {
			event := models.Event{}
			err := json.Unmarshal(ev.Body, &event)
			if err == nil {
				ch <- event
			} else {
				log.Println(err)
			}
		}
		log.Println("inside")

	}()

	return ch
}
