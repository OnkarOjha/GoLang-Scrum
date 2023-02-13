package main

import "fmt"

func main() {
	fmt.Println("Select sattement with Golang channels")
	ninja1 := make(chan interface{}) ; close(ninja1)
	ninja2 := make(chan interface{});close(ninja2)
	// go captainElect(ninja1 , "Ninja 1")
	// go captainElect(ninja2 , "Ninja 2")
	// select{
	// case message := <-ninja1:
	// 	fmt.Println(message)
	// case message := <-ninja2:
    //     fmt.Println(message)
	// }

	ninja1Count := 0
	ninja2Count := 0

	for i :=0 ; i<100 ;i++{
		select {
        case <-ninja1:
            ninja1Count++
        case <-ninja2:
			ninja2Count++
            
	}
   
}
fmt.Printf("ninja 1 count: %d\n", ninja1Count)
fmt.Printf("ninja 2 count: %d\n", ninja2Count)

}

func captainElect(ninja chan string , message string){
	ninja <- message
}

