package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	gettingReadyForMission()
	gettingReadyForMissionCond()
}

var ready bool

func gettingReadyForMissionCond() {
	cond := sync.NewCond(&sync.Mutex{})
	go gettingReadyWithCond(cond)

	cond.L.Lock()
	workIntervals := 0
	for !ready {
		// time.Sleep(5 * time.Second)
		workIntervals++
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Printf("We are now ready!! After %d Work intervals \n", workIntervals)
}

func gettingReadyForMission() {
	go gettingReady()

	workIntervals := 0
	for !ready {
		time.Sleep(5 * time.Second)
		workIntervals++
	}
	fmt.Printf("We are now ready!! After %d Work intervals \n", workIntervals)
}
func gettingReadyWithCond(cond *sync.Cond) {
	sleep()

	ready = true
	cond.Signal()
}
func gettingReady() {
	sleep()

	ready = true
}

func sleep() {
	rand.Seed(time.Now().UnixNano())
	someTime := time.Duration(1+rand.Intn(5)) * time.Second
	time.Sleep(someTime)
}
