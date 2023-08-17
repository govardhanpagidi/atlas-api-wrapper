package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/collection"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
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
	var model collection.InputModel
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseNamePathParam]
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.DatabaseName = &databaseName
	_, _ = logger.Debugf("Create Collection Request : %+v", model.String())
	startTime := time.Now()

	response := collection.Create(&model)
	elapsedTime := time.Since(startTime)
	logger.Debugf("Create collection REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// DeleteCollection handles DELETE requests deletes the requested collection
func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseNamePathParam]
	collectionName := vars[constants.CollectionNamePathParam]
	hostname := r.URL.Query().Get(constants.HostNamePathParam)
	username := r.URL.Query().Get(constants.UsernamePathParam)
	password := r.URL.Query().Get(constants.PasswordPathParam)

	model := collection.InputModel{
		DatabaseName:   &databaseName,
		HostName:       &hostname,
		Username:       &username,
		Password:       &password,
		CollectionName: &collectionName,
	}
	_, _ = logger.Debugf("Delete Collection Request : %+v", model.String())
	startTime := time.Now()

	response := collection.Delete(&model)
	elapsedTime := time.Since(startTime)
	logger.Debugf("Delete collection REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	responseHandler.Write(response, w, constants.ClusterHandler)
}
