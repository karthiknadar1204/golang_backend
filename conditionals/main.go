// if statements in Go do not use parentheses around the condition:
// if height > 4 {
//     fmt.Println("You are tall enough!")
// }

// if height > 6 {
//     fmt.Println("You are super tall!")
// } else if height > 4 {
//     fmt.Println("You are tall enough!")
// } else {
//     fmt.Println("You are not tall enough!")
// }

package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	// don't touch above this line

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
	z:=10
	if z>10{
		fmt.Println("z is greater than 10")
	}else if z<10{
		fmt.Println("z is less than 10")
	}else{
		fmt.Println("z is equal to 10")
	}
}
