package main

import "fmt"

func enqueue(q []int , num int)[]int{

	q = append(q , num)
	fmt.Println("enqueues value : ",num)
	return q

}

 func dequeue(q []int) ([]int ,  int){
	element := q[0]

	if len(q) ==1{
		var tmp = []int{}
		return tmp , element
	}

	return q[1:] , element
 }

func main(){
	fmt.Println("Implementation of queue using SLICE in go!!")
	var queue = make([]int , 0)

	queue = enqueue(queue , 10)
	fmt.Println("queue after pushing 10 ",queue)

	queue , element := dequeue(queue)
	fmt.Println("The dequeued element is: ",element)
	fmt.Println("The new queue is: ",queue)




}