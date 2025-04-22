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
	RestaurantId int       `json:"restaurantId"`
	Menu         MenuModel `json:"menu"`
}

type MenuModel struct {
	Name   NotEmptyString `json:"name"`
	Groups []GroupModel   `json:"groups"`
}

type GroupModel struct {
	GroupName NotEmptyString `json:"GroupName"`
	Meals     []MealModel    `json:"meals"`
}

type MealModel struct {
	Name        NotEmptyString    `json:"name"`
	Price       float32           `json:"price"`
	Categories  []CategoryModel   `json:"categories"`
	Ingredients []IngredientModel `json:"ingredients"`
}

type IngredientModel struct {
	Name NameValue `json:"name"`
}

type CategoryModel struct {
	Value NameValue `json:"value"`
}

type NotEmptyString struct {
	Value NameValue `json:"value"`
}

type NameValue struct {
	Value string `json:"value"`
}

// Consume implements interfaces.ServiceBusConsumer.
func (m *MenuChangesConsumer) Consume(body []byte) error {
	var msg MenuMessage
	json.Unmarshal(body, &msg)
	// fmt.Println(json)
	panic("unimplemented")
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
