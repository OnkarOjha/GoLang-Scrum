package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// URL definittion

// const url = "https://lco.dev"

const url ="https://www.incrementors.com/tools/dummy-content-generator/"

func main(){
	fmt.Println("THis is for Web Request")

	response  , err := http.Get(url)

	if err!=nil{
		panic(err)
	}

	fmt.Printf("REponse is of type: %T \n",response)

	//it is my responsibility to close this connection
	defer response.Body.Close()

	// reading response
	databytes , err := ioutil.ReadAll(response.Body)

	if err!=nil{
		panic(err)

	}

	content := string(databytes)

	fmt.Println(content)

}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )


// const myurl  = "http://localhost:8000/get"
// func main(){
// 	fmt.Println("Revision time boss , Get REQUEST!!")
// 	response , err := http.Get(myurl)

// 	if err!=nil{
// 		panic(err)
// 	}

// 	// now we will read the data from ioutil

// 	data , err := ioutil.ReadAll(response.Body)
// 	if err!=nil{
// 		panic(err)
// 	}

// 	fmt.Println(string(data))




	
// }