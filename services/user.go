package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wisnu-bdd/be-go-with-auth-template/models"
	"github.com/wisnu-bdd/be-go-with-auth-template/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionUser = "user"

func GetUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := utils.AccessCollection(collectionUser)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*models.User
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func InsertUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID() // if you update your model to use primitive.ObjectID
	user.CreatedAt = time.Now()

	collection := utils.AccessCollection(collectionUser)

	newUser, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted record with ID: ", newUser.InsertedID)
	return err
}

func GetUserByID(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := utils.AccessCollection(collectionUser)

	objID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return nil, errID
	}

	var result *models.User
	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
	// return &result, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := utils.AccessCollection(collectionUser)

	var result *models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
	// return &result, nil
}

func UpdateUserByID(id string, updatedUser *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid menu group ID: %v", err)
	}

	updatedUser.UpdatedAt = time.Now() // optional, track updates

	collection := utils.AccessCollection(collectionUser)

	update := bson.M{
		"$set": bson.M{
			"email":     updatedUser.Email,
			"password":  updatedUser.Password,
			"updatedAt": updatedUser.UpdatedAt,
		},
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return fmt.Errorf("update failed: %v", err)
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with id: %s", id)
	}

	fmt.Println("Updated user with ID:", id)
	return nil
}

func UpdateUserDetailsByEmail(email string, updatedUser *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updatedUser.UpdatedAt = time.Now() // optional, track updates

	collection := utils.AccessCollection(collectionUser)

	update := bson.M{
		"$set": bson.M{
			"email":     updatedUser.Email,
			"updatedAt": updatedUser.UpdatedAt,
		},
	}

	res, err := collection.UpdateOne(ctx, bson.M{"email": email}, update)
	if err != nil {
		return fmt.Errorf("update failed: %v", err)
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with email: %s", email)
	}

	fmt.Println("Updated user with email:", email)
	return nil
}

func UpdateUserPasswordByEmail(email string, updatedUser *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updatedUser.UpdatedAt = time.Now() // optional, track updates

	collection := utils.AccessCollection(collectionUser)

	update := bson.M{
		"$set": bson.M{
			"password":  updatedUser.Password,
			"updatedAt": updatedUser.UpdatedAt,
		},
	}

	res, err := collection.UpdateOne(ctx, bson.M{"email": email}, update)
	if err != nil {
		return fmt.Errorf("update failed: %v", err)
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("no document found with email: %s", email)
	}

	fmt.Println("Updated user with email:", email)
	return nil
}

func DeleteUserByID(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := utils.AccessCollection(collectionUser)

	objID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return errID
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no document found with id %s", id)
	}
	return nil
}
