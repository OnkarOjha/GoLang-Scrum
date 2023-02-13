package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Working with maps...")
	// mp := make(map[int]interface{})
	// for i:=0; i<10;i++{
	// 	go func(){
	// 		mp[0]= 1
	// 	}()
	// }
	// mp[0]=1
	// mp[1]=2

	//put values in map
	syncMap := sync.Map{}
	syncMap.Store(0, 1)
	syncMap.Store(1, 2)
	syncMap.Store(2, 3)
	syncMap.Store(3, 4)

	//getting values from map
	syncValue , syncOk := syncMap.Load(0)
	fmt.Println(syncValue,syncOk)

	// deleting values from map
	syncMap.Delete(0)
   

    //getting values from map
    syncValue, syncOk = syncMap.Load(1)
	fmt.Println(syncValue)

	// range in sync map
	syncMap.Range(func(key,value interface{})bool{
		fmt.Println(key,value,"|")
		return true
	})
	
}
