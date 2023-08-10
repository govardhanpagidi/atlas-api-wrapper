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
	util.SetupLogger("atlas-api-helper.handlers.databaseuser")
}

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

func DeleteDatabase(w http.ResponseWriter, r *http.Request) {
	setupDatabaseLog()
	vars := mux.Vars(r)
	// Read a specific parameter
	databaseName := vars[constants.DatabaseName]

	hostname := r.URL.Query().Get("HostName")
	username := r.URL.Query().Get("Username")
	password := r.URL.Query().Get("Password")
	response := database.Delete(&database.DeleteInputModel{
		DatabaseName: &databaseName,
		HostName:     &hostname,
		Username:     &username,
		Password:     &password,
	})
	responseHandler.Write(response, w, constants.ClusterHandler)
}
