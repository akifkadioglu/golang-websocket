package socket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

var UPGRADER = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	clients    = make(map[*Client]bool)
	broadcast  = make(chan []byte)
	register   = make(chan *Client)
	unregister = make(chan *Client)
)

func Run() {
	for {
		select {
		case client := <-register:
			clients[client] = true
		case client := <-unregister:
			client.mu.Lock()
			delete(clients, client)
			client.mu.Unlock()

		case message := <-broadcast:
			for client := range clients {
				client.mu.Lock()
				err := client.conn.WriteMessage(websocket.TextMessage, message)
				client.mu.Unlock()

				if err != nil {
					log.Printf("Error sending message to client: %v", err)
					unregister <- client
				}
			}
		}
	}
}
