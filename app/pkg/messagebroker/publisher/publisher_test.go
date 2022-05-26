package publisher_test

import (
	"testing"

	v1 "github.com/nndergunov/deliveryApp/app/pkg/api/v1"
	"github.com/nndergunov/deliveryApp/app/pkg/messagebroker/publisher"
	amqp "github.com/rabbitmq/amqp091-go"
)

const hostURL = "amqp://guest:guest@localhost:5672/"

func TestNewPublisher(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		postTheme string
		postData  string
	}{
		{
			name:      "order posted mock",
			postTheme: "orders",
			postData:  "new order posted",
		},
	}

	for _, currTest := range tests {
		test := currTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			// creating a listener to accept messages
			listener, err := amqp.Dial(hostURL)
			if err != nil {
				t.Fatal(err)
			}

			listenChannel, err := listener.Channel()
			if err != nil {
				t.Fatal(err)
			}

			err = listenChannel.ExchangeDeclare(test.postTheme, "fanout", true, false, false, false, nil)
			if err != nil {
				t.Fatal(err)
			}

			listenQueue, err := listenChannel.QueueDeclare("", false, false, true, false, nil)
			if err != nil {
				t.Fatal(err)
			}

			err = listenChannel.QueueBind(listenQueue.Name, "", test.postTheme, false, nil)
			if err != nil {
				t.Fatal(err)
			}

			msgChan, err := listenChannel.Consume(listenQueue.Name, "", true, false, false, false, nil)
			if err != nil {
				t.Fatal(err)
			}

			// creating publisher, the main test subject
			publish, err := publisher.NewEventPublisher(hostURL)
			if err != nil {
				t.Fatal(err)
			}

			err = publish.Publish(test.postTheme, test.postData)
			if err != nil {
				t.Fatal(err)
			}

			// receiving the message
			message := <-msgChan

			receiveData := new(string)

			err = v1.Decode(message.Body, receiveData)
			if err != nil {
				t.Fatal(err)
			}

			if test.postData != *receiveData {
				t.Errorf("Expected: %v; Got: %v", test.postData, receiveData)
			}
		})
	}
}
