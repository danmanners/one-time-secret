package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

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

	// Creating the struct
	secrets := make(map[string]string)

	// Offer up a pong response to check if things are responsive
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	//
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	// Create a secret from a data post
	r.Post("/secret", func(w http.ResponseWriter, r *http.Request) {
		s := uuid.New().String()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		fmt.Printf("Body: %s\nBodyLen: %d\n", body, len(body))

		if string(body) != "" {
			secrets[s] = string(body)
			fmt.Printf("Wrote '%s' to '%s'.\n", string(body), s)
			fmt.Fprintf(w, "%s\n", s)
		} else {
			fmt.Fprintf(w, "No data was sent\n")
		}

	})

	//Get Secret
	r.Get("/{secret}", func(w http.ResponseWriter, r *http.Request) {
		data := chi.URLParam(r, "secret")
		fmt.Printf("'%s' fetch was attempted.\n", data)
		s := secrets[data]
		if s != "" {
			w.Write([]byte(fmt.Sprintf("%s\n", s)))
		} else {
			w.Write([]byte(fmt.Sprintf("'%s' no longer or has never existed.\n", data)))
		}
		// Kill the secret
		delete(secrets, data)
	})

	listeningPort := functions.GetEnv("tapi_port", "3000")
	http.ListenAndServe(fmt.Sprintf(":%s", listeningPort), r)
}
