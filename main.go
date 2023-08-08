package main

import (
	"fmt"
	"github.com/atlas-api-helper/handlers"
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

	// A router using gorilla/mux
	r := mux.NewRouter()
	apiRouter := r.PathPrefix(api).Subrouter()

	apiRouter.HandleFunc(uri(constants.ClusterWithGroupIdAndName), handlers.GetCluster).Methods(http.MethodGet)
	//apiRouter.HandleFunc(uri(constants.ClusterWithGroupId), handlers.GetAllCluster).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.ClusterWithGroupIdAndName), handlers.DeleteCluster).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.Cluster), handlers.CreateCluster).Methods(http.MethodPost)
	//apiRouter.HandleFunc(uri(constants.Cluster), handlers.UpdateCluster).Methods(http.MethodPut)

	// Start the server on a given port
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func uri(handlerName string) string {
	return fmt.Sprintf("/%s", handlerName)
}
