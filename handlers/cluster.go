package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/cluster"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func setupClusterLog() {
	util.SetupLogger("atlas-api-helper.handlers.cluster")
}

// GetCluster handles GET requests to return the state of the cluster
func GetCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	Name := vars[constants.ClusterNamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)

	model := cluster.InputModel{ProjectId: &projectId, ClusterName: &Name, PrivateKey: &privateKey, PublicKey: &publicKey}
	_, _ = logger.Debugf("Get clusters request : %s", model.String())
	startTime := time.Now()
	response := cluster.Read(&model)
	elapsedTime := time.Since(startTime)
	logger.Debugf("Get all Clusters REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())

	responseHandler.Write(response, w, constants.ClusterHandler)
}

// GetAllCluster handles GET requests to return all the clusters along with cluster's advanced configuration
func GetAllCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()
	vars := mux.Vars(r)
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)
	projectId := vars[constants.ProjectIdPathParam]
	model := cluster.InputModel{ProjectId: &projectId, PrivateKey: &privateKey, PublicKey: &publicKey}
	_, _ = logger.Debugf("Get all clusters request : %+v", model.String())
	startTime := time.Now()
	response := cluster.List(&model)
	elapsedTime := time.Since(startTime)
	logger.Debugf("Get all Clusters REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// DeleteCluster handles the DELETE requests to terminate the cluster
func DeleteCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()

	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	name := vars[constants.ClusterNamePathParam]
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)
	model := cluster.InputModel{ProjectId: &projectId, ClusterName: &name, PrivateKey: &privateKey, PublicKey: &publicKey}

	_, _ = logger.Debugf("Delete cluster request : %+v", model.String())
	startTime := time.Now()
	response := cluster.Delete(&model)
	elapsedTime := time.Since(startTime)
	logger.Debugf("Delete Cluster REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// CreateCluster handles the POST requests to create the cluster with the provided TshirtSize
func CreateCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()
	var model cluster.InputModel
	vars := mux.Vars(r)
	projectId := vars[constants.ProjectIdPathParam]
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.ProjectId = &projectId

	_, _ = logger.Debugf("Create cluster request : %+v", model.String())

	startTime := time.Now()
	response := cluster.Create(r.Context(), &model)
	elapsedTime := time.Since(startTime)
	logger.Debugf("Create Cluster REST API  response:%+v and execution time:%s", response.String(), elapsedTime.String())
	responseHandler.Write(response, w, constants.ClusterHandler)
}
