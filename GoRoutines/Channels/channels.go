package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	fmt.Println("Channels in GoLang!!")

	links :=  []string{
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://linkedin.com",
	}

	// create a new channel to communicate between different go routines
	c := make(chan string)

	//loop over links
	for _ , link := range links{
		//order is maintained when we print out messages
		// one request at a time is made s it produces delay
		// go keyword will make new go routines for
		// each call to the function
		go checkLink(link , c)
	}
	// it's kinda blockage to code so we exit after printing one link
	
	for l := range c {
		//funtion literal
		go func(l string){
		time.Sleep(5*time.Second) 
		checkLink(l,c)
		}(l)
	}
	// fmt.Println("kyabbatay")

	// for {
	// 	time.Sleep(5*time.Second) 
	// 	go checkLink(<-c,c)
	// }


}
func checkLink(link string , c chan string){
	_, err := http.Get(link)
	if err!=nil{
		fmt.Println(link , "might be down!!")
		c<- link
		return
	}

	fmt.Println(link,"Link is UP1!!")
	c<- link
}