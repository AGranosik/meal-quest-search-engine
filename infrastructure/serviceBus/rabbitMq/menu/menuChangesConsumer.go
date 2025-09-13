package menuBusService

import (
	"encoding/json"
	"fmt"
	"main/infrastructure/database"
	"main/infrastructure/serviceBus/interfaces"
	"strconv"

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

	dbModel, _ := mapMenuQueueModelToMenu(msg.Message)

	result := m.database.Create(&dbModel)
	if result.Error != nil {
		fmt.Errorf(result.Error.Error())
		return result.Error
	}
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

func mapMenuQueueModelToMenu(dto MenuQueueModel) (database.Menu, error) {
	menu := database.Menu{
		RestaurantID: uint(dto.RestaurantId),
		Groups:       make([]database.Group, len(dto.Groups)),
	}

	for i, groupDTO := range dto.Groups {
		group, err := mapGroupQueueModelToGroup(groupDTO)
		if err != nil {
			return database.Menu{}, fmt.Errorf("error in group %d: %w", i, err)
		}
		menu.Groups[i] = group
	}

	return menu, nil
}

// Maps GroupQueueModel to Group
func mapGroupQueueModelToGroup(dto GroupQueueModel) (database.Group, error) {
	group := database.Group{
		Name:  dto.Name,
		Meals: make([]database.Meal, len(dto.Meals)),
	}

	for i, mealDTO := range dto.Meals {
		meal, err := mapMealQueueModelToMeal(mealDTO)
		if err != nil {
			return database.Group{}, fmt.Errorf("error in meal %d: %w", i, err)
		}
		group.Meals[i] = meal
	}

	return group, nil
}

// Maps MealQueueModel to Meal
func mapMealQueueModelToMeal(dto MealQueueModel) (database.Meal, error) {
	price, err := strconv.ParseFloat(dto.Price, 32)
	if err != nil {
		return database.Meal{}, fmt.Errorf("invalid price '%s': %w", dto.Price, err)
	}

	meal := database.Meal{
		Name:  dto.Name,
		Price: float32(price),
	}

	meal.Categories = mapCategoryQueueModels(dto.Categories)
	meal.Ingredients = mapIngredientQueueModels(dto.Ingredients)

	return meal, nil
}

// Maps slice of CategoryQueueModel to []Category
func mapCategoryQueueModels(dtos []CategoryQueueModel) []database.Category {
	categories := make([]database.Category, len(dtos))
	for i, dto := range dtos {
		categories[i] = database.Category{
			ID:   uint(dto.Id),
			Name: dto.Name,
		}
	}
	return categories
}

// Maps slice of IngredientQueueModel to []Ingredient
func mapIngredientQueueModels(dtos []IngredientQueueModel) []database.Ingredient {
	ingredients := make([]database.Ingredient, len(dtos))
	for i, dto := range dtos {
		ingredients[i] = database.Ingredient{
			Name: dto.Name,
		}
	}
	return ingredients
}
