package story

import (
	"log"
	"ttbackend/database"
	"ttbackend/models"

	"github.com/gofiber/fiber/v2"
)

// CreateStoryRequest is a representation of the request body for creating a story

func CreateStory(c *fiber.Ctx) error {
	var story models.Story

	if err := c.BodyParser(&story); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	result := database.Database.Db.First(&story, story.AuthorID)
	if result.Error != nil {
		return c.Status(400).SendString("Author not found")
	}

	result = database.Database.Db.Create(&story)
	if result.Error != nil {
		return c.Status(400).SendString("Failed to create story")
	}

	log.Printf("Story created: %v", story)
	return c.Status(201).JSON(story)
}
