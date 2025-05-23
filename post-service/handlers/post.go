package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/subhro01/reddit-clone/post-service/database"
    "github.com/subhro01/reddit-clone/post-service/models"
)

// GET /posts
func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	database.DB.Find(&posts)
	return c.JSON(posts)
}

// GET /posts/:id
func GetPostById(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	result := database.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	return c.JSON(post)
}

//POST /posts
func CreatePost(c *fiber.Ctx) error {
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(&post)
	return c.Status(201).JSON(post)
}

// PUT /posts/:id
func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	result := database.DB.First(&post, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}

	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	database.DB.Save(&post)
	return c.JSON(post)
}

// DELETE /posts/:id
func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	result := database.DB.Delete(&post, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	return c.SendStatus(204)
}