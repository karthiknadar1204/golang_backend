package main
import "fmt"

// function without parameters and return type
func simpleFunction(){
	fmt.Println("This is a simple function")
}

// function with parameters and return type
// here a and b are parameters and int is the return type

//func func_name(parameter_1 type, parameter_2 type) return_type{
//	body
//}
func add(a int, b int) int{
	return a + b
}

// function with multiple return types
func addAndSubtract(a int, b int) (int, int){
	return a + b, a - b
}

// function with named return types
// here sum and difference are named return types
//func func_name(parameter_1 type, parameter_2 type) (return_type_1 name, return_type_2 name) {
//	body
//}
//here basically,we are defining the return variables and then returning them and also defining the type of the return variables.
func addAndSubtractNamed(a int, b int) (sum int, difference int){
	sum = a + b
	difference = a - b
	return
}

func main(){
	// calling a function
	simpleFunction()
	fmt.Println("--------------------------------")
	// calling a function with parameters and return type
	fmt.Println(add(1, 2))
	fmt.Println("--------------------------------")
	// calling a function with multiple return types
	fmt.Println(addAndSubtract(1, 2))
	fmt.Println("--------------------------------")
	// calling a function with named return types
	fmt.Println(addAndSubtractNamed(1, 2))
}