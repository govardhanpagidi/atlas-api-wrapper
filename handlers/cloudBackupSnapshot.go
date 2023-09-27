package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/atlas-api-helper/resources/cloudBackupSnapshot"

	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/constants"
	responseHandler "github.com/atlas-api-helper/util/responsehandler"
	"github.com/gorilla/mux"
)

// setupClusterLog sets up the logger for the cluster API handlers
func setupCloudCkackupSnapshotLog() {
	util.SetupLogger("atlas-api-helper.handlers.cluster")
}

// GetAllBackupSnapshot handles GET requests to return all cloud backup snapshots of a cluster
// @Summary Get all cloud backup snapshots of a cluster
// @ID GetAllBackupSnapshot
// @Description Get all cloud backup snapshots of a cluster by project ID and cluster name
// @Tags Cloud Backup Snapshot
// @Accept json
// @Produce json
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param ClusterName path string true "Cluster name" default(s-aws-04-09-23-15-02-41-5e8de3e1042f5b33ab81f33a)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /api/project/{ProjectId}/cluster/{ClusterName}/snapshot [get]
func GetAllBackupSnapshot(w http.ResponseWriter, r *http.Request) {
	setupCloudCkackupSnapshotLog()
	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	clusterName := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

	model := cloudBackupSnapshot.InputModel{
		ProjectId:   &projectId,
		ClusterName: &clusterName,
		PrivateKey:  &privateKey,
		PublicKey:   &publicKey,
	}
	//log the input model
	util.Debugf(r.Context(), "Get cluster request : %+v", model.ToString())
	startTime := time.Now()

	//make the API call to read a cluster
	response := cloudBackupSnapshot.List(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Get Cluster REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterBackupSnapshotHandler)
}

// CreateBackupSnapshot handles POST requests to create a cloud backup snapshot of a cluster
// @Summary Create a cloud backup snapshot of a cluster
// @ID CreateBackupSnapshot
// @Description Create a cloud backup snapshot of a cluster by project ID and cluster name
// @Tags Cloud Backup Snapshot
// @Accept json
// @Produce json
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param ClusterName path string true "Cluster name" default(s-aws-04-09-23-15-02-41-5e8de3e1042f5b33ab81f33a)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param InputModel body cloudBackupSnapshot.InputModel true "body"
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /api/project/{ProjectId}/cluster/{ClusterName}/snapshot [post]
func CreateBackupSnapshot(w http.ResponseWriter, r *http.Request) {
	setupCloudCkackupSnapshotLog()
	// Fetch input parameters and create input model
	vars := mux.Vars(r)
	var model cloudBackupSnapshot.InputModel
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

	// Log the input model
	util.Debugf(r.Context(), "Create cloud backup snapshot request : %+v", model.ToString())

	startTime := time.Now()

	// Make the API call to create a cloud backup snapshot
	response := cloudBackupSnapshot.Create(r.Context(), &model)

	// Calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create Cloud Backup Snapshot REST API response:%+v and execution time:%s", response.String(), elapsedTime.String())

	// Write the response to the output
	responseHandler.Write(response, w, constants.ClusterBackupSnapshotHandler)
}
