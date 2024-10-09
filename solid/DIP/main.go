package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	init_db()
	init_route(app)
	app.Listen(":3000")
}
