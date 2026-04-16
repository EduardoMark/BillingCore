package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Consume(
	ctx context.Context,
	queue string,
	handle func(ctx context.Context, msg []byte) error,
) error {
	conn, ch, err := Connection()
	if err != nil {
		return fmt.Errorf("Consume error getting RabbitMQ connection: %w", err)
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
		return fmt.Errorf("Consume error declaring queue: %w", err)
	}

	msgs, err := ch.ConsumeWithContext(
		ctx,
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Consume error consuming messages: %w", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-msgs:
				if err := handle(ctx, msg.Body); err != nil {
					msg.Nack(false, true)
					continue
				}

				msg.Ack(false)
			}
		}
	}()

	<-ctx.Done()

	return nil
}
