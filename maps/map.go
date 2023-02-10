package main

import (
	"fmt"
	// "sort"
)

func main(){
	// mp := map[int]string{
	// 	90 : "wfe",
	// 	91 : "wfe",
	// 	92 : "wfe",
	// 	93 : "wfe",
	// 	94 : "wfe",
	// 	95 : "wfe",
	// }

	// fmt.Println(mp)

	// var mp = make(map[int]string , 2)

	// fmt.Println(len(mp))

	// mp[2] = "oko"
	// mp[3] = "oko"
	// mp[8] = "oko"
	// mp[1] = "rt" // overridjng the earlier initialised value

	// keys := make([]int , 0 , len(mp))

	// for k := range mp{ 
	// 	keys = append(keys , k)
	// }
	// sort.Ints(keys)
	// // fmt.Println(keys)


	// for _ , v := range keys{
	// 	fmt.Println(v , mp[v])
	// }
	
	// fmt.Println(mp)
	// fmt.Println(" length of map is " ,len(mp))
	// delete(mp , 2)
	// fmt.Println("after deletion: ",mp)

	// ma := map[int]string{}
	// fmt.Println(ma)
	mp := make(map[string]float64)
	mp["pi"] = 3.14
	
	fmt.Println(mp["pi"])

	v := mp["pi"]
	fmt.Println("The value of v : ",v)
	v = mp["pie"] // value assigned is 0
	fmt.Println("The value of v : ",v)

	fmt.Println(mp["pi"])

	delete(mp , "pi")
	fmt.Println(mp)

}