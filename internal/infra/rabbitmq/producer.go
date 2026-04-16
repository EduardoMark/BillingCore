package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Producer struct{}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) Publish(ctx context.Context, queue string, message any) error {
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("Publish error marshaling message: %w", err)
	}

	conn, ch, err := Connection()
	if err != nil {
		return fmt.Errorf("Publish error getting RabbitMQ connection: %w", err)
	}
	defer conn.Close()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)
	if err != nil {
		return fmt.Errorf("Publish error declaring queue: %w", err)
	}

	err = ch.PublishWithContext(
		ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msgBytes,
		},
	)
	if err != nil {
		return fmt.Errorf("Publish error publishing message: %w", err)
	}

	return nil
}
