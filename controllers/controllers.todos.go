package controllers

import (
	"todos-api/models"
	"todos-api/utils"

	"github.com/gofiber/fiber/v2"
)

var todos = []models.Todo{}

func Get(context *fiber.Ctx) error {
	context.JSON(todos)
	return nil
}

func Post(context *fiber.Ctx) error {
	todo := models.Todo{}

	if err := context.BodyParser(&todo); err != nil {
		return err
	}

	todos = append(todos, todo)
	context.Status(201).JSON(todo)
	return nil
}

func GetById(context *fiber.Ctx) error {
	id := context.Params("id")

	todo, _ := utils.Find(todos, func(todo interface{}) bool {
		return todo.(models.Todo).ID == id
	})

	if todo == nil {
		return context.Status(404).JSON(fiber.Map{"message": "Todo not found"})
	}

	context.Status(200).JSON(todo)

	return nil
}

func DeleteById(context *fiber.Ctx) error {
	id := context.Params("id")

	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			context.Status(204).JSON(fiber.Map{
				"message": "Todo deleted",
			})
			return nil
		}
	}

	context.Status(404).JSON(fiber.Map{
		"error": "Todo not found",
	})

	return nil
}
