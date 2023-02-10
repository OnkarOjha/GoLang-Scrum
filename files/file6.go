package main

import (
	"fmt"
	"os"
)

type Person struct{
	Name string
	Age int
	Password string
	Email string
}

func createFile(content Person){
file , err := os.Create("json.txt")
if err!= nil {
    fmt.Println(err)
    return
}
defer file.Close()

err := json.Unmarshal(content, &file)
if err!= nil {
    fmt.Println(err)
    return
}

length ,err := file.WriteString(content)
fmt.Println("length ",length)
	


}

func main(){
	fmt.Println("Marshalling and unmarshalling in files at go")
	p := Person{Name:"Onkar", Age:21 , Password:"Onkar@123",Email:"onkar@in"}
	// fmt.Println(p)
	createFile(p)
}