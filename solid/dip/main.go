package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	getDB()
	init_route(app)
	app.Listen(":3000")
}
