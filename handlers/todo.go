package handlers

import (
	"strconv"

	"github.com/NorousS/todo-fiber/configs"
	"github.com/NorousS/todo-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	result := configs.DB.Find(&todos)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch todos",
		})
	}
	return c.JSON(todos)
}

func GetTodoByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var todo models.Todo
	result := configs.DB.First(&todo, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}
	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	result := configs.DB.Create(&todo)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create todo",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var todo models.Todo
	result := configs.DB.First(&todo, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	configs.DB.Save(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	result := configs.DB.Delete(&models.Todo{}, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found or already deleted",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
