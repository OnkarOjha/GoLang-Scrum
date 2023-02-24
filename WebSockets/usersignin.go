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

var signedInClients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/ws", wsHandler)
    go handleMessages()

    log.Println("Server started on localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the user is signed in
    signedIn := true
    cookie, err := r.Cookie("signedIn")
    if err == nil && cookie.Value == "true" {
        signedIn = true
    }

    // If the user is not signed in, return 401 Unauthorized
    if !signedIn {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Upgrade the HTTP connection to a WebSocket connection
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading connection:", err)
        return
    }

    // Add the signed-in client to the list of clients
    signedInClients[conn] = true

    // Send a welcome message to the client
    conn.WriteMessage(websocket.TextMessage, []byte("Welcome to the chat room!"))

    // Listen for messages from the client
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            delete(signedInClients, conn)
            return
        }
        message := fmt.Sprintf("%s", p)
        log.Println("Received message:", message)

        // Broadcast the message to all signed-in clients
        broadcast <- []byte(message)
        if messageType == websocket.CloseMessage {
            delete(signedInClients, conn)
            return
        }
    }
}

func handleMessages() {
    for {
        // Get the next message from the broadcast channel
        message := <-broadcast

        // Send the message to all signed-in clients
        for client := range signedInClients {
            err := client.WriteMessage(websocket.TextMessage, message)
            if err != nil {
                log.Println("Error sending message to client:", err)
                client.Close()
                delete(signedInClients, client)
            }
        }
    }
}
