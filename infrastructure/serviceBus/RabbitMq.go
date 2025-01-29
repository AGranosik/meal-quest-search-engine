package serviceBus

import (
	"errors"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	connectionString string
	connection       *amqp.Connection
	channel          *amqp.Channel
}

func CreateRabbitMq() (ServiceBusProvider, error) {
	connectionString := os.Getenv("RABBITMQ")
	if connectionString == "" {
		return nil, errors.New("Connection string cannot be empty")
	}

	return &RabbitMq{
		connectionString: connectionString,
	}, nil
}

func (rabbit *RabbitMq) Start() (ServiceBusProvider, error) {
	conn, err := amqp.Dial(rabbit.connectionString)

	if err != nil {
		return rabbit, err
	}

	rabbit.connection = conn
	return rabbit, nil
}

func (rabbit *RabbitMq) Consume(f func()) (ServiceBusProvider, error) {
	return rabbit, nil
}

func (rabbit *RabbitMq) Stop() {
	rabbit.connection.Close()
	// return rabbit
}
