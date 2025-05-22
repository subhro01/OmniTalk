package main

import (
	"github.com/gofiber/fiber/v2"
	"post-service/database"
	"post-service/models"
)

func main() {
	app := fiber.New()

	database.ConnectDB()
	database.DB.AutoMigrate(&models.Post{})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from post service")
	})

	app.Listen(":8001")
}