package main

import configration "main/infrastructure/database/migrations/configurations"

func main() {
	configration.ConfigureDatabase()
	configration.MigrateModels()
}
