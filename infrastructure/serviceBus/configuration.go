package serviceBus

import (
	"main/infrastructure/serviceBus/interfaces"
	"main/infrastructure/serviceBus/rabbitMq"

	"gorm.io/gorm"
)

// refactor
func ConfigureServiceBusProvider(db *gorm.DB) {
	rabbit := rabbitMq.CreateRabbitMq()

	rabbit.Start()
	configureRestaurantChangesConsumption(rabbit, db)
}

func configureRestaurantChangesConsumption(serviceBus interfaces.ServiceBusProvider, db *gorm.DB) {
	consumer := rabbitMq.NewConsumer("restaurants.changes", "search-engine", db)
	serviceBus.Consume(consumer)
}
