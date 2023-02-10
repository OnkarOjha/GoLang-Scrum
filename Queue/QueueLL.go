package main

import (
	
	"container/list"
	"fmt"

)

func main(){
	fmt.Println("Queue using Linked List!!")

	queue := list.New()

	//Enqueue
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)
	fmt.Println("Queue: ",*queue)

	//Dequeue
	front := queue.Front()
	fmt.Println("Dequeued value is: ",front.Value)
	queue.Remove(front)
	fmt.Println("Queue: ",*queue)


}