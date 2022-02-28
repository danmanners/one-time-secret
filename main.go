package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/danmanners/one-time-secret/functions"
	"github.com/danmanners/one-time-secret/routes"
)

// Main Function
func main() {
	// Check if an AES string was provided or not

	// Set up the Chi Router
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.URLFormat,
		middleware.RealIP,
		middleware.URLFormat,
		render.SetContentType(render.ContentTypeJSON),
	)

	// Offer up a pong response to check if things are responsive
	r.Get("/ping", routes.Ping)

	// Create a secret from a data post
	r.Post("/secret", routes.CreateSecret)

	// Get Secret
	r.Get("/{secret}", routes.GetSecret)

	// Catchall
	r.Get("/", routes.Catchall)

	// Listening Port
	listeningPort := functions.GetEnv("ots_port", "3000")
	http.ListenAndServe(fmt.Sprintf(":%s", listeningPort), r)
}
