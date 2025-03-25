package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

var ctx context.Context
var err error

func main() {
	ctx = context.Background()
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {

	})

	if err = http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err.Error())
	}
}
