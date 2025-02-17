package serviceBus

import (
	"main/infrastructure/serviceBus/interfaces"
	"main/infrastructure/serviceBus/rabbitMq"
)

func ConfigureServiceBusProvider() {
	rabbit := rabbitMq.CreateRabbitMq()

	rabbit.Start()
	configureRestaurantChangesConsumption(rabbit)
}

func configureRestaurantChangesConsumption(serviceBus interfaces.ServiceBusProvider) {
	consumer := rabbitMq.NewConsumer("restaurants.changes", "search-engine")
	serviceBus.Consume(consumer)
}
