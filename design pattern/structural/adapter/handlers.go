package main

import (
	"adapter/service"
	srv_abstraction "adapter/service_abstraction"

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

	adapter := service.NewEmailService(db)
	client := srv_abstraction.New_email_service(adapter)
	client.SendRegistrationEmail(body.ID)

	return nil
}
