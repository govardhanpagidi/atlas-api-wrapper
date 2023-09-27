package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/atlas-api-helper/resources/cloudBackupSchedule"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/constants"
	responseHandler "github.com/atlas-api-helper/util/responsehandler"
	"github.com/gorilla/mux"
)

// UpdateClusterBackupPolicy handles the PUT requests to update the cluster backup policy
// @Summary UpdateClusterBackupPolicy handles the PUT requests to update the cluster backup policy
// @Description Update the cluster backup policy
// @ID UpdateClusterBackupPolicy
// @Tags Cloud Backup Schedule
// @Accept json
// @Produce json
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param ClusterName path string true "Cluster name" default(s-aws-04-09-23-15-02-41-5e8de3e1042f5b33ab81f33a)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param UpdateInputModel body cloudBackupSchedule.Model true "body"
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /api/project/{ProjectId}/clusters/{ClusterName}/backup/schedule [patch]
func UpdateClusterBackupPolicy(w http.ResponseWriter, r *http.Request) {
	//fetch all input parameters and create input model
	var model cloudBackupSchedule.Model

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	name := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)

	//decode the request body into input model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	model.ProjectId = &projectId
	model.ClusterName = &name
	model.PublicKey = &publicKey
	model.PrivateKey = &privateKey

	//log the input model
	util.Debugf(r.Context(), "update cloudBackup request : %+v", model.ToString())
	startTime := time.Now()

	//make the API call to update a cloud backup restore schedule
	response := cloudBackupSchedule.Update(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Update CloudBackup REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.CloudBackupHandler)
}

// GetCloudBackupSchedule handles GET requests to return the state of the cluster backup schedule
// @Summary Get the state of a cluster backup schedule
// @Description Get the state of a cluster backup schedule by project ID and cluster name
// @ID GetCloudBackupSchedule
// @Tags Cloud Backup Schedule
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
// @Router /api/project/{ProjectId}/clusters/{ClusterName}/backup/schedule [get]
func GetCloudBackupSchedule(w http.ResponseWriter, r *http.Request) {

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	name := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	model := cloudBackupSchedule.Model{ProjectId: &projectId, ClusterName: &name, PrivateKey: &privateKey, PublicKey: &publicKey}

	//log the input model
	util.Debugf(r.Context(), "Get Cloud Backup request : %+v", model.ToString())
	startTime := time.Now()

	//make the API call to read a cloud backup schedule
	response := cloudBackupSchedule.Read(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Get Cloud Backup REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.CloudBackupHandler)
}
