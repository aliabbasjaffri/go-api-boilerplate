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

type UserDao struct {
	Server     string
	Username   string
	Password   string
	Database   string
	Collection string
}

var mongoClient * mongo.Client

func ( T * UserDao) establishDBConnection() * mongo.Collection {
	//establish connection
	connectionString := fmt.Sprintf("mongodb://%v:%v@%v:27017", T.Username, T.Password, T.Server)
	fmt.Print(connectionString)
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	if mongoClient, err := mongo.Connect(context.TODO(), clientOptions); err != nil {
		log.Fatal(err)
	} else {
		// Check the connection
		if err := mongoClient.Ping(context.TODO(), nil); err != nil {
			fmt.Print("Unable to ping MongoDB server. Aborting connection request.")
			log.Fatal(err)
			return nil
		}
		fmt.Println("Connected to MongoDB!")
		return mongoClient.Database(T.Database).Collection(T.Collection)
	}
	return nil
}

func ( T * UserDao) closeDBConnection() {
	if err := mongoClient.Disconnect(context.TODO()); err != nil {
		fmt.Print("Unable to close Mongo DB connection")
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed!")
}

func ( T * UserDao) AddUser(user model.User)  {
	collection := T.establishDBConnection()
	_context, _ := context.WithTimeout(context.Background(), 5*time.Second)

	if insertResult, err := collection.InsertOne(_context, user); err != nil {
		fmt.Print("Error occurred during object insertion in DB")
		log.Fatal(err)
	} else {
		fmt.Println("User inserted successfully: ", insertResult.InsertedID)
	}

	T.closeDBConnection()
}

func ( T * UserDao) GetAllUsers() []* model.User {
	collection := T.establishDBConnection()

	var _findResults []* model.User
	_option := options.Find().SetLimit(5)

	if curr, err := collection.Find(context.TODO(), bson.D{}, _option); err != nil {
		fmt.Printf("No users found")
		log.Fatal(err)
	} else {
		for curr.Next(context.TODO()) {
			var obj model.User
			if err := curr.Decode(&obj); err != nil {
				log.Fatal(err)
			}
			_findResults = append(_findResults, &obj)
		}
		//if we come across an error in the cursor
		if err := curr.Err(); err != nil {
			log.Fatal(err)
		}
		curr.Close(context.TODO())
	}
	T.closeDBConnection()
	return _findResults
}

func ( T * UserDao) FindUser(_name string, _email string) []* model.User {
	collection := T.establishDBConnection()

	var _findResults []* model.User
	_filter := bson.D{
		{ "$or", bson.D{
			{"name", _name},
			{"email", _email},
			},
		},
	}
	_option := options.Find().SetLimit(3)

	if curr, err := collection.Find(context.TODO(), _filter, _option); err != nil {
		fmt.Printf(
			"No user found with name: %s or email: %s",
			_name,
			_email,
		)
		log.Fatal(err)
	} else {
		for curr.Next(context.TODO()) {
			var obj model.User
			if err := curr.Decode(&obj); err != nil {
				log.Fatal(err)
			}
			_findResults = append(_findResults, &obj)
		}
		//if we come across an error in the cursor
		if err := curr.Err(); err != nil {
			log.Fatal(err)
		}
		curr.Close(context.TODO())
	}
	T.closeDBConnection()
	return _findResults
}

func ( T * UserDao) UpdateUser(_email string, _age int) int {
	//open db connection, retrieve respective collection
	collection := T.establishDBConnection()

	var modifiedCount int
	//filter to search the user
	filter := bson.D{{"email", _email}}
	//update to apply on the filtered user
	update := bson.D{
		{"$set", bson.D{
				{"age", _age},
			},
		},
	}
	//perform update on collection
	if Result, err := collection.UpdateMany(context.TODO(), filter, update); err != nil {
		fmt.Printf(
			"User with email %s not found",
			_email,
		)
		log.Fatal(err)
	} else {
		//update user and close db connection
		fmt.Printf("Updated %v User \n", Result.ModifiedCount)
		modifiedCount = int(Result.ModifiedCount)
	}
	T.closeDBConnection()
	return modifiedCount
}

func ( T * UserDao) DeleteUser(_email string) int {
	collection := T.establishDBConnection()

	var deletedCount int
	filter := bson.D{{"email", _email}}
	if Result, err := collection.DeleteOne(context.TODO(), filter); err != nil {
		fmt.Printf(
			"User with email %s not found!",
			_email,
		)
		log.Fatal(err)
	} else {
		fmt.Printf("Users deleted: %v", Result.DeletedCount)
		deletedCount = int(Result.DeletedCount)
	}
	T.closeDBConnection()
	return deletedCount
}