package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/cluster"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
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
	response := cluster.Read(&cluster.InputModel{ProjectId: &projectId, ClusterName: &Name, PrivateKey: &privateKey, PublicKey: &publicKey})
	responseHandler.Write(response, w, constants.ClusterHandler)
}

// GetAllCluster handles GET requests to return all the clusters along with cluster's advanced configuration
func GetAllCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()
	vars := mux.Vars(r)
	publicKey := r.URL.Query().Get(constants.PublicKeyQueryParam)
	privateKey := r.URL.Query().Get(constants.PrivateKeyQueryParam)
	projectId := vars[constants.ProjectIdPathParam]
	response := cluster.List(&cluster.InputModel{ProjectId: &projectId, PrivateKey: &privateKey, PublicKey: &publicKey})
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
	response := cluster.Delete(&cluster.InputModel{ProjectId: &projectId, ClusterName: &name, PrivateKey: &privateKey, PublicKey: &publicKey})
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
	response := cluster.Create(r.Context(), &model)
	responseHandler.Write(response, w, constants.ClusterHandler)
}
