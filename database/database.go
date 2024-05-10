/* Connecting to the database - mongodb */

package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var userCollection *mongo.Collection

const mongoUri = "mongodb://localhost:27017/"
const databaseName = "todo"
const collectionName = "todolist"

// initialize database
func Init() {

	log.Printf("connecting to mongodb ...\n")

	//client options
	clientOptions := options.Client().ApplyURI(mongoUri)

	//create mongo client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("failed creating mongodb client instance :(\n )")
	}

	//ping
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	log.Printf("Connection successful\n")

	//creating mongodb collection instance
	userCollection = client.Database(databaseName).Collection(collectionName)

	log.Printf("User collection created: %v and is ready to use!\n ", userCollection)
}
