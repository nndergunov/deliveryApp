package subscriber

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Subscriber struct {
	conn    *amqp.Connection
	ch      *amqp.Channel
	msgChan chan []byte
}

func NewSubscriber(url string, msgChan chan []byte) (*Subscriber, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("connecting to messagebroker: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("opening channel: %w", err)
	}

	return &Subscriber{
		conn:    conn,
		ch:      ch,
		msgChan: msgChan,
	}, nil
}

func (s Subscriber) SubscribeToTheme(theme string) error {
	err := s.ch.ExchangeDeclare(
		theme,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("declaring exchange: %w", err)
	}

	queue, err := s.ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("declaring queue: %w", err)
	}

	err = s.ch.QueueBind(
		queue.Name,
		"",
		theme,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("binding queue: %w", err)
	}

	messages, err := s.ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("registering consumer: %w", err)
	}

	go func() {
		for delivery := range messages {
			s.msgChan <- delivery.Body
		}
	}()

	return nil
}
