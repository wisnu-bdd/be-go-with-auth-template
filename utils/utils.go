package utils

import (
	"github.com/wisnu-bdd/be-go-with-auth-template/db"

	"go.mongodb.org/mongo-driver/mongo"
)

func AccessCollection(collectionName string) *mongo.Collection {
	collection := db.Client.Database("bdd-it-product-dev-guide").Collection(collectionName)
	return collection
}