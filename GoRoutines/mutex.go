package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	fmt.Println("Playing with mutex and race condition")
// 	var score = []int{0}

// 	wg := &sync.WaitGroup{}
// 	mut := sync.RWMutex{}

// 	wg.Add(3)

// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("One Routine")
// 		m.Lock()
// 		score = append(score, 1)
// 		m.Unlock()
// 		defer wg.Done()
// 	}(wg, &mut)

// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("Second Routine")
// 		m.Lock()
// 		score = append(score, 2)
// 		m.Unlock()
// 		defer wg.Done()

// 	}(wg, &mut)

// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("Third Routine")
// 		m.Lock()
// 		score = append(score, 3)
// 		m.Unlock()
// 		defer wg.Done()

// 	}(wg, &mut)

// 	wg.Wait()
// 	fmt.Println(score)
// 	fmt.Println("Go routine finished")
// }

var count int
var mut sync.RWMutex
var wait sync.WaitGroup
var rwlock sync.RWMutex

func main() {
	// basics()
	readWrite()
}
func readWrite() {
	go write()
	go read()
	go read()
	

	
	go read()
	time.Sleep(5 * time.Second)
	fmt.Println("Done wiuth reading and writing")
}

func read() {

	rwlock.RLock()
	defer rwlock.RUnlock()

	fmt.Println("Read Locking...")
	time.Sleep(1* time.Second)

	fmt.Println("Reading Unlocked...")

}

func write() {

	rwlock.Lock()
	defer rwlock.Unlock()

	fmt.Println("Write Locking...")
	time.Sleep(1 * time.Second)

	fmt.Println("Writing Unlocked...")

}
func basics() {
	iterations := 1000
	wait.Add(iterations)
	for i := 0; i < iterations; i++ {
		go increment()
	}
	wait.Wait()

	fmt.Println("Resulted count is : ", count)
}
func increment() {
	mut.Lock()
	count++
	mut.Unlock()
	defer wait.Done()

}
