package main

import (
	cont "cont/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("File system as DB with CRUD")

	cont.People = append(cont.People, cont.Person{Id: 1, Name: "Onkar Ojha", Gender: "male", Email: "onkar123@gmail.com", Address: "mohali"})
	cont.People = append(cont.People, cont.Person{Id: 2, Name: "tech geek", Gender: "male", Email: "techgeek123@gmail.com", Address: "chandigarh"})

	http.HandleFunc("/create", cont.CreateClient)
	http.HandleFunc("/update/", cont.UpdateClient)
	http.HandleFunc("/delete/", cont.DeleteClient)
	http.HandleFunc("/read", cont.ReadClient)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
