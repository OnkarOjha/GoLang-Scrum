package main

import (
	"fmt"

	// "github.com/google/go-dap"
)

type node struct{
	data int 
	next *node
}
type linkedlist struct{
	length int
	head *node
	tail *node
}

//function to return length of the linkedlist

func (l linkedlist) Len() int{
	return l.length
}

// function to display the ll

func (l linkedlist) Display(){
	for l.head!=nil{
		fmt.Printf("%v -> ",l.head.data)
		l.head = l.head.next
	}
	fmt.Println()
}

// ll pushback function

func (l *linkedlist) Pushback(n *node){
	if l.head == nil{
		l.head = n
		l.tail = n
		l.length++
	}else{
		l.tail.next = n
		l.tail = n
		l.length++
	}
}


func main(){
	fmt.Println("playing with linked list this time !!!")
	node1 := &node{data : 20}
	node2:= &node{data : 30}
	node3:= &node{data : 40}
	node4:= &node{data : 50}
	node5:= &node{data : 60}
	// node1.next = node2
	// fmt.Print(node1)

	//initialisation of linkedlsit

	list :=linkedlist{}

	list.Pushback(node1)
	list.Pushback(node2)
	list.Pushback(node3)
	list.Pushback(node4)
	list.Pushback(node5)
	list.Display()


}