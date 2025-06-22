package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("start")
	fmt.Println(add(1, 2))
	// defer is used to execute a function after the main function is executed i.e at the end of the program
	//now if we have multiple defer statements,it executes like a stack like lifo state.
	defer fmt.Println("defer her ein middle of the program 1")
	defer fmt.Println("defer her ein middle of the program 2")
	defer fmt.Println("defer her ein middle of the program 3")
	fmt.Println("end hai lodu")
}
