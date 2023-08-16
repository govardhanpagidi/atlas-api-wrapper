package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/database"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
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
	response := database.Create(&model)
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
	response := database.Delete(&database.DeleteInputModel{
		DatabaseName: &databaseName,
		HostName:     &hostname,
		Username:     &username,
		Password:     &password,
	})
	responseHandler.Write(response, w, constants.ClusterHandler)
}
