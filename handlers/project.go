package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/resources/project"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/ResponseHandler"
	"github.com/atlas-api-helper/util/constants"
	"github.com/gorilla/mux"
	"net/http"
)

func setupLog() {
	util.SetupLogger("atlas-api-helper.handlers.project")
}

// GetProject handles GET requests to retrieve all projects
func GetProject(w http.ResponseWriter, r *http.Request) {
	setupLog()
	vars := mux.Vars(r)

	// Read a specific parameter
	projectID := vars[constants.ID]

	response := project.Read(r.Context(), &project.Model{Id: &projectID})
	responseHandler.CommonResponseHandler(response, w, "ProjectHandler")
	return
}

// GetAllProjects handles GET requests to retrieve all projects
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	setupLog()

	// Use the parameters as needed
	response := project.ReadAll(r.Context())
	responseHandler.CommonResponseHandler(response, w, "ProjectHandler")
	return
}

// DeleteProject handles GET requests to retrieve all projects
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	setupLog()
	queryParams := r.URL.Query()

	// Read a specific parameter
	projectID := queryParams.Get(constants.ID)
	response := project.Delete(r.Context(), &project.Model{Id: &projectID})
	responseHandler.CommonResponseHandler(response, w, "ProjectHandler")
	return
}

// CreateProject handles POST requests to create a new project
func CreateProject(w http.ResponseWriter, r *http.Request) {
	setupLog()
	var model project.Model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := project.Create(r.Context(), &model)
	responseHandler.CommonResponseHandler(response, w, "ProjectHandler")
	return
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	setupLog()
	var model project.Model
	err := json.NewDecoder(r.Body).Decode(&model)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := project.Update(r.Context(), &model)
	responseHandler.CommonResponseHandler(response, w, "ProjectHandler")
	return
}
