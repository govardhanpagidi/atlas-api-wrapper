package database_user

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/validator"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
)

var CreateRequiredFields = []string{constants.Username, constants.Password, constants.PublicKey, constants.PrivateKey, constants.ProjectID}
var ReadRequiredFields = []string{constants.ProjectID, constants.Username, constants.PublicKey, constants.PrivateKey}
var DeleteRequiredFields = []string{constants.ProjectID, constants.Username, constants.PublicKey, constants.PrivateKey}
var ListRequiredFields = []string{constants.ProjectID, constants.PublicKey, constants.PrivateKey, constants.PublicKey, constants.PrivateKey}

func setup() {
	util.SetupLogger("mongodb-atlas-database-user")
}

// validateModel to validate inputs to all actions
func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, " create databaseuser is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			HttpError:      configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	groupID, dbUser := setModel(inputModel)
	util.Debugf(ctx, "Arguments: Project ID: %s, Request %#+v", groupID, dbUser)

	request := client.DatabaseUsersApi.CreateDatabaseUser(context.Background(), groupID, dbUser)

	databaseUser, _, err := request.Execute()

	if err != nil {
		fmt.Println("Error creating database user:", err)
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.UserCreateError].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.UserCreateError].Message, *inputModel.Username),
		}
	}
	util.Debugf(ctx, "newUser: %+v", databaseUser)

	return atlasresponse.AtlasRespone{
		Response:       databaseUser,
		HttpStatusCode: configuration.GetConfig()[constants.UserCreateSuccess].Code,
		HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.UserCreateSuccess].Message, *inputModel.Username),
	}
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	if errEvent := validateModel(ReadRequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, " read database user is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			HttpError:      configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	groupID := *inputModel.ProjectId
	username := *inputModel.Username
	dbName := constants.DbuserDbName

	databaseUser, _, err := client.DatabaseUsersApi.GetDatabaseUser(context.Background(), groupID, dbName, username).Execute()

	if err != nil {
		util.Warnf(ctx, " Get Database User Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.UserNotFound].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.UserNotFound].Message, *inputModel.Username),
		}
	}

	currentModel := convertToModel(ctx, inputModel, databaseUser, groupID)
	return atlasresponse.AtlasRespone{
		Response:       currentModel,
		HttpStatusCode: configuration.GetConfig()[constants.FetchUser].Code,
		HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.FetchUser].Message, *inputModel.Username),
	}
}

func convertToModel(ctx context.Context, inputModel *InputModel, databaseUser *admin.CloudDatabaseUser, groupID string) Model {
	var currentModel Model
	currentModel.Username = inputModel.Username
	currentModel.DatabaseName = &databaseUser.DatabaseName
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
	util.Debugf(ctx, "databaseUser:%+v", databaseUser)

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
	util.Debugf(ctx, "currentModel.Roles:%+v", roles)
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
	return currentModel
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, "delete databaseUser is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			HttpError:      configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	groupID := *inputModel.ProjectId
	username := *inputModel.Username
	dbName := constants.DbuserDbName

	user, _, err := client.DatabaseUsersApi.DeleteDatabaseUser(ctx, groupID, dbName, username).Execute()

	if err != nil {
		util.Warnf(ctx, " Delete DatabaseUser Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.DeleteDatabaseUserError].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.DeleteDatabaseUserError].Message, *inputModel.Username),
		}
	}

	return atlasresponse.AtlasRespone{
		Response:       user,
		HttpStatusCode: configuration.GetConfig()[constants.DeleteDatabaseUserSuccess].Code,
		HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.DeleteDatabaseUserSuccess].Message, *inputModel.Username),
	}
}

// List handles listing database users
func List(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	if errEvent := validateModel(ListRequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, "list databaseUsers is failing with invalid parameters: %#+v", errEvent.Error())

		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			HttpError:      configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	groupID := *inputModel.ProjectId

	dbUserModels := make([]interface{}, 0)

	databaseUsers, _, err := client.DatabaseUsersApi.ListDatabaseUsers(ctx, groupID).Execute()

	if err != nil {
		util.Warnf(ctx, " list databaseUsers Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.UserListError].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.UserListError].Message, *inputModel.ProjectId),
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
		HttpStatusCode: configuration.GetConfig()[constants.UserListSuccess].Code,
		HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.UserListSuccess].Message, *inputModel.ProjectId),
	}
}

func setModel(inputModel *InputModel) (string, *admin.CloudDatabaseUser) {
	adminDefaultDbRole := admin.DatabaseUserRole{
		CollectionName: nil,
		DatabaseName:   constants.DbuserDbName,
		RoleName:       constants.DbAdminRoleName,
	}
	adminDefaultAtlasRole := admin.DatabaseUserRole{
		CollectionName: nil,
		DatabaseName:   constants.DbuserDbName,
		RoleName:       constants.AtlasAdminRole,
	}
	var roles []admin.DatabaseUserRole

	roles = append(roles, adminDefaultDbRole)
	roles = append(roles, adminDefaultAtlasRole)
	databaseUser := admin.CloudDatabaseUser{
		AwsIAMType:      nil,
		DatabaseName:    constants.DbuserDbName,
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
