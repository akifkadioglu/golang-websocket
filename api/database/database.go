package database

import (
	"github.com/akifkadioglu/golang-websocket/ent"
)

type Idatabase interface {
	main()
}

var client *ent.Client
var err error

type PostgreSQL struct{}

func Connection() {
	var database PostgreSQL
	Idatabase.main(database)
}
func DBManager() *ent.Client {
	return client
}
