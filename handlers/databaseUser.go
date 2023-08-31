package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	database_user "github.com/atlas-api-helper/resources/databaseUser"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/constants"
	responseHandler "github.com/atlas-api-helper/util/responsehandler"
	"github.com/gorilla/mux"
)

// setupDatabaseUserLog sets up the logger for the database user API handlers
func setupDatabaseUserLog() {
	util.SetupLogger("atlas-api-helper.handlers.databaseUser")
}

// GetDatabaseUser handles GET requests to retrieve one database user
// @Summary Get a database user
// @Description Get a database user with the specified ID
// @Tags Database User
// @Accept json
// @Produce json
// @Param id path string true "Database user ID"
// @Success 200 {object} database_user.Model
// @Failure 400 {object} atlasresponse.AtlasRespone
// @Failure 401 {object} atlasresponse.AtlasRespone
// @Failure 403 {object} atlasresponse.AtlasRespone
// @Failure 404 {object} atlasresponse.AtlasRespone
// @Failure 500 {object} atlasresponse.AtlasRespone
// @Router /database-users/{id} [get]
func GetDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	username := vars[constants.UsernamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)

	//create input model for get database user API
	model := database_user.InputModel{
		ProjectId:  &projectId,
		Username:   &username,
		PublicKey:  &publicKey,
		PrivateKey: &privateKey,
	}

	//log the input model
	util.Debugf(r.Context(), "Get databaseUser Request : %+v", model.String())
	startTime := time.Now()

	//make API call to fetch a database user
	response := database_user.Read(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Get databaseUser REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

// GetAllDatabaseUser handles GET requests to retrieve all database users
// @Summary Get all database users
// @Description Get all database users
// @Tags Database User
// @Accept json
// @Produce json
// @Success 200 {object} []database_user.Model
// @Failure 400 {object} atlasresponse.AtlasRespone
// @Failure 401 {object} atlasresponse.AtlasRespone
// @Failure 403 {object} atlasresponse.AtlasRespone
// @Failure 404 {object} atlasresponse.AtlasRespone
// @Failure 500 {object} atlasresponse.AtlasRespone
// @Router /database-users [get]
func GetAllDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)

	//create input model for list all database users API
	model := database_user.InputModel{
		ProjectId:  &projectId,
		PublicKey:  &publicKey,
		PrivateKey: &privateKey,
	}

	//log the input model
	util.Debugf(r.Context(), "Get all databaseUser Request : %+v", model.String())
	startTime := time.Now()

	//make API call to list all database users
	response := database_user.List(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Get all databaseUser REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

// DeleteDatabaseUser handles DELETE requests to delete one database user
// @Summary Delete a database user
// @Description Delete a database user with the specified ID
// @Tags Database User
// @Accept json
// @Produce json
// @Param id path string true "Database user ID"
// @Success 200 {object} atlasresponse.AtlasRespone
// @Failure 400 {object} atlasresponse.AtlasRespone
// @Failure 401 {object} atlasresponse.AtlasRespone
// @Failure 403 {object} atlasresponse.AtlasRespone
// @Failure 404 {object} atlasresponse.AtlasRespone
// @Failure 500 {object} atlasresponse.AtlasRespone
// @Router /database-users/{id} [delete]
func DeleteDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	username := vars[constants.UsernamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)

	//create input model for delete database user API
	model := database_user.InputModel{
		ProjectId:  &projectId,
		Username:   &username,
		PublicKey:  &publicKey,
		PrivateKey: &privateKey,
	}

	//log the input model
	util.Debugf(r.Context(), "Delete databaseUser Request : %+v", model.String())
	startTime := time.Now()

	//make API call to delete a database user
	response := database_user.Delete(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Delete databaseUser REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}

// CreateDatabaseUser handles POST requests to create one database user
// @Summary Create a database user
// @Description Create a new database user with the specified name and email
// @Tags Database User
// @Accept json
// @Produce json
// @Param name formData string true "Database user name"
// @Param email formData string true "Database user email"
// @Success 200 {object} database_user.Model
// @Failure 400 {object} atlasresponse.AtlasRespone
// @Failure 401 {object} atlasresponse.AtlasRespone
// @Failure 403 {object} atlasresponse.AtlasRespone
// @Failure 404 {object} atlasresponse.AtlasRespone
// @Failure 500 {object} atlasresponse.AtlasRespone
// @Router /database-users [post]
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

	//log the input model
	util.Debugf(r.Context(), "Create databaseUser Request : %+v", model.String())
	startTime := time.Now()

	//make API call to create a database user
	response := database_user.Create(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create databaseUser REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}
