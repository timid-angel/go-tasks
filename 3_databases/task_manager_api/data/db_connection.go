package services

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var TaskCollection *mongo.Collection
var DB_Name = "task_test"

func ConnectDB() error {
	db_string := os.Getenv("DB_CONNECTION_STRING")
	if db_string == "" {
		return ServiceError{message: "Error: DB connection string not found. Make sure the environment variables are set correctly."}
	}

	clientOptions := options.Client().ApplyURI(db_string)
	client, connectionErr := mongo.Connect(context.TODO(), clientOptions)
	if connectionErr != nil {
		return ServiceError{message: fmt.Sprintf("Error: %v", connectionErr.Error())}
	}

	pingErr := client.Ping(context.TODO(), nil)
	if pingErr != nil {
		return ServiceError{message: fmt.Sprintf("Error: %v", pingErr.Error())}
	}

	Client = client
	TaskCollection = client.Database("task_test").Collection("tasks")
	return nil
}
