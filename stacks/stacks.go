package main

import "fmt"

type stack []string

//is empty function for stack
func (s *stack) IsEmpty() bool{
	return len(*s) == 0
}


// push a new value onto the stack
func (s *stack) Push(str string){
	*s = append(*s ,str)
}

//remove and return top element of stack.
// return false if element not found
func (s *stack) Pop() (string,bool){
	if s.IsEmpty(){
		return "",false
	}else{
		//index of topmost element
		topIndex := len(*s) - 1
		//index into the slice and obtain the element
		element := (*s)[topIndex]
		//Remove topIndex from the stack by slicing it off.
		*s = (*s)[:topIndex]
		return element,true

	}
}

func main(){
	fmt.Println("playing with stacks in golang!!")

	// create an instance of this stack
	var stack stack

	stack.Push("Hy!")
	stack.Push("Onkar")
	stack.Push("GoLang!")

	// fmt.Println(stack)
	// stack.Pop()
	// fmt.Println(stack)

	for len(stack) >0 {
		s , boo := stack.Pop()
			if boo == true{
				fmt.Println(s)
			}
		}
	}




