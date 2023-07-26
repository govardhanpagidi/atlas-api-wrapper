// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package custom_db_role

import (
	"context"
	"github.com/atlas-api-helper/resources/profile"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"
	"github.com/aws/aws-sdk-go/aws"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
	"net/http"
)

func setup() {
	util.SetupLogger("mongodb-atlas-custom-db-role")
}

var CreateRequiredFields = []string{constants.ProjectID, constants.RoleName}
var ReadRequiredFields = []string{constants.ProjectID, constants.RoleName}
var UpdateRequiredFields = []string{constants.ProjectID, constants.RoleName}
var DeleteRequiredFields = []string{constants.ProjectID, constants.RoleName}
var ListRequiredFields = []string{constants.ProjectID}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", currentModel)
	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	atlasCustomDBRole := currentModel.ToCustomDBRole()
	customDBRole, response, err := client.CustomDatabaseRolesApi.CreateCustomDatabaseRole(ctx, *currentModel.ProjectId, atlasCustomDBRole).Execute()
	if err != nil {
		_, _ = logger.Warnf("CreateCustomDatabaseRole error: %v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: response.StatusCode,
			HttpError:      err.Error(),
		}
	}

	currentModel.completeByAtlasRole(*customDBRole)

	return atlasresponse.AtlasRespone{
		Response:       currentModel,
		HttpStatusCode: response.StatusCode,
		HttpError:      "",
	}
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", currentModel)
	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	atlasCustomDdRole, response, err := client.CustomDatabaseRolesApi.GetCustomDatabaseRole(ctx, *currentModel.ProjectId, *currentModel.RoleName).Execute()
	if err != nil {
		_, _ = logger.Warnf("GetCustomDatabaseRole error: %v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: response.StatusCode,
			HttpError:      err.Error(),
		}
	}

	currentModel.completeByAtlasRole(*atlasCustomDdRole)

	return atlasresponse.AtlasRespone{
		Response:       currentModel,
		HttpStatusCode: response.StatusCode,
		HttpError:      "",
	}
}

// Update handles the Update event from the Cloudformation service.
func Update(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", currentModel)
	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	var actions []admin.DatabasePrivilegeAction
	for _, a := range currentModel.Actions {
		actions = append(actions, a.toAtlasAction())
	}

	var inheritedRoles []admin.DatabaseInheritedRole
	for _, ir := range currentModel.InheritedRoles {
		inheritedRoles = append(inheritedRoles, ir.toAtlasInheritedRole())
	}

	inputCustomDBRole := admin.UpdateCustomDBRole{
		Actions:        actions,
		InheritedRoles: inheritedRoles,
	}

	atlasCustomDdRole, response, err := client.CustomDatabaseRolesApi.UpdateCustomDatabaseRole(ctx, *currentModel.ProjectId, *currentModel.RoleName, &inputCustomDBRole).Execute()

	if err != nil {
		_, _ = logger.Warnf("UpdateCustomDatabaseRole error: %v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: response.StatusCode,
			HttpError:      err.Error(),
		}
	}

	currentModel.completeByAtlasRole(*atlasCustomDdRole)

	return atlasresponse.AtlasRespone{
		Response:       currentModel,
		HttpStatusCode: response.StatusCode,
		HttpError:      "",
	}
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", currentModel)
	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	response, err := client.CustomDatabaseRolesApi.DeleteCustomDatabaseRole(ctx, *currentModel.ProjectId, *currentModel.RoleName).Execute()

	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: response.StatusCode,
			HttpError:      err.Error(),
		}
	}

	return atlasresponse.AtlasRespone{
		Response:       currentModel,
		HttpStatusCode: response.StatusCode,
		HttpError:      "",
	}
}

// List handles the List event from the Cloudformation service.
func List(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", currentModel)
	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	customDBRoleResponse, response, err := client.CustomDatabaseRolesApi.ListCustomDatabaseRoles(ctx, *currentModel.ProjectId).Execute()
	if err != nil {
		_, _ = logger.Warnf("ListCustomDatabaseRoles error: %v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: response.StatusCode,
			HttpError:      err.Error(),
		}
	}

	mm := make([]interface{}, 0)
	for _, customDBRole := range customDBRoleResponse {
		var m Model
		m.completeByAtlasRole(customDBRole)
		m.ProjectId = currentModel.ProjectId
		m.Profile = currentModel.Profile
		mm = append(mm, m)
	}

	return atlasresponse.AtlasRespone{
		Response:       mm,
		HttpStatusCode: response.StatusCode,
		HttpError:      "",
	}
}

func (m *Model) ToCustomDBRole() *admin.UserCustomDBRole {
	var actions []admin.DatabasePrivilegeAction

	for _, a := range m.Actions {
		actions = append(actions, a.toAtlasAction())
	}

	var inheritedRoles []admin.DatabaseInheritedRole
	for _, ir := range m.InheritedRoles {
		inheritedRoles = append(inheritedRoles, ir.toAtlasInheritedRole())
	}

	return &admin.UserCustomDBRole{
		Actions:        actions,
		InheritedRoles: inheritedRoles,
		RoleName:       *m.RoleName,
	}
}

func (a InheritedRole) toAtlasInheritedRole() admin.DatabaseInheritedRole {
	return admin.DatabaseInheritedRole{
		Db:   *a.Db,
		Role: *a.Role,
	}
}

func (a Action) toAtlasAction() admin.DatabasePrivilegeAction {
	var resources []admin.DatabasePermittedNamespaceResource
	for _, r := range a.Resources {
		resources = append(resources, r.toAtlasResource())
	}

	return admin.DatabasePrivilegeAction{
		Action:    *a.Action,
		Resources: resources,
	}
}

func (r Resource) toAtlasResource() admin.DatabasePermittedNamespaceResource {
	return admin.DatabasePermittedNamespaceResource{
		Cluster:    *r.Cluster,
		Collection: *r.Collection,
		Db:         *r.DB,
	}
}

func (m *Model) completeByAtlasRole(role admin.UserCustomDBRole) {
	var actions []Action
	for _, a := range role.Actions {
		actions = append(actions, atlasActionToModel(a))
	}

	var inheritedRoles []InheritedRole
	for _, ir := range role.InheritedRoles {
		inheritedRoles = append(inheritedRoles, atlasInheritedRoleToModel(ir))
	}

	m.Actions = actions
	m.InheritedRoles = inheritedRoles
	m.RoleName = &role.RoleName
}

func atlasActionToModel(action admin.DatabasePrivilegeAction) Action {
	var resources []Resource
	for _, r := range action.Resources {
		resources = append(resources, atlasResourceToModel(r))
	}

	return Action{
		Action:    &action.Action,
		Resources: resources,
	}
}

func atlasResourceToModel(resource admin.DatabasePermittedNamespaceResource) Resource {
	return Resource{
		Collection: &resource.Collection,
		DB:         &resource.Db,
		Cluster:    &resource.Cluster,
	}
}

func atlasInheritedRoleToModel(inheritedRole admin.DatabaseInheritedRole) InheritedRole {
	return InheritedRole{
		Db:   &inheritedRole.Db,
		Role: &inheritedRole.Role,
	}
}
