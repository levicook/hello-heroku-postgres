package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/levicook/go-detect"
)

var port string

func init() {
	port = detect.String(os.Getenv("PORT"), "5001")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, env := range os.Environ() {
			fmt.Fprintln(w, env)
		}
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
