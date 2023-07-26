package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/custom_db_role"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
)

func setupCustomDBRole() {
	util.SetupLogger("atlas-api-helper.handlers.databaseuser")
}

// GetCustomDbRole handles GET requests to retrieve all custom db roles
func GetCustomDbRole(w http.ResponseWriter, r *http.Request) {
	setupLog()
	vars := mux.Vars(r)

	// Read a specific parameter
	groupId := vars[constants.GroupID]
	roleName := vars[constants.RoleName]
	response := custom_db_role.Read(r.Context(), &custom_db_role.Model{ProjectId: &groupId, RoleName: &roleName})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}

func GetAllCustomDbRoles(w http.ResponseWriter, r *http.Request) {
	setupLog()
	vars := mux.Vars(r)

	// Read a specific parameter
	groupId := vars[constants.GroupID]
	response := custom_db_role.List(r.Context(), &custom_db_role.Model{ProjectId: &groupId})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}

func DeleteCustomDbRoles(w http.ResponseWriter, r *http.Request) {
	setupLog()

	vars := mux.Vars(r)
	// Read a specific parameter
	groupId := vars[constants.GroupID]
	roleName := vars[constants.RoleName]
	response := custom_db_role.Delete(r.Context(), &custom_db_role.Model{ProjectId: &groupId, RoleName: &roleName})
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}

func CreateCustomDbRole(w http.ResponseWriter, r *http.Request) {
	setupLog()
	var model custom_db_role.Model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := custom_db_role.Create(r.Context(), &model)
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}

func UpdateCustomDbRole(w http.ResponseWriter, r *http.Request) {
	setupLog()
	var model custom_db_role.Model
	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := custom_db_role.Update(r.Context(), &model)
	responseHandler.Write(response, w, constants.DatabaseUserHandlerName)
	return
}
