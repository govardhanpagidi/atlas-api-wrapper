package project

import (
	"context"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasResponse"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"
	"go.mongodb.org/atlas/mongodbatlas"
	"net/http"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Name}
var UpdateRequiredFields = []string{constants.ID}

type UpdateAPIKey struct {
	Key     string
	APIKeys *mongodbatlas.AssignAPIKey
}

// validateModel inputs based on the method
func validateModel(fields []string, model *Model) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, currentModel *Model) atlasResponse.AtlasRespone {
	_, _ = logger.Debugf("Create currentModel: %+v", currentModel)

	if err := validateModel(CreateRequiredFields, currentModel); err != nil {
		_, _ = logger.Warnf("Validation Error")
		return atlasResponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "validation Error",
		}
	}
	client, err := util.NewMongoDBClient(ctx)
	if err != nil {
		_, _ = logger.Warnf("Mongo Atlas Connection Error")
		return atlasResponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "mongo Atlas Connection Error",
		}
	}

	var projectOwnerID string
	if currentModel.ProjectOwnerId != nil {
		projectOwnerID = *currentModel.ProjectOwnerId
	}
	project, res, err := client.Projects.Create(ctx, &mongodbatlas.Project{
		Name:                      *currentModel.Name,
		OrgID:                     *currentModel.OrgId,
		WithDefaultAlertsSettings: currentModel.WithDefaultAlertsSettings,
	}, &mongodbatlas.CreateProjectOptions{ProjectOwnerID: projectOwnerID})

	if err != nil {
		_, _ = logger.Debugf("Create - error: %+v", err)
		return atlasResponse.AtlasRespone{
			Response:       project,
			HttpStatusCode: res.Response.StatusCode,
			HttpError:      err.Error(),
		}
	}

	// Add ApiKeys
	if len(currentModel.ProjectApiKeys) > 0 {
		for _, key := range currentModel.ProjectApiKeys {
			_, err = client.ProjectAPIKeys.Assign(context.Background(), project.ID, *key.Key, &mongodbatlas.AssignAPIKey{Roles: key.RoleNames})
			if err != nil {
				_, _ = logger.Warnf("Assign Key Error: %s", err)
				return atlasResponse.AtlasRespone{
					Response:       nil,
					HttpStatusCode: http.StatusBadRequest,
					HttpError:      err.Error(),
				}
			}
		}
	}

	// Add Teams
	if len(currentModel.ProjectTeams) > 0 {
		_, _, err = client.Projects.AddTeamsToProject(context.Background(), project.ID, readTeams(currentModel.ProjectTeams))
		if err != nil {
			_, _ = logger.Warnf("AddTeamsToProject Error: %s", err.Error())
			return atlasResponse.AtlasRespone{
				Response:       nil,
				HttpStatusCode: http.StatusBadRequest,
				HttpError:      err.Error(),
			}
		}
	}

	currentModel.Id = &project.ID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount
	progressEvent, res, err := updateProjectSettings(currentModel, client)
	if err != nil {
		return atlasResponse.AtlasRespone{
			Response:       progressEvent,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}
	model, res, err := getProjectWithSettings(client, currentModel)
	if err != nil {
		_, _ = logger.Warnf("getProject Error: %s", err.Error())
		return atlasResponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}
	errMsg := getErrorMessage(err)

	return atlasResponse.AtlasRespone{
		Response:       model,
		HttpStatusCode: http.StatusOK,
		HttpError:      errMsg,
	}
}

func getErrorMessage(err error) string {
	errorMsg := ""
	if err != nil {
		errorMsg = err.Error()
	}
	return errorMsg
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, currentModel *Model) atlasResponse.AtlasRespone {
	client, err := util.NewMongoDBClient(ctx)

	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", err.Error())
		return atlasResponse.AtlasRespone{Response: nil, HttpStatusCode: http.StatusInternalServerError, HttpError: err.Error()}
	}
	var res *mongodbatlas.Response
	model, res, err := getProjectWithSettings(client, currentModel)
	return atlasResponse.AtlasRespone{Response: model, HttpStatusCode: res.Response.StatusCode, HttpError: getErrorMessage(err)}
}

func ReadAll(ctx context.Context) atlasResponse.AtlasRespone {
	client, err := util.NewMongoDBClient(ctx)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", err.Error())
		return atlasResponse.AtlasRespone{Response: nil, HttpStatusCode: http.StatusInternalServerError, HttpError: err.Error()}
	}
	projects, res, err := ReadAllProjects(ctx, client)
	if err != nil {
		return atlasResponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.Response.StatusCode,
			HttpError:      err.Error(),
		}
	}
	return atlasResponse.AtlasRespone{
		Response:       projects,
		HttpStatusCode: res.Response.StatusCode,
		HttpError:      getErrorMessage(err),
	}
}

func ReadAllProjects(ctx context.Context, client *mongodbatlas.Client) (projectsToReturn []*mongodbatlas.Project, res *mongodbatlas.Response, err error) {
	projects, apiRes, err := client.Projects.GetAllProjects(ctx, nil)
	if err != nil {
		return nil, apiRes, err
	}
	for _, project := range projects.Results {
		projectsToReturn = append(projectsToReturn, project)
	}
	return projectsToReturn, apiRes, err
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(ctx context.Context, currentModel *Model) atlasResponse.AtlasRespone {
	ctx = context.WithValue(ctx, "", "")
	client, err := util.NewMongoDBClient(ctx)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", err.Error())
		return atlasResponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}
	var id string
	if currentModel.Id != nil {
		id = *currentModel.Id
	}

	var res *mongodbatlas.Response
	_, res, err = getProject(client, currentModel)
	if err != nil {
		_, _ = logger.Warnf("GetProjectError error: %v", err.Error())
		return atlasResponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.Response.StatusCode,
			HttpError:      err.Error(),
		}
	}
	_, _ = logger.Debugf("Deleting project with id(%s)", id)

	res, err = client.Projects.Delete(context.Background(), id)

	return atlasResponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: res.Response.StatusCode,
		HttpError:      getErrorMessage(err),
	}
}

func updateProjectSettings(currentModel *Model, client *mongodbatlas.Client) (updatedModel *Model, res *mongodbatlas.Response, err error) {
	if currentModel.ProjectSettings != nil {
		// Update project settings
		projectSettings := mongodbatlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: currentModel.ProjectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
			IsRealtimePerformancePanelEnabled:           currentModel.ProjectSettings.IsRealtimePerformancePanelEnabled,
			IsDataExplorerEnabled:                       currentModel.ProjectSettings.IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 currentModel.ProjectSettings.IsPerformanceAdvisorEnabled,
			IsSchemaAdvisorEnabled:                      currentModel.ProjectSettings.IsSchemaAdvisorEnabled,
			IsExtendedStorageSizesEnabled:               currentModel.ProjectSettings.IsExtendedStorageSizesEnabled,
		}

		_, res, err := client.Projects.UpdateProjectSettings(context.Background(), *currentModel.Id, &projectSettings)
		if err != nil {
			return currentModel, res, err
		}
	}
	return currentModel, res, nil
}

// Read project
func getProject(client *mongodbatlas.Client, currentModel *Model) (model *Model, res *mongodbatlas.Response, err error) {
	var project *mongodbatlas.Project
	if currentModel.Name != nil && len(*currentModel.Name) > 0 {
		project, res, err = getProjectByName(currentModel.Name, client)
	} else {
		project, res, err = getProjectByID(currentModel.Id, client)
	}
	if err != nil {
		return nil, res, err
	}
	currentModel.Name = &project.Name
	currentModel.OrgId = &project.OrgID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount
	currentModel.Id = &project.ID

	return currentModel, res, nil
}

// Read project
func getProjectWithSettings(client *mongodbatlas.Client, currentModel *Model) (model *Model, res *mongodbatlas.Response, err error) {
	currentModel, res, err = getProject(client, currentModel)
	if err != nil {
		return currentModel, res, err
	}
	model, res, err = readProjectSettings(client, *currentModel.Id, currentModel)

	return model, res, err
}

func getProjectByName(name *string, client *mongodbatlas.Client) (model *mongodbatlas.Project, res *mongodbatlas.Response, err error) {
	project, res, err := client.Projects.GetOneProjectByName(context.Background(), *name)
	return project, res, err
}

func getProjectByID(id *string, client *mongodbatlas.Client) (model *mongodbatlas.Project, res *mongodbatlas.Response, err error) {
	project, res, err := client.Projects.GetOneProject(context.Background(), *id)
	return project, res, err
}

func readProjectSettings(client *mongodbatlas.Client, id string, currentModel *Model) (model *Model, res *mongodbatlas.Response, err error) {
	// Get teams from project
	teamsAssigned, res, err := client.Projects.GetProjectTeamsAssigned(context.Background(), id)
	if err != nil {
		_, _ = logger.Warnf("ProjectId : %s, Error: %s", id, err)
		return nil, res, err
	}

	projectSettings, res, err := client.Projects.GetProjectSettings(context.Background(), id)
	if err != nil {
		_, _ = logger.Warnf("ProjectId : %s, Error: %s", id, err)
		return nil, res, err
	}
	// Set projectSettings
	currentModel.ProjectSettings = &ProjectSettings{
		IsCollectDatabaseSpecificsStatisticsEnabled: projectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
		IsRealtimePerformancePanelEnabled:           projectSettings.IsRealtimePerformancePanelEnabled,
		IsDataExplorerEnabled:                       projectSettings.IsDataExplorerEnabled,
		IsPerformanceAdvisorEnabled:                 projectSettings.IsPerformanceAdvisorEnabled,
		IsSchemaAdvisorEnabled:                      projectSettings.IsSchemaAdvisorEnabled,
		IsExtendedStorageSizesEnabled:               projectSettings.IsExtendedStorageSizesEnabled,
	}

	// Set teams
	var teams []ProjectTeam
	for _, team := range teamsAssigned.Results {
		if len(team.TeamID) > 0 {
			teams = append(teams, ProjectTeam{TeamId: &team.TeamID, RoleNames: team.RoleNames})
		}
	}

	currentModel.ProjectTeams = teams
	currentModel.ProjectApiKeys = nil // hack: cfn test. Extra APIKey(default) getting added and cfn test fails.
	return currentModel, res, err
}

// Get difference in Teams
func getChangeInTeams(currentTeams []ProjectTeam, oTeams []*mongodbatlas.Result) (newTeams []*mongodbatlas.ProjectTeam,
	changedTeams []*mongodbatlas.ProjectTeam, removeTeams []*mongodbatlas.ProjectTeam) {
	for _, nTeam := range currentTeams {
		if nTeam.TeamId != nil && len(*nTeam.TeamId) > 0 {
			matched := false
			for _, oTeam := range oTeams {
				if nTeam.TeamId != nil && *nTeam.TeamId == oTeam.TeamID {
					changedTeams = append(changedTeams, &mongodbatlas.ProjectTeam{TeamID: *nTeam.TeamId, RoleNames: nTeam.RoleNames})
					matched = true
					break
				}
			}
			// Add to newTeams
			if !matched {
				newTeams = append(newTeams, &mongodbatlas.ProjectTeam{TeamID: *nTeam.TeamId, RoleNames: nTeam.RoleNames})
			}
		}
	}

	for _, oTeam := range oTeams {
		if len(oTeam.TeamID) > 0 {
			matched := false
			for _, nTeam := range currentTeams {
				if nTeam.TeamId != nil && *nTeam.TeamId == oTeam.TeamID {
					matched = true
					break
				}
			}
			if !matched {
				removeTeams = append(removeTeams, &mongodbatlas.ProjectTeam{TeamID: oTeam.TeamID, RoleNames: oTeam.RoleNames})
			}
		}
	}
	return newTeams, changedTeams, removeTeams
}

// Get difference in ApiKeys
func getChangeInAPIKeys(groupID string, currentKeys []ProjectApiKey, oKeys []mongodbatlas.APIKey) (newKeys, changedKeys, removeKeys []UpdateAPIKey) {
	for _, nKey := range currentKeys {
		if nKey.Key != nil && len(*nKey.Key) > 0 {
			matched := false
			for _, oKey := range oKeys {
				if nKey.Key != nil && *nKey.Key == oKey.ID {
					changedKeys = append(changedKeys, UpdateAPIKey{Key: *nKey.Key, APIKeys: &mongodbatlas.AssignAPIKey{Roles: nKey.RoleNames}})
					matched = true
					break
				}
			}
			// Add to newKeys
			if !matched {
				newKeys = append(newKeys, UpdateAPIKey{Key: *nKey.Key, APIKeys: &mongodbatlas.AssignAPIKey{Roles: nKey.RoleNames}})
			}
		}
	}

	for _, oKey := range oKeys {
		if len(oKey.ID) > 0 {
			matched := false
			for _, nKey := range currentKeys {
				if nKey.Key != nil && *nKey.Key == oKey.ID {
					matched = true
					break
				}
			}
			if !matched {
				for _, role := range oKey.Roles {
					// Consider only current ProjectRoles
					if role.GroupID == groupID {
						removeKeys = append(removeKeys, UpdateAPIKey{Key: oKey.ID})
					}
				}
			}
		}
	}
	return newKeys, changedKeys, removeKeys
}

func readTeams(teams []ProjectTeam) []*mongodbatlas.ProjectTeam {
	var newTeams []*mongodbatlas.ProjectTeam
	for _, t := range teams {
		if t.TeamId != nil && len(*t.TeamId) > 0 {
			newTeams = append(newTeams, &mongodbatlas.ProjectTeam{TeamID: *t.TeamId, RoleNames: t.RoleNames})
		}
	}
	return newTeams
}

func Update(ctx context.Context, currentModel *Model) atlasResponse.AtlasRespone {

	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		_, _ = logger.Warnf("Validation Error")
		return atlasResponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "validation Error",
		}
	}

	client, pe := util.NewMongoDBClient(ctx)
	if pe != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", pe)
		return atlasResponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      pe.Error(),
		}
	}

	var projectID string
	if currentModel.Id != nil {
		projectID = *currentModel.Id
	}
	currentModel, res, err := getProject(client, currentModel)
	if err != nil {
		return atlasResponse.AtlasRespone{
			Response:       currentModel,
			HttpStatusCode: res.Response.StatusCode,
			HttpError:      err.Error(),
		}
	}

	if currentModel.ProjectTeams != nil {
		// Get teams from project
		teamsAssigned, res, errr := client.Projects.GetProjectTeamsAssigned(context.Background(), projectID)
		if errr != nil {
			_, _ = logger.Warnf("ProjectId : %s, Error: %s", projectID, errr)
			return atlasResponse.AtlasRespone{
				Response:       teamsAssigned,
				HttpStatusCode: res.Response.StatusCode,
				HttpError:      errr.Error(),
			}
		}
		newTeams, changedTeams, removeTeams := getChangeInTeams(currentModel.ProjectTeams, teamsAssigned.Results)

		// Remove Teams
		for _, team := range removeTeams {
			res, err = client.Teams.RemoveTeamFromProject(context.Background(), projectID, team.TeamID)
			if err != nil {
				_, _ = logger.Warnf("ProjectId : %s, Error: %s", projectID, err)
				return atlasResponse.AtlasRespone{
					Response:       nil,
					HttpStatusCode: res.Response.StatusCode,
					HttpError:      err.Error(),
				}
			}
		}
		// Add Teams
		if len(newTeams) > 0 {
			teamsAssigned, res, err = client.Projects.AddTeamsToProject(context.Background(), projectID, newTeams)
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return atlasResponse.AtlasRespone{
					Response:       teamsAssigned,
					HttpStatusCode: res.Response.StatusCode,
					HttpError:      err.Error(),
				}
			}
		}
		// Update Teams
		for _, team := range changedTeams {
			_, res, err = client.Teams.UpdateTeamRoles(context.Background(), projectID, team.TeamID, &mongodbatlas.TeamUpdateRoles{RoleNames: team.RoleNames})
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return atlasResponse.AtlasRespone{
					Response:       nil,
					HttpStatusCode: res.Response.StatusCode,
					HttpError:      err.Error(),
				}
			}
		}
	}

	if currentModel.ProjectApiKeys != nil {
		// Get APIKeys from project
		projectAPIKeys, res, errr := client.ProjectAPIKeys.List(context.Background(), projectID, &mongodbatlas.ListOptions{ItemsPerPage: 1000, IncludeCount: true})
		if err != nil {
			_, _ = logger.Warnf("ProjectId : %s, Error: %s", projectID, errr)
			return atlasResponse.AtlasRespone{
				Response:       nil,
				HttpStatusCode: res.Response.StatusCode,
				HttpError:      errr.Error(),
			}
		}

		// Get Change in ApiKeys
		newAPIKeys, changedKeys, removeKeys := getChangeInAPIKeys(*currentModel.Id, currentModel.ProjectApiKeys, projectAPIKeys)

		// Remove old keys
		for _, key := range removeKeys {
			res, err = client.ProjectAPIKeys.Unassign(context.Background(), projectID, key.Key)
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return atlasResponse.AtlasRespone{
					Response:       nil,
					HttpStatusCode: res.Response.StatusCode,
					HttpError:      err.Error(),
				}
			}
		}

		// Add Keys
		for _, key := range newAPIKeys {
			res, err = client.ProjectAPIKeys.Assign(context.Background(), projectID, key.Key, key.APIKeys)
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return atlasResponse.AtlasRespone{
					Response:       nil,
					HttpStatusCode: res.Response.StatusCode,
					HttpError:      err.Error(),
				}
			}
		}

		// Update Key Roles
		for _, key := range changedKeys {
			res, err = client.ProjectAPIKeys.Assign(context.Background(), projectID, key.Key, key.APIKeys)
			if err != nil {
				_, _ = logger.Warnf("Error: %s", err)
				return atlasResponse.AtlasRespone{
					Response:       nil,
					HttpStatusCode: res.Response.StatusCode,
					HttpError:      err.Error(),
				}
			}
		}
	}

	if currentModel.ProjectSettings != nil {
		// Update project settings
		projectSettings := mongodbatlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: currentModel.ProjectSettings.IsCollectDatabaseSpecificsStatisticsEnabled,
			IsRealtimePerformancePanelEnabled:           currentModel.ProjectSettings.IsRealtimePerformancePanelEnabled,
			IsDataExplorerEnabled:                       currentModel.ProjectSettings.IsDataExplorerEnabled,
			IsPerformanceAdvisorEnabled:                 currentModel.ProjectSettings.IsPerformanceAdvisorEnabled,
			IsSchemaAdvisorEnabled:                      currentModel.ProjectSettings.IsSchemaAdvisorEnabled,
		}
		_, res, err = client.Projects.UpdateProjectSettings(context.Background(), projectID, &projectSettings)
		if err != nil {
			_, _ = logger.Warnf("Update - error: %+v", err)
			return atlasResponse.AtlasRespone{
				Response:       nil,
				HttpStatusCode: res.Response.StatusCode,
				HttpError:      err.Error(),
			}
		}
	}

	toRet, res, err := getProjectWithSettings(client, currentModel)
	if err != nil {
		return atlasResponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.Response.StatusCode,
			HttpError:      err.Error(),
		}
	}

	return atlasResponse.AtlasRespone{
		Response:       toRet,
		HttpStatusCode: res.Response.StatusCode,
		HttpError:      getErrorMessage(err),
	}
}
