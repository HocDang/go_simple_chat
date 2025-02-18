package repository

import (
	"context"
	"fmt"

	"w3s/go-backend/domain"
	"w3s/go-backend/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	collection := tr.database.Collection(tr.collection)

	// if task.ID is empty, generate new ObjectID
	if task.ID == primitive.NilObjectID {
		task.ID = primitive.NewObjectID()
	}

	_, err := collection.InsertOne(c, task)

	return err
}

func (tr *taskRepository) Fetch(c context.Context) ([]domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	var tasks []domain.Task
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &tasks)
	return tasks, err
}

func (tr *taskRepository) GetByID(c context.Context, taskID string) (*domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	var task domain.Task

	taskObjectID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(c, bson.M{"_id": taskObjectID}).Decode(&task)
	if err != nil {
		return nil, fmt.Errorf("task not found")
	}

	return &task, nil
}
