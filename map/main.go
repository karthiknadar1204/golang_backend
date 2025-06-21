package main

import "fmt"

func main() {
	// make(map[KeyType]ValueType)
	scores := make(map[string]int)

	// Adding key-value pairs
	scores["Alice"] = 98
	scores["Bob"] = 85
	fmt.Println(scores) // Output: map[Alice:98 Bob:85]

	// Accessing values
	fmt.Println(scores["Alice"]) // Output: 98

	// Deleting a key
	delete(scores, "Bob")
	fmt.Println(scores) // Output: map[Alice:98]

	//Checking if a key exists
	value,exists:=scores["Alice"]
	if exists{
		fmt.Println(value)
	}

	
	
	
	
}