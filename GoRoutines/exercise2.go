package main

import (
	"fmt"
	"runtime"
)

// func main() {
// 	fmt.Println("program to store 100 numbers into a channel and read them back")

// 	c := make(chan int)

// 	go func() {

// 		for i := 0; i < 100; i++ {
// 			c <- i
// 		}
// 		close(c)

// 	}()
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(i, <-c)
// 	}
// 	fmt.Println("code exit!")

// }

func main() {

	fmt.Println("launching 10 go routines and 10 different values into it!!")
	c := make(chan int)
	a := 0
	for i := 0; i < 10; i++ {
		go func() {
			c <- a
			fmt.Println(<-c)
		}()
		a++
		fmt.Println("value of a:",a)
		fmt.Println("go routines: ", runtime.NumGoroutine())
	}
}
