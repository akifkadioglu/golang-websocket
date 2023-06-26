package database

import (
	"context"
	"log"

	"github.com/akifkadioglu/golang-websocket/ent"
	"github.com/akifkadioglu/golang-websocket/env"
	_ "github.com/lib/pq"
)

func (d PostgreSQL) main() {
	var dns = env.Getenv(env.DB_EXTERNAL_URL)

	client, err = ent.Open("postgres", dns)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
