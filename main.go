package main

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// db, err := createDbConnection()

	//error checking
	// failOnError(err, "Failed to connect to database")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"restaurant.changes",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	queue, err := ch.QueueDeclare(
		"search-engine",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		queue.Name,
		"",
		"restaurant.changes",
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")
	fmt.Println("Up & Running.")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func createDbConnection() (db *gorm.DB, err error) {
	databaseServer := os.Getenv("DATABASE_SERVER")
	searchEngineDatabase := os.Getenv("SEARCH_ENGINE_DATABASE")
	return gorm.Open(postgres.Open(fmt.Sprintf("%s/%s", databaseServer, searchEngineDatabase)), &gorm.Config{})
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
