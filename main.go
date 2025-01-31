package main

import (
	"fmt"
	"log"
	"meal-quest/search-engine/infrastructure/serviceBus"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// db, err := createDbConnection()

	rabbit := serviceBus.CreateRabbitMq()

	rabbit = rabbit.Start().
		WithExchange("restaurant.changes").
		WithQueue("search-engine", "restaurant.changes")

	rabbit.Consume()
	fmt.Println("Up & Running.")

	var forever chan struct{}

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
