package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type ToDo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	fmt.Println("Server is running on port", port)

	app := fiber.New()

	todos := []ToDo{}

	// Return all ToDos from memory
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Create a ToDo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &ToDo{}

		// if we can't parse the body into a ToDo, return the given error
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err})
		}

		// We don't want to create ToDo's with empty bodies because that's meaningless
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ToDo body is required"})
		}

		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)

	})

	// Mark a ToDo as completed
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "ToDo not found"})
	})

	// Delete a ToDo by ID
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		previousLength := len(todos)

		newTodos := make([]ToDo, 0, len(todos)-1)
		for j := range todos {
			if fmt.Sprint(todos[j].ID) != id {
				newTodos = append(newTodos, todos[j])
			}
		}

		todos = newTodos
		if len(todos) == previousLength {
			return c.Status(404).JSON(fiber.Map{"error": "No ToDo associated with the given ID"})
		} else {
			return c.Status(200).JSON(fiber.Map{"success": "true"})
		}

	})

	log.Fatal(app.Listen(":" + port))

}
