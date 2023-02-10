package main

import "fmt"

//will create one function passing interface as a parameter
func kuchBhi(i interface{}) {
	switch v := i.(type) {
	// TODO yaha pr T keyword ke jagah special keyword {type} use kra hai
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
		break
	case string:
		fmt.Printf("the given string is %s which is  %T\n", v,v)
		break
	default:
		fmt.Printf("The unkown type of variable given is %T\n", v)
	}
}

func main() {
	fmt.Println("Playing with Type Swithcing this time!!!")

	kuchBhi(21)
	kuchBhi("Onkar")
	kuchBhi(34.22)

}
