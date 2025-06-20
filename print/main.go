package main

import "fmt"

func main() {
	name := "John"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)  // Prints: Name: John, Age: 25
	fmt.Printf("Age: %d, Name: %s", age, name) 
}
