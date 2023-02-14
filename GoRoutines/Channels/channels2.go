package main

import (
	"fmt"
	"runtime"
	"sync"
	
	// "time"
)

func square(n float64) float64 {
	return n * n
}

// func main() {
// 	ch := make(chan float64)

// 	go func() {
// 		ch <- square(5.7)
// 	}()
// 	fmt.Println(<-ch)
// 	close(ch)
// 	fmt.Printf("%T \n" , <-ch)
// }

func main() {
	fmt.Println("CPU's : ", runtime.NumCPU())
	fmt.Println("GoRoutines : ", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	var mut sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mut.Lock()
			v := counter
			// time.Sleep(2 * time.Second)
			runtime.Gosched() // allow something else to run
			v++
			counter = v
			mut.Unlock()
			defer wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("Counter : ", counter)
	fmt.Println("GoRoutines : ", runtime.NumGoroutine())

}
