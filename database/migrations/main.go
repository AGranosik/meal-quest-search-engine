package main

import configration "search-engine/database/migrations/configurations"

func main() {
	configration.ConfigureDatabase()
	configration.MigrateModels()
}
