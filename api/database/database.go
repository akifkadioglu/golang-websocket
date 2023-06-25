package database

import (
	"github.com/akifkadioglu/golang-websocket/ent"
)

type Idatabase interface {
	main()
}

var client *ent.Client
var err error

type MySQL struct{}
type SQLite struct{}
type PostgreSQL struct{}

func Init() {

	var database MySQL
	Idatabase.main(database)
}
func DBManager() *ent.Client {
	return client
}
