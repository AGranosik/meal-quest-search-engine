package menuBusService

import (
	"encoding/json"
	"main/infrastructure/serviceBus/interfaces"

	"gorm.io/gorm"
)

type MenuChangesConsumer struct {
	exchangeName string
	queueName    string
	database     *gorm.DB
}

type MenuMessage struct {
	Message MenuQueueModel `json:"message"`
}

type MenuQueueModel struct {
	RestaurantId int               `json:"restaurantId"`
	Name         string            `json:"name"`
	Groups       []GroupQueueModel `json:"groups"`
}

type GroupQueueModel struct {
	Name  string           `json:"name"`
	Meals []MealQueueModel `json:"meals"`
}

type MealQueueModel struct {
	Categories  []CategoryQueueModel   `json:"categories"`
	Ingredients []IngredientQueueModel `json:"ingredients"`
	Name        string                 `json:"name"`
	Price       string                 `json:"price"`
}

type IngredientQueueModel struct {
	Name string `json:"name"`
}

type CategoryQueueModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Consume implements interfaces.ServiceBusConsumer.
func (m *MenuChangesConsumer) Consume(body []byte) error {
	var msg MenuMessage
	json.Unmarshal(body, &msg)

	return nil
}

// TODO: REFACTOR
func (m *MenuChangesConsumer) GetExchange() string {
	return m.exchangeName
}

// GetQueueName implements interfaces.ServiceBusConsumer.
func (m *MenuChangesConsumer) GetQueueName() string {
	return m.queueName
}

func NewConsumer(exchangeName string, queueName string, database *gorm.DB) interfaces.ServiceBusConsumer {
	return &MenuChangesConsumer{
		exchangeName: exchangeName,
		queueName:    queueName,
		database:     database,
	}
}
