package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/databaseUser"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func setupDatabaseUserLog() {
	util.SetupLogger("atlas-api-helper.handlers.databaseUser")
}

// GetDatabaseUser handles GET requests to retrieve one database user
func GetDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	username := vars[constants.UsernamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)

	model := database_user.InputModel{ProjectId: &projectId, Username: &username, PublicKey: &publicKey, PrivateKey: &privateKey}
	util.Debugf(r.Context(), "Get databaseUser Request : %+v", model.String())
	startTime := time.Now()

	//make API call to fetch a database user
	response := database_user.Read(r.Context(), &model)

	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Get databaseUser REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

// GetAllDatabaseUser handles GET requests to retrieve all database users
func GetAllDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)

	model := database_user.InputModel{ProjectId: &projectId, PublicKey: &publicKey, PrivateKey: &privateKey}
	util.Debugf(r.Context(), "Get all databaseUser Request : %+v", model.String())
	startTime := time.Now()

	//make API call to list all database users
	response := database_user.List(r.Context(), &model)

	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Get all databaseUser REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

// DeleteDatabaseUser handles DELETE requests to delete one database user
func DeleteDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	username := vars[constants.UsernamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)

	model := database_user.InputModel{ProjectId: &projectId, Username: &username, PublicKey: &publicKey, PrivateKey: &privateKey}
	util.Debugf(r.Context(), "Delete databaseUser Request : %+v", model.String())
	startTime := time.Now()

	//make API call to delete a database user
	response := database_user.Delete(r.Context(), &model)

	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Delete databaseUser REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

// CreateDatabaseUser handles POST requests to create one database user
func CreateDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	var model database_user.InputModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	model.ProjectId = &projectId

	util.Debugf(r.Context(), "Create databaseUser Request : %+v", model.String())
	startTime := time.Now()

	//make API call to create a database user
	response := database_user.Create(r.Context(), &model)

	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create databaseUser REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}
