package main

import (
	"net/http"

	"github.com/akifkadioglu/golang-websocket/database"
	"github.com/akifkadioglu/golang-websocket/env"
	"github.com/akifkadioglu/golang-websocket/routes"
	"github.com/akifkadioglu/golang-websocket/utils"
)

func main() {

	env.InitEnv(env.Local)
	database.Connection()
	utils.InitToken()
	r := routes.CreateServer()
	http.ListenAndServe(":10000", r)
}