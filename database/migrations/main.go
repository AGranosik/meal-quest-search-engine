package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbURL := "postgres://admin:S3cret@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.Exec("CREATE DATABASE search_engine;")

	searchEngine, err := gorm.Open(postgres.Open("postgres://admin:S3cret@localhost:5432/search_engine"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	searchEngine.Migrator()
	searchEngine.Exec("CREATE EXTENSION postgis;")
	db.AutoMigrate(&Restaurant)

}
