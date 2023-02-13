package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Playing with mutex and race condition")
	var score = []int{0}

	wg := &sync.WaitGroup{}
	mut := sync.RWMutex{}

	wg.Add(3)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One Routine")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		defer wg.Done()
	}(wg, &mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Second Routine")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		defer wg.Done()

	}(wg, &mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Third Routine")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		defer wg.Done()

	}(wg, &mut)

	wg.Wait()
	fmt.Println(score)
	fmt.Println("Go routine finished")
}
