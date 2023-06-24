package socket

import (
	"log"
	"net/http"

	"github.com/akifkadioglu/golang-websocket/utils"
)

func SocketHandler(w http.ResponseWriter, r *http.Request) {
	utils.UPGRADER.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := utils.UPGRADER.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	reader(conn)
}
