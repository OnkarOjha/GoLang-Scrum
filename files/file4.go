package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main(){
	filename , err := os.OpenFile("test4.txt" , os.O_APPEND|os.O_CREATE , 0644)

	//  filename := "test4.txt"

	// val:= "old\nfalcon\nsky\ncup\nforest\n"
	 
	len , err := filename.WriteString("BHAI SAHABFOURTH TIME")
	errorFunc(err)

	fmt.Println("The size: ", len)
	fmt.Println("Done")

	



	

	

}

func WriteFile(){
	fileName := "data.txt"

    val := "old\nfalcon\nsky\ncup\nforest\n"
    data := []byte(val)

    err := ioutil.WriteFile(fileName, data, 0644)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("done")
}

func errorFunc(err error){
	if err!=nil{
		panic(err)
	}
}