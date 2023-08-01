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

	apiRouter.HandleFunc(uri(constants.DatabaseuserGetHandler), handlers.GetDatabaseUser).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.DatabaseuserWithGroupId), handlers.GetAllDatabaseUser).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.DatabaseuserGetHandler), handlers.DeleteDatabaseUser).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.DatabaseuserHandler), handlers.CreateDatabaseUser).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.DatabaseuserHandler), handlers.UpdateDatabaseUser).Methods(http.MethodPut)

	apiRouter.HandleFunc(uri(constants.CustomDbRoleGetHandler), handlers.GetCustomDbRole).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.CustomDbRoleWithGroupId), handlers.GetAllCustomDbRoles).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.CustomDbRoleGetHandler), handlers.DeleteCustomDbRoles).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.CustomDbRole), handlers.CreateCustomDbRole).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.CustomDbRole), handlers.UpdateCustomDbRole).Methods(http.MethodPut)

	apiRouter.HandleFunc(uri(constants.ProjectInviteWithGroupIDAndInviteId), handlers.GetProjectInvitation).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.ProjectInviteWithGroupId), handlers.GetAllprojectInvites).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.ProjectInviteWithGroupIDAndInviteId), handlers.DeleteProjectInvites).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.ProjectInvite), handlers.CreateProjectInvite).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.ProjectInvite), handlers.UpdateProjectInvite).Methods(http.MethodPut)

	apiRouter.HandleFunc(uri(constants.ClusterWithGroupIdAndName), handlers.GetCluster).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.ClusterWithGroupId), handlers.GetAllCluster).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.ClusterWithGroupIdAndName), handlers.DeleteCluster).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.Cluster), handlers.CreateCluster).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.Cluster), handlers.UpdateCluster).Methods(http.MethodPut)

	// Start the server on a given port
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func uri(handlerName string) string {
	return fmt.Sprintf("/%s", handlerName)
}
