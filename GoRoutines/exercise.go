package main

import (
	"fmt"
	"sync"
	// "runtime"
)

//TODO EXERCISE 1
// func main() {
// 	fmt.Println("Exercise form udemy!!")
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	fmt.Println("Goroutines starting: ", runtime.NumGoroutine())
// 	go greeting(&wg)

// 	go greeter(&wg)
// 	fmt.Println("Goroutines middle: ", runtime.NumGoroutine())

// 	wg.Wait()
// 	fmt.Println("Goroutines at end: ", runtime.NumGoroutine())

// }

// func greeting(wg *sync.WaitGroup) {
// 	fmt.Println("Hello, Greetings!!")
// 	wg.Done()
// }
// func greeter(wg *sync.WaitGroup) {
// 	fmt.Println("Hello, Greeter!!")
// 	wg.Done()
// }

//TODO EXERCISE 2
// type person struct{
// 	firstName string
// }

// type human interface{
// 	speak()
// }
// // function having pointer receiver
// func (p *person)speak(){
// 	fmt.Println("hi! I am human")
// }
// //attaching the function to struct
// func saySomething(h human) {
// 	h.speak()
// }
// func main(){
// 	fmt.Println("This is exercise 2")
// 	p1 := person{
// 		firstName: "onkar",
//     }

// 	saySomething(&p1)
// }

//TODO EXERCISE 3

// func main() {
// 	fmt.Println("Exercise 3 from udemy!!")
// 	increment := 0
// 	loop := 10

// 	var wg sync.WaitGroup
// 	var mut sync.Mutex
// 	// var av atomic.Value
// 	wg.Add(loop)
// 	for i := 0; i < loop; i++ {
// 		go func() {
// 			mut.Lock()
// 			increment1 := increment
// 			// runtime.Gosched()
// 			increment1++
// 			increment = increment1
// 			fmt.Println("increment 1 : ", increment)
// 			mut.Unlock()
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("end value:", increment)
// }

//TODO EXERCISE 5

// // removing the above race condition with the help pf atomic
// func main() {
// 	loop := 10
// 	var wg sync.WaitGroup
// 	var increment int64
// 	wg.Add(loop)
// 	for i := 0; i < loop; i++ {
// 		atomic.AddInt64(&increment, 1)
// 		fmt.Println(atomic.LoadInt64(&increment))
// 		wg.Done()
// 	}
// 	wg.Wait()
// 	fmt.Println("End value: ", increment)
// }

//TODO EXERCISE 6

// func main() {
// 	fmt.Println("Printing CPU usage and core usig go run , go build")
// 	fmt.Println("The CPU runtime is: ", runtime.NumCPU())
// 	fmt.Println("The GOOS of os ", runtime.GOOS)
// 	fmt.Println("The runtime" ,runtime.GOARCH)

// }

// func main(){

// 	// c := make(chan int ,2)
// 	// cs := make(chan <-int)
// 	// cr := make(<-chan int)

// 	// // cs<-21

// 	// fmt.Printf("%T \n",c)
// 	// fmt.Printf("%T \n",cs)
// 	// fmt.Printf("%T \n",cr)
// 	// cr = c
// 	// cs = c
// 	// fmt.Printf("%T \n",cs)
// 	// fmt.Printf("%T \n",cr)

// 	// ch<-32
// 	// ch<-65

// 	// fmt.Println(<-ch)
// 	// fmt.Println(<-ch)
// 	// fmt.Println(<-cr)

// 	c := make(chan int)

// 	  go receive(c)

// 	//   send(c)
// 	for  v :=range c{
// 		fmt.Println("value of channel: ", v)
// 	}

// 	fmt.Println("program end")

// }

// func receive(c chan<- int){
//     for i:=0 ; i<100;i++{
// 		c<-i
// 	}
// 	close(c)

// }

// // func send(c <-chan int){
// // 	fmt.Println("value: ",<-c)

// // }

// func main() {
// 	fmt.Println("playing with select in channels!")

// 	eve := make(chan int)
// 	odd := make(chan int)
// 	quit := make(chan int)

// 	go send(eve, odd, quit)
// 	go receive(eve, odd, quit)
// }
// func receive(e, o, q <-chan int) {
// 	for {
// 		select {
// 		case v := <-e:
// 			fmt.Println(v)
// 		case v := <-o:
// 			fmt.Println(v)
// 		case v := <-q:
// 			fmt.Println(v)

// 			return
// 		}

// 	}

// }

// func send(e, o, q chan<- int) {
// 	for i := 0; i < 100; i++ {
// 		if i%2 == 0 {
// 			e <- i
// 		} else {
// 			o <- i
// 		}

// 	}
// 	// close(e)
// 	// close(o)
// 	q <- 0
// 	// close(q)
// }

func main() {
	c := make(chan string ,1)
	input := make(chan string ,1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		input <- "Onkar"
		
		// fmt.Println(<-input)
		//TODO channel to channel data transfer
		c <-<-input// TODO sirf buffered channel k case mein kaam krega
		// fmt.Println(<-c)
		//close(input)
		close(c)
		wg.Done()
	}()
	defer wg.Wait()
	go func() {

		for val := range c {
			fmt.Println("value i: ", val)
		}
		wg.Done()
	
	}()
}
