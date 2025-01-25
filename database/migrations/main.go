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
	tx := searchEngine.Exec("CREATE EXTENSION postgis;")

	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}
	tx = searchEngine.Exec("SELECT PostGIS_Version();")

	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	//restaurant
	tx = searchEngine.Exec(`
        CREATE TABLE IF NOT EXISTS restaurants (
		restaurant_id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		geom GEOGRAPHY(Point, 4326) NOT NULL
        );
    `)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}
}
