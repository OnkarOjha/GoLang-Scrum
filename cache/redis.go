package main

import (
	"fmt"
	"log"
	"github.com/gomodule/redigo/redis"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
type Podcast struct{
	Title string `redis:"title"`
	Creator string `redis:"creator"`
	Catoegory string `redis:"catoegory"`
	Fee float64 `redis:"fee"`
}

func main() {
	fmt.Println("working with cache and redis...")
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkError(err)
	defer conn.Close()
	_ ,err = conn.Do(

		"HMSET",
		"podcast:1",
		"title",
		"Tech Over Tea",
		"creator",
		"joe",
		"catorgory",
		"technology",
		"membership_fee",
		9.99,
	)
	checkError(err)
	title, err := redis.String(conn.Do("HGET", "podcast:1", "title"  ))
	checkError(err)
	fmt.Println("podcst title is: " , title)
	fee , err := redis.Float64(conn.Do("HGET", "podcast:1", "membership_fee"))
	checkError(err)
	fmt.Println("Membership fee is : ",fee)

	// TODO printing values in form of maps 
	values , err := redis.StringMap(conn.Do("HGETALL", "podcast:1"))
	checkError(err)
	fmt.Println("Values are : ", values) 

	// TODO printing valus in form of structs
	for k , v := range values { 
		fmt.Println("Key:",k)
		fmt.Println("Values: ",v)
	}	

	//TODO unmarshalling redis value on to a struct
	reply , err := redis.Values(conn.Do("HGETALL","podcast:1"))
	checkError(err)
	var podcast Podcast
	err = redis.ScanStruct(reply, &podcast)
	checkError(err)
	fmt.Println("podcast title is: ", podcast)
}
