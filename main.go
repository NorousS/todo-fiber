package main

import (
	"github.com/NorousS/todo-fiber/configs"
	"github.com/NorousS/todo-fiber/handlers"
	"github.com/NorousS/todo-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Инициализация БД
	configs.InitDB()
	configs.DB.AutoMigrate(&models.Todo{})

	app := fiber.New()

	// Роуты
	app.Get("/todos", handlers.GetTodos)
	app.Post("/todos", handlers.CreateTodo)

	app.Listen(":3000")
}
