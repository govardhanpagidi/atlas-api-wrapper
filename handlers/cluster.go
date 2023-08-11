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

func GetCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()
	vars := mux.Vars(r)
	groupId := vars[constants.GroupID]
	Name := vars[constants.Name]
	publicKey := r.URL.Query().Get(constants.PublicKey)
	privateKey := r.URL.Query().Get(constants.PrivateKey)
	response := cluster.Read(&cluster.InputModel{ProjectId: &groupId, ClusterName: &Name, PrivateKey: &privateKey, PublicKey: &publicKey})
	responseHandler.Write(response, w, constants.ClusterHandler)
}

func GetAllCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()
	vars := mux.Vars(r)
	publicKey := r.URL.Query().Get(constants.PublicKey)
	privateKey := r.URL.Query().Get(constants.PrivateKey)
	groupId := vars[constants.GroupID]
	response := cluster.List(&cluster.InputModel{ProjectId: &groupId, PrivateKey: &privateKey, PublicKey: &publicKey})
	responseHandler.Write(response, w, constants.ClusterHandler)
}

func DeleteCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()

	vars := mux.Vars(r)
	groupId := vars[constants.GroupID]
	Name := vars[constants.Name]
	publicKey := r.URL.Query().Get(constants.PublicKey)
	privateKey := r.URL.Query().Get(constants.PrivateKey)
	response := cluster.Delete(&cluster.InputModel{ProjectId: &groupId, ClusterName: &Name, PrivateKey: &privateKey, PublicKey: &publicKey})
	responseHandler.Write(response, w, constants.ClusterHandler)
}

func CreateCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()
	var model cluster.InputModel

	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := cluster.Create(r.Context(), &model)
	responseHandler.Write(response, w, constants.ClusterHandler)
}
