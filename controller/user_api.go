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
	log.Fatal(json.NewEncoder(response).Encode("User added successfully"))
}

func GetAllUsers(response http.ResponseWriter, _ * http.Request) {
	response.Header().Set("content-type", "application/json")
	users := daoObj.GetAllUsers()
	log.Fatal(json.NewEncoder(response).Encode(users))
}

func UpdateUser(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type", "application/json")
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		fmt.Print("Error occurred during the creation of object")
		log.Fatal(err)
	}

	updateCount := daoObj.UpdateUser(user.Email, user.Age)
	log.Fatal(json.NewEncoder(response).Encode(fmt.Sprintf("%v User/s updated successfully", updateCount)))
}

func DeleteUser(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type",  "application/json")
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		fmt.Print("Error occurred during param retrieval")
		log.Fatal(err)
	}

	deleteCount := daoObj.DeleteUser(user.Email)
	log.Fatal(json.NewEncoder(response).Encode(fmt.Sprintf("%v User/s deleted successfully", deleteCount)))
}