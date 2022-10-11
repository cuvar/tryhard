package main

import (
	"graphql/srv/graph"
	"graphql/srv/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)


	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
						// Check against your desired domains here
						return r.Host == "http://localhost:8080"
				},
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
		},
})

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
	log.Fatal(http.ListenAndServe(":"+port, router))

	// err := http.ListenAndServe(":8080", router)
	// if err != nil {
	// 	panic(err)
	// }
}
