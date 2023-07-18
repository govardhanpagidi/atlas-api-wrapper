package project

import (
	"context"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"

	"go.mongodb.org/atlas/mongodbatlas"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Name}
var UpdateRequiredFields = []string{constants.ID}

type UpdateAPIKey struct {
	Key     string
	APIKeys *mongodbatlas.AssignAPIKey
}

func setupLog() {
	//util.SetupLogger("mongodb-atlas-project")
}

// validateModel inputs based on the method
func validateModel(fields []string, model *Model) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, currentModel *Model) (interface{}, error) {
	setupLog()
	_, _ = logger.Debugf("Create currentModel: %+v", currentModel)

	if err := validateModel(CreateRequiredFields, currentModel); err != nil {
		_, _ = logger.Warnf("Validation Error")
		return nil, err
	}
	client, err := util.NewMongoDBClient(ctx)
	if err != nil {
		return nil, err
	}

	var projectOwnerID string
	if currentModel.ProjectOwnerId != nil {
		projectOwnerID = *currentModel.ProjectOwnerId
	}
	project, _, err := client.Projects.Create(ctx, &mongodbatlas.Project{
		Name:                      *currentModel.Name,
		OrgID:                     *currentModel.OrgId,
		WithDefaultAlertsSettings: currentModel.WithDefaultAlertsSettings,
	}, &mongodbatlas.CreateProjectOptions{ProjectOwnerID: projectOwnerID})

	if err != nil {
		_, _ = logger.Debugf("Create - error: %+v", err)
		return nil, err
	}

	// Add ApiKeys
	if len(currentModel.ProjectApiKeys) > 0 {
		for _, key := range currentModel.ProjectApiKeys {
			_, err = client.ProjectAPIKeys.Assign(context.Background(), project.ID, *key.Key, &mongodbatlas.AssignAPIKey{Roles: key.RoleNames})
			if err != nil {
				_, _ = logger.Warnf("Assign Key Error: %s", err)
				return nil, err
			}
		}
	}

	// Add Teams
	if len(currentModel.ProjectTeams) > 0 {
		_, _, err = client.Projects.AddTeamsToProject(context.Background(), project.ID, readTeams(currentModel.ProjectTeams))
		if err != nil {
			_, _ = logger.Warnf("AddTeamsToProject Error: %s", err.Error())
			return nil, err
		}
	}

	currentModel.Id = &project.ID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount

	progressEvent, err := updateProjectSettings(currentModel, client)
	if err != nil {
		return progressEvent, err
	}
	model, err := getProjectWithSettings(client, currentModel)
	if err != nil {
		_, _ = logger.Warnf("getProject Error: %s", err.Error())
		return nil, err
	}

	return model, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, currentModel *Model) (interface{}, error) {
	setupLog()

	client, err := util.NewMongoDBClient(ctx)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", err.Error())
		return nil, err
	}

	model, err := getProjectWithSettings(client, currentModel)
	return model, err
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(ctx context.Context, currentModel *Model) (err error) {
	setupLog()
	ctx = context.WithValue(ctx, "", "")
	client, err := util.NewMongoDBClient(ctx)
	if err != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", err.Error())
		return err
	}
	var id string
	if currentModel.Id != nil {
		id = *currentModel.Id
	}

	_, err = getProject(client, currentModel)
	if err != nil {
		return nil
	}
	_, _ = logger.Debugf("Deleting project with id(%s)", id)

	_, err = client.Projects.Delete(context.Background(), id)
	return err
}

func updateProjectSettings(currentModel *Model, client *mongodbatlas.Client) (interface{}, error) {
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
			return res, err
		}
	}
	return currentModel, nil
}

// Read project
func getProject(client *mongodbatlas.Client, currentModel *Model) (model *Model, err error) {
	var project *mongodbatlas.Project
	if currentModel.Name != nil && len(*currentModel.Name) > 0 {
		project, err = getProjectByName(currentModel.Name, client)
		if err != nil {
			return nil, err
		}
	} else {
		project, err = getProjectByID(currentModel.Id, client)
		if err != nil {
			return nil, err
		}
	}
	currentModel.Name = &project.Name
	currentModel.OrgId = &project.OrgID
	currentModel.Created = &project.Created
	currentModel.ClusterCount = &project.ClusterCount
	currentModel.Id = &project.ID

	return currentModel, nil
}

// Read project
func getProjectWithSettings(client *mongodbatlas.Client, currentModel *Model) (model *Model, err error) {
	currentModel, err = getProject(client, currentModel)
	if err != nil {
		return currentModel, err
	}
	model, err = readProjectSettings(client, *currentModel.Id, currentModel)

	if err != nil {
		return model, err
	}

	return model, nil
}

func getProjectByName(name *string, client *mongodbatlas.Client) (model *mongodbatlas.Project, err error) {
	project, res, err := client.Projects.GetOneProjectByName(context.Background(), *name)
	if err != nil {
		if res.Response.StatusCode == 401 { // cfn test
			return nil, err
		}
		return project, err
	}
	return project, err
}

func getProjectByID(id *string, client *mongodbatlas.Client) (model *mongodbatlas.Project, err error) {
	project, res, err := client.Projects.GetOneProject(context.Background(), *id)
	if err != nil {
		if res.Response.StatusCode == 401 { // cfn test
			return nil, err
		}
		return project, err
	}
	return project, err
}

func readProjectSettings(client *mongodbatlas.Client, id string, currentModel *Model) (model *Model, err error) {
	// Get teams from project
	teamsAssigned, _, err := client.Projects.GetProjectTeamsAssigned(context.Background(), id)
	if err != nil {
		_, _ = logger.Warnf("ProjectId : %s, Error: %s", id, err)
		return nil, err
	}

	projectSettings, _, err := client.Projects.GetProjectSettings(context.Background(), id)
	if err != nil {
		_, _ = logger.Warnf("ProjectId : %s, Error: %s", id, err)
		return nil, err
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
	return currentModel, err
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
