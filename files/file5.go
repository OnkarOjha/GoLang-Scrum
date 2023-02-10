package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("This is for the 5th time")

	// need to open a file with append or write command 

	// filename , err := os.OpenFile("test5.txt", os.O_APPEND|os.O_CREATE , 0644)

	// if err!=nil{
	// 	panic(err)
	// }

	fmt.Println("Enter Filename :")
	var filename string
	fmt.Scanln(&filename)

	//data to be taken from the user

	fmt.Println("Enter text:")
	inputReader := bufio.NewReader(os.Stdin)
	data , _ := inputReader.ReadString('\n')

	WriteFile(filename ,data )
	ReadFile(filename)

}
// now lets create two function 
//1. to write content in the file
//2. to read content form the file

func WriteFile(filename , data string ){

	fmt.Println("Writing content into the file...")

	// first of all create file
	file , err := os.Create(filename)

	if err!=nil{
		panic(err)
	}

	len , err := file.WriteString(data)

	if err!=nil{
		panic(err)
	}

	defer file.Close()

	fmt.Println("Filename is: ",file.Name())
	fmt.Println("Size of File is: ",len)



}

func ReadFile(filename string){
	fmt.Println("Reading file starts now!!")

	data , err := ioutil.ReadFile(filename)

	if err!=nil{
		panic(err)

	}

	fmt.Println("Name of File : ",filename)
	fmt.Println("Length of file : ",len(data))
	fmt.Println("Data of file is: ", data)

}