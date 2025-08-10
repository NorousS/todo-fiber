package handlers

import (
	"github.com/NorousS/todo-fiber/configs"
	"github.com/NorousS/todo-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	configs.DB.Find(&todos)
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	configs.DB.Create(&todo)
	return c.JSON(todo)
}

// Добавьте остальные методы (Update, Delete) по аналогии
