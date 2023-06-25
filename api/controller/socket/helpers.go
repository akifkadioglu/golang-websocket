package socket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}
type Registration struct {
	Client *Client
	Group  string
}

var UPGRADER = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Group string
	Data  []byte
}

var (
	groups     = make(map[string]map[*Client]bool)
	broadcast  = make(chan Message)
	register   = make(chan Registration)
	unregister = make(chan *Client)
)

func Run() {
	for {
		select {
		case registration := <-register:
			group := registration.Group
			client := registration.Client

			if _, ok := groups[group]; !ok {
				groups[group] = make(map[*Client]bool)
			}

			groups[group][client] = true

		case client := <-unregister:
			for group, clients := range groups {
				if _, ok := clients[client]; ok {
					delete(clients, client)
					if len(clients) == 0 {
						delete(groups, group)
					}
					break
				}
			}

		case message := <-broadcast:
			group := message.Group
			data := message.Data

			if clients, ok := groups[group]; ok {
				for client := range clients {
					err := client.conn.WriteMessage(websocket.TextMessage, data)
					if err != nil {
						log.Printf("Error sending message to client: %v", err)
						unregister <- client
					}
				}
			}
		}
	}
}
