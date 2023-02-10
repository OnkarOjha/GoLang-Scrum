package main

import (
	"bufio"
	"fmt"
	"os"

	// "golang.org/x/tools/godoc/redirect"
)

func main(){
 reader := bufio.NewReader(os.Stdin)
 fmt.Println("enter the rating for our Pizza: ")

 // commma ok syntax // error ok syntax

 input , _ := reader.ReadString('\n') // i want to read unless I found a \n 
 fmt.Println("Thanks for rating :",input)
 fmt.Printf("Type of this rating %T \n",input)
 
}