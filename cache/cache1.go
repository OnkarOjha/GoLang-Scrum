package main

import (
	"fmt"
	"time"
	"github.com/patrickmn/go-cache"
)
type Example struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Working with caching in golang...")

	newExample := &Example{
		Name: "onkar",
		Age:  18,
	}
	c := cache.New(3*time.Minute, 5*time.Minute)

	abcd , found := c.Get("abcd")
	if found {
		a := abcd.([]Example)
		a = append(a , *newExample)
		c.Set("abcd" ,a , cache.NoExpiration)
        fmt.Println("Stored from if part")
	}else{
		c.Set("abcd" , newExample , cache.NoExpiration)
        fmt.Println("Stored from else part")
	}

	abcd , found = c.Get("abcd")
	if found{
			fmt.Println(abcd)
	}
}
