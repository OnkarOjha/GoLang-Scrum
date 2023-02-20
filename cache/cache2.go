package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)
type Person struct{
	Name string
	Age int
	b bool
}
func main() {
	// start:= time.Now()
	fmt.Println("Working with cache from documentation")
	c := cache.New(5*time.Minute, 10*time.Minute)

	// // set the value of any key provided by you  with default expiration time
	// // or yaha pr hmne explicitly default time s e bhi chota time period provide kr diya jisse 5 minute toi hai hi pr ab 2 minute mei  hi expire kr jaynge
	// c.Set("key1","Onkar" , 5*time.Second)
	// // warna hum default time method ka v use kr ske the \

	// // c.Set("key1" , "Onkar", cache.DefaultExpiration)
	// time.Sleep(10*time.Second)

	// // get the value by providing the key
	// val , found := c.Get("key1")
	// if found{
	// 	fmt.Println(val)
	// }else{
	// 	fmt.Println("not found")
	// }
	// fmt.Println("Time ends at: ",time.Since(start))
	

	//sendig multiple values to cache
	c.Set("multiple",Person{Name : "Onkar" , Age : 21 , b : true},cache.DefaultExpiration)
	val , found := c.Get("multiple")
	if found{
       multiple := val.(Person)
	   fmt.Println(multiple)
    }else{
        fmt.Println("not found")
    }

	}



