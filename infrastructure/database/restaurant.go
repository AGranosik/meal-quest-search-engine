package database

const RestaurantTableName = "restaurants"

type Restaurant struct {
	RestaurantId uint   `gorm:"primaryKey"`
	Description  string `gorm:"size:3000;not null" json:"description"`
	StreetName   string `gorm:"size:430;not null" json:"streetname"`
	City         string `gorm:"size:420;not null" json:"city"`
	Name         string `gorm:"size:255;not null" json:"name"` // Maps to name VARCHAR(255) NOT NULL
	Geom         string `gorm:"type:GEOGRAPHY(Point,4326);not null" json:"geom"`
	Logo         []byte `json:"logo"`
}

func (Restaurant) TableName() string {
	return RestaurantTableName
}
