package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Hello mod in go lang!!!")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")


	// activating the server in Golang
	// but this can through some error so we counter error as well , Scene ho jayeg bhai
	// log.Fatal comes in rescue to encounter errors , Bachane ko aya kaun Kaun , MAIN HO DON!!
	log.Fatal(http.ListenAndServe(":3000" , r))

}

func greeter(){
	fmt.Println("hey there mod users")
}

func serveHome(w http.ResponseWriter , r *http.Request){
	//sending response to the incoming request
	w.Write([]byte("<h1> Welcome to GoLang Onkar!!! </h1>"))

	// go mod tidy just removes all those packages that you are not using 
	// go list to show the list of libraries into which the code is currently dependant at
	// go mod verify to verify all the stuffs
	//go mod why expalins about the heirarchy of  dependencies
	// go mod edit to edit the mod file
	


}