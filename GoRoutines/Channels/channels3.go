package main

import "fmt"

func main() {
	fmt.Println("Working with Channels")
	myCh := make(chan int)

	// myCh <-1
	// fmt.Println(<-myCh)
	// until we don't create go routines we can't access the channels
	
}
