package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

type application struct {
	auth struct {
		username string
		password string
	}
}

func main() {
	fmt.Println("workin with authentication")

	app := new(application)

	app.auth.username = os.Getenv("AUTH_USERNAME")
	app.auth.password = os.Getenv("AUTH_PASSWORD")

	if app.auth.username == "" {
		fmt.Println("basic auth username must be provided")
	}
	if app.auth.password == "" {
		fmt.Println("please provide a password")
	}

	//registering a new mux server
	mux := http.NewServeMux()
	mux.HandleFunc("/unprotected", app.unprotectedHandler)
	mux.Handle("/protected", app.basicAuth(http.HandlerFunc(app.protectedHandler)))

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting server on %s", srv.Addr)
	// err := srv.ListenAndServeTLS("./localhost.pem", "./localhost-key.pem")
	// log.Fatal(err)
	log.Fatal(http.ListenAndServe(":8000", mux))

}

// two handler funciton for protected and unprotected endpoints
func (app *application) protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a protected handler")
}

func (app *application) unprotectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a unprotected handler")
}

// implemtning the middleware
func (app *application) basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth() //returns the username and passeword provided by the user
		fmt.Println(string(username))
		fmt.Println(string(password))
		//TODO yahi pr password check krma hai ki sahi hai ya glt

		if userNameCheck(username)|| ok  || passwordCheck(password){
			// now we got the username and password now hash it
			usernamehash := sha256.Sum256([]byte(username))

			passwordhash := sha256.Sum256([]byte(password))

			expectedUsernamehash := sha256.Sum256([]byte(app.auth.username))
			expectedPasswordhash := sha256.Sum256([]byte(app.auth.password))
			// now after taking the both the entered and expected hash vaues of username and password we will match it with each other
			usernameMatch := (subtle.ConstantTimeCompare(usernamehash[:], expectedUsernamehash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordhash[:], expectedPasswordhash[:]) == 1)

			// now according to compared vlaues start the sevrer and pass value to  it
			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted" , charset="UTF-8"`)
		http.Error(w, "Invalid", http.StatusUnauthorized)
	})
}

func userNameCheck(username string) bool {
	if len(username) == 0 {
		
		return false
	}
	if !strings.ContainsAny(username, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}
	if username == "Username" || username == "username" {
		return false
	}

	return true
}

func passwordCheck(password string) bool {

	//1. contains atleast one upper case letter
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}
	// 2. contains atleast one lower case letter
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}
	// 3. contains atleast one number
	match, err := regexp.MatchString("\\d", password)
	if err != nil {
		return false
	}
	if !match {
		return  false
	}
	// 4. contains atleast one special character
	special, _ := regexp.MatchString(`[\p{P}\p{S}\p{Z}]`, password)
	if !special {
		return false

	}
	// 8. should be between 8 to 20 characters
	if utf8.RuneCountInString(password) < 8 {
		return false
	}

	return true
}