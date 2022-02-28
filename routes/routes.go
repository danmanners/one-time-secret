package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"

	"github.com/danmanners/one-time-secret/functions"
	variables "github.com/danmanners/one-time-secret/vars"
	"github.com/google/uuid"
)

// Basic ping/pong health check endpoint
func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string("pong")))
}

// Attempt to create a secret
func CreateSecret(w http.ResponseWriter, r *http.Request) {
	s := uuid.New().String()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	// If the string body is not empty, create it and add it to the secrets map.
	if string(body) != "" {
		key := functions.CheckForAesKey()
		variables.Secrets[s], _ = functions.Encrypt(
			string(body),
		)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("%s\n", s)))
	} else {
		// If the string body is empty, return that nothing was sent.
		fmt.Fprintf(w, "No data was sent\n")
	}
}

// Attempt to retrieve a secret
func GetSecret(w http.ResponseWriter, r *http.Request) {
	// Get the plaintext secret from the datastring
	data := chi.URLParam(r, "secret")

	//
	s := variables.Secrets[data]
	if s != "" {
		if u, err := url.ParseRequestURI(s); err == nil {

			// If the secret is a URL, redirect to that endpoint.
			http.Redirect(w, r, u.String(), http.StatusTemporaryRedirect)
		} else {

			// If the secret is NOT a URL, return it to the request.
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%s\n", s)))
		}

	} else {
		// If no secret is found, crash out and respond appropriately.
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("'%s' no longer or has never existed.\n", data)))
	}

	// Kill the secret
	delete(variables.Secrets, data)
}

// Catch all path function
func Catchall(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(string("Nothing at this path.\n")))
}
