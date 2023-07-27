package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/project_invitation"
	"github.com/atlas-api-helper/util"
	responseHandler "github.com/atlas-api-helper/util/Responsehandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
)

func setupProjectInvitation() {
	util.SetupLogger("atlas-api-helper.handlers.databaseuser")
}

// GetProjectInvitation handles GET requests to retrieve all GetProjectInvitation
func GetProjectInvitation(w http.ResponseWriter, r *http.Request) {
	setupLog()
	vars := mux.Vars(r)
	// Read a specific parameter
	groupId := vars[constants.GroupID]
	inviteId := vars[constants.InvitationID]
	response := project_invitation.Read(r.Context(), &project_invitation.Model{ProjectId: &groupId, Id: &inviteId})
	responseHandler.Write(response, w, constants.ProjectInviteHandlerName)
	return
}

func GetAllprojectInvites(w http.ResponseWriter, r *http.Request) {
	setupLog()
	vars := mux.Vars(r)

	// Read a specific parameter
	groupId := vars[constants.GroupID]
	response := project_invitation.List(r.Context(), &project_invitation.Model{ProjectId: &groupId})
	responseHandler.Write(response, w, constants.ProjectInviteHandlerName)
	return
}

func DeleteProjectInvites(w http.ResponseWriter, r *http.Request) {
	setupLog()

	vars := mux.Vars(r)
	// Read a specific parameter
	groupId := vars[constants.GroupID]
	inviteId := vars[constants.InvitationID]
	response := project_invitation.Delete(r.Context(), &project_invitation.Model{ProjectId: &groupId, Id: &inviteId})
	responseHandler.Write(response, w, constants.ProjectInviteHandlerName)
	return
}

func CreateProjectInvite(w http.ResponseWriter, r *http.Request) {
	setupLog()
	var model project_invitation.Model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := project_invitation.Create(r.Context(), &model)
	responseHandler.Write(response, w, constants.ProjectInviteHandlerName)
	return
}

func UpdateProjectInvite(w http.ResponseWriter, r *http.Request) {
	setupLog()
	var model project_invitation.Model
	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := project_invitation.Update(r.Context(), &model)
	responseHandler.Write(response, w, constants.ProjectInviteHandlerName)
	return
}
