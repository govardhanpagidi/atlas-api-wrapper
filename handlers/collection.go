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
// @Tags Collection
// @Accept json
// @Produce json
// @Param name formData string true "Collection name"
// @Param description formData string false "Collection description"
// @Success 200 {object} atlasresponse.AtlasRespone
// @Failure 400 {object} atlasresponse.AtlasRespone
// @Failure 401 {object} atlasresponse.AtlasRespone
// @Failure 403 {object} atlasresponse.AtlasRespone
// @Failure 404 {object} atlasresponse.AtlasRespone
// @Failure 500 {object} atlasresponse.AtlasRespone
// @Router /collections [post]
func CreateCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()

	//fetch all input parameters and create input model
	var model collection.InputModel
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseNamePathParam]

	//decode the request body into input model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.DatabaseName = &databaseName

	//log the input model
	util.Debugf(r.Context(), "Create Collection Request : %+v", model.String())
	startTime := time.Now()

	//make API call to create a collection
	response := collection.Create(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create collection REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// DeleteCollection handles DELETE requests deletes the requested collection
// @Summary Delete a collection
// @Description Delete a collection with the specified ID
// @Tags Collection
// @Accept json
// @Produce json
// @Param id path string true "Collection ID"
// @Success 200 {object} atlasresponse.AtlasRespone
// @Failure 400 {object} atlasresponse.AtlasRespone
// @Failure 401 {object} atlasresponse.AtlasRespone
// @Failure 403 {object} atlasresponse.AtlasRespone
// @Failure 404 {object} atlasresponse.AtlasRespone
// @Failure 500 {object} atlasresponse.AtlasRespone
// @Router /collections/{id} [delete]
func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseNamePathParam]
	collectionName := vars[constants.CollectionNamePathParam]
	hostname := r.URL.Query().Get(constants.HostNamePathParam)
	username := r.URL.Query().Get(constants.UsernamePathParam)
	password := r.URL.Query().Get(constants.PasswordPathParam)

	//create input model for delete collection API
	model := collection.DeleteInputModel{
		DatabaseName:   &databaseName,
		HostName:       &hostname,
		Username:       &username,
		Password:       &password,
		CollectionName: &collectionName,
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
	responseHandler.Write(response, w, constants.ClusterHandler)
}
