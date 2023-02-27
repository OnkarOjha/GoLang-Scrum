package main

import (
	"fmt"
	"net/http"
	"regexp"
)

var (
	listusersRe = regexp.MustCompile(`^\/users[\/]*$`)
	getusersRe = regexp.MustCompile(`^\/users[\/d+]*$`)
	createusersRe = regexp.MustCompile(`^\/users[\/]*$`)
)

type userHandler struct{}

// ? creating a user
type 

// ?creating a data store
type datastore struct{
	m map[string]user
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// restricitng the path type it can handle it will only attach with json type of response
	w.Header().Set("content-type", "application/json")

	switch{
	case r.Method == http.MethodGet && listusersRe.MatchString(r.URL.Path):
		h.List(w,r)
		return
	case r.Method == http.MethodGet && getusersRe.MatchString(r.URL.Path):
		h.Get(w,r)
		return
	case r.Method == http.MethodPost && createusersRe.MatchString(r.URL.Path):
		h.Create(w,r)
		return 
	default :
		notFound(w,r)
		return
	}
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 page not found"))
}	

// ^ ENDPOINTS
// GET /users
// GET /users/{id}
// POST /users

func main() {
	fmt.Println("working with JWT token")
	mux := http.NewServeMux()
	mux.Handle("/users/", &userHandler{})
	http.ListenAndServe(":8080", mux)
}
