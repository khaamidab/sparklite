package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()
	//run database
	configs.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	//routes
	routes.UserRoute(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7001"
	}

	app.Listen(":" + port)
}
