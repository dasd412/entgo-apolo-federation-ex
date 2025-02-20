package client

import (
	"db"
	"entgo.io/ent/dialect/sql"
	"log"
	"os"
	"time"
	"user/pkg/ent"
)

func NewEntClient(master *sql.Driver, replica *sql.Driver) *ent.Client {
	if os.Getenv("SERVICE_ENV") == "local" {
		return ent.NewClient(
			ent.Driver(
				&db.Drivers{
					Master:  master,
					Replica: replica,
				},
			),
			ent.Debug(),
			ent.Log(
				func(argument ...any) {
					start := time.Now()
					duration := time.Since(start)
					log.Printf("took: %v , entgo: %v ", duration, argument)
				},
			),
		)
	} else {
		return ent.NewClient(
			ent.Driver(&db.Drivers{Master: master, Replica: replica}),
		)
	}
}
