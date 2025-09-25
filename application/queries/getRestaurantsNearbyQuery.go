package queries

import (
	"encoding/base64"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RestaurantNerbyQueryModel struct {
	RestaurantId int     `json:"restaurantId"`
	Name         string  `json:"name"`
	Distance     float32 `json:"distance"`
	Logo         []byte  `json:"logo"`
	Description  string  `json:"description"`
}

type RestaurantNerbyDto struct {
	RestaurantId int     `json:"restaurantId"`
	Name         string  `json:"name"`
	Distance     float32 `json:"distance"`
	Logo         string  `json:"logo"`
	Description  string  `json:"description"`
}

type ResurantsNearbyQuery struct {
	Latx           string
	Laty           string
	PageSize       int
	PageNumber     int
	RestaurantName *string
}

func GetRestaurantNearby(query ResurantsNearbyQuery, db *gorm.DB) ([]RestaurantNerbyDto, error) {
	var restaurants []RestaurantNerbyQueryModel
	offset := (query.PageNumber - 1) * query.PageSize

	// Base query
	tx := db.Model(&RestaurantNerbyQueryModel{}).
		Table("restaurants").
		Select(`restaurant_id, name, logo, description,
		        ST_Distance(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography) AS distance`,
			query.Latx, query.Laty).
		Clauses(clause.OrderBy{
			Expression: clause.Expr{SQL: `geom <-> ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography`, Vars: []interface{}{
				query.Latx, query.Laty,
			}},
		}).
		Limit(query.PageSize).
		Offset(offset)

	// Optional name filter
	if query.RestaurantName != nil && *query.RestaurantName != "" {
		tx = tx.Where("name ILIKE ?", "%"+*query.RestaurantName+"%")
	}

	// Execute
	if err := tx.Scan(&restaurants).Error; err != nil {
		return []RestaurantNerbyDto{}, err
	}

	return MapRestaurantsToDto(restaurants), nil
}

func MapRestaurantsToDto(restaurants []RestaurantNerbyQueryModel) []RestaurantNerbyDto {
	dtos := make([]RestaurantNerbyDto, len(restaurants))
	for i, r := range restaurants {
		var logoBase64 string
		if len(r.Logo) > 0 {
			logoBase64 = base64.StdEncoding.EncodeToString(r.Logo)
		}

		dtos[i] = RestaurantNerbyDto{
			RestaurantId: int(r.RestaurantId),
			Name:         r.Name,
			Distance:     r.Distance,
			Logo:         logoBase64,
			Description:  r.Description,
		}
	}
	return dtos
}
