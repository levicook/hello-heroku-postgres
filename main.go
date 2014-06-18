package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/levicook/go-detect"
	_ "github.com/lib/pq"
)

var (
	port        string
	databaseUrl string
	database    *sql.DB
)

func init() {
	port = detect.String(os.Getenv("PORT"), "5001")
	databaseUrl = detect.String(os.Getenv("DATABASE_URL"), "") // todo is there a sensible default for this?

	db, err := sql.Open("postgres", databaseUrl)
	panicIf(err)
	database = db
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, env := range os.Environ() {
			fmt.Fprintln(w, env)
		}

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
