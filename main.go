package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
   // variables in go
   // var <variable-name> <variable-type> = <variable-value>
	var startup string = "Textio SMS service booting up..."
	var messages string = "Sending test message"
	var confirmation string = "Message sent!"

	fmt.Println(startup)
	fmt.Println(messages)
	fmt.Println(confirmation)

	// variables in go can also be declared without a type
	// <variable-name> := <variable-value>
	// this is called short declaration operator or the walrus operator
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."


	// variables can be declared and initialised in a single line
	// multiple variables can be declared in a single line
	// <variable-name1>, <variable-name2> := <variable-value1>, <variable-value2>
	averageOpenRate, displayMessage := 50, "hi there"

	userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, isSubscribed: %t, Message: %s",
		fname, lname, age, messageRate, isSubscribed, message)

	// if a variable is declared with an uppercase letter,it is considered a public variable and can be accessed from other packages,like it can be exported to other packages.
	var PublicVariable string = "This is a public variable"
	fmt.Println(PublicVariable)

	// if a variable is declared with a lowercase letter,it is considered a private variable and cannot be accessed from other packages,like it cannot be exported to other packages.
	var privateVariable string = "This is a private variable"
	fmt.Println(privateVariable)




	fmt.Println(userLog)
	fmt.Println(averageOpenRate, displayMessage)
}
