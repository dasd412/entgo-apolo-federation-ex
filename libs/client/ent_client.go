package client

import (
	"context"
	"entgo.io/ent"
)

type EntClient interface {
	Close() error
	Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error)
}

// todo *ent.Client에서 "user/pkg/ent" 의존성 제거 필요
//func NewEntClient(master *sql.Driver, replica *sql.Driver) *ent.Client {
//	if os.Getenv("SERVICE_ENV") == "local" {
//		return ent.NewClient(
//			ent.Driver(
//				&db.Drivers{
//					Master:  master,
//					Replica: replica,
//				},
//			),
//			ent.Debug(),
//			ent.Log(
//				func(argument ...any) {
//					start := time.Now()
//					duration := time.Since(start)
//					log.Printf("took: %v , entgo: %v ", duration, argument)
//				},
//			),
//		)
//	} else {
//		return ent.NewClient(
//			ent.Driver(&db.Drivers{Master: master, Replica: replica}),
//		)
//	}
//}
