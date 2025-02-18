package rabbitMq

import (
	"encoding/json"
	"fmt"
	"main/database"
	"main/infrastructure/serviceBus/interfaces"

	"gorm.io/gorm"
)

const (
	EXCHANGE_NAME = "restaurants.changes"
	QUEUE_NAME    = "search-engine"
)

type RestaurantChangesConsumer struct {
	exchangeName string
	queueName    string
	database     *gorm.DB
}
type RabbitMqMessage struct {
	Message RestaurantQueueModel `json:"message"`
}

type RestaurantQueueModel struct {
	Name  string  `json:"name"`
	XAxis float64 `json:"xAxis"`
	YAxis float64 `json:"yAxis"`
}

// can reate cfg struct later
func NewConsumer(exchangeName string, queueName string, database *gorm.DB) interfaces.ServiceBusConsumer {
	return &RestaurantChangesConsumer{
		exchangeName: exchangeName,
		queueName:    queueName,
		database:     database,
	}
}

func (consumer *RestaurantChangesConsumer) Consume(body []byte) error {
	var msg RabbitMqMessage
	json.Unmarshal(body, &msg)

	restaurantDb := convertToRestaurant(msg.Message)
	consumer.database.Create(&restaurantDb)
	return nil
}

func (consumer *RestaurantChangesConsumer) GetExchange() string {
	return consumer.exchangeName
}
func (consumer *RestaurantChangesConsumer) GetQueueName() string {
	return consumer.queueName
}

func convertToRestaurant(model RestaurantQueueModel) database.Restaurant {
	geom := fmt.Sprintf("SRID=4326;POINT(%f %f)", model.XAxis, model.YAxis)
	return database.Restaurant{
		Name: model.Name,
		Geom: geom,
	}
}
