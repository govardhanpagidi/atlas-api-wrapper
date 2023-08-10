package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/collection"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
)

func setupCollectionLog() {
	util.SetupLogger("atlas-api-helper.handlers.databaseuser")
}

func CreateCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()
	var model collection.InputModel

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := collection.Create(&model)
	responseHandler.Write(response, w, constants.ClusterHandler)
}

func DeleteCollection(w http.ResponseWriter, r *http.Request) {
	setupCollectionLog()
	vars := mux.Vars(r)
	// Read a specific parameter
	databaseName := vars[constants.DatabaseName]
	collectionName := vars[constants.CollectionName]
	hostname := r.URL.Query().Get("HostName")
	username := r.URL.Query().Get("Username")
	password := r.URL.Query().Get("Password")

	response := collection.Delete(&collection.InputModel{
		DatabaseName:   &databaseName,
		HostName:       &hostname,
		Username:       &username,
		Password:       &password,
		CollectionName: &collectionName,
	})
	responseHandler.Write(response, w, constants.ClusterHandler)
}
