package database

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/resources/cluster"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson"
	"strings"

	"github.com/atlas-api-helper/util/logger"

	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/validator"
)

// CreateRequiredFields is a slice of strings that contains the required fields for creating a resource
var CreateRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.DatabaseName, constants.Username, constants.Password, constants.PublicKey, constants.PrivateKey}

// ReadAllRequiredFields is a slice of strings that contains the required fields for creating a resource
var ReadAllRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.Username, constants.Password, constants.PublicKey, constants.PrivateKey}

// DeleteRequiredFields is a slice of strings that contains the required fields for deleting a resource
var DeleteRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.DatabaseName, constants.Username, constants.Password, constants.PublicKey, constants.PrivateKey}

// validateModel This method is used for validation of InputModel
func validateModel(fields []string, model interface{}) error {
	return validator.ValidateModel(fields, model)
}

// setup initializes logger
func setup() {
	util.SetupLogger("mongodb-atlas-database")
}

// Create This method is used to create a database and provided collection in the cluster
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger for the MongoDB Atlas Collection resource
	setup()

	// Validate the inputModel using the CreateRequiredFields and the validator package
	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " create database is failing with invalid parameters: %#+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	//fetch hostname from the cluster
	inputModel.HostName = getHostName(ctx, *inputModel)

	if inputModel.HostName == nil && *inputModel.HostName == constants.EmptyString {
		util.Warnf(ctx, "Cluster Hostname not set")
		message := fmt.Sprintf(configuration.GetConfig()[constants.ClusterNameNotSet].Message)
		return handleError(constants.ClusterNameNotSet, message, nil)
	}
	// Create a new MongoDB client using the inputModel's username, password, and hostname
	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	// Check if the hostname is assigned to the cluster and throw error
	if err != nil {
		// If there is an error creating the MongoDB client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", err.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, err)
	}

	// Set the collection name to "default" if it is not provided in the inputModel
	collectionName := constants.DefaultCollectionString
	if inputModel.CollectionName != nil {
		collectionName = *inputModel.CollectionName
	}

	// Create the collection in the database using the inputModel's database name and the collection name
	dbCreateErr := client.Database(*inputModel.DatabaseName).CreateCollection(context.Background(), collectionName, nil)

	if dbCreateErr != nil {
		// If there is an error creating the collection, log a warning and return an error response
		util.Warnf(ctx, " database Create database Error: %#+v", dbCreateErr.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.DatabaseError].Message, *inputModel.DatabaseName)
		return handleError(constants.MongoClientCreationError, message, dbCreateErr)
	}

	// Get the name of the created database
	dbName := client.Database(*inputModel.DatabaseName).Name()

	// If the collection is created successfully, return a success response
	return atlasresponse.AtlasResponse{
		Status: fmt.Sprintf(configuration.GetConfig()[constants.DatabaseSuccess].Message, dbName),
	}
}

// ReadAll This method is used to create a database and provided collection in the cluster
func ReadAll(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger for the MongoDB Atlas Collection resource
	setup()

	// Validate the inputModel using the CreateRequiredFields and the validator package
	if errEvent := validateModel(ReadAllRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " Read all databases is failing with invalid parameters: %#+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	//fetch hostname from the cluster
	inputModel.HostName = getHostName(ctx, *inputModel)

	// Check if the hostname is assigned to the cluster and throw error
	if inputModel.HostName == nil && *inputModel.HostName == constants.EmptyString {
		util.Warnf(ctx, "Cluster Hostname not set")
		message := fmt.Sprintf(configuration.GetConfig()[constants.ClusterNameNotSet].Message)
		return handleError(constants.ClusterNameNotSet, message, nil)
	}

	// Create a new MongoDB client using the inputModel's username, password, and hostname
	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		// If there is an error creating the MongoDB client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", err.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, err)
	}
	//databasesOptions := options.ListDatabasesOptions{}

	// Create the collection in the database using the inputModel's database name and the collection name
	databaseNames, dbCreateErr := client.ListDatabaseNames(ctx, bson.M{})
	if dbCreateErr != nil {
		// If there is an error creating the collection, log a warning and return an error response
		util.Warnf(ctx, " list all databases Error: %#+v", dbCreateErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, dbCreateErr)
	}

	formattedDbNames := []string{strings.Join(databaseNames, ", ")}
	// If the collection is created successfully, return a success response
	return atlasresponse.AtlasResponse{
		Response: formattedDbNames,
	}
}

// Delete method drops the database from the cluster
func Delete(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger for the MongoDB Atlas Collection resource
	setup()

	// Validate the inputModel using the DeleteRequiredFields and the validator package
	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " delete database is failing with invalid parameters: %#+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	//fetch hostname from the cluster
	inputModel.HostName = getHostName(ctx, *inputModel)

	// Check if the hostname is assigned to the cluster and throw error
	if inputModel.HostName == nil || *inputModel.HostName == constants.EmptyString {
		util.Warnf(ctx, "Cluster Hostname not set")
		message := fmt.Sprintf(configuration.GetConfig()[constants.ClusterNameNotSet].Message)
		return handleError(constants.ClusterNameNotSet, message, nil)
	}

	// Create a new MongoDB client using the inputModel's username, password, and hostname
	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		// If there is an error creating the MongoDB client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", err.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, err)
	}

	// Drop the database from the cluster using the inputModel's database name
	dbDeleteErr := client.Database(*inputModel.DatabaseName).Drop(context.Background())

	if dbDeleteErr != nil {
		// If there is an error dropping the database, log a warning and return an error response
		util.Warnf(ctx, "delete database Error: %#+v", dbDeleteErr.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.DatabaseDeleteError].Message, *inputModel.DatabaseName)
		return handleError(constants.DatabaseDeleteError, message, dbDeleteErr)
	}

	// Get the name of the deleted database
	dbName := client.Database(*inputModel.DatabaseName).Name()

	// If the database is dropped successfully, return a success response
	return atlasresponse.AtlasResponse{
		Status: fmt.Sprintf(configuration.GetConfig()[constants.DatabaseDeleteSuccess].Message, dbName),
	}
}

// getHostName retrieves the hostname of a MongoDB cluster based on the provided input.
// It creates a MongoDB client using the provided public key and private key, and then
// queries the cluster information to obtain the hostname.
//
// If the cluster hostname is available, it is returned as a string pointer. If any errors
// occur during the process, an empty string pointer is returned.
//
// ctx: The context for the operation.
// inputModel: The input model containing information needed to fetch the hostname.
//
// Returns:
// - A string pointer representing the hostname of the MongoDB cluster (or an empty string pointer on error).
func getHostName(ctx context.Context, inputModel InputModel) *string {
	hostName := ""
	// Create a MongoDB client using the public key and private key
	AdminClient, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))

	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return &hostName
	}

	// Query cluster information
	clusterModel, resp, err := cluster.ReadCluster(ctx, AdminClient, &cluster.Model{ProjectId: inputModel.ProjectId, Name: inputModel.ClusterName})
	if err != nil {
		util.Warnf(ctx, "error cluster get- err:%+v resp:%+v", err, resp)
		return &hostName
	}

	// Extract and return the cluster hostname if available
	if clusterModel.ConnectionStrings.StandardSrv != nil && *clusterModel.ConnectionStrings.StandardSrv != constants.EmptyString {
		parts := strings.SplitN(*clusterModel.ConnectionStrings.StandardSrv, "//", 2)
		return &parts[1]
	} else {
		util.Warnf(ctx, "Cluster Hostname not yet created")
		return &hostName
	}
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
