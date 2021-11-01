package db

import (
	"context"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CustomerCollection *mongo.Collection
)

func Init() {
	log.Println("Connect DB ")
	client, err := ConnectionWithAuth()
	if err != nil {
		log.Panic(err)
	}

	dbName := viper.GetString("DB.Name")
	customerCollectionName := viper.GetString("DB.Collection")

	CustomerCollection = client.Database(dbName).Collection(customerCollectionName)
}

func ConnectionWithAuth() (*mongo.Client, error) {
	auth := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		Username:      viper.GetString("DB.User"),
		Password:      viper.GetString("DB.Password"),
		AuthSource:    "admin",
	}
	mongoURI := "mongodb://" + viper.GetString("DB.IP") + ":" + viper.GetString("DB.Port")

	clientOptions := options.Client()
	clientOptions.ApplyURI(mongoURI)
	clientOptions.SetAuth(auth)
	clientOptions.SetDirect(true)

	err := clientOptions.Validate()
	if err != nil {
		log.Println("Invalid client options:", err)
		return nil, err
	}
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Println("cannot create client:", err)
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println("connection error", err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("ping error", err)
		return nil, err
	}
	log.Println("Ping DBUserPassword successful")
	return client, nil
}

func SetupCollections(client *mongo.Client) {
	dbName := viper.GetString("DB.Name")
	customerCollectionName := viper.GetString("DB.Collection")

	CustomerCollection = client.Database(dbName).Collection(customerCollectionName)
}
