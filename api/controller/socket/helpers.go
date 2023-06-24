package socket

import (
	"log"

	"github.com/gorilla/websocket"
)

func reader(conn *websocket.Conn) {
	for {
		message, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		if err := conn.WriteMessage(message, p); err != nil {
			log.Println(err)
			return
		}
	}
}
