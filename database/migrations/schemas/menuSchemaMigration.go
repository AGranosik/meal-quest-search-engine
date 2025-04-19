package schemas

import (
	"main/database"

	"gorm.io/gorm"
)

func MigrateMenu(searchEngine *gorm.DB) {
	searchEngine.AutoMigrate(
		&database.Ingredient{},
		&database.Category{},
		&database.Meal{},
		&database.Group{},
		&database.Menu{},
	)
}
