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
	var msg MenuQueueModel
	json.Unmarshal(body, &msg)

	return nil
}

// func convertToMenu(model MenuQueueModel) database.Menu {

// 	return database.Menu{
// 		ID:           model.Id.Value,
// 		RestaurantID: uint(model.RestaurantId),
// 	}
// }

// func convertGroups(queueGroups []GroupModel) database.Group {
// 	result := make([]database.Group, len(queueGroups))
// 	for i := 0; i < len(queueGroups); i++ {
// 		group := queueGroups[i]

// 	}
// }

// func convertMeals(queueMeals []MealModel) database.Meal {
// 	result := make([]database.Meal, len(queueMeals))
// 	for i := 0; i < len(queueMeals); i++ {

// 	}
// }

// func convertCategories(queueCategories []CategoryModel) database.Category{
// 	result := make([]database.Category, len(queueCategories))
// 	for i := 0; i < len(queueCategories); i++ {
// 		category := queueCategories[i]
// 		result[i] = database.Category{
// 			// ID: category.,
// 		}
// 	}
// }

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
