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
	publicKey := r.URL.Query().Get("publicKey")
	privateKey := r.URL.Query().Get("privateKey")
	response := cluster.Read(&cluster.InputModel{ProjectId: &groupId, ClusterName: &Name, PrivateKey: &privateKey, PublicKey: &publicKey})
	responseHandler.Write(response, w, constants.ClusterHandler)
	return
}

/*
	func GetAllCluster(w http.ResponseWriter, r *http.Request) {
		setupClusterLog()
		vars := mux.Vars(r)

		// Read a specific parameter
		groupId := vars[constants.GroupID]
		response := cluster.List(r.Context(), &cluster.Model{ProjectId: &groupId})
		responseHandler.Write(response, w, constants.ClusterHandler)
		return
	}
*/
func DeleteCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()

	vars := mux.Vars(r)
	// Read a specific parameter
	groupId := vars[constants.GroupID]
	Name := vars[constants.Name]
	publicKey := r.URL.Query().Get("publicKey")
	privateKey := r.URL.Query().Get("privateKey")
	response := cluster.Delete(&cluster.InputModel{ProjectId: &groupId, ClusterName: &Name, PrivateKey: &privateKey, PublicKey: &publicKey})
	responseHandler.Write(response, w, constants.ClusterHandler)
	return
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
	return
}

/*func UpdateCluster(w http.ResponseWriter, r *http.Request) {
	setupClusterLog()
	var model cluster.Model
	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := cluster.Update(r.Context(), &model)
	responseHandler.Write(response, w, constants.ClusterHandler)
	return
}*/
