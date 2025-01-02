package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/DakotaErickson/GoTaskApp/handlers"
	"github.com/DakotaErickson/GoTaskApp/repository"
)

func main() {
	fmt.Println("Hello World!")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB Atlas")

	collection := client.Database("golang_db").Collection("todos")

	// Create a repository instance
	todoRepo := repository.NewTodoRepository(collection)

	// Create a handler instance
	todoHandler := handlers.NewTodoHandler(*todoRepo)

	app := fiber.New()

	// Register routes using the handler
	app.Get("/api/todos", todoHandler.GetAll)
	app.Post("/api/todos", todoHandler.Create)
	app.Patch("/api/todos/:id", todoHandler.MarkComplete)
	app.Delete("/api/todos/:id", todoHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
