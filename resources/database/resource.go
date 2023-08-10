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

func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

func validateDeleteModel(fields []string, model *DeleteInputModel) error {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-database-user")
}

func Create(inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", inputModel)

	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}
	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}
	dbCreateErr := client.Database(*inputModel.DatabaseName).CreateCollection(context.Background(), *inputModel.CollectionName, nil)
	if dbCreateErr != nil {
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

func Delete(inputModel *DeleteInputModel) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", inputModel)

	if errEvent := validateDeleteModel(DeleteRequiredFields, inputModel); errEvent != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      errEvent.Error(),
		}
	}
	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}
	dbCreateErr := client.Database(*inputModel.DatabaseName).Drop(context.Background())
	if dbCreateErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusInternalServerError,
			HttpError:      dbCreateErr.Error(),
		}
	}
	dbName := client.Database(*inputModel.DatabaseName).Name()
	return atlasresponse.AtlasRespone{
		Response:       "Database Deleted Successfully:" + dbName,
		HttpStatusCode: http.StatusOK,
		HttpError:      "",
	}
}
