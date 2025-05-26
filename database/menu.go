package database

const MenuTableName = "menus"
const GroupTableName = "groups"
const MealTableName = "meals"
const CategoryTableName = "categories"
const IngredientsTableName = "ingredients"

type Menu struct {
	ID           uint `gorm:"primaryKey"`
	RestaurantID uint
	Restaurant   Restaurant
	Groups       []Group
}

type Group struct {
	ID     uint `gorm:"primaryKey"`
	MenuID uint
	Menu   Menu
	Name   string
	Meals  []Meal
}

type Meal struct {
	ID          uint `gorm:"primaryKey"`
	GroupID     uint
	Group       Group
	Name        string
	Price       float32      `gorm:"type:numeric(6,2)"`
	Categories  []Category   `gorm:"many2many:meal_categories;"`
	Ingredients []Ingredient `gorm:"many2many:meal_ingredients;"`
}

type Category struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"uniqueIndex"`
	Meals []Meal `gorm:"many2many:meal_categories;"`
}

type Ingredient struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Meals []Meal `gorm:"many2many:meal_ingredients;"`
}
