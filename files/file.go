package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("welcome to files in GoLang!!")
	content := "This needs to go to a file named - onkar.in"

	// os package for creating files
	file  ,err := os.Create("./mytryfile.txt")

	checkNilError(err)

	//io package for writing , reading , appending and all tasks

	length , err := io.WriteString(file ,content)
	checkNilError(err)

	fmt.Println("Length is : ",length)

	defer file.Close()

	readFile("./mytryfile.txt")
	
}

func readFile(filename string){
	databyte, err  := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
	
}

func checkNilError(err error){
	if err != nil {
		panic(err)

	}
}