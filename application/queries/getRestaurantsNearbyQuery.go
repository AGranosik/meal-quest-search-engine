package queries

import (
	"encoding/base64"

	"gorm.io/gorm"
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
	Latx       string
	Laty       string
	PageSize   int
	PageNumber int
}

func GetRestaurantNearby(query ResurantsNearbyQuery, db *gorm.DB) ([]RestaurantNerbyDto, error) {
	var restuarants []RestaurantNerbyQueryModel
	offset := (query.PageNumber - 1) * query.PageSize
	tx := db.Raw(`
			SELECT restaurant_id, name, logo, description, ST_Distance(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography) AS distance
			FROM public.restaurants
			ORDER BY geom <-> ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography
			LIMIT ? OFFSET ?
			`, query.Latx, query.Laty, query.Latx, query.Laty, query.PageSize, offset).
		Scan(&restuarants)

	if tx.Error != nil {
		return []RestaurantNerbyDto{}, tx.Error
	}
	return MapRestaurantsToDto(restuarants), tx.Error
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
