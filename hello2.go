package main

import "fmt"

const (
	tomato, apple int = iota + 1, iota + 2 // iota  = 0
	orange, chevy // 1
	ford, c //

)

func main() {
	fmt.Println(tomato) // 1
	fmt.Println(apple) // 2
	fmt.Println(orange) //2 
	fmt.Println(chevy) // 3
	fmt.Println(ford) //3
	fmt.Println(c) //3



	// fmt.Println(tomato , apple , orange , chevy , ford , v ,c, a)


}