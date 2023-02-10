package main

import (
	"fmt"
	"io"
	"net/http"
	// "os"
)


type logwriter struct{}

func main(){
	fmt.Println("this time for interface!!")


	resp , err := http.Get("https://www.google.com")

	if err!=nil{
		panic(err)
	}

	// fmt.Println(resp)

	// bs := make([]byte , 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logwriter{}


	io.Copy(lw, resp.Body)

}
func (logwriter) Write(b []byte) (int,error){
	fmt.Println(string(b))
	fmt.Println("The length of byte stream coming is: ",len(b))

	return len(b) , nil

}

