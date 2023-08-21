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

func validateModel(fields []string, model interface{}) error {
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
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)
	if err != nil {
		util.Warnf(ctx, "Create MongoDriver Error : %+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	database := client.Database(*inputModel.DatabaseName)
	var successCollections []*string
	var failedCollections []*string

	for _, collectionName := range inputModel.CollectionNames {
		dbCreateErr := database.CreateCollection(context.Background(), *collectionName, nil)
		if dbCreateErr != nil {
			util.Warnf(ctx, "Create Collection error : %+v", dbCreateErr.Error())
			failedCollections = append(failedCollections, collectionName)
		} else {
			successCollections = append(successCollections, collectionName)
		}
	}

	if len(successCollections) > 0 {
		successMessage := fmt.Sprintf("Successfully created collections: %s", util.ToStringSlice(successCollections))
		util.Debugf(ctx, successMessage)
	}

	if len(failedCollections) > 0 {
		errorMessage := fmt.Sprintf("Failed to create collections: %s", util.ToStringSlice(failedCollections))
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.CollectionError].Code,
			Message:        errorMessage,
		}
	}

	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.CollectionSuccess].Code,
		Message:        fmt.Sprintf(configuration.GetConfig()[constants.CollectionSuccess].Message, util.ToStringSlice(successCollections)),
	}
}

// Delete method drops the collection from the database
func Delete(ctx context.Context, inputModel *DeleteInputModel) atlasresponse.AtlasRespone {
	setup()
	if errEvent := validateModel(DeleteRequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, "delete collection is failing with invalid parameters : %+v", errEvent.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.InvalidInputParameter].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error()),
		}
	}

	client, err := util.MongoDriverClient(*inputModel.Username, *inputModel.Password, *inputModel.HostName)

	if err != nil {
		util.Warnf(ctx, "Create Mongo Driver Client Error : %+v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.MongoClientCreationError].Code,
			Message:        configuration.GetConfig()[constants.MongoClientCreationError].Message,
		}
	}

	database := client.Database(*inputModel.DatabaseName)
	dbCreateErr := database.Collection(*inputModel.CollectionName).Drop(context.Background())

	if dbCreateErr != nil {
		util.Warnf(ctx, "Drop Collection Error : %+v", dbCreateErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: configuration.GetConfig()[constants.CollectionDeleteError].Code,
			Message:        fmt.Sprintf(configuration.GetConfig()[constants.CollectionDeleteError].Message, *inputModel.CollectionName),
		}
	}

	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.CollectionDeleteSuccess].Code,
		Message:        fmt.Sprintf(configuration.GetConfig()[constants.CollectionDeleteSuccess].Message, *inputModel.CollectionName),
	}
}
