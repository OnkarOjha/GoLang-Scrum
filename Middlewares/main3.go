package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/", home)
	// err := http.ListenAndServe(":8080", handler{})
	// log.Fatal(err)

	//as we want all  our middlewares to work before the routes are being called
	//so  we need to call the middlewares before the routes
	// in simple language we can do this by creating a simple servermux
	mux := http.NewServeMux()
	// now use this mux to run our home route
	
	mux.HandleFunc("/", home)
	// mux.Handle("/product", logRequestMiddleware(http.HandlerFunc(productHandler)))
	mux.HandleFunc("/user", userHandler)
	// mux.Handler()

	// now place the middleware before SERVERMUX
	err := http.ListenAndServe(":8000" , logRequestMiddleware(secureHeadersMiddlewareMiddle(mux)))
	// err := http.ListenAndServe(":8000", mux)

	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("operating on home page")
	w.Write([]byte("kem cho!!!!!!!!"))
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("executing  the custom handler")
	w.Write([]byte("Data from custom handler"))
}

// logRequestMiddleware logs basic info of a HTTP request
// RemoteAddr: Network address that sent the request (IP:port)
// Proto: Protocol version
// Method: HTTP method
// URL: Request URL

func logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LOG %s - %s %s %s\n", r.RemoteAddr, r.Proto, r.Method, r.URL.Path)
		
		next.ServeHTTP(w, r)
		

	})
}

// secureHeadersMiddleware adds two basic security headers to each HTTP response
// X-XSS-Protection: 1; mode-block can help to prevent XSS attacks
// X-Frame-Options: deny can help to prevent clickjacking attacks

func secureHeadersMiddlewareMiddle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode-block")
		w.Header().Set("X-Frame-Options", "deny")

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
