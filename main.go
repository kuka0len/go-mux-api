package main

import (
	"net/http"

	"github.com/aaznadi/go-mux-api/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
