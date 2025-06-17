package main

import (
	"fmt"
	"log"
	"main/api"
	"main/infrastructure/serviceBus"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := createDbConnection()

	if err != nil {
		log.Panic(err)
	}
	serviceBus.ConfigureServiceBusProvider(db)
	fmt.Println("Up & Running.")

	// var forever chan struct{}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	// <-forever

	api.ConfigureRestaurantsEndpoints(db)
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
