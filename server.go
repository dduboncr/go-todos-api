package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func helloWorldHandler(context *fiber.Ctx) error {
	
	return context.SendString("Hello, World!")
}

func main() {
	app := fiber.New()

	// recover
	app.Use(recover.New())
	app.Use(cors.New())

	app.Get("/", helloWorldHandler)

	app.Listen(":3000")

}