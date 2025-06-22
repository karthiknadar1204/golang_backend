package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("start")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	//defer is used to close the response body after the main function is executed,reason is that if we dont close the response body,it will leak memory and it will cause a memory leak.
	//it ensures that the response body is closed after you have finished reading from it.
	//it is important to close resources like network connections and file handles to free up system resources.
	defer resp.Body.Close()

	//io.ReadAll is used to read the response body and it returns a slice of bytes and an error basically bytes of data.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	//string(data) is used to convert the slice of bytes to a string i.e it converts the bytes of data to a string.
	fmt.Println("data", string(data))
}
