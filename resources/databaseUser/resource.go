package database_user

import (
	"context"
	"fmt"

	"github.com/atlas-api-helper/util/logger"

	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/validator"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
)

// CreateRequiredFields is a slice of strings that contains the required fields for creating a resource
var CreateRequiredFields = []string{constants.Username, constants.Password, constants.PublicKey, constants.PrivateKey, constants.ProjectID}

// UpdateRequiredFields is a slice of strings that contains the required fields for creating a resource
var UpdateRequiredFields = []string{constants.Username, constants.PublicKey, constants.PrivateKey, constants.ProjectID}

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
func validateModel(fields []string, model interface{}) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger for the MongoDB Atlas Database User resource
	setup()

	// Validate the inputModel using the CreateRequiredFields and the validator package
	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " create databaseuser is failing with invalid parameters: %#+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	// Create a new MongoDB SDK client using the inputModel's public and private keys
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		// If there is an error creating the MongoDB SDK client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return handleError(constants.MongoClientCreationError, "", peErr)
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
		message := fmt.Sprintf(configuration.GetConfig()[constants.UserCreateError].Message, *inputModel.Username)
		return handleError(constants.UserCreateError, message, err)
	}

	// Log the newly created database user
	util.Debugf(ctx, "newUser: %+v", databaseUser)

	// If the database user is created successfully, return a success response with the database user as the response
	return atlasresponse.AtlasResponse{
		Response: databaseUser,
	}
}

// Update handles the Update roles for the specified database user
func Update(ctx context.Context, inputModel *UpdateInputModel) atlasresponse.AtlasResponse {
	// Set up the logger for the MongoDB Atlas Database User resource
	setup()

	// Validate the inputModel using the CreateRequiredFields and the validator package
	if errEvent := validateModel(UpdateRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " create databaseuser is failing with invalid parameters: %#+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	// Create a new MongoDB SDK client using the inputModel's public and private keys
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		// If there is an error creating the MongoDB SDK client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return handleError(constants.MongoClientCreationError, "", peErr)
	}

	updateInputModel := ConvertUpdateToInputModel(inputModel)
	// Set the groupID and dbUser variables using the inputModel
	groupID, dbUser := setModel(&updateInputModel)

	// Log the groupID and dbUser variables
	util.Debugf(ctx, "Arguments: Project ID: %s, Request %#+v", groupID, dbUser)

	// Create a new database user using the MongoDB SDK client and the groupID and dbUser variables
	request := client.DatabaseUsersApi.UpdateDatabaseUser(context.Background(), groupID, constants.DbuserDbName, *inputModel.Username, dbUser)

	databaseUser, _, err := request.Execute()

	if err != nil {
		// If there is an error creating the database user, log a warning and return an error response
		fmt.Println("Error creating database user:", err)
		message := fmt.Sprintf(configuration.GetConfig()[constants.UserCreateError].Message, *inputModel.Username)
		return handleError(constants.UserCreateError, message, err)
	}

	// Log the newly created database user
	util.Debugf(ctx, "newUser: %+v", databaseUser)

	// If the database user is created successfully, return a success response with the database user as the response
	return atlasresponse.AtlasResponse{
		Response: databaseUser,
	}
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger for the MongoDB Atlas Database User resource
	setup()

	// Validate the inputModel using the ReadRequiredFields and the validator package
	if errEvent := validateModel(ReadRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " read database user is failing with invalid parameters: %#+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	// Create a new MongoDB SDK client using the inputModel's public and private keys
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		// If there is an error creating the MongoDB SDK client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return handleError(constants.MongoClientCreationError, "", peErr)
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
		message := fmt.Sprintf(configuration.GetConfig()[constants.FetchuserError].Message, *inputModel.Username)
		return handleError(constants.FetchuserError, message, err)
	}

	// Convert the database user to a model using the convertToModel function
	currentModel := convertToModel(ctx, inputModel, databaseUser)

	// If the database user is retrieved successfully, return a success response with the currentModel as the response
	return atlasresponse.AtlasResponse{
		Response: currentModel,
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
func Delete(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger
	setup()

	// Validate the inputModel using the DeleteRequiredFields slice
	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, "delete databaseUser is failing with invalid parameters: %#+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	// Create a new MongoDB SDK client using the inputModel's PublicKey and PrivateKey fields
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	// If there was an error creating the client, log a warning and return an error response
	if peErr != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return handleError(constants.MongoClientCreationError, "", peErr)
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
		message := fmt.Sprintf(configuration.GetConfig()[constants.DeleteDatabaseUserError].Message, *inputModel.Username)
		return handleError(constants.MongoClientCreationError, message, err)
	}

	// Return a success response
	return atlasresponse.AtlasResponse{
		Status: fmt.Sprintf(configuration.GetConfig()[constants.DeleteDatabaseUserSuccess].Message, *inputModel.Username),
	}
}

// List handles listing database users
func List(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger
	setup()

	// Validate the inputModel using the ListRequiredFields slice
	if errEvent := validateModel(ListRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, "list databaseUsers is failing with invalid parameters: %#+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	// Create a new MongoDB SDK client using the inputModel's PublicKey and PrivateKey fields
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	// If there was an error creating the client, log a warning and return an error response
	if peErr != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
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
		message := fmt.Sprintf(configuration.GetConfig()[constants.UserListError].Message, *inputModel.ProjectId)
		return handleError(constants.MongoClientCreationError, message, err)
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
	return atlasresponse.AtlasResponse{
		Response: dbUserModels,
	}
}

// convert UI Model to MongoDB Model
func setModel(inputModel *InputModel) (string, *admin.CloudDatabaseUser) {

	// Create a slice of admin.DatabaseUserRole with the adminDefaultDbRole and adminDefaultAtlasRole variables
	var roles []admin.DatabaseUserRole

	// Iterate through the roles and provide comments for each role.
	for _, role := range inputModel.Roles {

		roles = append(roles, admin.DatabaseUserRole{
			CollectionName: role.CollectionName,
			DatabaseName:   *role.DatabaseName,
			RoleName:       *role.RoleName,
		})
	}

	// Create the databaseUser variable with the inputModel's values and the roles slice
	databaseUser := admin.CloudDatabaseUser{
		DatabaseName: constants.DbuserDbName,
		GroupId:      *inputModel.ProjectId,
		Password:     inputModel.Password,
		Roles:        roles,
		Username:     *inputModel.Username,
	}

	// Return the inputModel's ProjectId and the databaseUser variable
	return *inputModel.ProjectId, &databaseUser
}

// handleError is a helper method that logs an error and returns an error response
func handleError(code string, message string, err error) atlasresponse.AtlasResponse {
	// If there is an error, log a warning
	if err != nil {
		errMsg := fmt.Sprintf("%s error:%s", code, err.Error())
		_, _ = logger.Warn(errMsg)
	}
	// If the message is empty, use the message from the configuration
	if message == constants.EmptyString {
		message = configuration.GetConfig()[code].Message
	}
	// Return an error response
	return atlasresponse.AtlasResponse{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[code].Code,
		Message:        message,
		ErrorCode:      configuration.GetConfig()[code].ErrorCode,
	}
}

// ConvertUpdateToInputModel converts an UpdateInputModel to an InputModel.
func ConvertUpdateToInputModel(updateModel *UpdateInputModel) InputModel {
	return InputModel{
		Username:   updateModel.Username,
		Password:   nil, // Assuming you don't want to copy the password
		PublicKey:  updateModel.PublicKey,
		PrivateKey: updateModel.PrivateKey,
		ProjectId:  updateModel.ProjectId,
		Roles:      updateModel.Roles,
	}
}
