package main

import (
	"dip/repository"
	"dip/service"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	ID int `json:"id"`
}

func send_email(c *fiber.Ctx) error {
	body := new(RequestBody)
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	if err := c.JSON(fiber.Map{"id": body.ID}); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	service.NewEmailService(repository.NewUserDatabaseRepository(db)).SendRegistrationEmail(uint(body.ID))
	return nil
}
