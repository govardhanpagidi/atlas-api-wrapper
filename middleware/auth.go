package middleware

import (
	"context"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"net/http"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header value
		publicKey, privateKey, ok := r.BasicAuth()

		// Parse the Basic authentication credentials
		_, _ = logger.Debugf("Request with pub token: %s", publicKey)
		if !ok {
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
