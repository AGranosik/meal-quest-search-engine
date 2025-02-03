package rabbitMq

import (
	"log"
	"os"
	"search-engine/infrastructure/serviceBus"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqServiceBusProvider struct {
	connectionString string
	connection       *amqp.Connection
	channel          *amqp.Channel
}

// error handlling
// order should't affect creation
func CreateRabbitMq() serviceBus.ServiceBusProvider {
	connectionString := os.Getenv("RABBITMQ")
	if connectionString == "" {
		log.Panicln("connection string cannot be empty")
	}

	return &RabbitMqServiceBusProvider{
		connectionString: connectionString,
	}
}

func (rabbit *RabbitMqServiceBusProvider) Start() serviceBus.ServiceBusProvider {
	conn, err := amqp.Dial(rabbit.connectionString)

	if err != nil {
		log.Panicln(err)
	}

	rabbit.connection = conn
	ch, err := conn.Channel()

	if err != nil {
		log.Panicln(err)
	}
	rabbit.channel = ch
	return rabbit
}

func (rabbit *RabbitMqServiceBusProvider) WithExchange(exchangeName string) serviceBus.ServiceBusProvider {

	if rabbit.channel == nil {
		rabbit.Start()
	}
	err := rabbit.channel.ExchangeDeclare(
		exchangeName,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Panicln(err)
	}
	return rabbit
}
func (rabbit *RabbitMqServiceBusProvider) WithQueue(queueName string, exchange string) serviceBus.ServiceBusProvider {
	if rabbit.channel == nil {
		rabbit.Start()
	}

	queue, err := rabbit.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Panicln()
	}

	err = rabbit.channel.QueueBind(
		queue.Name,
		"",
		exchange,
		false,
		nil,
	)

	if err != nil {
		log.Panicln(err)
	}
	return rabbit
}

// some consumer entity to handle
// there should be some exchange interface which will have exchange name already
func (rabbit *RabbitMqServiceBusProvider) Consume() serviceBus.ServiceBusProvider {
	msgs, err := rabbit.channel.Consume(
		"search-engine",     // queue
		"search-engine-app", // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,
	)

	if err != nil {
		log.Panicln(err)
	}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()
	return rabbit
}

func (rabbit *RabbitMqServiceBusProvider) Stop() {
	rabbit.connection.Close()
	rabbit.channel.Close()
}
