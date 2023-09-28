package collection

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/resources/cluster"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// DeleteRequiredFields is a slice of strings that contains the required fields for deleting a resource
var DeleteRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.DatabaseName, constants.Username, constants.Password, constants.PublicKey, constants.PrivateKey}

// ReadAllRequiredFields is a slice of strings that contains the required fields for deleting a resource
var ReadAllRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.DatabaseName, constants.Username, constants.Password, constants.PublicKey, constants.PrivateKey}

// validateModel This function validates the given model against the given fields using the validator package
func validateModel(fields []string, model interface{}) error {
	// Call the ValidateModel function from the validator package with the given fields and model
	return validator.ValidateModel(fields, model)
}

// setup This function sets up the logger for the MongoDB Atlas Collection resource
func setup() {
	// Call the SetupLogger function from the util package with the logger name "mongodb-atlas-collection"
	util.SetupLogger("mongodb-atlas-collection")
}

// Create This method is used to create a collection in the database
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger for the MongoDB Atlas Collection resource
	setup()

	// Validate the inputModel using the CreateRequiredFields and the validator package
	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, "create collection is failing with invalid parameters : %+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	// fetch hostname from the cluster
	inputModel.HostName = getHostName(ctx, *inputModel.PublicKey, *inputModel.PrivateKey, *inputModel.ProjectId, *inputModel.ClusterName)

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
		util.Warnf(ctx, "Create MongoDriver Error : %+v", err.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, err)
	}

	// Get the database from the client using the inputModel's database name
	database := client.Database(*inputModel.DatabaseName)
	var successCollections []*string
	var failedCollections []*string

	// Iterate over the inputModel's collection names and create each collection in the database
	for _, collectionName := range inputModel.CollectionNames {
		dbCreateErr := database.CreateCollection(context.Background(), *collectionName, nil)
		if dbCreateErr != nil {
			// If there is an error creating the collection, log a warning and add the collection name to the failedCollections slice
			util.Warnf(ctx, "Create Collection error : %+v", dbCreateErr.Error())
			failedCollections = append(failedCollections, collectionName)
		} else {
			// If the collection is created successfully, add the collection name to the successCollections slice
			successCollections = append(successCollections, collectionName)
		}
	}

	// If there are any successfully created collections, log a debug message
	if len(successCollections) > 0 {
		successMessage := fmt.Sprintf("Successfully created collections: %s", util.ToStringSlice(successCollections))
		util.Debugf(ctx, successMessage)
	}

	// If there are any failed collections, return an error response
	if len(failedCollections) > 0 {
		errorMessage := fmt.Sprintf("Failed to create collections: %s", util.ToStringSlice(failedCollections))
		return handleError(constants.MongoClientCreationError, errorMessage, err)
	}

	// If all collections are created successfully, return a success response
	return atlasresponse.AtlasResponse{
		Status: fmt.Sprintf(configuration.GetConfig()[constants.CollectionSuccess].Message, util.ToStringSlice(successCollections)),
	}
}

// ReadAll method drops the collection from the database
func ReadAll(ctx context.Context, inputModel *DeleteInputModel) atlasresponse.AtlasResponse {
	// Set up the logger for the MongoDB Atlas Collection resource
	setup()

	// Validate the inputModel using the DeleteRequiredFields and the validator package
	if errEvent := validateModel(ReadAllRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, "Read all collection names is failing with invalid parameters : %+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}
	//fetch hostname from the cluster
	inputModel.HostName = getHostName(ctx, *inputModel.PublicKey, *inputModel.PrivateKey, *inputModel.ProjectId, *inputModel.ClusterName)

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
		util.Warnf(ctx, "Create Mongo Driver Client Error : %+v", err.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, err)
	}

	// Get the database from the client using the inputModel's database name
	database := client.Database(*inputModel.DatabaseName)

	collectionsOptions := options.ListCollectionsOptions{}

	// fetch the collection from the database
	collectionNames, dbCreateErr := database.ListCollectionNames(ctx, collectionsOptions)

	if dbCreateErr != nil {
		// If there is an error dropping the collection, log a warning and return an error response
		util.Warnf(ctx, "read all Collection Error : %+v", dbCreateErr.Error())
		return handleError(constants.CollectionListError, constants.EmptyString, dbCreateErr)
	}

	// If the collection is dropped successfully, return a success response
	return atlasresponse.AtlasResponse{Response: collectionNames}
}

// Delete method drops the collection from the database
func Delete(ctx context.Context, inputModel *DeleteInputModel) atlasresponse.AtlasResponse {
	// Set up the logger for the MongoDB Atlas Collection resource
	setup()

	// Validate the inputModel using the DeleteRequiredFields and the validator package
	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, "delete collection is failing with invalid parameters : %+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	//fetch hostname from the cluster
	inputModel.HostName = getHostName(ctx, *inputModel.PublicKey, *inputModel.PrivateKey, *inputModel.ProjectId, *inputModel.ClusterName)

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
		util.Warnf(ctx, "Create Mongo Driver Client Error : %+v", err.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, err)
	}

	// Get the database from the client using the inputModel's database name
	database := client.Database(*inputModel.DatabaseName)

	// Drop the collection from the database
	dbCreateErr := database.Collection(*inputModel.CollectionName).Drop(context.Background())

	if dbCreateErr != nil {
		// If there is an error dropping the collection, log a warning and return an error response
		util.Warnf(ctx, "Drop Collection Error : %+v", dbCreateErr.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.CollectionDeleteError].Message, *inputModel.CollectionName)
		return handleError(constants.CollectionDeleteError, message, dbCreateErr)
	}

	// If the collection is dropped successfully, return a success response
	return atlasresponse.AtlasResponse{
		Status: fmt.Sprintf(configuration.GetConfig()[constants.CollectionDeleteSuccess].Message, *inputModel.CollectionName),
	}
}

func handleError(code string, message string, err error) atlasresponse.AtlasResponse {
	if err != nil {
		errMsg := fmt.Sprintf("%s error:%s", code, err.Error())
		_, _ = logger.Warn(errMsg)
	}
	if message == constants.EmptyString {
		message = configuration.GetConfig()[code].Message
	}
	return atlasresponse.AtlasResponse{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[code].Code,
		Message:        message,
		ErrorCode:      configuration.GetConfig()[code].ErrorCode,
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
func getHostName(ctx context.Context, publicKey string, privateKey string, projectId string, clusterName string) *string {
	hostName := ""
	// Create a MongoDB client using the public key and private key
	AdminClient, peErr := util.NewMongoDBSDKClient(cast.ToString(publicKey), cast.ToString(privateKey))

	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return &hostName
	}

	// Query cluster information
	clusterModel, resp, err := cluster.ReadCluster(ctx, AdminClient, &cluster.Model{ProjectId: &projectId, Name: &clusterName})
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
