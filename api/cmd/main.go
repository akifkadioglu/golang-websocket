package main

import (
	"net/http"

	"github.com/akifkadioglu/golang-websocket/routes"
)

func main() {

	r := routes.CreateServer()
	http.ListenAndServe(":3000", r)
}
