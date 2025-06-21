package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// counter:=0
	// for counter<5{
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// for i:=0;i<5;i++{
	// 	fmt.Println(i)
	// }

	// for{
	// 	fmt.Println("Infinite loop")
	// 	counter++
	// 	if counter==10{
	// 		break
	// 	}
	// }




	//The range keyword is used to iterate over arrays, slices, maps, and strings.
	numbers:=[]int{1,2,3,4,5}
	for index,value:=range numbers{
		fmt.Println(index,value)
	}

	//The range keyword can also be used to iterate over maps.
	ages:=map[string]int{
		"John":20,
		"Jane":21,
		"Jim":22,
	}
	for name,age:=range ages{
		fmt.Println(name,age)
	}

	//The range keyword can also be used to iterate over strings.
	string:="Hello, World!"
	for index,value:=range string{
		fmt.Println(index,value)
	}


	//The break keyword is used to exit a loop.
	for i:=0;i<10;i++{
		if i==5{
			break
		}
	}
}