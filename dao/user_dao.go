package dao

import (
	"context"
	"fmt"
	"github.com/aliabbasjaffri/go-api-boilerplate/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func establishDBConnection() * mongo.Client {
	//the connection string should be passed as constructor
	//re-implement the logic in object oriented paradigm
	//replace the username and password with OS vars
	clientOptions := options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017")
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

func closeDBConnection(client * mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil { log.Fatal(err) }
	fmt.Println("Connection to MongoDB closed!")
}

func getUserCollection(client * mongo.Client) * mongo.Collection {
	return client.Database("go-api-db").Collection("users")
}

func AddUser(user model.User)  {
	dbClient := establishDBConnection()
	collection := getUserCollection(dbClient)
	_context, _ := context.WithTimeout(context.Background(), 5*time.Second)

	insertResult, err := collection.InsertOne(_context, user)
	if err != nil {
		fmt.Print("Error occurred during object insertion in DB")
		log.Fatal(err)
	}

	fmt.Println("User inserted successfully: ", insertResult.InsertedID)
	closeDBConnection(dbClient)
}

func GetAllUsers() []* model.User {
	dbClient := establishDBConnection()
	collection := getUserCollection(dbClient)

	var _findResults []* model.User
	_option := options.Find().SetLimit(5)
	curr, err := collection.Find(context.TODO(), bson.D{}, _option)
	if err != nil {
		fmt.Printf("No users found")
		log.Fatal(err)
	}

	for curr.Next(context.TODO()) {
		var obj model.User
		err := curr.Decode(&obj)

		if err != nil {
			log.Fatal(err)
		}
		_findResults = append(_findResults, &obj)
	}
	//if we come across an error in the cursor
	if err := curr.Err(); err != nil {
		log.Fatal(err)
	}

	curr.Close(context.TODO())
	closeDBConnection(dbClient)
	return _findResults
}

func FindUser(_name string, _email string) []* model.User {
	dbClient := establishDBConnection()
	collection := getUserCollection(dbClient)

	var _findResults []* model.User
	_filter := bson.D{
		{ "$or", bson.D{
			{"name", _name},
			{"email", _email},
			},
		},
	}
	_option := options.Find().SetLimit(3)

	curr, err := collection.Find(context.TODO(), _filter, _option)

	if err != nil {
		fmt.Printf(
			"No user found with name: %s or email: %s",
			_name,
			_email,
		)
		log.Fatal(err)
	}
	for curr.Next(context.TODO()) {
		var obj model.User
		err := curr.Decode(&obj)

		if err != nil {
			log.Fatal(err)
		}
		_findResults = append(_findResults, &obj)
	}

	//if we come across an error in the cursor
	if err := curr.Err(); err != nil {
		log.Fatal(err)
	}

	curr.Close(context.TODO())
	closeDBConnection(dbClient)
	return _findResults
}

func UpdateUser(_email string, _age int) {
	//open db connection, retrieve respective collection
	dbClient := establishDBConnection()
	collection := getUserCollection(dbClient)

	//filter to search the user
	filter := bson.D{{"email", _email}}

	//update to apply on the filtered user
	update := bson.D{
		{"$inc", bson.D{
				{"age", _age},
			},
		},
	}
	//perform update on collection
	Result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		fmt.Printf(
			"User with email %s not found",
			_email,
		)
		log.Fatal(err)
	}
	//update user and close db connection
	fmt.Printf("Updated %v documents", Result.ModifiedCount)
	closeDBConnection(dbClient)
}

func DeleteUser(_email string) {
	dbClient := establishDBConnection()
	collection := getUserCollection(dbClient)

	filter := bson.D{{"email", _email}}
	Result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		fmt.Printf(
			"User with email %s not found!",
			_email,
		)
		log.Fatal(err)
	}
	fmt.Printf("Users deleted: %v", Result.DeletedCount)
	closeDBConnection(dbClient)
}