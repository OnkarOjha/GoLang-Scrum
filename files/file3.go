package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Dynamic taking entry and printing in file")

	// ab yaha ek ek krke bulaoo 
	// data := "Bhai teesri baar try kr rha is to dynamic"
	// // filename := "test2.txt"
	// CreateFile("test2.txt" ,  data)
	// ReadFile("test2.txt")

	fmt.Println("Enter filename: ")
	var fileName string
	fmt.Scanln(&fileName)

	// user input for file content

	fmt.Println("Enter text:")
	inputReader := bufio.NewReader(os.Stdin)
	input , _ := inputReader.ReadString('\n')


	// now file has been created and read it now

	// cal both the functions after taking all the inputs

	CreateFile(fileName , input)
	ReadFile(fileName)


}
//creating a function that will dynamically take name of the file that
// the user needs and what type of text the user actually nees to put in
func CreateFile(filename  , data string){

	fmt.Print("Writing in the file\n")

	file , err := os.Create(filename)

	errorFunc(err)	

	defer file.Close()

	// now writing data to the file

	len , err := file.WriteString(data)

	errorFunc(err)

	// just check the name of the files 
	fmt.Println("\n File name: ",file.Name())
	fmt.Println("\n Size of the file: ", len)



}

// now create a function to read whatever you have written into the file

func ReadFile(filename string){
	fmt.Print("Now we will be reading what is written in the file..\n")


	// so for reading we have to check the io package

	data , err := ioutil.ReadFile(filename)
	errorFunc(err)

	fmt.Println("The name of the file is : ",filename)
	fmt.Println("The size of the file is : ",len(data))
	fmt.Println("\n Data: ",data)
}

func errorFunc(err error){
	if err!=nil{
		panic(err)
	}
}