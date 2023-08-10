package database_user

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
	"net/http"
)

var CreateRequiredFields = []string{constants.Username, constants.Password, constants.DatabaseName, constants.PublicKey, constants.PrivateKey, constants.ProjectID}
var ReadRequiredFields = []string{constants.ProjectID, constants.DatabaseName, constants.Username, constants.PublicKey, constants.PrivateKey}
var DeleteRequiredFields = []string{constants.ProjectID, constants.DatabaseName, constants.Username, constants.PublicKey, constants.PrivateKey}
var ListRequiredFields = []string{constants.ProjectID, constants.PublicKey, constants.PrivateKey, constants.PublicKey, constants.PrivateKey}

func setup() {
	util.SetupLogger("mongodb-atlas-database-user")
}

// validateModel to validate inputs to all actions
func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", inputModel)

	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}

	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)
	if peErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	groupID, dbUser := setModel(inputModel)
	_, _ = logger.Debugf("Arguments: Project ID: %s, Request %#+v", groupID, dbUser)
	request := client.DatabaseUsersApi.CreateDatabaseUser(context.Background(), groupID, dbUser)
	databaseUser, res, err := request.Execute()
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}
	if err != nil {
		fmt.Println("Error creating database user:", err)
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}
	_, _ = logger.Debugf("newUser: %+v", databaseUser)

	return atlasresponse.AtlasRespone{
		Response:       databaseUser,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

// Read handles the Read event from the Cloudformation service.
func Read(inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()

	if errEvent := validateModel(ReadRequiredFields, inputModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "Validation Error",
		}
	}

	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)
	if peErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	groupID := *inputModel.ProjectId
	username := *inputModel.Username
	dbName := *inputModel.DatabaseName

	databaseUser, res, err := client.DatabaseUsersApi.GetDatabaseUser(context.Background(), groupID, dbName, username).Execute()
	statuscode := http.StatusBadRequest
	if res != nil {
		statuscode = res.StatusCode
	}
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: statuscode,
			HttpError:      err.Error(),
		}
	}
	var currentModel Model
	currentModel.Username = inputModel.Username
	currentModel.DatabaseName = inputModel.DatabaseName
	currentModel.ProjectId = inputModel.ProjectId
	currentModel.DatabaseName = &databaseUser.DatabaseName

	if currentModel.LdapAuthType != nil {
		currentModel.LdapAuthType = databaseUser.LdapAuthType
	}
	if currentModel.AWSIAMType != nil {
		currentModel.AWSIAMType = databaseUser.AwsIAMType
	}
	if currentModel.X509Type != nil {
		currentModel.X509Type = databaseUser.X509Type
	}
	currentModel.Username = &databaseUser.Username
	_, _ = logger.Debugf("databaseUser:%+v", databaseUser)

	var roles []RoleDefinition

	for i := range databaseUser.Roles {
		r := databaseUser.Roles[i]
		role := RoleDefinition{
			CollectionName: r.CollectionName,
			DatabaseName:   &r.DatabaseName,
			RoleName:       &r.RoleName,
		}

		roles = append(roles, role)
	}
	currentModel.Roles = roles
	_, _ = logger.Debugf("currentModel.Roles:%+v", roles)
	var labels []LabelDefinition

	for i := range databaseUser.Labels {
		l := databaseUser.Labels[i]
		label := LabelDefinition{
			Key:   l.Key,
			Value: l.Value,
		}

		labels = append(labels, label)
	}
	currentModel.Labels = labels

	cfnid := fmt.Sprintf("%s-%s", *currentModel.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid
	return atlasresponse.AtlasRespone{
		Response:       currentModel,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}

	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)
	if peErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	groupID := *inputModel.ProjectId
	username := *inputModel.Username
	dbName := *inputModel.DatabaseName

	dbuser, res, err := client.DatabaseUsersApi.DeleteDatabaseUser(ctx, groupID, dbName, username).Execute()
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}
	return atlasresponse.AtlasRespone{
		Response:       dbuser,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

// List handles listing database users
func List(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()

	if errEvent := validateModel(ListRequiredFields, inputModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}

	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)
	if peErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	groupID := *inputModel.ProjectId

	dbUserModels := make([]interface{}, 0)

	databaseUsers, res, err := client.DatabaseUsersApi.ListDatabaseUsers(ctx, groupID).Execute()
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}

	if len(databaseUsers.Results) > 0 {
		for i := range databaseUsers.Results {
			var model Model

			databaseUser := databaseUsers.Results[i]
			model.DatabaseName = &databaseUser.DatabaseName
			model.LdapAuthType = databaseUser.LdapAuthType
			model.X509Type = databaseUser.X509Type
			model.Username = &databaseUser.Username
			model.ProjectId = inputModel.ProjectId
			var roles []RoleDefinition

			for i := range databaseUser.Roles {
				r := databaseUser.Roles[i]
				role := RoleDefinition{
					CollectionName: r.CollectionName,
					DatabaseName:   &r.DatabaseName,
					RoleName:       &r.RoleName,
				}

				roles = append(roles, role)
			}
			model.Roles = roles

			var labels []LabelDefinition

			for i := range databaseUser.Labels {
				l := databaseUser.Labels[i]
				label := LabelDefinition{
					Key:   l.Key,
					Value: l.Value,
				}
				labels = append(labels, label)
			}

			model.Labels = labels
			cfnid := fmt.Sprintf("%s-%s", databaseUser.Username, databaseUser.GroupId)

			model.UserCFNIdentifier = &cfnid
			dbUserModels = append(dbUserModels, model)
		}
	}

	return atlasresponse.AtlasRespone{
		Response:       dbUserModels,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

func setModel(inputModel *InputModel) (string, *admin.CloudDatabaseUser) {
	adminDefaultDbRole := admin.DatabaseUserRole{
		CollectionName: nil,
		DatabaseName:   "admin",
		RoleName:       "dbAdmin",
	}
	adminDefaultAtlasRole := admin.DatabaseUserRole{
		CollectionName: nil,
		DatabaseName:   "admin",
		RoleName:       "atlasAdmin",
	}
	var roles []admin.DatabaseUserRole

	roles = append(roles, adminDefaultDbRole)
	roles = append(roles, adminDefaultAtlasRole)
	databaseUser := admin.CloudDatabaseUser{
		AwsIAMType:      nil,
		DatabaseName:    *inputModel.DatabaseName,
		DeleteAfterDate: nil,
		GroupId:         *inputModel.ProjectId,
		Labels:          nil,
		LdapAuthType:    nil,
		Links:           nil,
		OidcAuthType:    nil,
		Password:        inputModel.Password,
		Roles:           roles,
		Scopes:          nil,
		Username:        *inputModel.Username,
		X509Type:        nil,
	}
	return *inputModel.ProjectId, &databaseUser
}
