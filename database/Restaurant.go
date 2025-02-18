package database

const RestaurantTableName = "restaurants"

type Restaurant struct {
	RestaurantId uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"size:255;not null" json:"name"` // Maps to name VARCHAR(255) NOT NULL
	Geom         string `gorm:"type:GEOGRAPHY(Point,4326);not null" json:"geom"`
}

func (Restaurant) TableName() string {
	return RestaurantTableName
}
