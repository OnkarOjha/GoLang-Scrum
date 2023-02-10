// package main

// import "fmt"

// func main() {
//     fmt.Println("Hello, World!")
// }
// import  rsc.io/quote
package main

import "fmt"

// import "rsc.io/quote"

func main() {

	//specifying the compiler that we have speciifed it with a integer type
    var number1 int =10
	fmt.Println(number1)



	// here we haven't specified the data type
	var number2 = 20
	fmt.Println(number2)

	// shorthand notation to assign the values to the variable
	number3 := 30
	fmt.Println(number3)

	// if no value is assigned to the variable and if the variable is created it wll
	// be assigned with default value
	var number4 int
	fmt.Println(number4)

	// we can change the value of a variable as well
	number4 = 100
	fmt.Println("The changed value",number4)
	//PS :- we can't change the data type of variable


	// decalring multiple varibales at once

	var name  , age = "Onkar" , 21
	fmt.Println("My name is ",name)
	fmt.Println("My age is",age)

}