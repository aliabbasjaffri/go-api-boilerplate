package main

import (
	"fmt"
	"github.com/aliabbasjaffri/go-api-boilerplate/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func main () {
	//MONGO_DB_URI := "localhost"
	//MONGO_DB_PORT := 27017
	//fmt.Println("hello world")
	fmt.Println("Starting the application...")

	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//client, _ = mongo.Connect(ctx, clientOptions)

	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/getusers", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/adduser", controller.CreateUser).Methods("POST")
	//router.HandleFunc("/persons", controller.GetAllUsers).Methods("GET")
	http.ListenAndServe(":9999", router)
}