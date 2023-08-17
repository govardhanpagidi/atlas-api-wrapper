package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/database"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
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
	var model database.InputModel

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, _ = logger.Debugf("Create database Request : %+v", model.String())
	startTime := time.Now()

	response := database.Create(&model)
	elapsedTime := time.Since(startTime)
	logger.Debugf("Create database REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	responseHandler.Write(response, w, constants.ClusterHandler)
}

// DeleteDatabase handles the DELETE calls to delete the requested database
func DeleteDatabase(w http.ResponseWriter, r *http.Request) {
	setupDatabaseLog()
	vars := mux.Vars(r)
	// Read a specific parameter
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
	_, _ = logger.Debugf("Delete database Request : %+v", model.String())
	startTime := time.Now()
	response := database.Delete(&model)
	elapsedTime := time.Since(startTime)
	logger.Debugf("Delete database REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	responseHandler.Write(response, w, constants.ClusterHandler)
}
