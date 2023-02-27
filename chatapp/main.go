package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Pages")
}

// a function that will continuously listen to incoming messages from the server.

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			isOpen =false
			connCount--
			log.Println("Error reading message:", err)
			return
		}
		// Broadcast message to all connected clients
		for client := range clients {
			if client == conn {
				continue
			}
			err = client.WriteMessage(messageType, p)
			if err != nil {
				// Handle error
				break
			}
		}

		if messageType != websocket.TextMessage {
			log.Println("Unexpected message type:", messageType)
			return
		}
		// print out that message for clarity
		fmt.Println("Message received:", string(p))

		// if err := conn.WriteMessage(websocket.TextMessage, p); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}
}

var clients = make(map[*websocket.Conn]bool)
var signedIn bool
var connCount int
var isOpen bool

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// ^ picking data from params
	token := r.URL.Query()["token"]
	tokenString := ""
	for _, v := range token {
		tokenString += v
	}


	//^ check if client is signed in or not
	cookie, err := r.Cookie("signedIn")
	fmt.Println("cookie: ",cookie)
    if err == nil && cookie.Value == "true" {
        signedIn = true
    }

    // If the user is not signed in, return 401 Unauthorized
    if !signedIn {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }
	
	//This will determine whether or not an incoming request from a different domain is allowed to connect,
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	//We can now start attempting to upgrade the incoming HTTP connection using the upgrader.Upgrade() function
	//which will take in the Response Writer and the pointer to the HTTP Request and return us with a pointer to a WebSocket connection,
	//or an error if it failed to upgrade.
	// ek new instance intialize kro Claims ka
	// lekin agr cookie ayi hai to sbse pehle JWT Token string utahynge
	
	// fmt.Println("claims", claims.Username)
	// uName = claims.Username

	// lekin agr cookie ayi hai to sbse pehle JWT Token string utahynge
	
	
	ws, err := upgrader.Upgrade(w, r, nil)
	isOpen = true
	if err != nil {
		log.Println(err)

	}
	connCount++
	if connCount > 2 {
		fmt.Println("Conncection exceeds 2")
		connCount--
		ws.Close()
		return
	}
	fmt.Println(connCount)
	fmt.Println("Client successfully connected")
	// claims := &Claims{}
	// fmt.Println("claims:", claims)
	clients[ws] = true
	// c, err := r.Cookie("token")

	// fmt.Println("cookie usernmae", c.Username)
	for isOpen{
		
		defer func() {
			delete(clients, ws)
			ws.Close()
		}()
		// whenever any client connected to our server it will recieve greeting messgae
		// ^ grabbing user name of the person signing in
	
		err = ws.WriteMessage(1, []byte("Hi " + uName))
		if err != nil {
			log.Println(err)
		}
		// listen indefinitely for new messages coming
		// through on our WebSocket connection
		reader(ws)
	}
}



//! START JWT Token

// create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// for the signin routes we need user password
var users = map[string]string{
	"onkar": "onkar@123",
	"aman":  "aman@123",
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
var uName string

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
	uName = creds.Username
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
	}else{
		signedIn = true
	}
	// *Finally, we set the client cookie for "token" as the JWT we just generated
	// *we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		
	})
	// sending the token requst
	w.Write([]byte(tokenString))

}

func Welcome(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
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
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		// ?agar match kr gya to return kro jwtKey jisem actual signature hai
		return jwtKey, nil
	})
	// fmt.Println("claims", claims.Username)
	// uName = claims.Username

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
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
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

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("excuting register")
    // Parse the JSON body of the request
    var data map[string]string
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Get the username and password from the JSON data
    username, ok := data["username"]
    if !ok {
        http.Error(w, "Missing username", http.StatusBadRequest)
        return
    }
    password, ok := data["password"]
    if !ok {
        http.Error(w, "Missing password", http.StatusBadRequest)
        return
    }

    // Store the username and password in the users map
    users[username] = password

    // Send a response indicating success
    fmt.Fprint(w, "User registered successfully")
}

// ! END Token

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}


func main() {
	fmt.Println("Go: WebSockets")
	setupRoutes()

	//^ setting up routes for  JWT client
	fmt.Println("Implementing authentication using JWT authentication")
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/register", RegisterHandler)


	log.Fatal(http.ListenAndServe(":8000", nil))

}
