package main

import (
	"todos-api/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)


func SetupRoutes(app *fiber.App) {
	app.Get("/todos", controllers.Get)
	app.Post("/todos", controllers.Post)
	app.Get("/todos/:id", controllers.GetById)
	app.Delete("/todos/:id", controllers.DeleteById)
}


func main() {

	app := fiber.New()

	SetupRoutes(app);

	app.Use(recover.New())
	app.Use(cors.New())

	app.Listen(":3000")
}
