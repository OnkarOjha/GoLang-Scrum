package main

import (
	"fmt"
	"log"
	"net/http"
)
func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"hello Home")
}
func setupRoutes() {
	http.HandleFunc("/", handleHome)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
	fmt.Println("Working with Youtube subscriber API")
	setupRoutes()
}
