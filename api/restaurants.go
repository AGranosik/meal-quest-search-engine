package api

import (
	"log"
	"main/application/queries"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func ConfigureRestaurantsEndpoints(db *gorm.DB) {
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/restaurant", func(c fiber.Ctx) error {
		latx := c.Query("x")
		laty := c.Query("y")
		pageSize, err := strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			return err
		}

		pageNumber, err := strconv.Atoi(c.Query("pageNumber"))
		if err != nil {
			return err
		}

		restaurantName := c.Query("name")

		result, err := queries.GetRestaurantNearby(queries.ResurantsNearbyQuery{
			Latx:           latx,
			Laty:           laty,
			PageSize:       pageSize,
			PageNumber:     pageNumber,
			RestaurantName: &restaurantName,
		}, db)

		if err != nil {
			return err
		}
		return c.JSON(result)
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
