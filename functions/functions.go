package main

import "fmt"

func main(){
	fmt.Println("These are Functions!!")
	greeter()

	result := adder(3,4)

	fmt.Println("REsults is: \n",result)

	proResult := proAdder(3,4,5,6,7,8,9)

	fmt.Println("Results of proAdder are: ", proResult)
}

func adder(valOne int , valTwo int) int {

	return valOne + valTwo

}

// i don't know how many values are there to come we need to add all the incoming values
func proAdder(values ...int)int{
	total := 0
	// values are now slices

	for _ , val := range values{
		total +=val
	}

	return total
}

func greeter(){
	fmt.Println("greetings GoLang!!")
	
}

