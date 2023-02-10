package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
)

func main(){

	fmt.Println("web request handling 3!!!")
	// PerformGetRequest()
	PerformPostRequest()
	// PerformFormRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"
	response  , err:= http.Get(myurl)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)
	// content is in byte format so type casting into string format
	content , _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))


	//premium method to get the data and convert it from byte to string

	var responseString strings.Builder

	
	byteCount , _ := responseString.Write(content)

	fmt.Println("bytecount is : ",byteCount)
	fmt.Println(responseString.String())

	

}

//POST REQUEST	

func PerformPostRequest(){
	const myurl = "http://localhost:8000/post"

	// fake JSON payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "GoLang Let's go!!",
			"price" :0,
			"platform" : "onkar.in"
		}
	`)

	response , err := http.Post(myurl , "application/json" , requestBody)

	if err!=nil{
		panic(err)
	}

	defer response.Body.Close()

	content , err := ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}

	fmt.Println(string(content))

	
}

//POST FORM REQUEST	

func PerformFormRequest(){
	const myurl = "http://localhost:8000/postform"

	//create FORMDATA

	data := url.Values{}
	data.Add("firstName" , "Onkar")
	data.Add("lastName" , "ojha")
	data.Add("email" , "onkar@chicmic.co.in")

	response , err := http.PostForm(myurl , data)
	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()
	content , _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}