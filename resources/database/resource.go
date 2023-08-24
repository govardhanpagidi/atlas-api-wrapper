package database

import (
	"context"
	"fmt"

	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/validator"
)

// CreateRequiredFields is a slice of strings that contains the required fields for creating a resource
var CreateRequiredFields = []string{constants.DatabaseName, constants.HostName, constants.Username, constants.Password}

// DeleteRequiredFields is a slice of strings that contains the required fields for deleting a resource
var DeleteRequiredFields = []string{constants.DatabaseName, constants.HostName, constants.Username, constants.Password}

// validateModel This method is used for validation of InputModel
func validateModel(fields []string, model interface{}) error {
	return validator.ValidateModel(fields, model)
}

// setup initializes logger
func setup() {
	util.SetupLogger("mongodb-atlas-database")
}

// Create This method is used to create a database and provided collection in the cluster
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	// Set up the logger for the MongoDB Atlas Collection resource
	setup()

	// Validate the inputModel using the CreateRequiredFields and the validator package
	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " create database is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	// Create a new MongoDB client using the inputModel's username, password, and hostname
	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		// If there is an error creating the MongoDB client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
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
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.DatabaseError].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.DatabaseError].Message, *inputModel.DatabaseName),
		}
	}

	// Get the name of the created database
	dbName := client.Database(*inputModel.DatabaseName).Name()

	// If the collection is created successfully, return a success response
	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.DatabaseSuccess].Code,
		Message:        fmt.Sprintf(configuration.GetConfig()[constants.DatabaseSuccess].Message, dbName),
	}
}

// Delete method drops the database from the cluster
func Delete(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	// Set up the logger for the MongoDB Atlas Collection resource
	setup()

	// Validate the inputModel using the DeleteRequiredFields and the validator package
	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		// If the inputModel is invalid, log a warning and return an error response
		util.Warnf(ctx, " delete database is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	// Create a new MongoDB client using the inputModel's username, password, and hostname
	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		// If there is an error creating the MongoDB client, log a warning and return an error response
		util.Warnf(ctx, " Create Mongo client Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	// Drop the database from the cluster using the inputModel's database name
	dbDeleteErr := client.Database(*inputModel.DatabaseName).Drop(context.Background())

	if dbDeleteErr != nil {
		// If there is an error dropping the database, log a warning and return an error response
		util.Warnf(ctx, "delete database Error: %#+v", dbDeleteErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.DatabaseDeleteError].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.DatabaseDeleteError].Message, *inputModel.DatabaseName),
		}
	}

	// Get the name of the deleted database
	dbName := client.Database(*inputModel.DatabaseName).Name()

	// If the database is dropped successfully, return a success response
	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.DatabaseDeleteSuccess].Code,
		Message:        fmt.Sprintf(configuration.GetConfig()[constants.DatabaseDeleteSuccess].Message, dbName),
	}
}
