package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket" // inbuilt golang package
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{conns: make(map[*websocket.Conn]bool)}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from", ws.RemoteAddr())
	// to keep the connection alive
	// we need to use mutex to avoid race condition
	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := ws.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// EOF means the connection is closed form the other side
				break
			}
			fmt.Println("read error", err)
			// if we use return then we will actually break the connection
			continue
		}
		// storin the msg
		msg := buffer[:n]
		fmt.Println(string(msg))
		ws.Write([]byte("thank you for the message!!"))

		// // boradcastrimfg messag
		// s.boradCast(msg)
	}
}

// for broadcasting messages
func (s *Server) boradCast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("write error", err)
			}
		}(ws)
	}
}

func (s *Server) orderBook(ws *websocket.Conn) {
	fmt.Println("new incoming request", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("order book -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep((time.Second * 2))
	}

}

func main() {
	fmt.Println("Working with Web Sockets")

	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/order" , websocket.Handler(server.orderBook))
	http.ListenAndServe(":8000", nil)
}
