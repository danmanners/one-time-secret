package routes

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/chi/v5"

	variables "github.com/danmanners/go-learning/vars"
	"github.com/google/uuid"
)

// Basic ping/pong health check endpoint
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
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
		variables.Secrets[s] = string(body)
		fmt.Printf("Source IP '%s' created '%s'.\n",
			net.ParseIP(strings.Split(r.RemoteAddr, ":")[0]),
			s,
		)
		fmt.Fprintf(w, "%s\n", s)
	} else {
		// If the string body is empty, return that nothing was sent.
		fmt.Fprintf(w, "No data was sent\n")
	}
}

// Attempt to retrieve a secret
func GetSecret(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "secret")
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
