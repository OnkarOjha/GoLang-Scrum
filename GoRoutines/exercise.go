package main

import (
	"fmt"
	
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
type person struct{
	firstName string
}

type human interface{
	speak() 
}
// function having pointer receiver
func (p *person)speak(){
	fmt.Println("hi! I am human")
}
//attaching the function to struct
func saySomething(h human) {
	h.speak()
}
func main(){
	fmt.Println("This is exercise 2")
	p1 := person{
		firstName: "onkar",
    }
	
	saySomething(&p1)
}