package main

import (
	"log"
	"net/http"

	"github.com/akifkadioglu/golang-websocket/database"
	"github.com/akifkadioglu/golang-websocket/env"
	"github.com/akifkadioglu/golang-websocket/routes"
	"github.com/akifkadioglu/golang-websocket/utils"
)

func main() {
	log.Println("*********************** STARTED ***********************")

	env.InitEnv(env.Prod)
	database.Connection()
	utils.InitToken()
	r := routes.CreateServer()
	http.ListenAndServe(":10000", r)
}
