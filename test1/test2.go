package main

import (
	
	"fmt"
	 "strings"

)

// func count(T any)(arr []T,f func(T) bool) int {
// 	count := 0

// 	for _ , s := range arr{
// 		if f(s) {
// 			count++
// 		}
// 	}

// 	return count
// }

func coun(T comparable)(arr []T,s bool) int {
	count := 0

	for _ , s := range arr{
		if s {
			count++
		}
	}

	return count
}

func main(){
	fmt.Println("The second test!!")

	arr := []byte{'c', 'd', 'b', 'd', 'a', 'a', 'd', 'b', 'd', 'c', 'e', 
	'c', 'h', 'e', 'b', 'c', 'f', 'a', 'e', 'd', 'b', 'b', 'b', 'h', 'f', 'f', 'b',
	 'h', 'g', 'e', 'g', 'g', 'b', 'c', 'b', 'c', 'h','z', 'b', 'h', 'a', 'b', 'e', 
	 'a', 'c', 'e', 'd', 'h', 'c', 'b', 'h', 'h', 'c', 'b', 'f', 'h', 'e', 'c', 'c', 'a', 
	 'h', 'g', 'a', 'g', 'a', 'e', 'c', 'f', 'h','q' ,'c', 'd', 'b', 'f', 'g', 'c', 'e', 'f',
	  'a', 'e', 'f', 'f', 'a', 'g', 'g', 'g', 'g', 'h','o' ,'a', 'a', 'e', 'b', 'g', 'f', 'c',
	   'g', 'a', 'a', 'h', 'a', 'f', 'f'}
	// fmt.Println(arr)
	
	fmt.Println(coun(arr , strings.Contains(arr , "a")))



	// mp:= make(map[string]int)

	// for _ , val:= range arr{
	// 	mp[string(val)]++
	// }

	// //choose key value pair
	// for k , val := range mp{
	// 	fmt.Print(k + " : ")
	// 	fmt.Println(val)
	// }
}