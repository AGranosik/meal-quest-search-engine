package main

import (
	"fmt"
	"log"
	"meal-quest/search-engine/database"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbURL := "postgres://admin:S3cret@localhost:5432/search_engine"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&database.Restaurant{})
	fmt.Println("Up & Running.")
}
