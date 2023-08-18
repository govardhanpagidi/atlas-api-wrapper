package collection

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/validator"
)

var CreateRequiredFields = []string{constants.CollectionName, constants.DatabaseName, constants.HostName, constants.Username, constants.Password}
var DeleteRequiredFields = []string{constants.CollectionName, constants.DatabaseName, constants.HostName, constants.Username, constants.Password}

func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

func setup() {
	util.SetupLogger("mongodb-atlas-collection")
}

// Create This method is used to create a collection in the database
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()

	if errEvent := validateModel(CreateRequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, "create collection is failing with invalid parameters : %+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		util.Warnf(ctx, "Create MongoDriver Error : %+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			HttpError:      configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	database := client.Database(*inputModel.DatabaseName)
	dbCreateErr := database.CreateCollection(context.Background(), *inputModel.CollectionName, nil)

	if dbCreateErr != nil {
		util.Warnf(ctx, "Create Collection error : %+v", dbCreateErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.CollectionError].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.CollectionError].Message, *inputModel.CollectionName),
		}
	}

	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.CollectionSuccess].Code,
		HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.CollectionSuccess].Message, *inputModel.CollectionName),
	}
}

// Delete method drops the collection from the database
func Delete(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, "delete collection is failing with invalid parameters : %+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		util.Warnf(ctx, "Create Mongo Driver Client Error : %+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			HttpError:      configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	database := client.Database(*inputModel.DatabaseName)
	dbCreateErr := database.Collection(*inputModel.CollectionName).Drop(context.Background())

	if dbCreateErr != nil {
		util.Warnf(ctx, "Drop Collection Error : %+v", dbCreateErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.CollectionDeleteError].Code,
			HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.CollectionDeleteError].Message, *inputModel.CollectionName),
		}
	}

	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.CollectionDeleteSuccess].Code,
		HttpError:      fmt.Sprintf(configuration.GetConfig()[constants.CollectionDeleteSuccess].Message, *inputModel.CollectionName),
	}
}
