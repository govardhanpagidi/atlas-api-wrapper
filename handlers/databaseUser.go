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
// @ID GetDatabaseUser
// @Tags Database User
// @Accept json
// @Produce json
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param Username path string true "Username" default(testUser)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Success 200 {object} database_user.Model
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /api/project/{ProjectId}/databaseUsers/{Username} [get]
func GetDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	username := vars[constants.Username]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

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
// @ID GetAllDatabaseUser
// @Tags Database User
// @Accept json
// @Produce json
// @Success 200 {object} []database_user.Model
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /api/project/{ProjectId}/databaseUsers [get]
func GetAllDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

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
// @ID DeleteDatabaseUser
// @Tags Database User
// @Accept json
// @Produce json
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param Username path string true "Username" default(testUser)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /api/project/{ProjectId}/databaseUsers/{Username} [delete]
func DeleteDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	username := vars[constants.Username]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

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
// @ID CreateDatabaseUser
// @Tags Database User
// @Accept json
// @Produce json
// @Success 200 {object} database_user.Model
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param InputModel body database_user.InputModel true "body"
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/databaseUsers [post]
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
	projectId := vars[constants.ProjectID]
	model.ProjectId = &projectId
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	model.PublicKey = &publicKey
	model.PrivateKey = &privateKey
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

// UpdateDatabaseUser handles POST requests to create one database user
// @Summary Update roles and permission for a database user
// @Description Update roles and permission for a database user
// @ID UpdateDatabaseUser
// @Tags Database User
// @Accept json
// @Produce json
// @Success 200 {object} database_user.Model
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param Username path string true "Username" default(testUser)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param InputModel body database_user.UpdateInputModel true "body"
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/databaseUsers/{Username} [patch]
func UpdateDatabaseUser(w http.ResponseWriter, r *http.Request) {
	setupDatabaseUserLog()

	//fetch all input parameters and create input model
	var model database_user.UpdateInputModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	userName := vars[constants.Username]
	model.ProjectId = &projectId
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	model.PublicKey = &publicKey
	model.PrivateKey = &privateKey
	model.Username = &userName
	//log the input model
	util.Debugf(r.Context(), "Create databaseUser Request : %+v", model.String())
	startTime := time.Now()

	//make API call to create a database user
	response := database_user.Update(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create databaseUser REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
}
