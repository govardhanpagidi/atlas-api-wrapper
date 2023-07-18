package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/atlas-api-helper/resources/project"
	"github.com/atlas-api-helper/util/constants"
)

// GetProjectHandler handles GET requests to retrieve all projects
func GetProjectHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// Read a specific parameter
	projectID := queryParams.Get(constants.ID)

	// Use the parameters as needed
	//TODO: return the http status code from resource
	response, err := project.Read(r.Context(), &project.Model{Id: &projectID})
	if err != nil {
		fmt.Println("CreateProjectHandler error:", err)
	}
	res, err := json.Marshal(response)
	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = w.Write(res)
	return
}

// DeleteProjectHandler handles GET requests to retrieve all projects
func DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	// Read a specific parameter
	projectID := queryParams.Get(constants.ID)

	// Use the parameters as needed
	//TODO: return the http status code from resource
	err := project.Delete(r.Context(), &project.Model{Id: &projectID})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("CreateProjectHandler error:", err)
		return
	}
	return
}

// CreateProjectHandler handles POST requests to create a new project
func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {

	var model project.Model
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//TODO: return the http status code from resource
	response, err := project.Create(r.Context(), &model)
	if err != nil {
		fmt.Println("CreateProjectHandler error:", err)
	}
	res, err := json.Marshal(response)
	_, err = w.Write(res)
	return
}
