package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Working with Channels")
	myCh := make(chan int, 4) // creating a buffered channel so that it can consume any numbre of values
	chBool := make(chan bool, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(chBool chan bool, ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		fmt.Println(isChannelOpen)
		fmt.Println("value is : ", val)
		// fmt.Println("2nd value is: ", <-myCh)
		// fmt.Println("3rd value is: ", <-myCh)
		// fmt.Println("4th value is: ", <-myCh)
		// fmt.Println("5th value is: ", <-myCh)
		// fmt.Println("6th value is: ", <-myCh)
		// fmt.Println("7th value is: ", <-myCh)
		// fmt.Println("8th value is: ", <-myCh)
		// fmt.Println("9th value is: ", <-myCh)
		// fmt.Println("10th value is: ", <-myCh)
		// fmt.Println("11th value is: ", <-myCh)

		// for i := 0; i < cap(myCh); i++ {
		// 	fmt.Printf(" %d value is %d \n", i, <-myCh)
		// }
		for message := range myCh {
			fmt.Println(message)
		}
		// chBool <- false // this will through deadlock so we need to create buffred channel
		fmt.Println("Boolean value is: ", <-chBool)
		fmt.Println("Boolean value capacity is: ", cap(chBool))

		fmt.Println(cap(myCh))

		defer wg.Done()
	}(chBool, myCh, wg)

	go func(chBool chan bool, ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 6
		myCh <- 7
		myCh <- 8
		myCh <- 9
		myCh <- 10
		myCh <- 11
		myCh <- 12
		myCh <- 13

		chBool <- true
		close(chBool)

		// myCh <- 0
		close(myCh)
		defer wg.Done()
	}(chBool, myCh, wg)
	// myCh <-1
	// fmt.Println(<-myCh)
	// until we don't create go routines we can't access the channels
	wg.Wait()
}
