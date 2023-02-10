package main

import  "fmt"

// jwtToken := 3000000

const LoginToken string  = "eknfw" // public token because of the capital "L"

func main(){

	// general initialization
	//  pehlaArray := [...]string{"Java" , "C++" , "Python"}

	// fmt.Println(pehlaArray[0])
	// fmt.Println(pehlaArray[1])
	// fmt.Println(pehlaArray[2])


	// initializing array
	// var pehlaArray[3] int

	// pehlaArray[0] = 2
	// pehlaArray[1] =3
	// pehlaArray[2] = 4

	// fmt.Println(pehlaArray)



	//intializing specific indexes of array
	// pehlaArray := [...]int{0:7 , 2:17}
	// fmt.Println(pehlaArray)
	// // fmt.Println(len(pehlaArray))

	// for i:= 0 ;i <len(pehlaArray);i++{
	// 	fmt.Println(pehlaArray[i])
	// }



	// multi deminsional array

	// pehlaArray := [2][2]int{{1,2} , {3,4}}

	// for i:=0 ; i<2;i++{
	// 	for j:=0 ;j<2;j++{
	// 		fmt.Println(pehlaArray[i][j])
	// 	}
	// }




	//slice
	// numbers := []int{1,2,3,4,5}

	// fmt.Println("Numbers : ",numbers)
	// fmt.Println(len(numbers))

	// lexer just protects you from following rules
	// lexer manually puts semi colon in the code when it goes for lexical analysis
	//  functions treated as types we can pass and return function as well

	// var username string = "omnkar"

	// fmt.Printf("variable is of type %T \n", username)

	// var smallVal uint8 = 255

	// fmt.Println(smallVal)
	// fmt.Printf("variable is of type %T \n", smallVal)

	// var smallfloat float32 = 255.47856264

	// fmt.Println(smallfloat)
	// fmt.Printf("variable is of type %T \n", smallfloat)

	// var webstie ="onkar.com"
	// fmt.Println(webstie)

	// no var style

	// numberOfUser := 3000000
	// fmt.Println(numberOfUser) 

	fmt.Println(LoginToken)
	fmt.Printf("Type of LoginToken is %T \n",LoginToken)

	 


}