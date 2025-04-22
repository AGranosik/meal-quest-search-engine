package serviceBus

import (
	"main/infrastructure/serviceBus/interfaces"
	"main/infrastructure/serviceBus/rabbitMq"
	menuBusService "main/infrastructure/serviceBus/rabbitMq/menu"

	"gorm.io/gorm"
)

func ConfigureServiceBusProvider(db *gorm.DB) {
	rabbit := rabbitMq.CreateRabbitMq()

	rabbit.Start()
	configureRestaurantChangesConsumption(rabbit, db)
	configureMenuChangesConsumption(rabbit, db)
}

// TODO: vhosts or sth different
func configureRestaurantChangesConsumption(serviceBus interfaces.ServiceBusProvider, db *gorm.DB) {
	consumer := rabbitMq.NewConsumer("restaurants.changes", "se-restaurants", db)
	serviceBus.Consume(consumer)
}

func configureMenuChangesConsumption(serviceBus interfaces.ServiceBusProvider, db *gorm.DB) {
	consuemr := menuBusService.NewConsumer("menus.changes", "se-menus", db)
	serviceBus.Consume(consuemr)
}
