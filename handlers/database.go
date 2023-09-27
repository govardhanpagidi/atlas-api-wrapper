package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/atlas-api-helper/resources/database"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/constants"
	responseHandler "github.com/atlas-api-helper/util/responsehandler"
	"github.com/gorilla/mux"
)

// setupDatabaseLog sets up the logger for the database API handlers
func setupDatabaseLog() {
	util.SetupLogger("atlas-api-helper.handlers.database")
}

// CreateDatabase handles POST calls to create a new database using the provided parameters
// @Summary Create a new database
// @Description Create a new database with the specified name and owner
// @ID CreateDatabase
// @Tags Database
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param Authorization header string true "authorization" default()
// @Param ProjectId path string true "ProjectId" default(projectId)
// @Param ClusterName path string true "ClusterName" default(clusterId)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param InputModel body database.InputModel true "body"
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/cluster/{ClusterName}/database [post]
func CreateDatabase(w http.ResponseWriter, r *http.Request) {
	setupDatabaseLog()

	//fetch all input parameters and create input model
	var model database.InputModel
	vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	projectId := vars[constants.ProjectID]
	clusterName := vars[constants.ClusterName]
	username := r.Context().Value(constants.Username).(string)
	password := r.Context().Value(constants.Password).(string)
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

	model.Username = &username
	model.Password = &password
	model.ProjectId = &projectId
	model.ClusterName = &clusterName
	model.PrivateKey = &privateKey
	model.PublicKey = &publicKey

	//log the input model
	util.Debugf(r.Context(), "Create database Request : %+v", model.String())
	startTime := time.Now()

	//make API call to create a database
	response := database.Create(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create database REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// ReadAllDatabase handles GET calls to list all databases using the provided parameters
// @Summary lists all database
// @Description read all the databases using hostname,username and password
// @ID ReadAllDatabase
// @Tags Database
// @Produce json
// @Param Authorization header string true "authorization" default()
// @Param ProjectId path string true "ProjectId" default(projectId)
// @Param ClusterName path string true "ClusterName" default(clusterId)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Security BasicAuth
// @Success 200 {object} atlasresponse.AtlasResponse
// @Failure 400 {object} atlasresponse.AtlasResponse
// @Failure 401 {object} atlasresponse.AtlasResponse
// @Failure 403 {object} atlasresponse.AtlasResponse
// @Failure 404 {object} atlasresponse.AtlasResponse
// @Failure 500 {object} atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/cluster/{ClusterName}/database [get]
func ReadAllDatabase(w http.ResponseWriter, r *http.Request) {
	setupDatabaseLog()

	vars := mux.Vars(r)
	//fetch all input parameters and create input model
	projectId := vars[constants.ProjectID]
	clusterName := vars[constants.ClusterName]
	username := r.Context().Value(constants.Username).(string)
	password := r.Context().Value(constants.Password).(string)
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

	//create input model for delete database API
	model := database.InputModel{
		Username:    &username,
		Password:    &password,
		ProjectId:   &projectId,
		ClusterName: &clusterName,
		PrivateKey:  &privateKey,
		PublicKey:   &publicKey,
	}

	//log the input model
	util.Debugf(r.Context(), "ReadAll database Request : %+v", model.String())
	startTime := time.Now()

	//make API call to create a database
	response := database.ReadAll(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Read all databases REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// DeleteDatabase handles the DELETE calls to delete the requested database
// @Summary Delete a database
// @Description Delete a database with the specified ID
// @ID DeleteDatabase
// @Tags Database
// @Accept json
// @Produce json
// @Param DatabaseName path string true "databaseName" default(testDatabase)
// @Security BasicAuth
// @Param ProjectId path string true "ProjectId" default(projectId)
// @Param ClusterName path string true "ClusterName" default(clusterId)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param Authorization header string true "authorization" default()
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/cluster/{ClusterName}/database/{DatabaseName} [delete]
func DeleteDatabase(w http.ResponseWriter, r *http.Request) {
	setupDatabaseLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseName]
	username := r.Context().Value(constants.Username).(string)
	password := r.Context().Value(constants.Password).(string)
	projectId := vars[constants.ProjectID]
	clusterName := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

	//create input model for delete database API
	model := database.InputModel{
		DatabaseName: &databaseName,
		Username:     &username,
		Password:     &password,
		ProjectId:    &projectId,
		ClusterName:  &clusterName,
		PublicKey:    &publicKey,
		PrivateKey:   &privateKey,
	}

	//log the input model
	util.Debugf(r.Context(), "Delete database Request : %+v", model.String())
	startTime := time.Now()

	//make API call to delete a database
	response := database.Delete(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Delete database REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterHandler)
}
