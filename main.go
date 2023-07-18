package main

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/handlers"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	api = "/api"
)

func main() {

	// Create a new router using gorilla/mux
	r := mux.NewRouter()
	apiRouter := r.PathPrefix(api).Subrouter()
	apiRouter.Use(DigestAuth)

	// Defining the REST API endpoints and their corresponding project handlers
	apiRouter.HandleFunc(uri(constants.ProjectHandler), handlers.CreateProjectHandler).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.ProjectHandler), handlers.GetProjectHandler).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.ProjectHandler), handlers.DeleteProjectHandler).Methods(http.MethodDelete)

	// Start the server on port 8080
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func uri(path string) string {
	return fmt.Sprintf("%s", path)
}

func DigestAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header value
		//authHeader := r.Header.Get("Authorization")
		publicKey, privateKey, ok := r.BasicAuth()
		// Parse the Digest authentication credentials
		//publicKey, privateKey, ok := parseDigestCredentials(authHeader)
		fmt.Printf("Request with pub token: %s", publicKey)
		if !ok {
			w.Header().Set("WWW-Authenticate", `Digest realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		reqContext := context.WithValue(r.Context(), constants.PubKey, publicKey)
		reqContext = context.WithValue(reqContext, constants.PvtKey, privateKey)
		// Authentication successful, call the next handler
		r = r.WithContext(reqContext)
		next.ServeHTTP(w, r)
	})
}
