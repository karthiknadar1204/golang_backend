// Slices
// Slices are dynamic, flexible arrays in Go. Unlike arrays, slices can grow or shrink in size. A slice is a reference to an underlying array.




// Creating Slices
// There are several ways to create slices:

// 1. Using make()
// package main
// import "fmt"
// func main() {
// 	// make([]type, length, capacity)
// 	slice1 := make([]int, 3) // length=3, capacity=3
// 	slice2 := make([]int, 3, 5) // length=3, capacity=5
// 	fmt.Printf("slice1: %v, length: %d, capacity: %d\n", slice1, len(slice1), cap(slice1))
// 	fmt.Printf("slice2: %v, length: %d, capacity: %d\n", slice2, len(slice2), cap(slice2))
// }





// 2. Using Slice Literal
// package main
// import "fmt"
// func main() {
// 	// Direct initialization
// 	fruits := []string{"apple", "banana", "orange"}
// 	fmt.Println(fruits) // Output: [apple banana orange]
// }




// 3. From Array
package main
import "fmt"
func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	// Creating slice from array
	slice := arr[1:4] // includes elements from index 1 to 3
	fmt.Println(slice) // Output: [2 3 4]

	//adding elements to the slice
	slice = append(slice, 6)
	fmt.Println(slice) // Output: [2 3 4 6]

	//adding multiple elements to the slice
	slice = append(slice, 7, 8, 9)
	fmt.Println(slice) // Output: [2 3 4 6 7 8 9]
}




// Slice Operations
// 1. Appending Elements
// package main
// import "fmt"
// func main() {
// 	slice := []int{1, 2, 3}
// 	// Append single element
// 	slice = append(slice, 4)
// 	fmt.Println(slice) // Output: [1 2 3 4]
// 	// Append multiple elements
// 	slice = append(slice, 5, 6, 7)
// 	fmt.Println(slice) // Output: [1 2 3 4 5 6 7]
// 	// Append another slice
// 	other := []int{8, 9}
// 	slice = append(slice, other...)
// 	fmt.Println(slice) // Output: [1 2 3 4 5 6 7 8 9]
// }




// 2. Slicing
// package main
// import "fmt"
// func main() {
// 	numbers := []int{0, 1, 2, 3, 4, 5}
// 	// Slice syntax: slice[startIndex:endIndex]
// 	fmt.Println(numbers[1:4]) // Output: [1 2 3]
// 	fmt.Println(numbers[:3]) // Output: [0 1 2]
// 	fmt.Println(numbers[3:]) // Output: [3 4 5]
// 	fmt.Println(numbers[:]) // Output: [0 1 2 3 4 5]
// }





// 3. Copying Slices
// package main
// import "fmt"
// func main() {
// 	src := []int{1, 2, 3}
// 	dst := make([]int, len(src))
// 	// Copy elements from src to dst
// 	copied := copy(dst, src)
// 	fmt.Printf("src: %v\n", src) // Output: src: [1 2 3]
// 	fmt.Printf("dst: %v\n", dst) // Output: dst: [1 2 3]
// 	fmt.Printf("copied elements: %d\n", copied) // Output: copied elements: 3
// }
// Important Concepts
// 1. Length vs Capacity
// Length: Number of elements in the slice (len())
// Capacity: Number of elements in underlying array from slice's first element (cap())
// The cap() method returns the capacity of a slice, which is the maximum number of elements the slice can hold before it needs to be resized.
// Key points about capacity:
// Initial Capacity: When you create a slice using make(), you can specify the initial capacity If not specified, capacity equals length.

// Capacity Growth: When you append elements beyond current capacity, Go creates a new array with doubled capacity The old array is garbage collected.






// 2. Nil Slices
// package main
// import "fmt"
// func main() {
// 	var slice []int
// 	fmt.Println(slice == nil) // Output: true
// 	fmt.Println(len(slice)) // Output: 0
// 	fmt.Println(cap(slice)) // Output: 0
// 	// Appending to nil slice
// 	slice = append(slice, 1)
// 	fmt.Println(slice) // Output: [1]
// }
	
// 3. Memory Efficiency
// When working with large slices, it's important to be aware of memory usage:

// package main
// // import "fmt"
// func main() {
// 	// Create a large slice
// 	original := make([]int, 1000000)
// 	// Taking a small slice
// 	// This still holds reference to large underlying array
// 	small := original[len(original)-3:]
// 	// Better approach: Create a new slice with copy
// 	better := make([]int, len(small))
// 	copy(better, small)
// 	// Now original can be garbage collected
// 	original = nil
// }




// 5. Using with Different Types
// package main

// import (
//     "fmt"
//     "slices"
// )

// func main() {
//     // Integer slices
//     ints1 := []int{1, 2, 3}
//     ints2 := []int{1, 2, 3}
//     fmt.Println("Integers:", slices.Equal(ints1, ints2))

//     // Float slices
//     floats1 := []float64{1.1, 2.2, 3.3}
//     floats2 := []float64{1.1, 2.2, 3.3}
//     fmt.Println("Floats:", slices.Equal(floats1, floats2))

//     // Byte slices
//     bytes1 := []byte("hello")
//     bytes2 := []byte("hello")
//     fmt.Println("Bytes:", slices.Equal(bytes1, bytes2))
// }


// 4. Comparing Custom Types
// package main

// import (
//     "fmt"
//     "slices"
// )

// type Person struct {
//     Name string
//     Age  int
// }

// func main() {
//     people1 := []Person{
//         {Name: "Alice", Age: 25},
//         {Name: "Bob", Age: 30},
//     }
    
//     people2 := []Person{
//         {Name: "Alice", Age: 25},
//         {Name: "Bob", Age: 30},
//     }

//     people3 := []Person{
//         {Name: "Alice", Age: 25},
//         {Name: "Bob", Age: 31},  // Different age
//     }

//     fmt.Println(slices.Equal(people1, people2))  // true
//     fmt.Println(slices.Equal(people1, people3))  // false
// }
