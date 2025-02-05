package rabbitMq

import "main/infrastructure/serviceBus/interfaces"

const (
	EXCHANGE_NAME = "restaurant.changes"
	QUEUE_NAME    = "search-engine"
)

type RestaurantChangesConsumer struct {
	exchangeName string
	queueName    string
}

// can reate cfg struct later
func NewConsumer(exchangeName string, queueName string, busService interfaces.ServiceBusProvider) interfaces.ServiceBusConsumer {
	return &RestaurantChangesConsumer{
		exchangeName: exchangeName,
		queueName:    queueName,
	}
}

func (consumer *RestaurantChangesConsumer) Consume(body []byte) error {
	return nil
}

func (consumer *RestaurantChangesConsumer) GetExchange() string {
	return consumer.exchangeName
}
func (consumer *RestaurantChangesConsumer) GetQueueName() string {
	return consumer.queueName
}
