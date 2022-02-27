package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	"github.com/danmanners/go-learning/functions"
	"github.com/danmanners/go-learning/routes"
)

// Main Function
func main() {
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

	//Get Secret
	r.Get("/{secret}", routes.GetSecret)

	listeningPort := functions.GetEnv("tapi_port", "3000")
	http.ListenAndServe(fmt.Sprintf(":%s", listeningPort), r)
}
