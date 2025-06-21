// We use structs in Go to represent structured data. It's often convenient to group different
//  types of variables together. For example, if we want to represent a car we could do the following:

// type car struct {
// 	make    string
// 	model   string
// 	doors   int
// 	mileage int
// }

// This creates a new struct type called car. All cars have a make, model, doors and mileage.

// We can create a new car by using the new keyword:

// myCar := car{
// 	make:    "Toyota",
// 	model:   "Camry",
// 	doors:   4,
// 	mileage: 10000,
// }

package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var karthik person
	fmt.Println(karthik)// "" "" 0->when you declare a variable in golang,thn whatever is its default value for the respective data type,it gets stored.
	karthik.FirstName="Karthik"
	karthik.LastName="luma"
	karthik.Age=20
	fmt.Println(karthik)

}



//NESTED STRUCTURE
// type car struct {
// 	make       string
// 	model      string
// 	doors      int
// 	mileage    int
// 	frontWheel wheel
// 	backWheel  wheel
// }

// type wheel struct {
// 	radius   int
// 	material string
// }

// The fields of a struct can be accessed using the dot (.) operator.
// myCar := car{}  ->instantiate a struct here, this is an empty struct,like object of a class
// myCar.frontWheel.radius = 5

// package main

// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }





// type rect struct {
// 	width int
// 	height int
//   }
  
//   // area has a receiver of (r rect)
//   func (r rect) area() int {
// 	return r.width * r.height
//   }
  
//   var r = rect{
// 	width: 5,
// 	height: 10,
//   }
  
//   fmt.Println(r.area())
  // prints 50