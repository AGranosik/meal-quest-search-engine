package rabbitMq

import (
	"search-engine/infrastructure/serviceBus"
)

const (
	EXCHANGE_NAME = "restaurant.changes"
	QUEUE_NAME    = "search-engine"
)

type RestaurantChangesConsumer struct {
	exchangeName string
	queueName    string
}

// can reate cfg struct later
func NewConsumer(exchangeName string, queueName string, busService serviceBus.ServiceBusProvider) serviceBus.ServiceBusConsumer {
	return &RestaurantChangesConsumer{
		exchangeName: exchangeName,
		queueName:    queueName,
	}
}

func (consumer *RestaurantChangesConsumer) Consume(body []byte) error {
	return nil
}
