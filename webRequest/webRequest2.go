// package main
	
// import (
//     "fmt"
//     // "net"
//     "net/url"
// )

//  const myurl string = "https://lco.dev.3000/learn?coursename=reactjs&paymentid=32424er"

// func main(){
// 	fmt.Println("Welcome to handling URL's")
// 	fmt.Println(myurl)

// 	//parsing

// 	result , err := url.Parse(myurl)

// 	if err!=nil{
// 		panic(err)
// 	}

// 	// fmt.Println(result.Scheme)
// 	// fmt.Println(result.Host)
// 	// fmt.Println(result.Path)
// 	// fmt.Println(result.RawQuery)
// 	// fmt.Println(result.Port())

// 	//store everything into varivbl

// 	qparams := result.Query()

// 	fmt.Printf("The type of qparams is: %T \n", qparams)

// 	fmt.Println(qparams["coursename"])
// 	fmt.Println(qparams["paymentid"])

// 	for _ , val := range qparams{
// 		fmt.Println("params is :  ",val)
// 	}







// 	// now take out the part of url
// 	partsOFUrl := &url.URL{
// 		Scheme : "https",
// 		Host: "lco.dev",
// 		Path : "/tutcss",
// 		RawPath:  "user=hitesh",
// 	}
// 	//construct a url out of it
// 	anotherUrl := partsOFUrl.String()
// 	fmt.Println(anotherUrl)
// }


package main

import (
	"fmt"
	"net/url"
)

// const myurl string ="https://stackoverflow.com/questions/33885235/should-a-go-package-ever-use-log-fatal-and-when"
const myurl string = "https://lco.dev.3000/learn?coursename=reactjs&paymentid=32424er"

func main(){
	fmt.Println("Constructing and Deconstructing with URL!!")

	//pehle url ko parse krnge
	// parsedUrl , err := url.Parse(myurl)

	// if err!=nil{
	// 	panic(err)
	// }

	// ab url  ko query krnge

	// queryUrl := parsedUrl.Query()
	// if err!=nil{
	// 	panic(err)
	// }

	// fmt.Printf("The type of url is %T \n",queryUrl)
	// fmt.Println("Scheme is: ",parsedUrl.Scheme)
	// fmt.Println("Host is: ",parsedUrl.Host)
	// fmt.Println("Port is: ",parsedUrl.Port())
	// fmt.Println("Path is: ",parsedUrl.Path)
	// fmt.Println("Raw Query is: ",parsedUrl.RawQuery)

	// fmt.Println("Scheme is: ",parsedUrl.Scheme)

	// for _ , val := range queryUrl{
	// 	fmt.Println("Params are: ",val)
	// }



	// constructing URL

	partsOfUrl := &url.URL{
		Scheme: "https",
		User:   url.UserPassword("user", "password"),
		Host: "google.com",

	}
	fmt.Println(partsOfUrl.Redacted())// encrypt krta hai xxxx form mein
	partsOfUrl.User = url.UserPassword("me", "newerPassword")
	fmt.Println(partsOfUrl.Redacted())
	partsOfUrl.User = url.UserPassword("me", "newer")
	fmt.Println(partsOfUrl.Redacted())

	newUrl := partsOfUrl.String()

	fmt.Println("New Url is: ", newUrl)
	fmt.Println("New Url is: ", partsOfUrl)




}