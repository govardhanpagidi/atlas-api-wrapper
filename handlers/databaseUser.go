package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/databaseUser"
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

	groupId := vars[constants.GroupID]
	databaseName := vars[constants.DatabaseName]
	username := vars[constants.Username]
	publicKey := r.URL.Query().Get("publicKey")
	privateKey := r.URL.Query().Get("privateKey")
	response := database_user.Read(&database_user.InputModel{ProjectId: &groupId, DatabaseName: &databaseName, Username: &username, PublicKey: &publicKey, PrivateKey: &privateKey})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

func GetAllDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()
	vars := mux.Vars(r)

	groupId := vars[constants.GroupID]
	publicKey := r.URL.Query().Get("publicKey")
	privateKey := r.URL.Query().Get("privateKey")
	response := database_user.List(r.Context(), &database_user.InputModel{ProjectId: &groupId, PublicKey: &publicKey, PrivateKey: &privateKey})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

func DeleteDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	vars := mux.Vars(r)
	// Read a specific parameter
	groupId := vars[constants.GroupID]
	databaseName := vars[constants.DatabaseName]
	username := vars[constants.Username]
	publicKey := r.URL.Query().Get("publicKey")
	privateKey := r.URL.Query().Get("privateKey")
	response := database_user.Delete(r.Context(), &database_user.InputModel{ProjectId: &groupId, DatabaseName: &databaseName, Username: &username, PublicKey: &publicKey, PrivateKey: &privateKey})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

func CreateDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()
	var model database_user.InputModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := database_user.Create(&model)
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}
