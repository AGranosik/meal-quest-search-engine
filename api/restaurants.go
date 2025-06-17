package api

import (
	"log"
	"main/database"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func ConfigureRestaurantsEndpoints(db *gorm.DB) {
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/restaurant", func(c fiber.Ctx) error {
		latx := c.Query("x")
		laty := c.Query("y")

		log.Println(latx)
		log.Println(laty)

// SELECT 
// *,
//   ST_Distance(geom, ST_SetSRID(ST_MakePoint(18.786309, 53.447191), 4326)::geography) AS distance_m
// FROM 
//   public.restaurants
// ORDER BY 
//   geom <-> ST_SetSRID(ST_MakePoint(18.786309, 53.447191), 4326)::geography
// LIMIT 5;
		var restuarants []database.Restaurant
		err := db.Table("restaurants").
			Where(`ST_X(geom::geometry) = ? and ST_Y(geom::geometry) = ?`, latx, laty).
			Scan(&restuarants).Error
		log.Println(err)
		// Send a string response to the client
		return c.JSON(restuarants)
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
