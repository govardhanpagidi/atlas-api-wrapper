package main

import (
	"fmt"
	"github.com/atlas-api-helper/handlers"
	"github.com/atlas-api-helper/middleware"
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
	apiRouter.Use(middleware.BasicAuth)

	// REST API endpoints and their corresponding handlers
	apiRouter.HandleFunc(uri(constants.ProjectHandler), handlers.CreateProject).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.ProjectHandlerWithId), handlers.GetProject).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.ProjectHandler), handlers.DeleteProject).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.ProjectHandler), handlers.UpdateProject).Methods(http.MethodPatch)
	apiRouter.HandleFunc(uri(constants.ProjectHandler), handlers.GetAllProjects).Methods(http.MethodGet)
	// Start the server on a given port
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func uri(handlerName string) string {
	return fmt.Sprintf("/%s", handlerName)
}
