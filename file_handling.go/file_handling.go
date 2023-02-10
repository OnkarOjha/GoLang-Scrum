package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){
	fmt.Println("Playing with files!!!")
	content := "This needs to go in a file - Onkar.in"
	// os package for creating file
	file , err := os.Create("./myfilego.txt")
	CheckErr(err)
	// io package for writing something in the file
	len ,err := io.WriteString(file  , content)

	CheckErr(err)

	fmt.Println("Length is : ",len)
	defer file.Close()

	ReadFile("./myfilego.txt")
}

// reading the file

func ReadFile(filename string){
	// ioutil for reading and other utlity function with the files
	// while reading file we will be getting data byte
	databyte , err := ioutil.ReadFile(filename)
	CheckErr(err)

	fmt.Println("The content is: ",string(databyte))


}
func CheckErr(err error){
	if err!=nil{
		panic(err)
	}
}