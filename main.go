package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/atlas-api-helper/docs"
	"github.com/atlas-api-helper/handlers"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	api       = "/api"
	port      = "8080"
	emptyPath = "/"
)

// @title MMC Atlas API Helper
// @version 1.0
// @description MMC Atlas API Helper
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
func main() {
	configuration.GetInstance()
	// A router using gorilla/mux
	r := mux.NewRouter()
	r.Use(util.TraceIDMiddleware)
	apiRouter := r.PathPrefix(api).Subrouter()

	// Serve the Swagger documentation at /docs/swagger.json
	r.PathPrefix(constants.SwaggerDocsPath).Handler(httpSwagger.Handler(
		httpSwagger.URL(constants.SwaggerDocsJsonPath), // The URL pointing to API definition
	))
	apiRouter.HandleFunc(uri(constants.CreateBackupSnapshotReqURI), handlers.CreateBackupSnapshot).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.GetAllBackupSnapshotReqURI), handlers.GetAllBackupSnapshot).Methods(http.MethodGet)

	apiRouter.HandleFunc(uri(constants.CreateRestoreClusterReqURI), handlers.CreateRestoreJob).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.GetRestoreJobReqURI), handlers.GetRestoreJob).Methods(http.MethodGet)

	apiRouter.HandleFunc(uri(constants.CreateBackupSnapshotReqURI), handlers.CreateBackupSnapshot).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.GetAllBackupSnapshotReqURI), handlers.GetAllBackupSnapshot).Methods(http.MethodGet)

	apiRouter.HandleFunc(uri(constants.CreateRestoreClusterReqURI), handlers.CreateRestoreJob).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.GetRestoreJobReqURI), handlers.GetRestoreJob).Methods(http.MethodGet)

	apiRouter.HandleFunc(uri(constants.GetClusterReqURI), handlers.GetCluster).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.CreateOrGetClusterReqURI), handlers.GetAllClusters).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.DeleteClusterReqURI), handlers.DeleteCluster).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.CreateOrGetClusterReqURI), handlers.CreateCluster).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.CreateOrGetClusterReqURI), handlers.UpdateCluster).Methods(http.MethodPatch)

	apiRouter.HandleFunc(uri(constants.DeleteOrGetDatabaseUserReqURI), handlers.GetDatabaseUser).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.CreateOrGetDatabaseUserReqURI), handlers.GetAllDatabaseUser).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.DeleteOrGetDatabaseUserReqURI), handlers.DeleteDatabaseUser).Methods(http.MethodDelete)
	apiRouter.HandleFunc(uri(constants.CreateOrGetDatabaseUserReqURI), handlers.CreateDatabaseUser).Methods(http.MethodPost)
	apiRouter.HandleFunc(uri(constants.DeleteOrGetDatabaseUserReqURI), handlers.UpdateDatabaseUser).Methods(http.MethodPatch)

	apiRouter.HandleFunc(uri(constants.CloudBackupScheduleReqURI), handlers.GetCloudBackupSchedule).Methods(http.MethodGet)
	apiRouter.HandleFunc(uri(constants.CloudBackupScheduleReqURI), handlers.UpdateClusterBackupPolicy).Methods(http.MethodPatch)

	subRouter := apiRouter.PathPrefix(emptyPath).Subrouter()
	subRouter.Use(util.BasicAuth)
	subRouter.HandleFunc(uri(constants.CreateDatabaseReqURI), handlers.CreateDatabase).Methods(http.MethodPost)
	subRouter.HandleFunc(uri(constants.GetDatabaseReqURI), handlers.ReadAllDatabase).Methods(http.MethodGet)
	subRouter.HandleFunc(uri(constants.DeleteDatabaseReqURI), handlers.DeleteDatabase).Methods(http.MethodDelete)

	subRouter.HandleFunc(uri(constants.CreateCollectionReqURI), handlers.CreateCollection).Methods(http.MethodPost)
	subRouter.HandleFunc(uri(constants.GetCollectionReqURI), handlers.ListCollection).Methods(http.MethodGet)
	subRouter.HandleFunc(uri(constants.DeleteCollectionReqURI), handlers.DeleteCollection).Methods(http.MethodDelete)

	// Start the server on a given port
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func uri(handlerName string) string {
	return fmt.Sprintf("/%s", handlerName)
}
