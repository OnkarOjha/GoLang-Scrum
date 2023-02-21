package main

import (
	// "errors"
	"fmt"
	// "os"
	"time"

	"github.com/patrickmn/go-cache"
)

type Person struct {
	Name string
	Age  int
	b    bool
}

func main() {
	// start:= time.Now()
	fmt.Println("Working with cache from documentation")
	// c := cache.New(5*time.Minute, 10*time.Minute)

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

	// //sendig multiple values to cache
	// c.Set("multiple",Person{Name : "Onkar" , Age : 21 , b : true},cache.DefaultExpiration)
	// val , found := c.Get("multiple")
	// if found{
	//    multiple := val.(Person)
	//    fmt.Println(multiple)
	// }else{
	//     fmt.Println("not found")
	// }
	// c.Delete("multiple")
	// _, found = c.Get("multiple")
	// if found{
	// 	multiple := val.(Person)
	// 	fmt.Println(multiple)
	// }else{
	// 	fmt.Println("not found")
	// }
	// items := map[string]cache.Item{
	// 	"key1": cache.Item{
	// 		Object:     "Onkar",
	// 		Expiration: time.Now().Add(10 * time.Second).UnixNano(),
	// 	},
	// 	"key2": cache.Item{
	// 		Object:     21,
	// 		Expiration: time.Now().Add(20 * time.Second).UnixNano(),
	// 	},
	// 	"key3": cache.Item{
	// 		Object:     true,
	// 		Expiration: time.Now().Add(30 * time.Second).UnixNano(),
	// 	},
	// }
	// // now create the default time out and expire items
	// c := cache.NewFrom(5*time.Minute, 10*time.Minute, items)

	// // val, found := c.Get("key1")

	// // if found {
	// // 	fmt.Println("value of key1", val)
	// // } else {
	// // 	fmt.Println("not found")
	// // }
	// // calling then add fucntion to add an existing key value pair
	// // err := c.Add("key1", "Onkar" ,10*time.Second )
	// // if err!=nil{
	// // 	fmt.Println("Error adding key",err)
	// // }

	// //decrementing values
	// err := c.Decrement("key2" , 10)
	// if err!=nil{
	// 	fmt.Println("While decrementing key2",err)
	// }else{
	// 	fmt.Println("key decrement success")
	// 	val , found := c.Get("key2")
	// 	if found{
	// 		fmt.Println("key value", val)
	// 	}
	// }

	// itemCount := c.ItemCount()
	// fmt.Println("Items in the cache", itemCount)

	// var i map[string]Item
	// i = c.Items()
	// fmt.Println("Items in the cache", i)
	// // //set the value of any key provided by you  with default expiration time

	// Create a new cache instance with a default expiration time of 5 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)
	c.Set("my_key", "Onkar" , cache.DefaultExpiration)
	// saving data into file using savefile method 
	err := c.SaveFile("my_file.txt")
	if err!=nil{
		fmt.Println("error in saving data to file",err)
		return
	}
	// // Load a cache dump from a file
	// file, err := os.Create("cache_dump.gob")
	// if err != nil {
	// 	fmt.Println("Error Creating cache dump file:", err)
	// 	return
	// }
	// defer file.Close()
	// // saving the data in the file
	// err = c.Save(file)
	// if err != nil {
	// 	fmt.Println("Error saving cache dump:", err)
	// 	return 
	// }

	// // clear the cache to simulate a fresh cache dump
	// c.Flush()

	// // Load a cache dump from a file
	// file , err := os.Open("my_file.txt")
	// if err!= nil {
    //     fmt.Println("Error Opening cache dump file:", err)
    //     return
    // }
	// defer file.Close()
	err = c.LoadFile("my_file.txt")
	if err!= nil {
        fmt.Println("Error loading cache dump:", err)
        return 
    }
	  // Retrieve the cached data
	  value, found := c.Get("my_key")
	  if found {
		  fmt.Println("Cached value for 'my_key':", value)
	  } else {
		  fmt.Println("Item 'my_key' not found in cache.")
	  }


	 
}

// rewriting the add function
// func (c Cache) Add(key string, x interface{}, d time.Duration) error {
// 	// now set value for already existing key
// 	_, found := c.items[key]
// 	if found {
// 		return errors.New("Key already exists")
// 	}
// 	return c.Set(key , x , d)
// }
// package main

// import (
//     // "errors"
//     "time"
// 	"fmt"
//     "github.com/patrickmn/go-cache"
// )

// func main() {
//     // Create a new cache instance with a default expiration time of 5 minutes
//     c := cache.New(5*time.Minute, 10*time.Minute)

//     // Add a key-value pair to the cache with an expiration time of 1 minute
//     err := c.Add("key1", "value1", 1*time.Minute)
//     if err != nil {
//         fmt.Println("Error adding to cache:", err)
//     }

//     // Add the same key-value pair to the cache again with the same expiration time
//     err = c.Add("key1", "value1", 1*time.Minute)
//     if err != nil {
//         fmt.Println("Error adding to cache:", err)
//     }
// }

// // Add adds a value to the cache, but only if the key does not already exist.
// // Returns an error if the key already exists in the cache.
// // func (c cache) Add(k string, x interface{}, d time.Duration) error {
// //     _, found := c.items[k]
// //     if found {
// //         return errors.New("Key already exists in cache")
// //     }
// //     return c.Set(k, x, d)
// // }
