package cloudBackupSnapshot

import (
	"context"
	"fmt"

	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
)

// Failure Define a Failure constant in the constants package
const Failure = "FAILED"

// RequiredFields Define the required fields for the input model
var RequiredFields = []string{constants.ProjectID, constants.PrivateKey, constants.PublicKey, constants.ClusterName}

// Set up the logger
func setup() {
	util.SetupLogger("mongodb-atlas-cloud-backup-snapshot")
}

// Validate the input model
func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

// Create a cloud backup snapshot
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger
	setup()

	// Validate the input model
	modelValidation := validateModel(RequiredFields, inputModel)
	if modelValidation != nil {
		util.Warnf(ctx, "create cloud backup snapshot is failing with invalid parameters : %+v", modelValidation.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, modelValidation.Error())
		return handleError(constants.InvalidInputParameter, message, modelValidation)
	}

	// Create a MongoDB SDK client
	client, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))
	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
	}

	// Set the snapshot request parameters
	clusterName := cast.ToString(inputModel.ClusterName)
	projectID := cast.ToString(inputModel.ProjectId)
	retentionInDaysStr := cast.ToString(inputModel.RetentionInDays)
	retentionInDays := util.ParseInt(retentionInDaysStr)
	snapshotRequest := admin.DiskBackupOnDemandSnapshotRequest{
		RetentionInDays: &retentionInDays,
		Description:     &inputModel.Description,
	}

	// Take a snapshot of the cloud backup
	snapshot, _, err := client.CloudBackupsApi.TakeSnapshot(context.Background(), projectID, clusterName, &snapshotRequest).Execute()
	if err != nil {
		message := fmt.Sprintf(configuration.GetConfig()[constants.CreateSnapshotError].Message, *inputModel.ClusterName)
		return handleError(constants.CreateSnapshotError, message, err)
	}

	// Return the snapshot in an AtlasResponse
	return atlasresponse.AtlasResponse{
		Response: snapshot,
	}
}

// List cloud backup snapshots
func List(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	// Set up the logger
	setup()

	// Validate the input model
	modelValidation := validateModel(RequiredFields, inputModel)
	if modelValidation != nil {
		util.Warnf(ctx, "list cloud backup snapshots is failing with invalid parameters : %+v", modelValidation.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, modelValidation.Error())
		return handleError(constants.InvalidInputParameter, message, modelValidation)
	}

	// Create a MongoDB SDK client
	client, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))
	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
	}

	// Check if the project exists
	_, _, projectErr := client.ProjectsApi.GetProject(context.Background(), cast.ToString(inputModel.ProjectId)).Execute()
	if projectErr != nil {
		util.Warnf(ctx, "Get Project error: %v", projectErr.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ResourceDoesNotExist].Message, constants.Project, *inputModel.ProjectId)
		return handleError(constants.ResourceDoesNotExist, message, projectErr)
	}

	// Get the cluster information
	cluster, res, err := client.MultiCloudClustersApi.GetCluster(ctx, *inputModel.ProjectId, *inputModel.ClusterName).Execute()
	if err != nil {
		util.Warnf(ctx, "error cluster get- err:%+v resp:%+v", err, res)
		message := fmt.Sprintf(configuration.GetConfig()[constants.ResourceDoesNotExist].Message, constants.Cluster, *inputModel.ClusterName)
		return handleError(constants.ResourceDoesNotExist, message, err)
	}

	// List the cloud backup snapshots
	if len(cluster.ReplicationSpecs) > 1 {
		// List replica set backups
		snapshots, _, err := client.CloudBackupsApi.ListReplicaSetBackups(context.Background(), *inputModel.ProjectId, *inputModel.ClusterName).Execute()
		if err != nil {
			util.Warnf(ctx, "List ReplicaSet Backups error: %v", projectErr.Error())
			message := fmt.Sprintf(configuration.GetConfig()[constants.ListReplicaSetBackupError].Message, *inputModel.ProjectId)
			return handleError(constants.ListReplicaSetBackupError, message, projectErr)
		}

		return atlasresponse.AtlasResponse{Response: snapshots}
	} else {
		// List sharded cluster backups
		snapshots, _, err := client.CloudBackupsApi.ListShardedClusterBackups(context.Background(), *inputModel.ProjectId, *inputModel.ClusterName).Execute()
		if err != nil {
			util.Warnf(ctx, "List sharded Cluster Backups error", projectErr.Error())
			message := fmt.Sprintf(configuration.GetConfig()[constants.ListShardedClusterBackupError].Message, *inputModel.ProjectId)
			return handleError(constants.ListShardedClusterBackupError, message, projectErr)
		}

		return atlasresponse.AtlasResponse{
			Response: snapshots,
		}
	}
}

// Handle errors and return an AtlasResponse
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
