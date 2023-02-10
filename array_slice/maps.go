package main

import (
	"fmt"

	// "golang.org/x/text/language"
)

func main(){
	fmt.Println("Welcome to maps!!")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["C++"] = "C plus plus"
	languages["RB"] = "Ruby"

	fmt.Println("List of all languages: ",languages)
	fmt.Println("JS shorts for: ",languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ",languages)

	//loops are interesting here actually
	for key , value := range languages{
		fmt.Printf("For key %v , value is %v \n",key,value)

	}
}