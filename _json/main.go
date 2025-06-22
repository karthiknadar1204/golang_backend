package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}
	fmt.Println("user", user)

	//marshalling->convrt struct to json
	//this is the process of converting a struct into a json encoded byte-array like returns slice of bytes and error.
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	//string(jsonData) is used to convert the slice of bytes to a string i.e it converts the bytes of data to a string.
	fmt.Println("jsonData", string(jsonData))



	//unmarshalling->convert json to struct
	//this is the process of converting a json encoded byte-array into a struct like returns slice of bytes and error.
	var user2 User
	//now unmarshall takes two arguments,first is the slice of bytes and second is the address of the struct getting created and nly returns error.
	err = json.Unmarshal(jsonData, &user2)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("user2", user2)
}
