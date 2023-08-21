package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/collection"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func setupCollectionLog() {
	util.SetupLogger("atlas-api-helper.handlers.collection")
}

// CreateCollection handles POST requests to create a new collection using the credentials
func CreateCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()

	//fetch all input parameters and create input model
	var model collection.InputModel
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseNamePathParam]
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.DatabaseName = &databaseName

	util.Debugf(r.Context(), "Create Collection Request : %+v", model.String())
	startTime := time.Now()

	//make API call to create a collection
	response := collection.Create(r.Context(), &model)

	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create collection REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// DeleteCollection handles DELETE requests deletes the requested collection
func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseNamePathParam]
	collectionName := vars[constants.CollectionNamePathParam]
	hostname := r.URL.Query().Get(constants.HostNamePathParam)
	username := r.URL.Query().Get(constants.UsernamePathParam)
	password := r.URL.Query().Get(constants.PasswordPathParam)

	model := collection.DeleteInputModel{
		DatabaseName:   &databaseName,
		HostName:       &hostname,
		Username:       &username,
		Password:       &password,
		CollectionName: &collectionName,
	}

	util.Debugf(r.Context(), "Delete Collection Request : %+v", model.String())
	startTime := time.Now()

	//make API call to delete a collection
	response := collection.Delete(r.Context(), &model)

	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Delete collection REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())
	responseHandler.Write(response, w, constants.ClusterHandler)
}
