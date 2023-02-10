package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("welcome to my pizza app")
	fmt.Println("Please rate my pizza between 1 and 5: ")
	reader := bufio.NewReader(os.Stdin) // reader stores the reference of budio reader input
	 
	input , _ := reader.ReadString('\n') // "_" to hold the errors

	fmt.Println("thanks for rating: " , input)

	numRating , err := strconv.ParseFloat(strings.TrimSpace(input) , 64) // string conversiom library
	// strings.TrimSpace(input) isne trim kr diya \n jo sath mei ja ra tha 
	if err != nil {
		fmt.Println(err)
		// panic(err) // panic will end the program right there as soon as error is there

	}else {
		fmt.Println("added 1 to your rating: ", numRating +1)
	}




}