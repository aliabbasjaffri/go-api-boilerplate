package controller
import (
	dao "api/v1/dao"
	//"context"
	"encoding/json"
	"net/http"
	//"time"
)

func CreateUser(response http.ResponseWriter, request * http.Request) {
	response.Header().Set("content-type", "application/json")
	var user string
	_ = json.NewDecoder(request.Body).Decode(&user)
	//_context, _ := context.WithTimeout(context.Background(), 5*time.Second)
	dao.AddUser("ali", 29,"ali@jaffri.com" )
	json.NewEncoder(response).Encode("User added successfully")
}