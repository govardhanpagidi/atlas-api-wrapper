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

// GetDatabaseUser handles GET requests to retrieve one database user
func GetDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	databaseName := vars[constants.DatabaseNamePathParam]
	username := vars[constants.UsernamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)
	response := database_user.Read(&database_user.InputModel{ProjectId: &projectId, DatabaseName: &databaseName, Username: &username, PublicKey: &publicKey, PrivateKey: &privateKey})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

// GetAllDatabaseUser handles GET requests to retrieve all database users
func GetAllDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)
	response := database_user.List(r.Context(), &database_user.InputModel{ProjectId: &projectId, PublicKey: &publicKey, PrivateKey: &privateKey})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

// DeleteDatabaseUser handles DELETE requests to delete one database user
func DeleteDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	databaseName := vars[constants.DatabaseNamePathParam]
	username := vars[constants.UsernamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)
	response := database_user.Delete(r.Context(), &database_user.InputModel{ProjectId: &projectId, DatabaseName: &databaseName, Username: &username, PublicKey: &publicKey, PrivateKey: &privateKey})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

// CreateDatabaseUser handles POST requests to create one database user
func CreateDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()
	var model database_user.InputModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	model.ProjectId = &projectId
	response := database_user.Create(&model)
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}
