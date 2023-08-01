package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/database_user"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
)

func setupDatabaseUserLog() {
	util.SetupLogger("atlas-api-helper.handlers.databaseuser")
}

// GetDatabaseUser handles GET requests to retrieve all database users
func GetDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()
	vars := mux.Vars(r)

	// Read a specific parameter
	groupId := vars[constants.GroupID]
	databaseName := vars[constants.DatabaseName]
	username := vars[constants.Username]
	response := database_user.Read(r.Context(), &database_user.Model{ProjectId: &groupId, DatabaseName: &databaseName, Username: &username})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}

func GetAllDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()
	vars := mux.Vars(r)

	// Read a specific parameter
	groupId := vars[constants.GroupID]
	response := database_user.List(r.Context(), &database_user.Model{ProjectId: &groupId})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}

func DeleteDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	vars := mux.Vars(r)
	// Read a specific parameter
	groupId := vars[constants.GroupID]
	databaseName := vars[constants.DatabaseName]
	username := vars[constants.Username]
	response := database_user.Delete(r.Context(), &database_user.Model{ProjectId: &groupId, DatabaseName: &databaseName, Username: &username})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}

func CreateDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()
	var model database_user.Model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := database_user.Create(r.Context(), &model)
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}

func UpdateDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()
	var model database_user.Model
	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := database_user.Update(r.Context(), &model)
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}
