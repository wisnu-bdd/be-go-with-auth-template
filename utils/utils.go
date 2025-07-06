package utils

import (
	"github.com/wisnu-bdd/be-go-with-auth-template/db"

	"go.mongodb.org/mongo-driver/mongo"
)

func AccessCollection(collectionName string) *mongo.Collection {
	collection := db.Client.Database("<cluster-name>").Collection(collectionName)
	return collection
}