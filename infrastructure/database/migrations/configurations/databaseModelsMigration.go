package configration

import (
	"fmt"
	"log"
	"main/infrastructure/database"
	"main/infrastructure/database/migrations/schemas"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigrateModels() {
	databaseServer := os.Getenv("DATABASE_SERVER")
	searchEngineDatabase := os.Getenv("SEARCH_ENGINE_DATABASE")
	searchEngine, err := gorm.Open(postgres.Open(fmt.Sprintf("%s/%s", databaseServer, searchEngineDatabase)), &gorm.Config{})

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
	//creates street name twice for some reason
	tx = searchEngine.Exec(fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
		restaurant_id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description VARCHAR(3000) NOT NULL,
		streetName VARCHAR(430) NOT NULL,
		city VARCHAR(420) NOT NULL,
		geom GEOGRAPHY(Point, 4326) NOT NULL
        );
    `, database.RestaurantTableName))
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}

	schemas.MigrateMenu(searchEngine)
}
