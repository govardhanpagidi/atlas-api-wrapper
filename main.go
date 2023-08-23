package main

import (
	"fmt"
	"github.com/atlas-api-helper/handlers"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	api  = "/api"
	port = "8080"
)

func main() {
	configuration.GetInstance()
	// A router using gorilla/mux
	r := mux.NewRouter()
	r.Use(util.TraceIDMiddleware)
	apiRouter := r.PathPrefix(api).Subrouter()

	apiRouter.HandleFunc(uri(constants.GetClusterReqURI), handlers.GetCluster).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.CreateOrGetClusterReqURI), handlers.GetAllClusters).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.DeleteClusterReqURI), handlers.DeleteCluster).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.CreateOrGetClusterReqURI), handlers.CreateCluster).Methods(http.MethodPost)

	apiRouter.HandleFunc(uri(constants.DeleteOrGetDatabaseUserReqURI), handlers.GetDatabaseUser).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.CreateOrGetDatabaseUserReqURI), handlers.GetAllDatabaseUser).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.DeleteOrGetDatabaseUserReqURI), handlers.DeleteDatabaseUser).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.CreateOrGetDatabaseUserReqURI), handlers.CreateDatabaseUser).Methods(http.MethodPost)

	apiRouter.HandleFunc(uri(constants.CreateDatabaseReqURI), handlers.CreateDatabase).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.DeleteDatabaseReqURI), handlers.DeleteDatabase).Methods(http.MethodDelete)

	apiRouter.HandleFunc(uri(constants.CreateCollectionReqURI), handlers.CreateCollection).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.DeleteCollectionReqURI), handlers.DeleteCollection).Methods(http.MethodDelete)

	// Start the server on a given port
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func uri(handlerName string) string {
	return fmt.Sprintf("/%s", handlerName)
}
