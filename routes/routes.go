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

func CreateSecret(w http.ResponseWriter, r *http.Request) {
	s := uuid.New().String()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	if string(body) != "" {
		variables.Secrets[s] = string(body)
		fmt.Printf("Source IP '%s' created '%s'.\n",
			net.ParseIP(strings.Split(r.RemoteAddr, ":")[0]),
			s,
		)
		fmt.Fprintf(w, "%s\n", s)
	} else {
		fmt.Fprintf(w, "No data was sent\n")
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func GetSecret(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "secret")
	s := variables.Secrets[data]
	if s != "" {
		if u, err := url.ParseRequestURI(s); err == nil {
			// If the secret is a URL, redirect to that mo'fo.

			http.Redirect(w, r, u.String(), http.StatusTemporaryRedirect)
		} else {
			// If the secret is _NOT_ a URL, return it to the request.
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%s\n", s)))
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("'%s' no longer or has never existed.\n", data)))
	}
	// Kill the secret
	delete(variables.Secrets, data)
}
