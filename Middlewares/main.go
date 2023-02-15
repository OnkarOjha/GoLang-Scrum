package main

import (
	
	"fmt"
	"net/http"

)

func main() {
	http.Handle("/", handler)
	http.HandleAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing the handler")
	w.Write([]byte("OK"))
}

