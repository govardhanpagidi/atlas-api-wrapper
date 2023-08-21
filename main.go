package main

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/handlers"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	api  = "/api"
	port = "8080"
)

func main() {
	configuration.GetInstance()
	// A router using gorilla/mux
	r := mux.NewRouter()
	r.Use(TraceIDMiddleware)
	apiRouter := r.PathPrefix(api).Subrouter()

	apiRouter.HandleFunc(uri(constants.ClusterStatusReqURI), handlers.GetCluster).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.ClusterReqURI), handlers.GetAllClusters).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.ClusterWithClusterNameReqURI), handlers.DeleteCluster).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.ClusterReqURI), handlers.CreateCluster).Methods(http.MethodPost)

	apiRouter.HandleFunc(uri(constants.DatabaseUserWithUsernameReqURI), handlers.GetDatabaseUser).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.DatabaseUserReqURI), handlers.GetAllDatabaseUser).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.DatabaseUserWithUsernameReqURI), handlers.DeleteDatabaseUser).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.DatabaseUserReqURI), handlers.CreateDatabaseUser).Methods(http.MethodPost)

	apiRouter.HandleFunc(uri(constants.DatabaseReqURI), handlers.CreateDatabase).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.DatabaseDeleteReqURI), handlers.DeleteDatabase).Methods(http.MethodDelete)

	apiRouter.HandleFunc(uri(constants.CollectionReqURI), handlers.CreateCollection).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.CollectionDeleteReqURI), handlers.DeleteCollection).Methods(http.MethodDelete)

	// Start the server on a given port
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func uri(handlerName string) string {
	return fmt.Sprintf("/%s", handlerName)
}

func TraceIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generate a trace ID
		traceID := fmt.Sprintf("TraceID-%d", time.Now().UnixNano())

		// Add the trace ID to the request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, constants.TraceID, traceID)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
