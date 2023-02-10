package main

import "fmt"

func main(){
	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("\nTwo")


	 fmt.Println("hELLO WORLD")

	 myDefer()
}

func myDefer(){
	for i := 0; i<5 ;i++{
		defer fmt.Print(i)
	}
}