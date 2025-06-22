package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() Todo {
	fmt.Println("start")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error", err)
		return Todo{}
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Println("error", res.StatusCode)
		return Todo{}
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return
	// }
	// fmt.Println("data", string(data))

	//now i want the response from the api endpoint to get stored in the tdo variable.
	var todo Todo
	//Decoder decodes the json into normal objects,in Decode,we pass the address of the variable to store the decoded data.
	//This method is preferred when working with large json files because they allow streaming and efficient decoding of large json payloads instea dof loading the entire json payload into the memory at once.
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error", err)
		return Todo{}
	}
	fmt.Println("todo", todo)
	return todo
}

func performPostRequest() {
	fmt.Println("start")
	todo := Todo{
		UserID:    23,
		Title:     "new todo",
		Completed: true,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	//convert json data to string.
	jsonString := string(jsonData)
	fmt.Println("jsonData", jsonString)

	//convert string data to reader:
	jsonReader := strings.NewReader(jsonString)

	//post takes three things
	//1) url
	//2) content type
	//3) body
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("res", res.Status)

	//read the response body
	//io.ReadAll is used to read the response body and it returns a slice of bytes and an error basically bytes of data.
	data, _ := io.ReadAll(res.Body)
	//string(data) is used to convert the slice of bytes to a string i.e it converts the bytes of data to a string.
	fmt.Println("data", string(data))
	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error", err)
	// 	return
	// }
	// fmt.Println("data", string(data))
}

func performPutRequest() {
	fmt.Println("start")
	todo := Todo{
		UserID:    25,
		Title:     "new todo",
		Completed: true,
	}

	//convert todo struct to json bytes
	// HTTP requests send data as strings/bytes over the network
	//Go structs are in-memory objects that can't be transmitted directly
	//We need to serialize (convert) the struct into a format that can be sent i.e is convert to json.
	jsonData, err := json.Marshal(todo)
	// jsonData is now: []byte(`{"userId":25,"title":"new todo","completed":true}`)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	// Step 2: Convert bytes to string for HTTP body
	jsonString := string(jsonData)
	// jsonString is now: `{"userId":25,"title":"new todo","completed":true}`

	//convert string to reader
	jsonReader := strings.NewReader(jsonString)

	// Step 3: Create HTTP request with JSON body
	//NewRequest creates a new request with the given method, url and body.
	//strings.NewReader is used to create a new reader from the string.
	//req is a pointer to the request object.
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", jsonReader)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	//send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("res", res.Status)
	data, _ := io.ReadAll(res.Body)
	fmt.Println("data", string(data))
}

func performDeleteRequest() {
	fmt.Println("start")
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("res", res.Status)
	data, _ := io.ReadAll(res.Body)
	fmt.Println("data", string(data))
}

func main() {
	// getResult := performGetRequest()
	// fmt.Println("getResult", getResult)
	// performPostRequest()
	// performPutRequest()
	performDeleteRequest()
}
