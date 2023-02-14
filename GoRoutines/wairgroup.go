package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("GoRoutines: \t", runtime.NumGoroutine())

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
	}()
	fmt.Println("GoRoutines: \t", runtime.NumGoroutine())

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
	}()
	fmt.Println("GoRoutines: \t", runtime.NumGoroutine())

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 3")
	}()

	wg.Wait()
	fmt.Println("All goroutines complete")

	fmt.Println("OS \t", runtime.GOOS)
	fmt.Println("ARCH \t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("GoRoutines: \t", runtime.NumGoroutine())

}
