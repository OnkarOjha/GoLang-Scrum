package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 3")
	}()

	wg.Wait()
	fmt.Println("All goroutines complete")
}
