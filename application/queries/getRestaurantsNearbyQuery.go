package queries

import "gorm.io/gorm"

type RestaurantNerbyDto struct {
	RestaurantId int     `json: "restaurantId"`
	Name         string  `json: "name"`
	Distance     float32 `json: "distance"`
}

type ResurantsNearbyQuery struct {
	Latx string
	Laty string
}

func GetRestaurantNearby(query ResurantsNearbyQuery, db *gorm.DB) ([]RestaurantNerbyDto, error) {
	var restuarants []RestaurantNerbyDto
	tx := db.Raw(`
			SELECT restaurant_id, name, ST_Distance(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography) AS distance
			FROM public.restaurants
			ORDER BY geom <-> ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography
			LIMIT 5
			`, query.Latx, query.Laty, query.Latx, query.Laty).
		Scan(&restuarants)

	return restuarants, tx.Error
}
