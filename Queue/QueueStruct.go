package main

import "fmt"

type Queue struct{
	Elements []int
	Size int
}

func (q *Queue) Enqueue(num int){
	if q.GetLen() == q.Size{
		fmt.Println("OVERFLOW!!")
		return
	}
	q.Elements = append(q.Elements, num)
}
func (q *Queue) GetLen()int{
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool{
	return len(q.Elements) == 0
}
func (q *Queue) Dequeue() int{
	if q.IsEmpty(){
		fmt.Println("UNDEFLOW!!")
		return 0
	}

	elements := q.Elements[0]

	if q.GetLen() == 1{
		q.Elements = nil
		return elements
	}

	q.Elements = q.Elements[1:]
	return elements
}

func main(){
	fmt.Println("Queue using structure!!")
	queue := Queue{Size : 3}
	fmt.Println(queue.Elements)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	fmt.Println(queue.Elements)
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println(queue.Elements)




}