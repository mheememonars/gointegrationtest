package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Find ...
func Find(query map[string]interface{}, collection *mongo.Collection, elements interface{}) error {
	bsonQuery := bson.M{}
	for k, v := range query {
		bsonQuery[k] = v
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bsonQuery)
	if err != nil {
	}

	err = cursor.All(ctx, elements)
	if err != nil {
	}
	defer cursor.Close(ctx)

	return err
}

// FindOne ...
func FindOne(query map[string]interface{}, collection *mongo.Collection, result interface{}) error {
	bsonQuery := bson.M{}
	for k, v := range query {
		bsonQuery[k] = v
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bsonQuery).Decode(result)
	if err != nil {
		log.Println(err)
	}

	return err
}

// Save ...
func Save(e interface{}, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return collection.InsertOne(ctx, e)
}
