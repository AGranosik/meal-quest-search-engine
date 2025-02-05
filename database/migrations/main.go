package main

import configration "main/database/migrations/configurations"

func main() {
	configration.ConfigureDatabase()
	configration.MigrateModels()
}
