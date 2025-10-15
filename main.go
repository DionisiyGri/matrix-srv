package main

import (
	"league-matrix/handlers"
	"league-matrix/internal/matrix"
	"league-matrix/parser"
	"log"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

const (
	defaulPort = ":8080"
)

func main() {
	h := handlers.New(parser.New(), matrix.NewCSVMatrixer())
	http.HandleFunc("/echo", h.Echo)
	http.HandleFunc("/invert", h.Invert)

	log.Printf("starting server on %s", defaulPort)
	http.ListenAndServe(defaulPort, nil)
}
