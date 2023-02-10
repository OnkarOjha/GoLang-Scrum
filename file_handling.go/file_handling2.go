package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	// "strings"
)

func main(){
	fmt.Println("Dynamic file handling")


	var fileName string
	fmt.Println("Enter the name of the file: ")
	fmt.Scanln(&fileName)

	fmt.Println("Enter the content: ")
	inputReader := bufio.NewReader(os.Stdin)
	input , err := inputReader.ReadString('\n')
	errCheck(err)

	CreateFile(fileName , input)
	ReadFile(fileName)
}

func CreateFile(filename , input  string){
	fmt.Println("Writing into the file!")
	file , err := os.Create(filename)
	errCheck(err)

	defer file.Close()

	//writing into the file
	len , err := file.WriteString(input)

	fmt.Println("Size of file is : ",len)
	fmt.Println("The name of the file is : ",file.Name())

}

func ReadFile(filename string){
	// now we will be reading from the file

	bytedata , err := ioutil.ReadFile(filename)
	errCheck(err)

	fmt.Println("the data of the file is: ",string(bytedata))
	fmt.Println("the name of the file is: ",filename)
}

func errCheck(err error){
	if err!=nil{
		panic(err)
	}
}