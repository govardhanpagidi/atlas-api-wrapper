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

// CreateRequiredFields is a slice of strings that contains the required fields for creating a resource
var CreateRequiredFields = []string{constants.Username, constants.Password, constants.PublicKey, constants.PrivateKey, constants.ProjectID}

// ReadRequiredFields is a slice of strings that contains the required fields for reading a resource
var ReadRequiredFields = []string{constants.ProjectID, constants.Username, constants.PublicKey, constants.PrivateKey}

// DeleteRequiredFields is a slice of strings that contains the required fields for deleting a resource
var DeleteRequiredFields = []string{constants.ProjectID, constants.Username, constants.PublicKey, constants.PrivateKey}

// ListRequiredFields is a slice of strings that contains the required fields for listing resources
var ListRequiredFields = []string{constants.ProjectID, constants.PublicKey, constants.PrivateKey}

// setup This function sets up the logger for the MongoDB Atlas Database User resource
func setup() {
	// Call the SetupLogger function from the util package with the logger name "mongodb-atlas-database-user"
	util.SetupLogger("mongodb-atlas-database-user")
}

// validateModel to validate inputs to all actions
func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	// Set up the logger for the MongoDB Atlas Database User resource
	setup()

	// Validate the inputModel using the CreateRequiredFields and the validator package
	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " create databaseuser is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	// Create a new MongoDB SDK client using the inputModel's public and private keys
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		// If there is an error creating the MongoDB SDK client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	// Set the groupID and dbUser variables using the inputModel
	groupID, dbUser := setModel(inputModel)

	// Log the groupID and dbUser variables
	util.Debugf(ctx, "Arguments: Project ID: %s, Request %#+v", groupID, dbUser)

	// Create a new database user using the MongoDB SDK client and the groupID and dbUser variables
	request := client.DatabaseUsersApi.CreateDatabaseUser(context.Background(), groupID, dbUser)

	databaseUser, _, err := request.Execute()

	if err != nil {
		// If there is an error creating the database user, log a warning and return an error response
		fmt.Println("Error creating database user:", err)
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.UserCreateError].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.UserCreateError].Message, *inputModel.Username),
		}
	}

	// Log the newly created database user
	util.Debugf(ctx, "newUser: %+v", databaseUser)

	// If the database user is created successfully, return a success response with the database user as the response
	return atlasresponse.AtlasRespone{
		Response:       databaseUser,
		HttpStatusCode: configuration.GetConfig()[constants.UserCreateSuccess].Code,
	}
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	// Set up the logger for the MongoDB Atlas Database User resource
	setup()

	// Validate the inputModel using the ReadRequiredFields and the validator package
	if errEvent := validateModel(ReadRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " read database user is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	// Create a new MongoDB SDK client using the inputModel's public and private keys
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		// If there is an error creating the MongoDB SDK client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	// Set the groupID, username, and dbName variables using the inputModel and the constants package
	groupID := *inputModel.ProjectId
	username := *inputModel.Username
	dbName := constants.DbuserDbName

	// Get the database user using the MongoDB SDK client and the groupID, dbName, and username variables
	databaseUser, _, err := client.DatabaseUsersApi.GetDatabaseUser(context.Background(), groupID, dbName, username).Execute()

	if err != nil {
		// If there is an error getting the database user, log a warning and return an error response
		util.Warnf(ctx, " Get Database User Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.UserNotFound].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.UserNotFound].Message, *inputModel.Username),
		}
	}

	// Convert the database user to a model using the convertToModel function
	currentModel := convertToModel(ctx, inputModel, databaseUser)

	// If the database user is retrieved successfully, return a success response with the currentModel as the response
	return atlasresponse.AtlasRespone{
		Response:       currentModel,
		HttpStatusCode: configuration.GetConfig()[constants.FetchUser].Code,
	}
}

// convertToModel converts a CloudDatabaseUser to a Model
func convertToModel(ctx context.Context, inputModel *InputModel, databaseUser *admin.CloudDatabaseUser) Model {
	// Create a new Model
	var currentModel Model

	// Set the currentModel's Username, DatabaseName, and ProjectId fields using the inputModel and the databaseUser
	currentModel.Username = inputModel.Username
	currentModel.DatabaseName = &databaseUser.DatabaseName
	currentModel.ProjectId = inputModel.ProjectId
	currentModel.DatabaseName = &databaseUser.DatabaseName

	// If the currentModel's LdapAuthType field is not nil, set it to the databaseUser's LdapAuthType field
	if currentModel.LdapAuthType != nil {
		currentModel.LdapAuthType = databaseUser.LdapAuthType
	}

	// If the currentModel's AWSIAMType field is not nil, set it to the databaseUser's AwsIAMType field
	if currentModel.AWSIAMType != nil {
		currentModel.AWSIAMType = databaseUser.AwsIAMType
	}

	// If the currentModel's X509Type field is not nil, set it to the databaseUser's X509Type field
	if currentModel.X509Type != nil {
		currentModel.X509Type = databaseUser.X509Type
	}

	// Set the currentModel's Username field to the databaseUser's Username field
	currentModel.Username = &databaseUser.Username

	// Log the databaseUser
	util.Debugf(ctx, "databaseUser:%+v", databaseUser)

	// Create a new slice of RoleDefinition structs
	var roles []RoleDefinition

	// Loop through the databaseUser's Roles slice and append a new RoleDefinition struct to the roles slice for each role
	for i := range databaseUser.Roles {
		r := databaseUser.Roles[i]
		role := RoleDefinition{
			CollectionName: r.CollectionName,
			DatabaseName:   &r.DatabaseName,
			RoleName:       &r.RoleName,
		}

		roles = append(roles, role)
	}

	// Set the currentModel's Roles field to the roles slice
	currentModel.Roles = roles

	// Log the currentModel's Roles field
	util.Debugf(ctx, "currentModel.Roles:%+v", roles)

	// Create a new slice of LabelDefinition structs
	var labels []LabelDefinition

	// Loop through the databaseUser's Labels slice and append a new LabelDefinition struct to the labels slice for each label
	for i := range databaseUser.Labels {
		l := databaseUser.Labels[i]
		label := LabelDefinition{
			Key:   l.Key,
			Value: l.Value,
		}

		labels = append(labels, label)
	}

	// Set the currentModel's Labels field to the labels slice
	currentModel.Labels = labels

	// Return the currentModel
	return currentModel
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	// Set up the logger
	setup()

	// Validate the inputModel using the DeleteRequiredFields slice
	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, "delete databaseUser is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	// Create a new MongoDB SDK client using the inputModel's PublicKey and PrivateKey fields
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	// If there was an error creating the client, log a warning and return an error response
	if peErr != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	// Set the groupID, username, and dbName variables
	groupID := *inputModel.ProjectId
	username := *inputModel.Username
	dbName := constants.DbuserDbName

	// Call the MongoDB SDK client's DeleteDatabaseUser method with the groupID, dbName, and username variables
	_, _, err := client.DatabaseUsersApi.DeleteDatabaseUser(ctx, groupID, dbName, username).Execute()

	// If there was an error deleting the database user, log a warning and return an error response
	if err != nil {
		util.Warnf(ctx, " Delete DatabaseUser Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.DeleteDatabaseUserError].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.DeleteDatabaseUserError].Message, *inputModel.Username),
		}
	}

	// Return a success response
	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.DeleteDatabaseUserSuccess].Code,
		Message:        fmt.Sprintf(configuration.GetConfig()[constants.DeleteDatabaseUserSuccess].Message, *inputModel.Username),
	}
}

// List handles listing database users
func List(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	// Set up the logger
	setup()

	// Validate the inputModel using the ListRequiredFields slice
	if errEvent := validateModel(ListRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, "list databaseUsers is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	// Create a new MongoDB SDK client using the inputModel's PublicKey and PrivateKey fields
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	// If there was an error creating the client, log a warning and return an error response
	if peErr != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	// Set the groupID variable
	groupID := *inputModel.ProjectId

	// Create a slice of empty interfaces to hold the database user models
	dbUserModels := make([]interface{}, 0)

	// Call the MongoDB SDK client's ListDatabaseUsers method with the groupID variable
	databaseUsers, _, err := client.DatabaseUsersApi.ListDatabaseUsers(ctx, groupID).Execute()

	// If there was an error listing the database users, log a warning and return an error response
	if err != nil {
		util.Warnf(ctx, " list databaseUsers Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.UserListError].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.UserListError].Message, *inputModel.ProjectId),
		}
	}

	// If there are database users, loop through them and create a model for each one
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
			dbUserModels = append(dbUserModels, model)
		}
	}

	// Return a success response with the database user models
	return atlasresponse.AtlasRespone{
		Response:       dbUserModels,
		HttpStatusCode: configuration.GetConfig()[constants.UserListSuccess].Code,
	}
}

// convert UI Model to MongoDB Model
func setModel(inputModel *InputModel) (string, *admin.CloudDatabaseUser) {
	// Create the adminDefaultDbRole variable with the constants.DbuserDbName and constants.DbAdminRoleName values
	adminDefaultDbRole := admin.DatabaseUserRole{
		CollectionName: nil,
		DatabaseName:   constants.DbuserDbName,
		RoleName:       constants.DbAdminRoleName,
	}

	// Create the adminDefaultAtlasRole variable with the constants.DbuserDbName and constants.AtlasAdminRole values
	adminDefaultAtlasRole := admin.DatabaseUserRole{
		CollectionName: nil,
		DatabaseName:   constants.DbuserDbName,
		RoleName:       constants.AtlasAdminRole,
	}

	// Create a slice of admin.DatabaseUserRole with the adminDefaultDbRole and adminDefaultAtlasRole variables
	var roles []admin.DatabaseUserRole
	roles = append(roles, adminDefaultDbRole)
	roles = append(roles, adminDefaultAtlasRole)

	// Create the databaseUser variable with the inputModel's values and the roles slice
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

	// Return the inputModel's ProjectId and the databaseUser variable
	return *inputModel.ProjectId, &databaseUser
}
