package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

// Mongo - MongoDB Struct with Client
type Mongo struct {
	Client mongo.Client
	DB     mongo.Database
}

// Init : Helper function to initialize MongoDB
func (m *Mongo) Init() {
	mongoURI := os.Getenv("MONGODB_URI")

	if mongoURI == "" {
		log.Fatal("MongoDB URI Env not provided")
	}

	// Parse the mongoDB URI to Connection string to get the Database Name
	cs, err := connstring.Parse(mongoURI)

	if err != nil {
		log.Fatal("Error while parsing MongoDB URI. Make sure it is right..")
	}

	// Create a new Mongo Client
	// Set retry writes to false because mongoDB deployment doesn't support it
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI + "?retryWrites=false"))

	if err != nil {
		log.Fatal("Error while initializing MongoDB server. Make sure the URI is correct..")
	}

	// Connect to MongoDB server with context
	err = client.Connect(context.TODO())

	if err != nil {
		log.Fatal("Error while connecting to MongoDB server. Make sure it is running..")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB Successfully")

	// Set Client and DB to Mongo Structure
	m.Client = *client
	m.DB = *client.Database(cs.Database)
}
