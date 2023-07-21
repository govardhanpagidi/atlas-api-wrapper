package database_user

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/resources/profile"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"
	"github.com/aws/aws-sdk-go/aws"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
	"net/http"
	"time"
)

var CreateRequiredFields = []string{constants.DatabaseName, constants.ProjectID, constants.Roles, constants.Username}
var ReadRequiredFields = []string{constants.ProjectID, constants.DatabaseName, constants.Username}
var UpdateRequiredFields = []string{constants.DatabaseName, constants.ProjectID, constants.Roles, constants.Username}
var DeleteRequiredFields = []string{constants.ProjectID, constants.DatabaseName, constants.Username}
var ListRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-database-user")
}

// validateModel to validate inputs to all actions
func validateModel(fields []string, model *Model) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v, prevModel: %#+v", currentModel)

	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil {
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

	groupID, dbUser, err := setModel(currentModel)
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: 0,
			HttpError:      "",
		}
	}
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
	_, _ = logger.Debugf("newUser: %s", databaseUser)
	cfnid := fmt.Sprintf("%s-%s", *currentModel.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid

	return atlasresponse.AtlasRespone{
		Response:       databaseUser,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()

	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "Validation Error",
		}
	}

	// Create atlas client
	if currentModel.Profile == nil {
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

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName

	databaseUser, res, err := client.DatabaseUsersApi.GetDatabaseUser(ctx, groupID, dbName, username).Execute()
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}

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

// Update handles the Update event from the Cloudformation service.
func Update(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	if errEvent := validateModel(UpdateRequiredFields, currentModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil {
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

	groupID, dbUser, err := setModel(currentModel)
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}
	databaseUser, res, err := client.DatabaseUsersApi.UpdateDatabaseUser(ctx, groupID, *currentModel.DatabaseName, *currentModel.Username, dbUser).Execute()
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}

	cfnid := fmt.Sprintf("%s-%s", *currentModel.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid

	return atlasresponse.AtlasRespone{
		Response:       databaseUser,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil {
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

	groupID := *currentModel.ProjectId
	username := *currentModel.Username
	dbName := *currentModel.DatabaseName

	dbuser, res, err := client.DatabaseUsersApi.DeleteDatabaseUser(ctx, groupID, dbName, username).Execute()
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}

	cfnid := fmt.Sprintf("%s-%s", *currentModel.Username, groupID)
	currentModel.UserCFNIdentifier = &cfnid
	return atlasresponse.AtlasRespone{
		Response:       dbuser,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

// List handles listing database users
func List(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()

	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil {
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

	groupID := *currentModel.ProjectId

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
			model.ProjectId = currentModel.ProjectId
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

func getDBUser(roles []admin.DatabaseUserRole, groupID string, currentModel *Model, labels []admin.ComponentLabel, scopes []admin.UserScope) *admin.CloudDatabaseUser {
	parsedTime, err := time.Parse("2006-01-02 15:04:05", *currentModel.DeleteAfterDate)
	if err != nil {
		_, _ = logger.Debugf(" currentModel: %#+v", currentModel)
		return nil
	}
	return &admin.CloudDatabaseUser{
		AwsIAMType:      currentModel.AWSIAMType,
		DatabaseName:    *currentModel.DatabaseName,
		DeleteAfterDate: &parsedTime,
		GroupId:         groupID,
		Labels:          labels,
		LdapAuthType:    currentModel.LdapAuthType,
		Password:        currentModel.Password,
		Roles:           roles,
		Scopes:          scopes,
		Username:        *currentModel.Username,
		X509Type:        currentModel.X509Type,
	}
}

func setModel(currentModel *Model) (string, *admin.CloudDatabaseUser, error) {
	var roles []admin.DatabaseUserRole
	for i := range currentModel.Roles {
		r := currentModel.Roles[i]
		role := admin.DatabaseUserRole{}
		if r.CollectionName != nil {
			role.CollectionName = r.CollectionName
		}
		if r.DatabaseName != nil {
			role.DatabaseName = *r.DatabaseName
		}

		if r.RoleName != nil {
			role.RoleName = *r.RoleName
		}
		roles = append(roles, role)
	}
	_, _ = logger.Debugf("roles: %#+v", roles)

	var labels []admin.ComponentLabel
	for i := range currentModel.Labels {
		l := currentModel.Labels[i]
		label := admin.ComponentLabel{
			Key:   l.Key,
			Value: l.Value,
		}
		labels = append(labels, label)
	}
	_, _ = logger.Debugf("labels: %#+v", labels)

	var scopes []admin.UserScope
	for i := range currentModel.Scopes {
		s := currentModel.Scopes[i]
		scope := admin.UserScope{
			Name: *s.Name,
			Type: *s.Type,
		}
		scopes = append(scopes, scope)
	}
	_, _ = logger.Debugf("scopes: %#+v", scopes)

	groupID := *currentModel.ProjectId
	_, _ = logger.Debugf("groupID: %#+v", groupID)

	none := "NONE"
	if currentModel.LdapAuthType == nil {
		currentModel.LdapAuthType = &none
	}
	if currentModel.AWSIAMType == nil {
		currentModel.AWSIAMType = &none
	}
	if currentModel.X509Type == nil {
		currentModel.X509Type = &none
	}
	if currentModel.Password == nil {
		if (currentModel.LdapAuthType == &none) && (currentModel.AWSIAMType == &none) && (currentModel.X509Type == &none) {
			err := fmt.Errorf("password cannot be empty if not LDAP or IAM or X509: %v", currentModel)
			return "", nil, err
		}
		s := ""
		currentModel.Password = &s
	}
	if (currentModel.X509Type != &none) || (currentModel.DeleteAfterDate == nil) {
		s := ""
		currentModel.DeleteAfterDate = &s
	}
	_, _ = logger.Debugf("Check Delete after date here::???????")
	user := getDBUser(roles, groupID, currentModel, labels, scopes)
	return groupID, user, nil
}
