package main

import (
	"fmt"
	"github.com/aliabbasjaffri/go-api-boilerplate/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main () {
	fmt.Println("Starting the API")

	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/getusers", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/adduser", controller.CreateUser).Methods("POST")
	router.HandleFunc("/updateuser", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteuser", controller.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9999", router))
}