package rabbitMq

import "search-engine/infrastructure/serviceBus"

const (
	EXCHANGE_NAME = "restaurant.changes"
	QUEUE_NAME    = "search-engine"
)

type RestaurantChangesConsumer struct {
}

func NewConsumer() *serviceBus.ServiceBusConsumer {
	return &RestaurantChangesConsumer{}
}

func (consumer *serviceBus.ServiceBusConsumer) Consume() {

}
