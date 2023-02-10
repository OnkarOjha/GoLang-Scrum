/* Three dots/ellipses notation and variadic function example
   (in go command ... means all files in the directory and subdirectories)
*/
package main

import (
	"fmt"
)

// show is a variadic function, it can have none or many arguments
func show(s ...string) {
	println()
	for i, val := range s {
		fmt.Printf("%d. %s\n", i, val)
	}
}
func main() {
	fmt.Println("Arrays are value types, fixed-length sequences of items of the same type.")
	fmt.Printf("Slices are reference types, can be resized by append.\n\n")

	seasons := []string{"Spring", "Summer", "Autumn", "Winter" , "sdfjgsdhfg"} // slice
	fmt.Println(cap(seasons))
	choices := [2]string{"Good", "Bad"}
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday","djbfgff" , "dfbsf"}
	fmt.Println("sdfsdf" ,cap(days))

	dsun := make([]string, 7, 7) //slice

	// Let's populate dsun with Sunday first
	for i, d := range days {
		if i == 6 {
			dsun[0] = d
		} else {
			dsun[i+1] = d
		}
	}

	fmt.Println("Length of slice seasons is", len(seasons), ", capacity is ", cap(seasons))
	fmt.Println("Length of slice dsun is", len(dsun), ", capacity is ", cap(dsun))
	fmt.Println("Length of array days is", len(days))
	fmt.Println("Length of array choices is", len(choices))

	show("a", "b")
	show(seasons...)
	// show(days...) //cannot use days (type [7]string) as type []string in argument to show
	// show(choices...) //cannot use choices (type [2]string) as type []string in argument to show
	show(dsun...)
	show("The", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog")
}
