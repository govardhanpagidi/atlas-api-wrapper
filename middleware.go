package main

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/util/constants"
	"net/http"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header value
		publicKey, privateKey, ok := r.BasicAuth()

		// Parse the Basic authentication credentials
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
