package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/DakotaErickson/GoTaskApp/models"
)

type TodoRepository struct {
	collection *mongo.Collection
}

func NewTodoRepository(collection *mongo.Collection) *TodoRepository {
	return &TodoRepository{collection: collection}
}

func (r *TodoRepository) GetTodos(ctx context.Context) ([]models.Todo, error) {
	var todos []models.Todo
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	result, err := r.collection.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	return todo, nil
}

func (r *TodoRepository) MarkTodoComplete(ctx context.Context, objectId primitive.ObjectID) error {
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err
}

func (r *TodoRepository) DeleteTodo(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	_, err = r.collection.DeleteOne(ctx, filter)
	return err
}
