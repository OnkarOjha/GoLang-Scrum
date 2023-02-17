package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// mux := http.NewServeMux()
	http.Handle("/user", logRequestMiddleware(http.HandlerFunc(productHandler)))
	http.HandleFunc("/user/hello", userHandler)
	//err := http.ListenAndServe(":4000", nil)
	err := http.ListenAndServe(":4000", nil)
	//err := http.ListenAndServe(":4000", nil)

	log.Fatal(err)

}

func logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LOG %s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

	})
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing productHandler")

	w.Write([]byte("Hello from productHandler"))
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing userHandler")

	w.Write([]byte("Hello from userHandler"))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing Home handler")

	w.Write([]byte("Hello Home handler"))
}
