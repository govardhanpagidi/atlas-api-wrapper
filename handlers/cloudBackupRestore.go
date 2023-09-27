package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/atlas-api-helper/resources/cloudBackupRestore"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/constants"
	responseHandler "github.com/atlas-api-helper/util/responsehandler"
	"github.com/gorilla/mux"
)

// setupClusterLog sets up the logger for the cluster API handlers
func setupCloudRestoreLog() {
	util.SetupLogger("atlas-api-helper.handlers.cloudbackuprestore")
}

// CreateRestoreJob @Summary Create Restore Job
// @Description Creates a restore job for the specified cluster.
// @ID CreateRestoreJob
// @Tags Cluster Restore
// @Accept json
// @Produce json
// @Param ClusterName path string true "The name of the cluster to restore."
// @Param ProjectId path string true "The ID of the project that contains the cluster to restore."
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param InputModel body cloudBackupRestore.InputModel true "body"
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /api/project/{ProjectId}/cluster/{ClusterName}/restore [post]
func CreateRestoreJob(w http.ResponseWriter, r *http.Request) {
	setupCloudRestoreLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	var model cloudBackupRestore.InputModel
	projectId := vars[constants.ProjectID]
	clusterName := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

	//decode the request body into input model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.ProjectId = &projectId
	model.ClusterName = &clusterName
	model.PublicKey = &publicKey
	model.PrivateKey = &privateKey

	//log the input model

	util.Debugf(r.Context(), "Create restore job request : %+v", model.String())
	startTime := time.Now()

	//make the API call to create a cluster backup restore job
	response := cloudBackupRestore.Create(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create restore job REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterBackupRestoreHandler)
}

// GetRestoreJob returns the restore job for the specified cluster.
//
// @Summary Get Restore Job
// @Description Returns the restore job for the specified cluster.
// @ID GetRestoreJob
// @Tags Cluster Restore
// @Accept json
// @Produce json
// @Param ClusterName path string true "The name of the cluster to restore."
// @Param ProjectId path string true "The ID of the project that contains the cluster to restore."
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param JobId query string true "The JobID of the cluster to restore."
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /api/project/{ProjectId}/cluster/{ClusterName}/restore [get]
// GetRestoreJob returns the status restore job for a given cluster
func GetRestoreJob(w http.ResponseWriter, r *http.Request) {
	setupCloudRestoreLog()
	// Fetch input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	clusterName := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	jobId := r.URL.Query().Get(constants.JobId)

	model := cloudBackupRestore.InputModel{
		ProjectId:   &projectId,
		ClusterName: &clusterName,
		PrivateKey:  &privateKey,
		PublicKey:   &publicKey,
		JobId:       &jobId,
	}

	// Log the input model
	util.Debugf(r.Context(), "Get restore job request : %+v", model.String())

	startTime := time.Now()

	// Make the API call to read a cloud backup restore job
	response := cloudBackupRestore.Read(r.Context(), &model)

	// Calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Get restore job REST API response:%+v and execution time:%s", response.String(), elapsedTime.String())

	// Write the response to the output
	responseHandler.Write(response, w, constants.ClusterBackupRestoreHandler)
}
