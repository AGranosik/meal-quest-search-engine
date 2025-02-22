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

func (rabbit *RabbitMqServiceBusProvider) Consume(consumer interfaces.ServiceBusConsumer) interfaces.ServiceBusProvider {

	channel := rabbit.configureExchange(consumer.GetExchange())
	rabbit.configureQueue(channel, consumer.GetQueueName(), consumer.GetExchange())

	err := channel.Qos(1, 0, true)

	if err != nil {
		log.Panicln(err)
	}

	msgs, err := channel.Consume(
		consumer.GetQueueName(), // queue
		"search-engine-app",     // consumer
		false,                   // auto-ack
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
			err = consumer.Consume(d.Body)
			if err == nil {
				d.Ack(false)
			} else {
				d.Nack(false, false)
			}
		}
	}()

	return rabbit
}

func (rabbit *RabbitMqServiceBusProvider) Stop() {
	rabbit.connection.Close()
}

func (rabbit *RabbitMqServiceBusProvider) configureExchange(exchangeName string) *amqp.Channel {
	channel, err := rabbit.connection.Channel()
	if err != nil {
		log.Panicln(err)
	}

	createExchange(exchangeName, channel)
	createExchange(exchangeName+".dlx", channel)
	return channel
}

func (rabbit *RabbitMqServiceBusProvider) configureQueue(channel *amqp.Channel, queueName string, exchange string) {

	dlxQueue := queueName + ".dlx"
	dlxEx := exchange + ".dlx"

	createQueue(channel, dlxQueue, dlxEx, nil)
	createQueue(channel, queueName, exchange, amqp.Table{
		"x-dead-letter-exchange": dlxEx,
	})
}

func createExchange(name string, channel *amqp.Channel) {
	err := channel.ExchangeDeclare(
		name,
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
}

func createQueue(channel *amqp.Channel, queueName string, exchange string, args amqp.Table) {
	_, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		args,
	)

	if err != nil {
		log.Panicln(err)
	}

	err = channel.QueueBind(
		queueName,
		"",
		exchange,
		false,
		nil,
	)
	if err != nil {
		log.Panicln(err)
	}
}
