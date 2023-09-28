package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/atlas-api-helper/resources/cluster"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/constants"
	responseHandler "github.com/atlas-api-helper/util/responsehandler"
	"github.com/gorilla/mux"
)

// setupClusterLog sets up the logger for the cluster API handlers
func setupClusterLog() {
	util.SetupLogger("atlas-api-helper.handlers.cluster")
}

// GetCluster handles GET requests to return the state of the cluster
// @Summary Get the state of a cluster
// @Description Get the state of a cluster by project ID and cluster name
// @ID GetCluster
// @Tags Cluster
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
// @Router /project/{ProjectId}/cluster/{ClusterName}/status [get]
func GetCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	name := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	model := cluster.InputModel{ProjectId: &projectId, ClusterName: &name, PrivateKey: &privateKey, PublicKey: &publicKey}

	//log the input model
	util.Debugf(r.Context(), "Get cluster request : %+v", model.String())
	startTime := time.Now()

	//make the API call to read a cluster
	response := cluster.Read(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Get Cluster REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// GetAllClusters handles GET requests to return all the clusters along with cluster's advanced configuration
// @Summary Get all clusters
// @Description Get all clusters along with their advanced configuration by project ID
// @ID GetAllClusters
// @OperationId GetAllClusters
// @Tags Cluster
// @Accept json
// @Produce json
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Success 200 {object} []cluster.Model
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/cluster [get]
func GetAllClusters(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	projectId := vars[constants.ProjectID]
	model := cluster.InputModel{ProjectId: &projectId, PrivateKey: &privateKey, PublicKey: &publicKey}

	//log the input model
	util.Debugf(r.Context(), "Get all clusters request : %+v", model.String())
	startTime := time.Now()

	//make the API call to read all clusters
	response := cluster.List(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Get all Clusters REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// DeleteCluster handles DELETE requests to delete a cluster
// @Summary Delete a cluster
// @Description Delete a cluster by project ID and cluster name
// @ID DeleteCluster
// @Tags Cluster
// @Accept json
// @Produce json
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param ClusterName path string true "Cluster name" default()
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param RetainBackup query string true "retainBackup" default(true)
// @Success 200 {object}  atlasresponse.AtlasResponse
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/cluster/{ClusterName} [delete]
func DeleteCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	name := vars[constants.ClusterName]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	retainBackup := r.URL.Query().Get(constants.RetainBackup)
	model := cluster.InputModel{ProjectId: &projectId, ClusterName: &name, PrivateKey: &privateKey, PublicKey: &publicKey, RetainBackup: &retainBackup}

	//log the input model
	util.Debugf(r.Context(), "Delete cluster request : %+v", model.String())
	startTime := time.Now()

	//make the API call to delete a cluster
	response := cluster.Delete(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Delete Cluster REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// CreateCluster handles the POST requests to create the cluster with the provided TshirtSize
// @Summary CreateCluster handles the POST requests to create the cluster with the provided TshirtSize
// @Description Create the cluster based on the provided TshirtSize configuration
// @ID CreateCluster
// x-go-name CreateCluster
// @Tags Cluster
// @Accept json
// @Produce json
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param InputModel body cluster.InputModel true "body"
// @Success 200 {object} cluster.Model
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/cluster [post]
func CreateCluster(w http.ResponseWriter, r *http.Request) {
	//fetch all input parameters and create input model
	var model cluster.InputModel
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	//decode the request body into input model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.ProjectId = &projectId
	model.PublicKey = &publicKey
	model.PrivateKey = &privateKey

	//log the input model
	util.Debugf(r.Context(), "Create cluster request : %+v", model.String())
	startTime := time.Now()

	//make the API call to create a cluster
	response := cluster.Create(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create Cluster REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// UpdateCluster handles the PUT requests to create the cluster with the provided MongoDbVersion and MongoDBMajorVersion
// @Summary UpdateCluster handles the Put requests to create the provided MongoDbVersion and MongoDBMajorVersion
// @Description Update the cluster with MongoDBMajorVersion
// @ID UpdateCluster
// @Tags Cluster
// @Accept json
// @Produce json
// @Param ProjectId path string true "Project ID" default(<projectID>)
// @Param ClusterName path string true "Cluster name" default()
// @Param x-mongo-publickey header string true "Public Key" default(<publicKey>)
// @Param x-mongo-privatekey header string true "Private Key" default(<privateKey>)
// @Param InputModel body cluster.UpdateInputModel true "body"
// @Success 200 {object} cluster.Model
// @Failure 400 {object}  atlasresponse.AtlasResponse
// @Failure 401 {object}  atlasresponse.AtlasResponse
// @Failure 403 {object}  atlasresponse.AtlasResponse
// @Failure 404 {object}  atlasresponse.AtlasResponse
// @Failure 500 {object}  atlasresponse.AtlasResponse
// @Router /project/{ProjectId}/cluster/{ClusterName} [patch]
func UpdateCluster(w http.ResponseWriter, r *http.Request) {
	//fetch all input parameters and create input model
	var model cluster.UpdateInputModel
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectID]
	publicKey := r.Header.Get(constants.PublicKeyHeader)
	privateKey := r.Header.Get(constants.PrivateKeyHeader)
	clusterName := vars[constants.ClusterName]

	//decode the request body into input model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.PublicKey = &publicKey
	model.PrivateKey = &privateKey
	model.ProjectId = &projectId
	model.ClusterName = &clusterName

	//log the input model
	util.Debugf(r.Context(), "Create cluster request : %+v", model.ToString())
	startTime := time.Now()

	//make the API call to create a cluster
	response := cluster.Update(r.Context(), &model)

	//calculate the elapsed time and log the response
	elapsedTime := time.Since(startTime)
	util.Debugf(r.Context(), "Create Cluster REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	//write the response to the output
	responseHandler.Write(response, w, constants.ClusterHandler)
}
