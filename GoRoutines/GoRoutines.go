package main

import (
	"fmt"
	"net/http"
	"sync"
)
var signals = []string{}
var waitGroup sync.WaitGroup // pointer to wait group
var mut sync.Mutex // pointer to mutex

func main() {
	// go greeter("hello")
	// greeter("hello go")
	// // we never waited for the hello to be printed
	websiteList := []string{
		"http://google.com",
		"http://facebook.com",
		"http://youtube.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://linkedin.com",
	}

	for _, web := range websiteList {

		go getStatusCode(web)
		waitGroup.Add(1) // only 1 go routine fired up!
    }
	waitGroup.Wait() // this will not allow main method to finish
	// until it recieves a response signal from the method called
	fmt.Println(signals)
}



func getStatusCode(endpoint string) {
	// this will send signal to wait in the main
	defer waitGroup.Done()


	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops endpoint", err)

	} else {
		mut.Lock() // jab tk append na ho jay band kr do
		signals = append(signals, endpoint)
		mut.Unlock() // ab append bnd ho gya ab unlock kr do
		fmt.Printf("%d status code for website %s", res.StatusCode, endpoint)
		fmt.Println()
	}
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(5 * time.Second)
// 		fmt.Println(s)
// 	}
// }
