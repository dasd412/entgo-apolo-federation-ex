package main

import (
	"context"
	"entgo.io/contrib/entgql"
	"errorwrapper"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"order/pkg/ent"
	"order/pkg/graph/resolver"
	"strconv"
	"time"
)

func main() {
	ctx := context.Background()

	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
		ent.Debug(),
		ent.Log(
			func(argument ...any) {
				start := time.Now()
				duration := time.Since(start)
				log.Printf("took: %v , entgo: %v ", duration, argument)
			},
		),
	)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	port := "8082"

	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	server := handler.New(resolver.NewSchema(client))
	server.AddTransport(transport.Websocket{}) // Add WebSocket first. Here there is no config, see below for examples.
	server.AddTransport(transport.Options{})   // If you are using the playground, it's smart to add Options and GET.
	server.AddTransport(transport.GET{})       // ...
	server.AddTransport(transport.POST{})      // ... Make sure this is after the WebSocket transport!

	server.Use(entgql.Transactioner{TxOpener: client})
	server.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		return errorwrapper.WrapError(ctx, err)
	})

	corsWrapper := cors.AllowAll().Handler

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", corsWrapper(server))

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
