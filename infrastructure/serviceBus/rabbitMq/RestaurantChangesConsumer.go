package rabbitMq

import (
	"encoding/json"
	"main/infrastructure/serviceBus/interfaces"
)

const (
	EXCHANGE_NAME = "restaurants.changes"
	QUEUE_NAME    = "search-engine"
)

type RestaurantChangesConsumer struct {
	exchangeName string
	queueName    string
}
type RabbitMqMessage struct {
	Message RestaurantQueueModel `json:"message"`
}

type RestaurantQueueModel struct {
	Name string `json:"name"`
}

// can reate cfg struct later
func NewConsumer(exchangeName string, queueName string) interfaces.ServiceBusConsumer {
	return &RestaurantChangesConsumer{
		exchangeName: exchangeName,
		queueName:    queueName,
	}
}

func (consumer *RestaurantChangesConsumer) Consume(body []byte) error {
	var msg RabbitMqMessage
	json.Unmarshal(body, &msg)
	return nil
}

func (consumer *RestaurantChangesConsumer) GetExchange() string {
	return consumer.exchangeName
}
func (consumer *RestaurantChangesConsumer) GetQueueName() string {
	return consumer.queueName
}
