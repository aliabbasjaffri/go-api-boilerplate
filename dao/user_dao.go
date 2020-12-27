package dao

import (
	. "../model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func EstablishDBConnection() * mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func CloseDBConnection(client * mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func GetUserCollection(client * mongo.Client) * mongo.Collection {
	return client.Database("go-api-db").Collection("users")
}

func AddUser(name string, age int, emailAddress string) {
	dbClient := EstablishDBConnection()
	collection := GetUserCollection(dbClient)

	user := User{Name: name, Age: age, Email: emailAddress}

	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User inserted successfully: ", insertResult.InsertedID)
	CloseDBConnection(dbClient)
}

//func AddMultipleUsers () {
//
//}

func UpdateUser(_email string, _age int) {
	dbClient := EstablishDBConnection()
	collection := GetUserCollection(dbClient)

	filter := bson.D{{"email", _email}}
	update := bson.D{
		{
			"$inc", bson.D{
				{"age", _age},
			},
		},
	}

	Result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v documents", Result.ModifiedCount)
	CloseDBConnection(dbClient)
}

