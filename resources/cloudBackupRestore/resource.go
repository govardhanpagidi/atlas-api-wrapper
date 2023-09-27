package cloudBackupRestore

import (
	"context"
	"encoding/json"
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

// CreateRequiredFields Required fields for creating a backup restore job
var CreateRequiredFields = []string{constants.ProjectID, constants.PrivateKey, constants.PublicKey, constants.ClusterName}

// ReadRequiredFields Required fields for reading a backup restore job
var ReadRequiredFields = []string{constants.ProjectID, constants.PrivateKey, constants.PublicKey, constants.ClusterName, constants.JobId}

// Constants for CRUD operations
const (
	AlreadyExists = "already exists"
	DoesntExists  = "does not exist"
	CREATE        = "CREATE"
	READ          = "READ"
	UPDATE        = "UPDATE"
	DELETE        = "DELETE"
	LIST          = "LIST"
)

// Set up the logger
func setup() {
	util.SetupLogger("mongodb-atlas-cloud-backup-snapshot")
}

// Validate the input model
func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

// Create a backup restore job
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	setup()

	// Validate the input model
	modelValidation := validateModel(CreateRequiredFields, inputModel)
	if modelValidation != nil {
		util.Warnf(ctx, "create cluster backup restore job is failing with invalid parameters : %+v", modelValidation.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, modelValidation.Error())
		return handleError(constants.InvalidInputParameter, message, modelValidation)
	}

	// Create a MongoDB SDK client
	client, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))

	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
	}

	// Extract input parameters
	clusterName := cast.ToString(inputModel.ClusterName)
	projectID := cast.ToString(inputModel.ProjectId)

	pointInTimeUtcSeconds := cast.ToInt(inputModel.PointInTimeUtcSeconds)
	oplogTs := cast.ToInt(inputModel.OpLogTs)
	oplogInc := cast.ToInt(inputModel.OpLogInc)
	deliveryType := cast.ToString(inputModel.DeliveryType)
	targetClusterName := cast.ToString(inputModel.TargetClusterName)
	targetProjectID := cast.ToString(inputModel.TargetProjectId)

	// Validate input parameters for automated and point-in-time delivery types
	if deliveryType == constants.Automated || deliveryType == constants.PointInTime {
		if targetClusterName == "" || targetProjectID == "" {
			return handleError(constants.InvalidTargerClusterNameAndProjectId, constants.EmptyString, nil)
		}

		if deliveryType == constants.PointInTime && (pointInTimeUtcSeconds > 0 || (oplogTs > 0 && oplogInc > 0)) {
			return handleError(constants.InvalidPointInTimeError, constants.EmptyString, nil)
		}
	}

	// Create a restore job request
	restoreJobRequest := admin.DiskBackupSnapshotRestoreJob{
		DeliveryType: deliveryType,
	}

	// Add the optional fields to the restore job request if they are not nil
	fmt.Printf("TargetGroupId: %s\n", targetClusterName)
	if inputModel.TargetProjectId != nil && *inputModel.TargetProjectId != "" {
		restoreJobRequest.TargetGroupId = targetProjectID
	} else {
		restoreJobRequest.TargetGroupId = projectID
	}
	if inputModel.TargetClusterName != nil && *inputModel.TargetClusterName != "" {
		restoreJobRequest.TargetClusterName = targetClusterName
	} else {
		restoreJobRequest.TargetClusterName = clusterName
	}
	if inputModel.SnapshotId != nil && *inputModel.SnapshotId != "" {
		restoreJobRequest.SnapshotId = inputModel.SnapshotId
	}
	if inputModel.OpLogTs != nil && *inputModel.OpLogInc != "" {
		restoreJobRequest.OplogTs = &oplogTs
	}
	if inputModel.OpLogInc != nil && *inputModel.OpLogInc != "" {
		restoreJobRequest.OplogInc = &oplogInc
	}
	if inputModel.PointInTimeUtcSeconds != nil && *inputModel.PointInTimeUtcSeconds != "" {
		restoreJobRequest.PointInTimeUTCSeconds = &pointInTimeUtcSeconds
	}

	restoreJobRequestJSON, err := json.MarshalIndent(restoreJobRequest, "", "  ")
	if err != nil {
		return handleError("error", constants.EmptyString, err)
	}
	fmt.Printf(" \n Restore job request: \n %s\n", restoreJobRequestJSON)

	// Create the restore job
	job, _, err := client.CloudBackupsApi.CreateBackupRestoreJob(ctx, projectID, clusterName, &restoreJobRequest).Execute()

	if err != nil {
		message := fmt.Sprintf(configuration.GetConfig()[constants.CreateRestoreJobError].Message, *inputModel.ClusterName)
		return handleError(constants.CreateRestoreJobError, message, err)
	}

	// Create a response with the job
	response := atlasresponse.AtlasResponse{
		Response: job,
	}

	return response
}

// Read a backup restore job
func Read(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasResponse {
	setup()

	// Validate the input model
	modelValidation := validateModel(ReadRequiredFields, inputModel)
	if modelValidation != nil {
		util.Warnf(ctx, "read cluster backup restore job is failing with invalid parameters : %+v", modelValidation.Error())
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

	// Extract input parameters
	clusterName := cast.ToString(inputModel.ClusterName)
	projectID := cast.ToString(inputModel.ProjectId)
	jobID := cast.ToString(inputModel.JobId)

	// Get the restore job
	response := atlasresponse.AtlasResponse{}
	job, _, err := client.CloudBackupsApi.GetBackupRestoreJob(ctx, projectID, clusterName, jobID).Execute()
	if err != nil {
		message := fmt.Sprintf(configuration.GetConfig()[constants.GetRestoreJobError].Message, *inputModel.ClusterName)
		return handleError(constants.GetRestoreJobError, message, err)
	}
	response.Response = job

	// Set the status of the response based on the status of the job
	if job.Cancelled != nil && *job.Cancelled {
		response.Status = constants.Cancelled
		response.Message = constants.RestoreJobCancelled
	} else if job.FinishedAt != nil {
		response.Status = constants.Success
		response.Message = constants.RestoreJobSuccess
	} else if job.Expired != nil && *job.Expired {
		response.Status = constants.Expired
		response.Message = constants.RestoreJobExpired
	} else if job.Failed != nil && *job.Failed {
		response.Status = constants.Failed
		response.Message = constants.RestoreJobFailed
	} else {
		response.Status = constants.InProgress
		response.Message = constants.RestoreJobInProgress
	}

	return response
}

// Handle errors
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
