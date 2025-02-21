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

	err = channel.ExchangeDeclare(
		exchangeName+".dlx",
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

// refactor
func (rabbit *RabbitMqServiceBusProvider) configureQueue(channel *amqp.Channel, queueName string, exchange string) {

	// _, err := channel.QueueDeclare(
	// 	queueName,
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )

	// if err != nil {
	// 	log.Panicln(err)
	// }

	// err = channel.QueueBind(
	// 	queueName,
	// 	"",
	// 	exchange,
	// 	false,
	// 	nil,
	// )

	// if err != nil {
	// 	log.Panicln(err)
	// }

	_, err := channel.QueueDeclare(
		queueName+".dlx",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Panicln(err)
	}

	err = channel.QueueBind(
		queueName+".dlx",
		"",
		exchange+".dlx",
		false,
		nil,
	)
	if err != nil {
		log.Panicln(err)
	}
	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange": exchange + ".dlx",
		},
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
			// deadletter-exchange to ten sam exchangr wiec się zapętla
			err = consumer.Consume(d.Body)

			err = d.Nack(false, false)

			if err != nil {
				log.Panicln(err)
			}
			// if err == nil {
			// 	d.Ack(false)
			// } else {
			// 	d.Nack(false, false)
			// }
		}
	}()

	return rabbit
}

func (rabbit *RabbitMqServiceBusProvider) Stop() {
	rabbit.connection.Close()
}
