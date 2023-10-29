package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func main() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Define sample user data
	users := []User{
		{Username: "user1", Password: "password1"},
		{Username: "user2", Password: "password2"},
		{Username: "user3", Password: "password3"},
	}

	// Insert sample user data into database
	collection := client.Database("mydb").Collection("users")
	for _, user := range users {
		_, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			log.Fatal("Failed to insert user:", err)
		}
	}

	fmt.Println("Inserted sample user data into database!")
}
