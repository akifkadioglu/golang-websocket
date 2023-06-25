package socket

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	group := chi.URLParam(r, "name")
	
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

	registration := Registration{
		Client: client,
		Group:  group,
	}

	register <- registration

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from client:", err)
			break
		}
		msg := Message{
			Group: group,
			Data:  message,
		}

		broadcast <- msg
	}

	unregister <- client
}
