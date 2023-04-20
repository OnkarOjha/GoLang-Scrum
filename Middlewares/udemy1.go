package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	http.ListenAndServe(":8080", router)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Welcome!\n"))
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("Onkar")
	w.Write([]byte("Hello, " + name + "!\n"))
}
