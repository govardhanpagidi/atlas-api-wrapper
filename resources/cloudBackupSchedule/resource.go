package cloudBackupSchedule

import (
	"context"
	"errors"
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

var RequiredFields = []string{constants.ProjectID, constants.ClusterName}

// validateModel This function validates the given model against the given fields using the validator package
func validateModel(fields []string, model interface{}) error {
	// Call the ValidateModel function from the validator package with the given fields and model
	return validator.ValidateModel(fields, model)
}

// setup This function sets up the logger for the MongoDB Atlas Collection resource
func setup() {
	// Call the SetupLogger function from the util package with the logger name "mongodb-atlas-collection"
	util.SetupLogger("mongodb-atlas-cloud-backup-schedule")
}

func Update(ctx context.Context, inputModel *Model) atlasresponse.AtlasResponse {
	setup()

	// Validate required fields in the request
	modelValidation := validateModel(RequiredFields, inputModel)
	if modelValidation != nil {
		util.Warnf(ctx, "Update cloud backup is failing with invalid parameters : %+v", modelValidation.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, modelValidation.Error())
		return handleError(constants.InvalidInputParameter, message, modelValidation)
	}

	//Create a mongo client using public key and private key
	client, peErr := util.NewMongoDBSDKClient(cast.ToString(*inputModel.PublicKey), cast.ToString(*inputModel.PrivateKey))

	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
	}
	backupPolicy, _, err := client.CloudBackupsApi.GetBackupSchedule(ctx, *inputModel.ProjectId, *inputModel.ClusterName).Execute()
	if err != nil {
		util.Warnf(ctx, "Failed to fetch backup policy for the cluster: %s", *inputModel.ClusterName)
		message := fmt.Sprintf(configuration.GetConfig()[constants.GetBackupScheduleError].Message, *inputModel.ClusterName)
		return handleError(constants.GetBackupScheduleError, message, err)
	}

	if !isPolicySchedulePresent(backupPolicy) {
		util.Warnf(ctx, "Error - Read policy backup schedule for cluster(%s)", *inputModel.ClusterName)
		message := fmt.Sprintf(configuration.GetConfig()[constants.GetPolicyScheduleError].Message, *inputModel.ClusterName)
		return handleError(constants.GetPolicyScheduleError, message, err)
	}

	return cloudBackupScheduleCreateOrUpdate(ctx, inputModel, client)
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, inputModel *Model) atlasresponse.AtlasResponse {
	// logger setup
	setup()
	util.Debugf(ctx, "Get the current snapshot schedule and retention settings for the cluster:%+v", *inputModel.ClusterName)
	// Validate required fields in the request
	if errEvent := validateModel(RequiredFields, inputModel); errEvent != nil {
		util.Warnf(ctx, "Red cloud backup is failing with invalid parameters : %+v", errEvent.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, errEvent.Error())
		return handleError(constants.InvalidInputParameter, message, errEvent)
	}

	//Create a mongo client using public key and private key
	client, peErr := util.NewMongoDBSDKClient(cast.ToString(*inputModel.PublicKey), cast.ToString(*inputModel.PrivateKey))

	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
	}

	// API call to Get the cloud backup schedule
	backupPolicy, _, err := client.CloudBackupsApi.GetBackupSchedule(ctx, *inputModel.ProjectId, *inputModel.ClusterName).Execute()
	if err != nil {
		util.Warnf(ctx, "Failed to fetch backup policy for the cluster: %s", *inputModel.ClusterName)
		message := fmt.Sprintf(configuration.GetConfig()[constants.GetBackupScheduleError].Message, *inputModel.ClusterName)
		return handleError(constants.GetBackupScheduleError, message, err)
	}
	util.Debugf(ctx, "Read() end currentModel:%+v", inputModel)
	// check the policy backup schedule is present for the cluster
	if !isPolicySchedulePresent(backupPolicy) {
		util.Warnf(ctx, "Error - Read policy backup schedule for cluster(%s)", *inputModel.ClusterName)
		message := fmt.Sprintf(configuration.GetConfig()[constants.GetPolicyScheduleError].Message, *inputModel.ClusterName)
		return handleError(constants.GetPolicyScheduleError, message, err)
	}
	// Response
	return atlasresponse.AtlasResponse{Response: backupPolicy}
}

// handles the Create/Update event from the Cloudformation service.
func cloudBackupScheduleCreateOrUpdate(ctx context.Context, currentModel *Model, client *admin.APIClient) atlasresponse.AtlasResponse {
	projectID := currentModel.ProjectId
	clusterName := currentModel.ClusterName

	if err := validateExportDetails(currentModel); err != nil {
		util.Warnf(ctx, "Validate Export Details error: %v", err.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ResourceDoesNotExist].Message, *currentModel.ClusterName)
		return handleError(constants.ResourceDoesNotExist, message, err)
	}

	if err := validatePolicies(currentModel); err != nil {
		util.Warnf(ctx, "Validate Policies error: %v", err.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ValidateExportDetails].Message, *currentModel.ClusterName)
		return handleError(constants.ValidateExportDetails, message, err)
	}

	cloudBackupScheduleRequest := modelToCloudBackupSchedule(currentModel)
	// API call to Create/Update cloud backup schedule
	clusterBackupScheduled, _, err := client.CloudBackupsApi.UpdateBackupSchedule(context.Background(), *projectID, *clusterName, cloudBackupScheduleRequest).Execute()
	if err != nil {
		util.Warnf(ctx, "Update Backup Schedule error: %v", err.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.UpdateBackupScheduleError].Message, *currentModel.ClusterName)
		return handleError(constants.UpdateBackupScheduleError, message, err)
	}
	// Response
	return atlasresponse.AtlasResponse{Response: clusterBackupScheduled}
}

// handleError is a helper method that logs an error and returns an error response
func handleError(code string, message string, err error) atlasresponse.AtlasResponse {
	// If there is an error, log a warning
	if err != nil {
		errMsg := fmt.Sprintf("%s error:%s", code, err.Error())
		_, _ = logger.Warn(errMsg)
	}
	// If the message is empty, use the message from the configuration
	if message == constants.EmptyString {
		message = configuration.GetConfig()[code].Message
	}
	// Return an error response
	return atlasresponse.AtlasResponse{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[code].Code,
		Message:        message,
		ErrorCode:      configuration.GetConfig()[code].ErrorCode,
	}
}

// function to converts 'currentModel' model class to mongodb 'CloudProviderSnapshotBackupPolicy' class.
func modelToCloudBackupSchedule(currentModel *Model) (out *admin.DiskBackupSnapshotSchedule) {
	out = &admin.DiskBackupSnapshotSchedule{}

	if currentModel.AutoExportEnabled != nil {
		out.AutoExportEnabled = currentModel.AutoExportEnabled
		if *currentModel.AutoExportEnabled && currentModel.Export != nil {
			out.Export = expandExport(*currentModel.Export)
		}
	}
	if currentModel.ReferenceHourOfDay != nil {
		out.ReferenceHourOfDay = currentModel.ReferenceHourOfDay
	}
	if currentModel.ReferenceMinuteOfHour != nil {
		out.ReferenceMinuteOfHour = currentModel.ReferenceMinuteOfHour
	}
	if currentModel.RestoreWindowDays != nil {
		out.RestoreWindowDays = currentModel.RestoreWindowDays
	}
	if currentModel.UseOrgAndGroupNamesInExportPrefix != nil {
		out.UseOrgAndGroupNamesInExportPrefix = currentModel.UseOrgAndGroupNamesInExportPrefix
	}
	if currentModel.Policies != nil {
		out.Policies = expandPolicies(currentModel.Policies)
	}
	if currentModel.UpdateSnapshots != nil {
		out.UpdateSnapshots = currentModel.UpdateSnapshots
	}
	if currentModel.CopySettings != nil || len(currentModel.CopySettings) > 0 {
		out.CopySettings = expandCopySettings(currentModel.CopySettings)
	}
	if currentModel.DeleteCopiedBackups != nil || len(currentModel.DeleteCopiedBackups) > 0 {
		out.DeleteCopiedBackups = expandDeleteCopiedBackups(currentModel.DeleteCopiedBackups)
	}
	return out
}

// function to converts model 'ApiDeleteCopiedBackupsView' class to mongodb 'DeleteCopiedBackup' class.
func expandDeleteCopiedBackups(deleteCopiedBackups []ApiDeleteCopiedBackupsView) []admin.DeleteCopiedBackups {
	cloudDeleteCopiedBackups := make([]admin.DeleteCopiedBackups, 0)
	for _, deleteCopiedBackup := range deleteCopiedBackups {
		copiedSetting := admin.DeleteCopiedBackups{
			CloudProvider:     deleteCopiedBackup.CloudProvider,
			RegionName:        deleteCopiedBackup.RegionName,
			ReplicationSpecId: deleteCopiedBackup.ReplicationSpecId,
		}
		cloudDeleteCopiedBackups = append(cloudDeleteCopiedBackups, copiedSetting)
	}
	return cloudDeleteCopiedBackups
}

// function to converts model 'ApiAtlasDiskBackupCopySettingView' class to mongodb 'CopySetting' class.
func expandCopySettings(copySettings []ApiAtlasDiskBackupCopySettingView) []admin.DiskBackupCopySetting {
	cloudCopySettings := make([]admin.DiskBackupCopySetting, 0)
	for _, copySetting := range copySettings {
		backupSetting := admin.DiskBackupCopySetting{
			CloudProvider:     copySetting.CloudProvider,
			RegionName:        copySetting.RegionName,
			ReplicationSpecId: copySetting.ReplicationSpecId,
			ShouldCopyOplogs:  copySetting.ShouldCopyOplogs,
			Frequencies:       copySetting.Frequencies,
		}
		cloudCopySettings = append(cloudCopySettings, backupSetting)
	}
	return cloudCopySettings
}

// function to converts model 'ApiPolicyView' class to mongodb 'Policy' class.
func expandPolicies(policies []ApiPolicyView) []admin.AdvancedDiskBackupSnapshotSchedulePolicy {
	schedulePolicies := make([]admin.AdvancedDiskBackupSnapshotSchedulePolicy, 0)
	for _, s := range policies {
		policy := admin.AdvancedDiskBackupSnapshotSchedulePolicy{
			Id:          s.ID,
			PolicyItems: expandPolicyItems(s.PolicyItems),
		}
		schedulePolicies = append(schedulePolicies, policy)
	}
	return schedulePolicies
}

// function to converts model 'ApiPolicyItemView' class to mongodb 'PolicyItem' class.
func expandPolicyItems(cloudPolicyItems []ApiPolicyItemView) []admin.DiskBackupApiPolicyItem {
	policyItems := make([]admin.DiskBackupApiPolicyItem, 0)
	for _, policyItem := range cloudPolicyItems {
		cPolicyItem := admin.DiskBackupApiPolicyItem{
			Id:                policyItem.ID,
			FrequencyInterval: cast.ToInt(policyItem.FrequencyInterval),
			FrequencyType:     cast.ToString(policyItem.FrequencyType),
			RetentionUnit:     cast.ToString(policyItem.RetentionUnit),
			RetentionValue:    cast.ToInt(policyItem.RetentionValue),
		}
		policyItems = append(policyItems, cPolicyItem)
	}
	return policyItems
}

// function to converts model 'Export' class to mongodb 'Export' class.
func expandExport(export Export) *admin.AutoExportPolicy {
	var exportArg admin.AutoExportPolicy

	if export.ExportBucketId != nil {
		exportArg.ExportBucketId = export.ExportBucketId
	}
	if export.FrequencyType != nil {
		exportArg.FrequencyType = export.FrequencyType
	}
	return &exportArg
}

func validatePolicies(currentModel *Model) error {
	if currentModel.Policies == nil || len(currentModel.Policies) == 0 {
		msg := "validation error: policies cannot be empty"
		return errors.New(msg)
	}
	for _, policy := range currentModel.Policies {
		if policy.PolicyItems == nil || len(policy.PolicyItems) == 0 {
			msg := "validation error: policy items cannot be empty"
			return errors.New(msg)
		}
		for _, policyItem := range policy.PolicyItems {
			if policyItem.FrequencyInterval == nil || policyItem.FrequencyType == nil ||
				policyItem.RetentionUnit == nil || policyItem.RetentionValue == nil {
				err := errors.New("validation error: All values from PolicyItem should be set when `PolicyItems` is set")
				_, _ = logger.Warnf("Update - error: %+v", err)
				return err
			}
		}
	}
	return nil
}

func validateExportDetails(currentModel *Model) error {
	if currentModel.AutoExportEnabled != nil && *currentModel.AutoExportEnabled && currentModel.Export != nil {
		if (currentModel.Export.FrequencyType) == nil {
			err := errors.New("error updating cloud backup schedule: FrequencyType should be set when `Export` is set")
			_, _ = logger.Warnf("Update - error: %+v", err)
			return err
		}
	}
	return nil
}

func isPolicySchedulePresent(backupPolicy *admin.DiskBackupSnapshotSchedule) bool {
	return (backupPolicy.Policies != nil || len(backupPolicy.Policies) > 0) && len(backupPolicy.Policies[0].PolicyItems) > 0
}
