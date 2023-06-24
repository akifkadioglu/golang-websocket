package routes

import (
	"github.com/akifkadioglu/golang-websocket/controller/home"
	"github.com/akifkadioglu/golang-websocket/controller/socket"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateServer() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	go socket.Run()

	r.Get("/", home.Home)
	r.HandleFunc("/socket", socket.SocketHandler)
	return r
}
