package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// urls := []string{
	// 	"https://api.openweathermap.org/data/2.5/weather?q=delhi&appid=4bcee6ca2132a9808f9b4fe9f0290ea2",
	// 	"https://api.openweathermap.org/data/2.5/weather?q=mohali&appid=4bcee6ca2132a9808f9b4fe9f0290ea2",
	// 	"https://api.openweathermap.org/data/2.5/weather?q=amritsar&appid=4bcee6ca2132a9808f9b4fe9f0290ea2",
	// }

	url1 := "https://api.openweathermap.org/data/2.5/weather?q=delhi&appid=4bcee6ca2132a9808f9b4fe9f0290ea2"
	url2 := "https://api.openweathermap.org/data/2.5/weather?q=mohali&appid=4bcee6ca2132a9808f9b4fe9f0290ea2"
	url3 := "https://api.openweathermap.org/data/2.5/weather?q=amritsar&appid=4bcee6ca2132a9808f9b4fe9f0290ea2"
	wg.Add(3)
	go func() {
		res, err := http.Get(url1)
		if err != nil {
			panic(err)
			
		}
		fmt.Println("response: ", res)
		fmt.Printf("%s: %s\n", url1, res.Status)
		defer wg.Done()
	}()
	go func() {
		res, err := http.Get(url2)
		if err != nil {
			panic(err)
			
		}
		fmt.Println("response: ", res)
		fmt.Printf("%s: %s\n", url2, res.Status)
		defer wg.Done()

	}()
	go func() {
		res, err := http.Get(url3)
		if err != nil {
			panic(err)
			
		}
		fmt.Println("response: ", res)
		fmt.Printf("%s: %s\n", url3, res.Status)
		defer wg.Done()

	}()

	// for _, url := range urls {
	// 	wg.Add(1)
	// 	go func(url string) {
	// 		defer wg.Done()
	// 		res, err := http.Get(url)
	// 		if err != nil {
	// 			fmt.Println("Error occurred: ", err)
	// 			return
	// 		}
	// 		defer res.Body.Close()
	// 		fmt.Println("result: ", res)
	// 		fmt.Println("")

	// 		fmt.Printf("%s: %s\n", url, res.Status)
	// 		fmt.Println("")
	// 	}(url)
	// }

	//^ execution without funciton
	// for _, url := range urls {
	// 	res, err := http.Get(url)
	// 	if err != nil {
	// 		fmt.Println("Error occurred: ", err)
	// 		return
	// 	}
	// 	defer res.Body.Close()
	// 	fmt.Println("result: ", res)
	// 	fmt.Println("")
	// 	fmt.Printf("%s: %s\n", url, res.Status)
	// }

	wg.Wait()
	fmt.Println("All API calls completed.")
}