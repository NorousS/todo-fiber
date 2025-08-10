package main

import (
	"github.com/NorousS/todo-fiber/configs"
	"github.com/NorousS/todo-fiber/handlers"
	"github.com/NorousS/todo-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitDB()
	configs.DB.AutoMigrate(&models.Todo{})

	app := fiber.New()

	// Роуты
	app.Get("/todos", handlers.GetAllTodos)
	app.Get("/todos/:id", handlers.GetTodoByID)
	app.Post("/todos", handlers.CreateTodo)
	app.Put("/todos/:id", handlers.UpdateTodo)
	app.Delete("/todos/:id", handlers.DeleteTodo)

	app.Listen(":3000")
}
