package Controllers

import (
	"fmt"
	"log"
	"time"

	"github.com/Anurag340/telemetryCollector/Database"
	"github.com/Anurag340/telemetryCollector/Models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func IngestData(c *fiber.Ctx) error{

	var validate = validator.New()

	var entry Models.LogEntry
	if err := c.BodyParser(&entry); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",})
	}
	fmt.Printf("📥 Received log: %+v\n", entry.Source)


	// Validate the entry
	if err:= validate.Struct(entry)  ; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Println("Processing log entry...")
	time.Sleep(20 * time.Second)

	// Save the entry to the database
	if err:= Database.DB.Create(&entry).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save entry",
		})
	}
	// Log the entry
	fmt.Println("Entry done",entry.Source)
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Entry saved successfully",
	})

}