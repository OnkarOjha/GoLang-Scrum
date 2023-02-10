package main


import "fmt"
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

	elem1 := &List[string]{val : "2"}
	elem2 := &List[string]{val : "3"}

	elem1.next = elem2

	current := elem1

	for current!=nil{
		fmt.Print(current.val)

		current = current.next
	}

	q := `kya haal hai bhai
	egegege
	
	gger
	`

	qs:= "onkar ojha"

	fmt.Println("%c",qs[0])

	fmt.Println(q)
}
