package router

import (
	"github.com/DakotaErickson/GoTaskApp/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, todoHandler handlers.TodoHandler) {
	todoGroup := app.Group("/api/todos")
	todoGroup.Get("/", todoHandler.GetAll)
	todoGroup.Post("/", todoHandler.Create)
	todoGroup.Patch("/:id", todoHandler.MarkComplete)
	todoGroup.Delete("/:id", todoHandler.Delete)

	// Add routes for future resources below
}
