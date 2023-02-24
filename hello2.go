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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, home!")
}

func reader(conn *websocket.Conn) {

	for {
		messageType, mess, err := conn.ReadMessage()
		if err != nil {

			fmt.Println(err)

		}

		fmt.Println(string(mess))

		if err := conn.WriteMessage(messageType, mess); err != nil {

			log.Println(err)
			return
		}

	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Hello, endpoint!")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	reader(ws)
}

func setupRoutes() {

	http.HandleFunc("/", home)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {

	fmt.Println("welcome to websockets")
	setupRoutes()
	http.ListenAndServe(":8080", nil)

}






http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	isOpen := true
	if err != nil {
		fmt.Println(err)
		return
	}

	connCount++

	if connCount > 2 {
		fmt.Println("Conncection exceeds 2")
		connCount--
		conn.Close()
		return
	}
	fmt.Println(connCount)
	clients[conn] = true
	fmt.Println(clients)

	for isOpen {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			isOpen = false
			connCount--
			delete(clients, conn)
			fmt.Println("Error from main routine", err)

			break
		}
		fmt.Println("Message sent:", string(msg))
		broadcast <- msg
		connection <- conn
	}
})