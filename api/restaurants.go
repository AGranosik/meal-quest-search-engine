package api

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func ConfigureRestaurantsEndpoints() {
	app := fiber.New()

	// Define a route for the GET method on the root path '/'
	app.Get("/restaurant", func(c fiber.Ctx) error {
		latx := c.Query("x")
		laty := c.Query("y")

		log.Println(latx)
		log.Println(laty)
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
