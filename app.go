package main

import (
	controller "api/v1/controller"
	"fmt"
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
	router.Host("localhost")
	router.HandleFunc("/person", controller.CreateUser).Methods("POST")
	http.ListenAndServe(":9999", router)
}