package rabbitMq

import (
	"encoding/json"
	"fmt"
	"main/infrastructure/database"
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
	ResutaurantId int               `json:"restaurantId"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	LogoData      []byte            `json:"logoData"`
	Address       AddressQueueModel `json:"address"`
}

type AddressQueueModel struct {
	StreetName string  `json:"streetName"`
	City       string  `json:"city"`
	XAxis      float64 `json:"xAxis"`
	YAxis      float64 `json:"yAxis"`
}

// can reate cfg struct later
func NewConsumer(exchangeName string, queueName string, database *gorm.DB) interfaces.ServiceBusConsumer {
	return &RestaurantChangesConsumer{
		exchangeName: exchangeName,
		queueName:    queueName,
		database:     database,
	}
}

//TODO: logo

func (consumer *RestaurantChangesConsumer) Consume(body []byte) error {
	var msg RabbitMqMessage
	json.Unmarshal(body, &msg)

	restaurantDb := convertToRestaurant(msg.Message)
	result := consumer.database.Create(&restaurantDb)
	if result.Error != nil {
		fmt.Errorf(result.Error.Error())
		return result.Error
	}
	return nil
}

func (consumer *RestaurantChangesConsumer) GetExchange() string {
	return consumer.exchangeName
}
func (consumer *RestaurantChangesConsumer) GetQueueName() string {
	return consumer.queueName
}

func convertToRestaurant(model RestaurantQueueModel) database.Restaurant {
	geom := fmt.Sprintf("SRID=4326;POINT(%f %f)", model.Address.XAxis, model.Address.YAxis)
	return database.Restaurant{
		RestaurantId: uint(model.ResutaurantId),
		Description:  model.Description,
		StreetName:   model.Address.StreetName,
		City:         model.Address.City,
		Name:         model.Name,
		Geom:         geom,
	}
}
