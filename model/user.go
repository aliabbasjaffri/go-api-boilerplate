package model

type User struct {
	Name  string        `json:"name"`
	Age   int           `json:"age"`
	Email string        `json:"email"`
	//Address * Address   `json:"address"`
}

//type Address struct {
//
//}