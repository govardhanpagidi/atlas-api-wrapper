// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package project_invitation

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

var CreateRequiredFields = []string{constants.ProjectID, constants.Username}
var ReadRequiredFields = []string{constants.ProjectID, constants.ID}
var UpdateRequiredFields = []string{constants.ProjectID, constants.ID}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ID}
var ListRequiredFields = []string{constants.ProjectID}

func validateModel(fields []string, model *Model) error {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-project-invitation")
}

func Create(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()

	_, _ = logger.Debug("Create() currentModel:%+v", currentModel)

	// Validation
	if err := validateModel(CreateRequiredFields, currentModel); err != nil {
		_, _ = logger.Warnf("Validation Error")
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "validation Error",
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	invitationReq := &admin.GroupInvitationRequest{
		Roles:    currentModel.Roles,
		Username: currentModel.Username,
	}

	invitation, res, err := client.ProjectsApi.CreateProjectInvitation(ctx, *currentModel.ProjectId, invitationReq).Execute()
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}
	currentModel.Id = invitation.Id

	// Response
	return atlasresponse.AtlasRespone{
		Response:       invitationToModel(currentModel, invitation),
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

func Read(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf("Read() currentModel:%+v", currentModel)
	// Validation
	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "validation Error",
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	invitation, res, err := client.ProjectsApi.GetProjectInvitation(ctx, *currentModel.ProjectId, *currentModel.Id).Execute()
	if err != nil {
		_, _ = logger.Warnf("Read - error: %+v", err)

		// if invitation already accepted
		if res.StatusCode == 404 {
			if alreadyAccepted, _ := validateProjectInvitationAlreadyAccepted(context.Background(), client, *currentModel.Username, *currentModel.ProjectId); alreadyAccepted {
				return atlasresponse.AtlasRespone{
					Response:       nil,
					HttpStatusCode: http.StatusBadRequest,
					HttpError:      "invitation has been already accepted",
				}
			}
		}
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}

	return atlasresponse.AtlasRespone{
		Response:       invitationToModel(currentModel, invitation),
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

func Update(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()

	_, _ = logger.Warnf("Update() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "validation Error",
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	invitationReq := &admin.GroupInvitationRequest{
		Roles: currentModel.Roles, Username: currentModel.Username,
	}

	invitation, res, err := client.ProjectsApi.UpdateProjectInvitation(ctx, *currentModel.ProjectId, invitationReq).Execute()
	if err != nil {
		_, _ = logger.Warnf("Update - error: %+v", err)
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}
	_, _ = logger.Debugf("%s invitation updated", *currentModel.Id)

	// Response
	return atlasresponse.AtlasRespone{
		Response:       invitationToModel(currentModel, invitation),
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

func Delete(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()

	_, _ = logger.Debugf("Delete() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "validation Error",
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}
	_, res, err := client.ProjectsApi.DeleteProjectInvitation(ctx, *currentModel.ProjectId, *currentModel.Id).Execute()
	if err != nil {
		_, _ = logger.Warnf("Delete - error: %+v", err)
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}
	_, _ = logger.Debugf("deleted invitation with Id :%s", *currentModel.Id)

	// Response
	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

func List(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()

	_, _ = logger.Debugf("List() currentModel:%+v", currentModel)

	// Validation
	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "validation Error",
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	invitations, res, err := client.ProjectsApi.ListProjectInvitations(ctx, *currentModel.ProjectId).Execute()
	if err != nil {
		_, _ = logger.Warnf("List - error: %+v", err)
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}

	var invites []interface{}
	// iterate invites
	for i := range invitations {
		invite := invitationToModel(currentModel, &invitations[i])
		invites = append(invites, invite)
	}

	// Response
	return atlasresponse.AtlasRespone{
		Response:       invites,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

func invitationToModel(currentModel *Model, invitation *admin.GroupInvitation) Model {
	expiresAt := invitation.ExpiresAt.Format("2006-01-02 15:04:05")
	createdAt := invitation.CreatedAt.Format("2006-01-02 15:04:05")
	out := Model{
		Profile:         currentModel.Profile,
		ProjectId:       currentModel.ProjectId,
		Username:        invitation.Username,
		Id:              invitation.Id,
		Roles:           invitation.Roles,
		ExpiresAt:       &expiresAt,
		CreatedAt:       &createdAt,
		InviterUsername: invitation.InviterUsername,
	}

	return out
}

func validateProjectInvitationAlreadyAccepted(ctx context.Context, client *admin.APIClient, username, projectID string) (bool, error) {
	user, _, err := client.MongoDBCloudUsersApi.GetUserByUsername(ctx, username).Execute()
	if err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}
	for _, role := range user.Roles {
		if role.GetGroupId() == projectID {
			return true, nil
		}
	}

	return false, nil
}
