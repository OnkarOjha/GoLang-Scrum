package main

import (
	"fmt"
	"time"
)

// func main() {
// 	even := make(chan int)
// 	odd := make(chan int)
// 	fanin := make(chan int)

// 	go send(even, odd)

// 	go receive(even, odd, fanin)

// 	for v := range fanin {
// 		fmt.Println(v)
// 	}
// 	fmt.Println("Exit!!")
// }

// //send channel
// func send(even, odd chan<- int) { // sending channel to send data , jah "pe" data send kr skte hai
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			even <- i
// 		} else {
// 			odd <- i
// 		}
// 	}
// 	close(even)
// 	close(odd)
// }

// //receive channel
// func receive(even, odd <-chan int, fanin chan<- int) { // receiving channel from which we can receive or extract data , jaha "se" data le skte hai
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		for v := range even {
// 			fanin <- v
// 		}
// 		wg.Done()

// 	}()

// 	go func(){
// 		for v := range odd {
//             fanin<-v
//         }
//         wg.Done()
// 	}()

// 	wg.Wait()
// 	close(fanin)
// }

// func main() {
// 	c := fanin(boring("Onkar"), boring("ojha"))
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(<-c)
// 	}
// 	fmt.Println("You are both  boring and I m leaving!!")
// }

// func boring(msg string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 0; ; i++ {
// 			c <- fmt.Sprintf("%s %d", msg, i)
// 			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// 		}
// 	}()

// 	return c

// }
// func fanin(input1 , input2 <-chan string) <-chan string {
// 	c := make(chan string)
// 	go func(){
// 		for{
// 			c <-<-input1
// 		}
// 	}()
// 	go func(){
// 		for{
// 			c <-<-input1
// 		}
// 	}()

// 	return c
// }

//TODO FANOUT
// func main() {
// 	c1 := make(chan int)
// 	c2 := make(chan int)

// 	go populate(c1)

// 	go fanout(c1, c2)

// 	for v := range c2 {
// 		fmt.Println(v)
// 	}
// 	fmt.Println("exitr!!!")
// }

// func populate(c chan int) {
// 	for i := 0; i < 100; i++ {
// 		c <- i
// 	}
// }

// func fanout(c1, c2 chan int) {
// 	var wg sync.WaitGroup
// 	const goroutines = 10
// 	for i := 0; i < goroutines; i++ {
// 		go func() {
// 			for v := range c1 {

// 				wg.Add(1)
// 				go func(v2 int) {
// 					c2 <- timeConsumingWork(v2)
// 					wg.Done()
// 				}(v)
// 			}
// 			wg.Done()
// 		}()
// 	}

// 	wg.Wait()
// 	close(c2)
// }
// func timeConsumingWork(v int) int {
// 	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
// 	return v + rand.Intn(1000)
// }

///TODO FANIN AGAIN
func main() {
	c1 := generate("Onkar")
	c2 := generate("Ojha")

	fanin := make(chan string)

	go func() {
		for {
			select {
			case str1 := <-c1:
				fanin <- str1
			case str2 := <-c2:
				fanin <- str2
			}
		}
	}()

	go func() {
		for {
			fmt.Println(<-fanin)
		}
	}()
}

func generate(msg string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- msg
			time.Sleep(time.Duration(100*time.Millisecond))
		}
	}()
	return channel
}
