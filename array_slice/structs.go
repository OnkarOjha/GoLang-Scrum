package main

import "fmt"

func main(){
	fmt.Println("Welcome to structs")

	// NO INHERITANCE IN GOLANG  , NO SUPER OR PARENT

	onkar := User{"Onkar", "onkar@chicmic.co.in" , true , 21}
	fmt.Println(onkar)
	fmt.Printf("Onkar details are : %+v \n" , onkar)
	fmt.Printf("Onkar  Name is %v and Email is %v  and he is %v and his age is %v  \n" , onkar.Name , onkar.Email , onkar.Status , onkar.Age)
	
	// s := &User{
	// 	Name: "Onkar",
	// 	Email : "onkar@chicmic.co.in",
	// 	Status : true,
	// 	Age : 21,
	// }
	
	// m := Map(s)
	
	// fmt.Printf("%#v \n" , m["Name"])
	// fmt.Printf("%#v \n" , m["Email"])
	
	// fmt.Printf("%#v \n" , m["Status"])
	
	// fmt.Printf("%#v \n" , m["Age"])


}


type User  struct{
	Name string 
	Email string
	Status bool
	Age int
}



