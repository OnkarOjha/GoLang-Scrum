package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("Use of select statement !!!")
	one := make(chan string)
	two := make(chan string)

	go func(){
		time.Sleep(1 * time.Second)
		one <- "Hy One"
	}()

	go func(){
		time.Sleep(1 * time.Second)
		two <- "Hy Two!!"
	}()
		
	// jo pehle aya usko select kr lega

		
		select{
		case pehla := <-one:
			fmt.Println(pehla)
		case dusra := <-two:
			fmt.Println(dusra)

		}

	

}