package comment

import (
	"log"
	"ttbackend/database"
	"ttbackend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateComment(c *fiber.Ctx) error {
	var comment models.Comment

	if err := c.BodyParser(&comment); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	// Check if the user exists
	var user models.User
	if result := database.Database.Db.First(&user, comment.UserID); result.Error != nil {
		return c.Status(400).SendString("User not found")
	}

	// Check if the story exists
	var story models.Story
	if result := database.Database.Db.First(&story, comment.StoryID); result.Error != nil {
		return c.Status(400).SendString("Story not found")
	}

	// Create the comment
	if result := database.Database.Db.Create(&comment); result.Error != nil {
		return c.Status(400).SendString("Failed to create comment")
	}

	log.Printf("Comment created: %v", comment)
	return c.Status(201).JSON(comment)
}
