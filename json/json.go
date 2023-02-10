package main

import (
	"encoding/json"
	"fmt"
)
type Ninja struct{
	Name string `json:"full_name"` 
	Weapon Weapon
	Level int
}

type Weapon struct{
	Name string `json:"Weapon_name"`
	
}

func main(){
	fmt.Println("JSON handling this time!!!")

	sFrom := `{"full_name" : "Onkar1","Weapon" : {"Weapon_name":"Chaku"},"Level" : 1}`
	fmt.Println(sFrom)
	var onkar Ninja
	err := json.Unmarshal([]byte(sFrom), &onkar)
	if err!=nil{
		panic(err)
		// fmt.Println("sad man")
	}
	fmt.Println(onkar)

	sTo , err := json.Marshal(onkar)
	if err!=nil{
		panic(err)
	}

	fmt.Println(sTo)
	fmt.Printf("%T \n",sTo)
	fmt.Printf("%s \n",sTo)
}