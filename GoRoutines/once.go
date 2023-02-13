package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionComplete bool
var once sync.Once

func main() {
	fmt.Println("Playing with once ")

	var wg sync.WaitGroup
	wg.Add(100)

	for  i:=0;i<100;i++{

		go func(){
			if foundTresure(){
				once.Do(markMissionComplete)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	checkMissionComplete()
}


func checkMissionComplete() {

	if missionComplete {
		fmt.Println("Mission Complete")
	} else {
		fmt.Println("Mission Not Complete")
	}
}

func markMissionComplete() {
	missionComplete = true
	for i:=0;i<5;i++{
		fmt.Println("Mission Complete is true")
	}
	

}

func foundTresure() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}
