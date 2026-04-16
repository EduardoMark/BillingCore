package rabbitmq

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connection() (*amqp.Connection, *amqp.Channel, error) {
	url := os.Getenv("RABBITMQ_AMQP_URL")
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, fmt.Errorf("Publish error connecting to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("Publish error creating channel: %w", err)
	}

	return conn, ch, nil
}
