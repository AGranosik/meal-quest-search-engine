package database

type Restaurant struct {
	RestaurantId uint `gorm:"primaryKey"`
	XCoordinate  float64
	YCoordinate  float64
}
