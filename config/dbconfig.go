package config

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getDBURI() string {
	dbUri := "mongodb://localhost:27017"
	return dbUri
}

var (
	client *mongo.Client
	err    error
)

func ConnectDB() *mongo.Client {
	if client == nil {
		client, err = mongo.Connect(context.TODO(), options.Client().
			ApplyURI(getDBURI()))
		if err != nil {
			panic(err)
		}

		pigError := client.Ping(context.Background(), nil)

		if pigError != nil {
			log.Fatal(pigError)
		}

	}
	fmt.Println("Connection Established....")
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	var local = "pre_order_food"
	log.Info("Client Status : ", client)
	if client == nil {
		ConnectDB()
	}
	return client.Database(local).Collection(collectionName)
}

func CreateUniqueIndex(collection *mongo.Collection, fieldName string) error {
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			fieldName: 1,
		},
		Options: options.Index().SetUnique(true),
	}

	// Create the index
	indexName, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return err
	}

	fmt.Printf("Created unique index: %s\n", indexName)
	return nil
}
