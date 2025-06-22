package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello world")
	time.Sleep(3000 * time.Millisecond)
}

func sayHello2() {
	fmt.Println("Hello world 2")
}

func main() {
	fmt.Println("go routines testing")
	go sayHello()
	go sayHello2()
	time.Sleep(4000 * time.Millisecond)
}
