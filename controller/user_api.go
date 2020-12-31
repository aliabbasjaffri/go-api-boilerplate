package controller

import (
	"encoding/json"
	"fmt"
	"github.com/aliabbasjaffri/go-api-boilerplate/dao"
	"github.com/aliabbasjaffri/go-api-boilerplate/model"
	"log"
	"net/http"
	"os"
)

var daoObj = dao.UserDao{}

func init() {
	daoObj.Server = os.Getenv("MONGO_SERVER")
	daoObj.Username = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	daoObj.Password = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	daoObj.Database = os.Getenv("MONGO_INITDB_DATABASE")
	daoObj.Collection = os.Getenv("MONGO_INITDB_DATABASE_COLL")
}

func CreateUser(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type", "application/json")
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		fmt.Print("Error occurred during conversion of JSON to User object")
		log.Fatal(err)
	}

	daoObj.AddUser(user)
	if err := json.NewEncoder(response).Encode("User added successfully"); err != nil {
		fmt.Print("Unable to encode response message")
		log.Fatal(err)
	}
}

func GetAllUsers(response http.ResponseWriter, _ * http.Request) {
	response.Header().Set("content-type", "application/json")
	users := daoObj.GetAllUsers()
	if err := json.NewEncoder(response).Encode(users); err != nil {
		fmt.Print("Unable to encode response message")
		log.Fatal(err)
	}
}

func UpdateUser(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type", "application/json")
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		fmt.Print("Error occurred during the creation of object")
		log.Fatal(err)
	}

	updateCount := daoObj.UpdateUser(user.Email, user.Age)
	if err := json.NewEncoder(response).Encode(
		fmt.Sprintf("%v User/s updated successfully", updateCount)); err != nil {
		fmt.Print("Unable to encode response message")
		log.Fatal(err)
	}
}

func DeleteUser(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type",  "application/json")
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		fmt.Print("Error occurred during param retrieval")
		log.Fatal(err)
	}

	deleteCount := daoObj.DeleteUser(user.Email)
	if err := json.NewEncoder(response).Encode(
		fmt.Sprintf("%v User/s deleted successfully", deleteCount)); err != nil {
		fmt.Print("Unable to encode response message")
		log.Fatal(err)
	}
}