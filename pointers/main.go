//Pointers are used to indirectly refer to the value stored in the variable,rather than the value itself.
//They provide a way to work with memory directly which can be used in tasks including efficient memory management and sharing data between functions
//declaring a pointer

package main

import "fmt"

func main(){
	var num int;
	num=2

	//this means ki pointer jo hai,woh jiss bhi data block ka address store karega,uss data block ke andar ka data is int type ka hai.
	var ptr *int;
	ptr=&num;
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr=10
	fmt.Println(num)
}






