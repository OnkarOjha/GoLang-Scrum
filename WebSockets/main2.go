package main

import (
	"fmt"
	"log"
	"net/http"

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

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	//This will determine whether or not an incoming request from a different domain is allowed to connect,
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	//We can now start attempting to upgrade the incoming HTTP connection using the upgrader.Upgrade() function
	//which will take in the Response Writer and the pointer to the HTTP Request and return us with a pointer to a WebSocket connection,
	//or an error if it failed to upgrade.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)

	}
	fmt.Println("Client successfully connected")
	clients[ws] = true
	defer func() {
		delete(clients, ws)
		ws.Close()
	}()
	// whenever any client connected to our server it will recieve greeting messgae
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Go: WebSockets")
	setupRoutes()

	log.Fatal(http.ListenAndServe(":8000", nil))

}
