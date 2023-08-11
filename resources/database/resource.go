package database

import (
	"context"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"
	"net/http"
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
	_, _ = logger.Debugf(" currentModel: %#+v", inputModel)

	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		_, _ = logger.Debugf(" database Create Valitaion Error: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}
	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)
	if err != nil {
		_, _ = logger.Debugf(" Create Mongo client Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}
	dbCreateErr := client.Database(*inputModel.DatabaseName).CreateCollection(context.Background(), *inputModel.CollectionName, nil)
	if dbCreateErr != nil {
		_, _ = logger.Debugf(" database Create database Error: %#+v", dbCreateErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusInternalServerError,
			HttpError:      dbCreateErr.Error(),
		}
	}
	dbName := client.Database(*inputModel.DatabaseName).Name()
	return atlasresponse.AtlasRespone{
		Response:       "Database Created Successfully:" + dbName,
		HttpStatusCode: http.StatusOK,
		HttpError:      "",
	}
}

// Delete method drops the database from the cluster
func Delete(inputModel *DeleteInputModel) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", inputModel)

	if errEvent := validateDeleteModel(DeleteRequiredFields, inputModel); errEvent != nil {
		_, _ = logger.Debugf(" database delete Valitaion Error: %#+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}
	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)
	if err != nil {
		_, _ = logger.Debugf(" Create Mongo client Error: %#+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}
	dbDeleteErr := client.Database(*inputModel.DatabaseName).Drop(context.Background())
	if dbDeleteErr != nil {
		_, _ = logger.Debugf(" database Delete database Error: %#+v", dbDeleteErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusInternalServerError,
			HttpError:      dbDeleteErr.Error(),
		}
	}
	dbName := client.Database(*inputModel.DatabaseName).Name()
	return atlasresponse.AtlasRespone{
		Response:       "Database Deleted Successfully:" + dbName,
		HttpStatusCode: http.StatusOK,
		HttpError:      "",
	}
}
