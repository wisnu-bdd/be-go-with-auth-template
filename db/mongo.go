package db

import (
	"context"
	"time"

	"github.com/wisnu-bdd/be-go-with-auth-template/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToMongo() error {
	// Load credentials
	MongoConnectionString := config.MongoConnectionString

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(MongoConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions) // 🔁 Update later
	if err != nil {
		return err
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	Client = client

	return nil
}
