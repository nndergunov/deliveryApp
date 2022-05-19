package publisher

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type Publisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewPublisher(url string) (*Publisher, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("connecting to messagebroker: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("opening channel: %w", err)
	}

	return &Publisher{
		conn: conn,
		ch:   ch,
	}, nil
}

func (p Publisher) Publish(theme string, data []byte) error {
	err := p.ch.ExchangeDeclare(
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

	err = p.ch.Publish(
		theme,
		"",
		false,
		false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            data,
		})
	if err != nil {
		return fmt.Errorf("declaring exchange: %w", err)
	}

	return nil
}
