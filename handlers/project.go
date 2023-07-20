package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/logger"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/atlas-api-helper/resources/project"
	"github.com/atlas-api-helper/util/constants"
)

func setupLog() {
	util.SetupLogger("atlas-api-helper.handlers.project")
}

// GetProject handles GET requests to retrieve all projects
func GetProject(w http.ResponseWriter, r *http.Request) {
	setupLog()
	vars := mux.Vars(r)

	// Read a specific parameter
	projectID := vars["Id"]

	// Use the parameters as needed
	//TODO: return the http status code from resource
	response := project.Read(r.Context(), &project.Model{Id: &projectID})
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(response)
	if response.HttpError != "" {
		_, _ = logger.Debugf("CreateProjectHandler error:%s", response.HttpError)
		//http.Error(w, response.HttpError, http.StatusNotFound)
		_, _ = w.Write(res)
		return
	}

	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, _ = w.Write(res)
	return
}

// GetAllProjects handles GET requests to retrieve all projects
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	setupLog()

	// Use the parameters as needed
	response := project.ReadAll(r.Context())
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(response)
	if response.HttpError != "" {
		_, _ = logger.Debugf("GetAllProjects error:%s", response.HttpError)
		//http.Error(w, response.HttpError, http.StatusNotFound)
		_, _ = w.Write(res)
		return
	}

	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, _ = w.Write(res)
	return
}

// DeleteProject handles GET requests to retrieve all projects
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	setupLog()
	queryParams := r.URL.Query()

	// Read a specific parameter
	projectID := queryParams.Get(constants.ID)

	w.Header().Set("Content-Type", "application/json")
	// Use the parameters as needed
	//TODO: return the http status code from resource
	err := project.Delete(r.Context(), &project.Model{Id: &projectID})
	if err.HttpError != "" {
		w.WriteHeader(http.StatusNotFound)
		_, _ = logger.Debugf("CreateProjectHandler error: %s", err.HttpError)
	}
	res, _ := json.Marshal(err)
	_, _ = w.Write(res)
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
	w.Header().Set("Content-Type", "application/json")

	//TODO: return the http status code from resource
	response := project.Create(r.Context(), &model)
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}
	res, _ := json.Marshal(response)
	if response.HttpError != "" {
		_, _ = logger.Debugf("CreateProjectHandler error:%s", errorMsg)
		_, err = w.Write(res)
		return
	}

	_, err = w.Write(res)
	return
}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	setupLog()
	var model project.Model
	err := json.NewDecoder(r.Body).Decode(&model)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := project.Update(r.Context(), &model)
	res, err := json.Marshal(response)
	_, err = w.Write(res)
	return
}
