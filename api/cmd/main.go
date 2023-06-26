package main

import (
	"net/http"

	"github.com/akifkadioglu/golang-websocket/database"
	"github.com/akifkadioglu/golang-websocket/env"
	"github.com/akifkadioglu/golang-websocket/routes"
	"github.com/akifkadioglu/golang-websocket/utils"
)

func main() {

	env.InitEnv(env.Prod)
	database.Connection()
	utils.InitToken()
	r := routes.CreateServer()
	http.ListenAndServe("172.168.10.243:3000", r)
}
