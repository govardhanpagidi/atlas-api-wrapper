package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/atlas-api-helper/resources/collection"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/constants"
	responseHandler "github.com/atlas-api-helper/util/responsehandler"
	"github.com/gorilla/mux"
)

// setupCollectionLog sets up the logger for the collection API handlers
func setupCollectionLog() {
	util.SetupLogger("atlas-api-helper.handlers.collection")
}

// CreateCollection handles POST requests to create a new collection
// @Summary Create a new collection
// @Description Create a new collection with the specified name and description
// @ID CreateCollection
// @Tags Collection
// @Accept json
// @Produce json
// @Param InputModel body collection.InputModel true "body"
// @Param DatabaseName path string true "databaseName" default(testDatabase)
// @Param ProjectId path string true "ProjectId" default(projectId)
// @Param ClusterName path string true "ClusterName" default(clusterId)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Security BasicAuth
// @Param Authorization header string true "authorization" default()
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @securityDefinitions.basic BasicAuth
// @Router /project/{ProjectId}/cluster/{ClusterName}/database/{DatabaseName}/collections [post]
func CreateCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()

	//fetch all input parameters and create input model
	var model collection.InputModel
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseName]
	username := r.Context().Value(constants.Username).(string)
	password := r.Context().Value(constants.Password).(string)
	projectId := vars[constants.ProjectID]
	clusterName := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	//decode the request body into input model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.DatabaseName = &databaseName
	model.Username = &username
	model.Password = &password
	model.ProjectId = &projectId
	model.ClusterName = &clusterName
	model.PublicKey = &publicKey
	model.PrivateKey = &privateKey

	//log the input model
	util.Debugf(r.Context(), "Create Collection Request : %+v", model.String())
	startTime := time.Now()

	//make API call to create a collection
	response := collection.Create(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create collection REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.Collectionhandler)
}

// ListCollection handles LIST requests deletes the requested collection
// @Summary Lists all the collection
// @Description Lists all the collections in the database
// @ID ListCollection
// @Tags Collection
// @Produce json
// @Param DatabaseName path string true "databaseName" default(testDatabase)
// @Param ProjectId path string true "ProjectId" default(projectId)
// @Param ClusterName path string true "ClusterName" default(clusterId)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param Authorization header string true "authorization" default()
// @Security BasicAuth
// @Success 200 {object} atlasresponse.AtlasResponse
// @Failure 400 {object} atlasresponse.AtlasResponse
// @Failure 401 {object} atlasresponse.AtlasResponse
// @Failure 403 {object} atlasresponse.AtlasResponse
// @Failure 404 {object} atlasresponse.AtlasResponse
// @Failure 500 {object} atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/cluster/{ClusterName}/database/{DatabaseName}/collections [get]
func ListCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseName]
	username := r.Context().Value(constants.Username).(string)
	password := r.Context().Value(constants.Password).(string)
	projectId := vars[constants.ProjectID]
	clusterName := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

	//create input model for delete collection API
	model := collection.DeleteInputModel{
		DatabaseName: &databaseName,
		Username:     &username,
		Password:     &password,
		ProjectId:    &projectId,
		ClusterName:  &clusterName,
		PublicKey:    &publicKey,
		PrivateKey:   &privateKey,
	}

	//log the input model
	util.Debugf(r.Context(), "List all Collections Request : %+v", model.String())
	startTime := time.Now()

	//make API call to delete a collection
	response := collection.ReadAll(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "List all collections REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.Collectionhandler)
}

// DeleteCollection handles DELETE requests deletes the requested collection
// @Summary Delete a collection
// @Description Delete a collection with the specified ID
// @ID DeleteCollection
// @Tags Collection
// @Accept json
// @Produce json
// @Param DatabaseName path string true "databaseName" default(testDatabase)
// @Param CollectionName path string true "collectionName" default(default)
// @Param ProjectId path string true "ProjectId" default(projectId)
// @Param ClusterName path string true "ClusterName" default(clusterId)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Security BasicAuth
// @Param Authorization header string true "authorization" default()
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/cluster/{ClusterName}/database/{DatabaseName}/collection/{CollectionName} [delete]
func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseName]
	collectionName := vars[constants.CollectionName]
	username := r.Context().Value(constants.Username).(string)
	password := r.Context().Value(constants.Password).(string)
	projectId := vars[constants.ProjectID]
	clusterName := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

	//create input model for delete collection API
	model := collection.DeleteInputModel{
		DatabaseName:   &databaseName,
		Username:       &username,
		Password:       &password,
		CollectionName: &collectionName,
		ProjectId:      &projectId,
		ClusterName:    &clusterName,
		PublicKey:      &publicKey,
		PrivateKey:     &privateKey,
	}

	//log the input model
	util.Debugf(r.Context(), "Delete Collection Request : %+v", model.String())
	startTime := time.Now()

	//make API call to delete a collection
	response := collection.Delete(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Delete collection REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.Collectionhandler)
}
