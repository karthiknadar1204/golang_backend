package main
import "fmt"

//now here we are returning a float32 and an error
//error is a built in type in go
//error is a type that represents a value that can be used to indicate that an error has occurred
//error can be explicitly returned from a function using return statement,panic keyword.
func divide(a, b float32) (float32,error) {
	if b == 0 {
		//fmt.Errorf is a function that returns a new error with the given format string and arguments
		return 0,fmt.Errorf("division by zero cannot be done")
	}
	return a / b,nil
}

func main() {
	fmt.Println("Hello, World!")
	//here we are calling the divide function and storing the result in ans and error in err
	//if the division by zero occurs, the error will be returned and the program will exit
	ans,err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(ans)
	}
}