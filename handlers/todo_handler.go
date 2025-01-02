package handlers

import (
	"context"
	"net/http"

	"github.com/DakotaErickson/GoTaskApp/models"
	"github.com/DakotaErickson/GoTaskApp/repository"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	repo repository.TodoRepository
}

func NewTodoHandler(repo repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: repo}
}

func (h *TodoHandler) GetAll(c *fiber.Ctx) error {
	todos, err := h.repo.GetTodos(context.Background())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(todos)
}

func (h *TodoHandler) Create(c *fiber.Ctx) error {
	todo := &models.Todo{}
	if err := c.BodyParser(todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	createdTodo, err := h.repo.CreateTodo(context.Background(), todo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(createdTodo)
}

func (h *TodoHandler) MarkComplete(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	err = h.repo.MarkTodoComplete(context.Background(), objectId)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Todo updated successfully"})
}

func (h *TodoHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.repo.DeleteTodo(context.Background(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Todo deleted successfully"})
}
