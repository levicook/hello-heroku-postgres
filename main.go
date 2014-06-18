package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	port     string
	database *sql.DB
)

func init() {
	port = os.Getenv("PORT")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	panicIf(err)
	database = db
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var pgPid int
		panicIf(database.QueryRow("SELECT pg_backend_pid()").Scan(&pgPid))
		fmt.Fprintf(w, "pg_backend_pid=%v", pgPid)
	})

	panicIf(http.ListenAndServe(":"+port, nil))
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
