package main

import (
	dbConfiguration "meal-quest/search-engine/database/migrations/configurations"
)

func main() {

	dbConfiguration.ConfigureDatabase()
	dbConfiguration.MigrateModels()
}
