package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// fmt.Println("Working with maps...")
	// // mp := make(map[int]interface{})
	// // for i:=0; i<10;i++{
	// // 	go func(){
	// // 		mp[0]= 1
	// // 	}()
	// // }
	// // mp[0]=1
	// // mp[1]=2

	// //put values in map
	// syncMap := sync.Map{}
	// syncMap.Store(0, 1)
	// syncMap.Store(1, 2)
	// syncMap.Store(2, 3)
	// syncMap.Store(3, 4)

	// //getting values from map
	// syncValue , syncOk := syncMap.Load(0)
	// fmt.Println(syncValue,syncOk)

	// // deleting values from map
	// syncMap.Delete(0)

	// //getting values from map
	// syncValue, syncOk = syncMap.Load(1)
	// fmt.Println(syncValue)

	// // range in sync map
	// syncMap.Range(func(key,value interface{})bool{
	// 	fmt.Println(key,value,"|")
	// 	return true
	// })
	// instead of creating amutex to lock andf then unlock component s for addidingn we have mutex
	// var sum int64
	// fmt.Println(sum)
	// atomic.AddInt64(&sum, 1)
	// fmt.Println(sum)

	// var mut sync.Mutex
	// mut.Lock()
	// sum += 1
	// mut.Unlock()
	// fmt.Println(sum)

	// sum += 1
	// fmt.Println(sum)

	// // To subtract a signed positive constant value c from x, do AddUint64(&x, ^uint64(c-1)).
	// usum := uint64(sum)
	// c := 2
	// atomic.AddUint64(&usum, ^uint64(c-1))
	// fmt.Println(sum)
	// fmt.Printf("%T \n", usum)

	var av atomic.Value
	name := nameStruct{name: "Onkar"}
	av.Store(name)
	fmt.Println(av.Load().(nameStruct).name)

	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		w := av.Load().(nameStruct)
		w.name = "Not Onkar"
		av.Store(w) 
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(av.Load().(nameStruct).name)
}
type nameStruct struct {
	name string
}
