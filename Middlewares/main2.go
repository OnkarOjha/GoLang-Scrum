package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// myhandler := http.HandlerFunc(handler)
	// // http.HandleFunc("/", handler)
	// http.Handle("/", middleware(myhandler))
	// http.Handle("/cookies", setTimeCookie(myhandler))
	// http.Handle("/filter", filterContentType(http.HandlerFunc(postHandler)))

	myHandler := http.HandlerFunc(postHandler)
	// here we are chaning middlewares but we don't to chain it like this 
	// if we are having 100 middlewares then we can't use it like m1(m2(m3(m4....)))
	// chain := filterContentType(setTimeCookie(myHandler))
	chain := CreateChain(filterContentType , setTimeCookie).Then(myHandler)
	http.Handle("/", chain)


	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing Handler")
	w.Write([]byte("Hy home Handler"))
}

func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Running before Handler")
		w.Write([]byte("Hijacking Handler"))
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Running after Handler")
	})

}

// Write a middleware that makes sure request has Header "Content-Type" application/json

func filterContentType(handler http.Handler) http.Handler {
	fmt.Println("Running filter content")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("405 - Header Content-Type incorrect"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Write a middleware that adds current server time to the reponse cookie

func setTimeCookie(handler http.Handler) http.Handler {
	fmt.Println("Running cookies")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cookie here is a struct
		cookie := http.Cookie{
			Name:  "Server Time(UTC)",
			Value: strconv.Itoa(int(time.Now().Unix())),
		}
		fmt.Println(cookie)

		http.SetCookie(w, &cookie)
		handler.ServeHTTP(w, r)
	})

}

// Now let us create a handler to handle the POST request, and then we will use the middlewares:

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
		return
	}

	var p Person
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&p)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 -Internal Server Error"))
		return 
	}
	defer r.Body.Close()

	fmt.Println("Got First Name: ", p.FirstName)
	fmt.Println("Got Last Name: ", p.LastName)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("201 - Created"))

}

 //TODO let's add a better way to handle middlewares
 type Middleware func(http.Handler) http.Handler
 type Chain []Middleware
 // returning  a slice of middlewares
 func CreateChain(middlewares ...Middleware) Chain{
	var slice Chain
	return append(slice , middlewares...)
 }
 // function for adding middlewares to the chain

 func (c Chain) Then(originalHandler http.Handler) http.Handler {
	// if originalHandler == nil{
	// 	originalHandler = http.DefaultServeMux
	// }
	for i := range c{
		// Same as to m1(m2(m3(originalHandler)))
		originalHandler = c[len(c)-1-i](originalHandler)
	}
	return originalHandler
 }
