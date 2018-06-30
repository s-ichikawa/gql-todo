package main

import (
	"github.com/s-ichikawa/gql-todo/graph"
	"net/http"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@tcp(127.0.0.1:33306)/gql_todo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	server := graph.Server{
		DB: db,
	}
	server.Run()

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
