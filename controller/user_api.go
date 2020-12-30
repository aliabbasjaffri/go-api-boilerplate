package controller

import (
	"encoding/json"
	"fmt"
	"github.com/aliabbasjaffri/go-api-boilerplate/dao"
	"github.com/aliabbasjaffri/go-api-boilerplate/model"
	"log"
	"net/http"
)

func CreateUser(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type", "application/json")
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		fmt.Print("Error occurred during conversion of JSON to User object")
		log.Fatal(err)
	}

	dao.AddUser(user)
	json.NewEncoder(response).Encode("User added successfully")
}

func GetAllUsers(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type", "application/json")
	users := dao.GetAllUsers()
	json.NewEncoder(response).Encode(users)
}

func UpdateUser(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type", "application/json")
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		fmt.Print("Error occurred during the creation of object")
		log.Fatal(err)
	}

	updateCount := dao.UpdateUser(user.Email, user.Age)
	json.NewEncoder(response).Encode(fmt.Sprintf("%v User/s updated successfully", updateCount))
}

func DeleteUser(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type",  "application/json")
	var user model.User

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		fmt.Print("Error occurred during param retrieval")
		log.Fatal(err)
	}

	deleteCount := dao.DeleteUser(user.Email)
	json.NewEncoder(response).Encode(fmt.Sprintf("%v User/s deleted successfully", deleteCount))
}