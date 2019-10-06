package main

import (
	"context"
	"encoding/json"
	"fmt"
	event "github.com/ega-forever/otus_go/integration_tests/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"os"
	"sync"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/streadway/amqp"
)

var amqpDSN = os.Getenv("TESTS_AMQP_DSN")

func init() {
	if amqpDSN == "" {
		amqpDSN = "amqp://guest:guest@localhost:5672/"
	}
}

const (
	queueName                 = "ToNotificationTest"
	notificationsExchangeName = "platform"
	notificationsExchangeKey  = "events"
)

type notifyTest struct {
	conn          *amqp.Connection
	ch            *amqp.Channel
	messages      [][]byte
	messagesMutex *sync.RWMutex
	stopSignal    chan struct{}
	rpcURI        string
	eventId       int64
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (test *notifyTest) startConsuming(interface{}) {
	test.messages = make([][]byte, 0)
	test.messagesMutex = new(sync.RWMutex)
	test.stopSignal = make(chan struct{})

	var err error

	test.conn, err = amqp.Dial(amqpDSN)
	panicOnErr(err)

	test.ch, err = test.conn.Channel()
	panicOnErr(err)

	// Consume
	_, err = test.ch.QueueDeclare(queueName, true, true, true, false, nil)
	panicOnErr(err)

	err = test.ch.QueueBind(queueName, notificationsExchangeKey, notificationsExchangeName, false, nil)
	panicOnErr(err)

	events, err := test.ch.Consume(queueName, "", true, true, false, false, nil)
	panicOnErr(err)

	go func(stop <-chan struct{}) {
		for {
			select {
			case <-stop:
				return
			case eventData := <-events:
				test.messagesMutex.Lock()
				test.messages = append(test.messages, eventData.Body)
				test.messagesMutex.Unlock()
			}
		}
	}(test.stopSignal)
}

func (test *notifyTest) stopConsuming(interface{}, error) {
	test.stopSignal <- struct{}{}

	panicOnErr(test.ch.Close())
	panicOnErr(test.conn.Close())
	test.messages = nil
}

func (test *notifyTest) iAbleToCreateRecord(text string, port string) error {
	test.rpcURI = viper.GetString("REST_HOST") + ":" + port
	conn, err := grpc.Dial(test.rpcURI, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := event.NewEventServiceClient(conn)

	timestamp := time.Now().Unix() * 1000

	req := &event.CreateEventReq{
		Event: &event.Event{Text: text, Timestamp: timestamp},
	}

	resp, err := client.CreateEvent(context.Background(), req)

	if resp != nil {
		test.eventId = resp.GetEvent().Id
	}

	return err
}

func (test *notifyTest) isRecordCreated() error {

	conn, err := grpc.Dial(test.rpcURI, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := event.NewEventServiceClient(conn)

	req := &event.GetEventReq{
		Id: test.eventId,
	}

	_, err = client.GetEvent(context.Background(), req)

	return err
}

func (test *notifyTest) iReceiveEventWithText(text string) error {
	time.Sleep(20 * time.Second)

	test.messagesMutex.RLock()
	defer test.messagesMutex.RUnlock()

	for _, msg := range test.messages {

		ev := event.Event{}

		err := json.Unmarshal(msg, &ev)

		if err != nil {
			return err
		}

		if ev.Text == text {
			return nil
		}
	}

	return fmt.Errorf("event with text '%s' was not found in %s", text, test.messages)
}

func FeatureContext(s *godog.Suite) {

	viper.SetDefault("REST_HOST", "localhost")
	viper.AutomaticEnv()

	test := new(notifyTest)

	s.BeforeScenario(test.startConsuming)

	s.Step(`^I create new event "([^"]*)" and request to port "([^"]*)"$`, test.iAbleToCreateRecord)
	s.Step(`^the record should be created`, test.isRecordCreated)
	s.Step(`^I receive event with text "([^"]*)"$`, test.iReceiveEventWithText)

	s.AfterScenario(test.stopConsuming)
}
