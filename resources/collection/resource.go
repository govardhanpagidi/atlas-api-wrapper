package collection

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
var DeleteRequiredFields = []string{constants.CollectionName, constants.DatabaseName, constants.HostName, constants.Username, constants.Password}

func validateModel(fields []string, model *InputModel) error {
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
	database := client.Database(*inputModel.DatabaseName)
	dbCreateErr := database.CreateCollection(context.Background(), *inputModel.CollectionName, nil)
	if dbCreateErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusInternalServerError,
			HttpError:      dbCreateErr.Error(),
		}
	}
	dbName := client.Database(*inputModel.DatabaseName).Name()
	return atlasresponse.AtlasRespone{
		Response:       "Collection Created Successfully:" + dbName,
		HttpStatusCode: http.StatusOK,
		HttpError:      "",
	}
}

func Delete(inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf(" currentModel: %#+v", inputModel)

	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
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
	database := client.Database(*inputModel.DatabaseName)
	dbCreateErr := database.Collection(*inputModel.CollectionName).Drop(context.Background())
	if dbCreateErr != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusInternalServerError,
			HttpError:      dbCreateErr.Error(),
		}
	}
	dbName := client.Database(*inputModel.DatabaseName).Name()
	return atlasresponse.AtlasRespone{
		Response:       "Collection Deleted Successfully:" + dbName,
		HttpStatusCode: http.StatusOK,
		HttpError:      "",
	}
}
