package main

import (
	"net/http"
	"encoding/json"
)

//User type: name, age, gender, email, phone number
//`json: field` is added to direct encoding/json on how the marshalling is to be done - to lowercase
type User struct {
	Name string `json:"name"`
	Age int	`json:"age"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Phone int `json:"phone"`
}

//Marshals the User type
func (u User)  ToJSON() []byte{
	ToJSON, err := json.Marshal(u)
	if err != nil{
		panic(err)
	}
	return ToJSON
}

//Unmarshal the User type
func FromJSON(data []byte) User{
	user := User{}
	err := json.Unmarshal(data, &user)
	if err != nil{
		panic(err)
	}
	return user

}

//Handles the API functionality
func RegistrationHandler(w http.ResponseWriter, r *http.Request){

}