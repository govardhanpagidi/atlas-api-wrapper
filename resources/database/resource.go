package database

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"
)

var CreateRequiredFields = []string{constants.CollectionName, constants.DatabaseName, constants.HostName, constants.Username, constants.Password}
var DeleteRequiredFields = []string{constants.DatabaseName, constants.HostName, constants.Username, constants.Password}

// validateModel This method is used for validation of InputModel
func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

// validateDeleteModel This method is used for validation of DeleteInputModel
func validateDeleteModel(fields []string, model *DeleteInputModel) error {
	return validator.ValidateModel(fields, model)
}

// setup initializes logger
func setup() {
	util.SetupLogger("mongodb-atlas-database-user")
}

// Create This method is used to create a database and provided collection in the cluster
func Create(inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()

	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		_, _ = logger.Warnf(" create database is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		_, _ = logger.Warnf(" Create Mongo client Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			HttpError:      configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	dbCreateErr := client.Database(*inputModel.DatabaseName).CreateCollection(context.Background(), *inputModel.CollectionName, nil)

	if dbCreateErr != nil {
		_, _ = logger.Warnf(" database Create database Error: %#+v", dbCreateErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.DatabaseError].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.DatabaseError].Message, *inputModel.DatabaseName),
		}
	}

	dbName := client.Database(*inputModel.DatabaseName).Name()

	return atlasresponse.AtlasRespone{
		Response:       fmt.Sprintf(configuration.GetConfig()[constants.DatabaseSuccess].Message, dbName),
		HttpStatusCode: configuration.GetConfig()[constants.DatabaseSuccess].Code,
		HttpError:      "",
	}
}

// Delete method drops the database from the cluster
func Delete(inputModel *DeleteInputModel) atlasresponse.AtlasRespone {
	setup()

	if errEvent := validateDeleteModel(DeleteRequiredFields, inputModel); errEvent != nil {
		_, _ = logger.Warnf(" delete database is failing with invalid parameters: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		_, _ = logger.Warnf(" Create Mongo client Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			HttpError:      configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	dbDeleteErr := client.Database(*inputModel.DatabaseName).Drop(context.Background())

	if dbDeleteErr != nil {
		_, _ = logger.Warnf("delete database Error: %#+v", dbDeleteErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.DatabaseDeleteError].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.DatabaseDeleteError].Message, *inputModel.DatabaseName),
		}
	}

	dbName := client.Database(*inputModel.DatabaseName).Name()

	return atlasresponse.AtlasRespone{
		Response:       fmt.Sprintf(configuration.GetConfig()[constants.DatabaseDeleteSuccess].Message, dbName),
		HttpStatusCode: configuration.GetConfig()[constants.DatabaseDeleteSuccess].Code,
		HttpError:      "",
	}
}
