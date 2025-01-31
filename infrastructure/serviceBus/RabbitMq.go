package serviceBus

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMq struct {
	connectionString string
	connection       *amqp.Connection
	channel          *amqp.Channel
}

// error handlling
// order should't affect creation
func CreateRabbitMq() ServiceBusProvider {
	connectionString := os.Getenv("RABBITMQ")
	if connectionString == "" {
		log.Panicln("connection string cannot be empty")
	}

	return &RabbitMq{
		connectionString: connectionString,
	}
}

func (rabbit *RabbitMq) Start() ServiceBusProvider {
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

func (rabbit *RabbitMq) WithExchange(exchangeName string) ServiceBusProvider {

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
func (rabbit *RabbitMq) WithQueue(queueName string, exchange string) ServiceBusProvider {
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
func (rabbit *RabbitMq) Consume() ServiceBusProvider {
	msgs, err := rabbit.channel.Consume(
		"search-engine",     // queue
		"search-engine-app", // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
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

func (rabbit *RabbitMq) Stop() {
	rabbit.connection.Close()
	rabbit.channel.Close()
}
