package main

import "fmt"

func main(){
	fmt.Println("Welcome to class of a pointer")

	// var ptr *int // creating a pointer 
	// fmt.Println("value of pointer is",ptr)

	myNumber := 23
	var ptr = &myNumber // creating a pointer as well as referencing it to soem value
	fmt.Println("Value of the actual pointer is", ptr) // value or adddress of actual pointer
	fmt.Println("Value of the actual pointer is", *ptr) // value of the memory location in which it is pointing out too

	//we can add multiply and do whatever artithmetic function that we can think of using the "*" and the pointers
	*ptr = *ptr *2
	fmt.Println("New value is : ", myNumber)


}