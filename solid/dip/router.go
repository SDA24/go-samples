package main

import "github.com/gofiber/fiber/v2"

func init_route(app *fiber.App) {
	app.Get("/", send_email)
}
