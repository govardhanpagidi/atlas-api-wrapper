package handlers

import (
	"encoding/json"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/logger"
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
	queryParams := r.URL.Query()

	// Read a specific parameter
	projectID := queryParams.Get(constants.ID)

	// Use the parameters as needed
	//TODO: return the http status code from resource
	response, err := project.Read(r.Context(), &project.Model{Id: &projectID})
	if err != nil {
		_, _ = logger.Debugf("CreateProjectHandler error:%s", err.Error())
		return
	}
	res, err := json.Marshal(response)
	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write(res)
	return
}

// DeleteProject handles GET requests to retrieve all projects
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	setupLog()
	queryParams := r.URL.Query()

	// Read a specific parameter
	projectID := queryParams.Get(constants.ID)

	// Use the parameters as needed
	//TODO: return the http status code from resource
	err := project.Delete(r.Context(), &project.Model{Id: &projectID})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = logger.Debugf("CreateProjectHandler error: %s", err.Error())
		return
	}
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
	//TODO: return the http status code from resource
	response, err := project.Create(r.Context(), &model)
	if err != nil {
		_, _ = logger.Debugf("CreateProjectHandler error:%s", err.Error())
		return
	}
	res, err := json.Marshal(response)
	_, err = w.Write(res)
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
	response, err := project.Update(r.Context(), &model)
	if err != nil {
		_, _ = logger.Debugf("CreateProjectHandler error:%s", err.Error())
		return
	}
	res, err := json.Marshal(response)
	_, err = w.Write(res)
	return
}
