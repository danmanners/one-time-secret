package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/danmanners/go-learning/functions"
)

// Main Function
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Offer up a pong response to check if things are responsive
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	//
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	listeningPort := functions.GetEnv("tapi_port", "3000")
	http.ListenAndServe(fmt.Sprintf(":%s", listeningPort), r)
}
