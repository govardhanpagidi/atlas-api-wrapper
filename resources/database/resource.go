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

var CreateRequiredFields = []string{constants.DatabaseName, constants.HostName, constants.Username, constants.Password}
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
	setup()

	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, " create database is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}
	collectionName := "default"
	if inputModel.CollectionName != nil {
		collectionName = *inputModel.CollectionName
	}

	dbCreateErr := client.Database(*inputModel.DatabaseName).CreateCollection(context.Background(), collectionName, nil)

	if dbCreateErr != nil {
		util.Warnf(ctx, " database Create database Error: %#+v", dbCreateErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.DatabaseError].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.DatabaseError].Message, *inputModel.DatabaseName),
		}
	}

	dbName := client.Database(*inputModel.DatabaseName).Name()

	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.DatabaseSuccess].Code,
		Message:        fmt.Sprintf(configuration.GetConfig()[constants.DatabaseSuccess].Message, dbName),
	}
}

// Delete method drops the database from the cluster
func Delete(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()

	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, " delete database is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		util.Warnf(ctx, " Create Mongo client Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	dbDeleteErr := client.Database(*inputModel.DatabaseName).Drop(context.Background())

	if dbDeleteErr != nil {
		util.Warnf(ctx, "delete database Error: %#+v", dbDeleteErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.DatabaseDeleteError].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.DatabaseDeleteError].Message, *inputModel.DatabaseName),
		}
	}

	dbName := client.Database(*inputModel.DatabaseName).Name()

	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.DatabaseDeleteSuccess].Code,
		Message:        fmt.Sprintf(configuration.GetConfig()[constants.DatabaseDeleteSuccess].Message, dbName),
	}
}
