package main

import(
	"fmt"
	"sort"
)

type Person struct{
	Name string
	age int
}

func main(){
	fmt.Println("Playing with slice sorting this time!!!")
	 person := []Person{
		{"onkar" , 2},
		{"Shyam" , 23},
		{"Ram" , 23},
		{"David" , 21},
	 }

	 sort.SliceStable(person , func(i, j int) bool {
		return person[i].age < person[j].age
	 })

	// sort.Slice(person , func(i, j int) bool {
	// 	return person[i].age < person[j].age
	//  })

	 for _ , v:= range person{
		fmt.Println(v)
	 }




}