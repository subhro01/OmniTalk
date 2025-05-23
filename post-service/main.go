package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/subhro01/reddit-clone/post-service/database"
	"github.com/subhro01/reddit-clone/post-service/handlers"
)

func main() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to DB
	database.ConnectDB()

	// Fiber app
	app := fiber.New()

	// Define routes
	app.Get("/posts", handlers.GetPosts)
	app.Get("/posts/:id", handlers.GetPostById)
	app.Post("/posts", handlers.CreatePost)
	app.Put("/posts/:id", handlers.UpdatePost)
	app.Delete("/posts/:id", handlers.DeletePost)

	// Get port from env or fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	fmt.Println("Post Service running on port " + port)
	log.Fatal(app.Listen(":" + port))
}
