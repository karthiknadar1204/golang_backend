package main

import "fmt"

func main() {
	// 1. Basic Array Declaration
	// Fixed-size array of 5 integers
	var numbers [5]int
	fmt.Println("Empty array:", numbers) // Output: [0 0 0 0 0]



	// 2. Array with Initial Values
	names := [3]string{"John", "Jane", "Bob"}
	fmt.Println("Names:", names) // Output: [John Jane Bob]



	// 3. Array with Size Inference
	// Let compiler count the elements
	scores := [...]int{95, 87, 98}
	fmt.Println("Scores:", scores) // Output: [95 87 98]



	// 4. Accessing and Modifying Elements
	fmt.Println("First name:", names[0]) // Output: John
	names[1] = "Janet"
	fmt.Println("Modified names:", names) // Output: [John Janet Bob]



	// 5. Array Length
	fmt.Println("Array length:", len(names)) // Output: 3




	// 6. Iterating Over Arrays
	fmt.Println("\nIterating over scores:")
	// Using traditional for loop
	for i := 0; i < len(scores); i++ {
		fmt.Printf("Score %d: %d\n", i, scores[i])
	}


	// Using range
	fmt.Println("\nUsing range:")
	for index, value := range scores {
		fmt.Printf("scores[%d] = %d\n", index, value)
	}




	// 7. Multi-dimensional Arrays
	// 2D array (3x2)
	matrix := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println("\nMatrix:", matrix)

	// Accessing elements in 2D array
	fmt.Println("Matrix[1][1]:", matrix[1][1]) // Output: 4






	// 8. Comparing Arrays
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{3, 2, 1}

	fmt.Println("\nArray Comparisons:")
	fmt.Println("arr1 == arr2:", arr1 == arr2) // Output: true
	fmt.Println("arr1 == arr3:", arr1 == arr3) // Output: false



	// 9. Copying Arrays
	var copyArr [3]int
	copyArr = arr1 // Creates a new copy
	fmt.Println("\nCopied array:", copyArr)
}





// Array Characteristics
// Fixed size (cannot grow or shrink)
// Comparison available (using ==)
// Assignment creates a copy


// Iteration
// Can use traditional for loop
// Can use range for cleaner iteration
// range provides both index and value