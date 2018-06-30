package main

import (
	"github.com/s-ichikawa/gql-todo/graph"
	"net/http"
	"github.com/vektah/gqlgen/handler"
	"fmt"
	"log"
)

func main() {
	app := &graph.MyApp{}
	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", handler.GraphQL(graph.MakeExecutableSchema(app)))

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
