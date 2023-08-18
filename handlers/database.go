package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/database"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func setupDatabaseLog() {
	util.SetupLogger("atlas-api-helper.handlers.database")
}

// CreateDatabase handles POST calls to create a new database using the provided parameters
func CreateDatabase(w http.ResponseWriter, r *http.Request) {
	setupDatabaseLog()

	//fetch all input parameters and create input model
	var model database.InputModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	util.Debugf(r.Context(), "Create database Request : %+v", model.String())
	startTime := time.Now()

	//make API call to create a database
	response := database.Create(r.Context(), &model)

	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create database REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	responseHandler.Write(response, w, constants.ClusterHandler)
}

// DeleteDatabase handles the DELETE calls to delete the requested database
func DeleteDatabase(w http.ResponseWriter, r *http.Request) {
	setupDatabaseLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	databaseName := vars[constants.DatabaseNamePathParam]

	hostname := r.URL.Query().Get(constants.HostNamePathParam)
	username := r.URL.Query().Get(constants.UsernamePathParam)
	password := r.URL.Query().Get(constants.PasswordPathParam)

	model := database.DeleteInputModel{
		DatabaseName: &databaseName,
		HostName:     &hostname,
		Username:     &username,
		Password:     &password,
	}

	util.Debugf(r.Context(), "Delete database Request : %+v", model.String())
	startTime := time.Now()

	//make API call to delete a database
	response := database.Delete(r.Context(), &model)

	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Delete database REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	responseHandler.Write(response, w, constants.ClusterHandler)
}