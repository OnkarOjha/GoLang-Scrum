package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Working on redis for caching....")

	//create a redsi client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// set a cache value
	err := client.Set("key1", "onkar", 5*time.Second).Err()
	if err!=nil{
		panic(err)
	}

	// now retrive value from redis
	value , err := client.Get("key1").Result()
	if err == redis.Nil {
		fmt.Println("Key doesn't exits")
	}else if err!=nil{
		panic(err)
	}else{
		fmt.Println("value : ", value)
	}

	// wait for the key to expire
    time.Sleep(6 * time.Second)

	// check if the key has expired
    exists, err := client.Exists("mykey").Result()
    if err != nil {
        panic(err)
    }

	if exists == 1 {
        fmt.Println("Key still exists")
    } else {
        fmt.Println("Key has expired")
    }
}
