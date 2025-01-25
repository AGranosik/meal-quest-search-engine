package configration

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConfigureDatabase() {
	databaseServer := os.Getenv("DATABASE_SERVER")
	searchEngineDatabase := os.Getenv("SEARCH_ENGINE_DATABASE")
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("%s/postgres", databaseServer)), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.Exec(fmt.Sprintf("CREATE DATABASE %s;", searchEngineDatabase))
}
