package socket

import (
	"fmt"
	"log"
	"net/http"
)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	UPGRADER.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := UPGRADER.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		conn: conn,
	}

	register <- client

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from client:", err)
			break
		}
		fmt.Println("Received message:", string(message))
		broadcast <- message
	}

	unregister <- client

}
