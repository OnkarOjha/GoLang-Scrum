package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

// create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// for the signin routes we need user password
var users = map[string]string{
	"onkar": "onkar@123",
	"user2": "password2",
}

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}




// ^Create a struct that will be encoded to a JWT.
// ^We add jwt.RegisteredClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	expectedPassword, ok := users[creds.Username]

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("The token string is: ", tokenString)

	// *Finally, we set the client cookie for "token" as the JWT we just generated
	// *we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

func Welcome(w http.ResponseWriter, r *http.Request) {
	c , err := r.Cookie("token")
	if err!=nil{
		// aagr cookie not found ayi
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// agr koi or dikt ayi hai to kch aur show rknge
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// lekin agr cookie ayi hai to sbse pehle JWT Token string utahynge
	tknStr := c.Value

	// ek new instance intialize kro Claims ka 
	claims := &Claims{}

	// !ab parse kro JWT STRING ko and store the results in claims
	// !Note that we are passing the key  as well in this method
	// !if the token is invalid (if it has expired according to the expiry time given)
	// !then we will return an "Unauthorized" error
	tkn , err := jwt.ParseWithClaims(tknStr , claims , func(token *jwt.Token)(interface{} , error){
		// ?agar match kr gya to return kro jwtKey jisem actual signature hai
		fmt.Println("  mftgjyelkigy:",token)
		return jwtKey , nil
	})
	fmt.Println("token: ", tkn)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
// ^ now we will be creating a new funciton
// ^ this function will basically be a refresh function
// ^ jisme ki agar expiry time ek token ka khtm ho rha hai to to ye us token
// ^ ko lega and uska time fr s increse krega jo ki hume manually nhi krna pdega

func Refresh(w http.ResponseWriter, r *http.Request) {
	// ? START code same as that of welcome fucntion
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	fmt.Println("token: ", tkn)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// ? END code ends now we wilkl actually satt with the refreshing work


	// ^ jo new token hoga  wo turnt issue nhi hoga 
	// ^ we will be attaching a 30 sec time wihtin which a new token weilll be issued
	// ^ otherwise error

	if time.Until(claims.ExpiresAt.Time) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// ab 30 sec baad new token generate kro
	expiresTime := time.Now().Add(5 * time.Minute)
	// ab is 5 minute ko claims ka new time bnao
	claims.ExpiresAt = jwt.NewNumericDate(expiresTime)
	// ab jo new claim bna new expiry time ke sath uspe naya signature lgao
	// mtlb basically new expiry time ke sath new token bnao
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// encoding gho gy token ki
	// ab ek complete JWT token string bnao
	tokenString , err := token.SignedString(jwtKey)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Set the new token as the users `token` cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expiresTime,
	})


}

// ^ now we will implement the logout functionality
func Logout(w http.ResponseWriter, r *http.Request) {
	// immediately clear the token cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}

func main() {
	// we will implement these handlers in the next sections
	fmt.Println("Implementing authentication using JWT authentication")
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8080", nil))
}
