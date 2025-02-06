package rabbitMq

import (
	"log"
	"main/infrastructure/serviceBus/interfaces"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMqServiceBusProvider struct {
	connectionString string
	connection       *amqp.Connection
}

func CreateRabbitMq() interfaces.ServiceBusProvider {
	connectionString := os.Getenv("RABBITMQ")
	if connectionString == "" {
		log.Panicln("connection string cannot be empty")
	}

	return &RabbitMqServiceBusProvider{
		connectionString: connectionString,
	}
}

func (rabbit *RabbitMqServiceBusProvider) Start() interfaces.ServiceBusProvider {
	conn, err := amqp.Dial(rabbit.connectionString)

	if err != nil {
		log.Panicln(err)
	}

	rabbit.connection = conn
	return rabbit
}

func (rabbit *RabbitMqServiceBusProvider) configureExchange(exchangeName string) *amqp.Channel {
	channel, err := rabbit.connection.Channel()

	if err != nil {
		log.Panicln(err)
	}
	err = channel.ExchangeDeclare(
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
	return channel
}
func (rabbit *RabbitMqServiceBusProvider) configureQueue(channel *amqp.Channel, queueName string, exchange string) {

	queue, err := channel.QueueDeclare(
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

	err = channel.QueueBind(
		queue.Name,
		"",
		exchange,
		false,
		nil,
	)

	if err != nil {
		log.Panicln(err)
	}
}

func (rabbit *RabbitMqServiceBusProvider) Consume(consumer interfaces.ServiceBusConsumer) interfaces.ServiceBusProvider {
	//instead exchange - queue there could be some kind of 'group' and 'type'
	channel := rabbit.configureExchange(consumer.GetExchange())
	rabbit.configureQueue(channel, consumer.GetQueueName(), consumer.GetExchange())

	msgs, err := channel.Consume(
		consumer.GetQueueName(), // queue
		"search-engine-app",     // consumer
		true,                    // auto-ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-wait
		nil,
	)

	if err != nil {
		log.Panicln(err)
	}

	go func() {
		for d := range msgs {
			consumer.Consume(d.Body)
		}
	}()

	return rabbit
}

func (rabbit *RabbitMqServiceBusProvider) Stop() {
	rabbit.connection.Close()
}
