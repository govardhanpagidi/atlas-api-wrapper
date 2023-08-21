package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/atlas-api-helper/resources/cluster"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
)

// setupClusterLog sets up the logger for the cluster API handlers
func setupClusterLog() {
	util.SetupLogger("atlas-api-helper.handlers.cluster")
}

// GetCluster handles GET requests to return the state of the cluster
func GetCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	name := vars[constants.ClusterNamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)
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
func GetAllClusters(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)
	projectId := vars[constants.ProjectIdPathParam]
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

// DeleteCluster handles the DELETE requests to terminate the cluster
func DeleteCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()

	//fetch all input parameters and create input model
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	name := vars[constants.ClusterNamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)
	model := cluster.InputModel{ProjectId: &projectId, ClusterName: &name, PrivateKey: &privateKey, PublicKey: &publicKey}

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
func CreateCluster(w http.ResponseWriter, r *http.Request) {
	//fetch all input parameters and create input model
	var model cluster.InputModel
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]

	//decode the request body into input model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.ProjectId = &projectId

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
