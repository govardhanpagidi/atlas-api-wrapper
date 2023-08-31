// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/atlas-api-helper/util/logger"

	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/configuration"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/validator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
)

// defaultLabel is the default label for all resources created by the provider
var defaultLabel = Labels{Key: aws.String("Infrastructure Tool"), Value: aws.String("MongoDB Atlas CloudFormation Provider")}

// CreateRequiredFields is a list of required fields for creating a MongoDB Atlas cluster
var CreateRequiredFields = []string{constants.ProjectID, constants.PrivateKey, constants.PublicKey, constants.TshirtSize, constants.CloudProvider}

// ReadRequiredFields is a list of required fields for reading a MongoDB Atlas cluster
var ReadRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.PublicKey, constants.PrivateKey}

// DeleteRequiredFields is a list of required fields for deleting a MongoDB Atlas cluster
var DeleteRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.PublicKey, constants.PrivateKey}

// ListRequiredFields is a list of required fields for listing MongoDB Atlas clusters
var ListRequiredFields = []string{constants.ProjectID, constants.PublicKey, constants.PrivateKey}

const (
	AlreadyExists = "already exists"
	DoesntExists  = "does not exist"
	CREATE        = "CREATE"
	READ          = "READ"
	UPDATE        = "UPDATE"
	DELETE        = "DELETE"
	LIST          = "LIST"
)

// setup initializes logger
func setup() {
	util.SetupLogger("mongodb-atlas-cluster")
}

// validateModel inputs based on the method
func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()

	// Validate required fields in the request
	modelValidation := validateModel(CreateRequiredFields, inputModel)
	if modelValidation != nil {
		util.Warnf(ctx, "create cluster is failing with invalid parameters : %+v", modelValidation.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, modelValidation.Error())
		return handleError(constants.InvalidInputParameter, message, modelValidation)
	}

	//Create a mongo client using public key and private key
	client, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))

	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
	}
	projectIpList, _, ipListError := client.ProjectIPAccessListApi.ListProjectIpAccessLists(context.Background(), *inputModel.ProjectId).Execute()
	if ipListError != nil {
		util.Warnf(ctx, "Error When fetching Project IP AccessList :%v", ipListError.Error())
		return handleError(constants.ProjectIpAccessListError, constants.EmptyString, ipListError)
	}

	hasPublicAccess := checkIfProjectHasPublicAccess(projectIpList.Results)

	//check if project already exists
	_, _, projectErr := client.ProjectsApi.GetProject(context.Background(), cast.ToString(inputModel.ProjectId)).Execute()

	if projectErr != nil {
		util.Warnf(ctx, "Get Project error: %v", projectErr.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ResourceDoesNotExist].Message, constants.Project, *inputModel.ProjectId)
		return handleError(constants.ResourceDoesNotExist, message, nil)
	}

	//load cluster configuration based on TshirtSize from config.json
	currentModel, err := loadClusterConfiguration(ctx, *inputModel)

	if err != nil {
		util.Warnf(ctx, "Create Current Model error: %v", err.Error())
		return handleError(constants.ClusterModelError, constants.EmptyString, err)
	}
	currentModel.validateDefaultLabel()

	//list all private endpoints for the specific project
	endPoints, _, endPointErr := client.PrivateEndpointServicesApi.ListPrivateEndpointServices(ctx, *inputModel.ProjectId, *inputModel.CloudProvider).Execute()

	if endPointErr != nil {
		util.Warnf(ctx, "Get PrivateEndpoint error: %v", endPointErr.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ListEndpointError].Message, *inputModel.ProjectId)
		return handleError(constants.ListEndpointError, message, err)
	}

	//check if at least one private endpoint is attached to the project
	count := len(endPoints)
	if count == 0 {
		util.Warnf(ctx, "PrivateEndpoint Not Configured for ProjectId %s error: %v", *inputModel.ProjectId, errors.New(configuration.GetConfig()[constants.NoEndpointConfigured].Message))
		message := fmt.Sprintf(configuration.GetConfig()[constants.NoEndpointConfigured].Message, *inputModel.ProjectId)
		return handleError(constants.NoEndpointConfigured, message, err)
	}

	//load all region names from the private endpoints
	util.Debugf(ctx, "Cluster create projectId: %s, clusterName: %s", *inputModel.ProjectId, *currentModel.Name)
	var endpointRegions []string
	for _, endPoint := range endPoints {
		endpointRegions = append(endpointRegions, *endPoint.RegionName)
	}

	//load all region names from config.json
	var clusterAdvancedConfigRegions []string
	for _, specs := range currentModel.ReplicationSpecs {
		for _, advancedConfig := range specs.AdvancedRegionConfigs {
			clusterAdvancedConfigRegions = append(clusterAdvancedConfigRegions, *advancedConfig.RegionName)
		}

	}
	if len(clusterAdvancedConfigRegions) == 0 {
		util.Warnf(ctx, "No advancedCluster configuration is provided for the cluster")
		return handleError(constants.NoAdvancedClusterConfiguration, constants.EmptyString, err)
	}

	//compare if the regions from json matches the regions from private endpoints
	isEndPointConfigured := checkIfEndpointRegionIsSameAsClusterRegion(endpointRegions, clusterAdvancedConfigRegions)

	if !isEndPointConfigured {
		message := fmt.Sprintf(configuration.GetConfig()[constants.NoEndpointConfigured].Message, *inputModel.ProjectId)
		return handleError(constants.NoEndpointConfiguredForRegion, message, err)
	}
	// Prepare cluster request
	clusterRequest, err := createClusterRequest(ctx, &currentModel)

	if err != nil {
		message := fmt.Sprintf(configuration.GetConfig()[constants.NoAdvancedClusterConfiguration].Message)
		return handleError(constants.ClusterRequestError, message, err)
	}

	// Create Cluster
	cluster, _, err := client.MultiCloudClustersApi.CreateCluster(ctx, cast.ToString(inputModel.ProjectId), clusterRequest).Execute()
	if err != nil {
		return handleError(constants.ClusterCreateError, constants.EmptyString, err)
	}

	model := &Model{}
	mapClusterToModel(ctx, model, cluster)
	model.PublicAccessEnabled = &hasPublicAccess
	// Fetch advanced cluster config
	processArgs, _, clusterErr := client.ClustersApi.GetClusterAdvancedConfiguration(context.Background(), *model.ProjectId, *model.Name).Execute()

	if clusterErr != nil {
		return handleError(constants.ClusterAdvancedListError, constants.EmptyString, err)
	}

	model.AdvancedSettings = flattenProcessArgs(processArgs)

	currentModel.StateName = cluster.StateName

	return atlasresponse.AtlasRespone{
		Response:       model,
		HttpStatusCode: configuration.GetConfig()[constants.ClusterCreateSuccess].Code,
	}
}

func checkIfProjectHasPublicAccess(results []admin.NetworkPermissionEntry) bool {
	for _, result := range results {
		if result.IpAddress != nil && *result.IpAddress == constants.PublicIp {
			return true
		}
		if result.CidrBlock != nil && *result.CidrBlock == constants.PublicIp {
			return true
		}
	}
	return false
}

// handleError is a helper method that logs an error and returns an error response
func handleError(code string, message string, err error) atlasresponse.AtlasRespone {
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
	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[code].Code,
		Message:        message,
	}
}

// checkIfEndpointRegionIsSameAsClusterRegion checks if the regions from config.json match the regions from private endpoints
func checkIfEndpointRegionIsSameAsClusterRegion(endpoints, advancedCluster []string) bool {
	// Create a map to store values from the first slice
	seen := make(map[string]bool)
	for _, item := range endpoints {
		seen[item] = true
	}

	// Check if any value from the second slice exists in the map
	for _, item := range advancedCluster {
		if seen[item] {
			return true
		}
	}

	return false
}

// loadClusterConfiguration loads the config.json file from project path
func loadClusterConfiguration(ctx context.Context, model InputModel) (Model, error) {
	var currentModel Model
	var ClusterConfig map[string]Model

	err := util.LoadConfigFromFile(constants.ClusterConfigLocation, &ClusterConfig)
	if ClusterConfig == nil || err != nil {
		return currentModel, fmt.Errorf("failed to load cluster configuration from file: %s", constants.ClusterConfigLocation)
	}

	// Extract the key for the current cluster configuration
	key := extractClusterKey(model)

	// Get the cluster configuration for the current key
	clusterConfig, ok := ClusterConfig[key]

	// Log the selected cluster configuration
	util.Debugf(ctx, "Selected Cluster Configuration : %+v  for the T-shirt Size :%s", clusterConfig.ToString(), *model.TshirtSize)

	// If the cluster configuration is found, set it as the current model
	if ok {
		currentModel = clusterConfig
	} else {
		// If the cluster configuration is not found, return an error
		return currentModel, fmt.Errorf("provided Cluster Size:%s and cloudProvider:%s is Invalid: ", *model.TshirtSize, *model.CloudProvider)
	}

	// If the cluster name is provided, set it as the current model's name
	if model.ClusterName != nil {
		currentModel.Name = model.ClusterName
	} else {
		// If the cluster name is not provided, generate a new name for the cluster
		currentModel.Name = generateClusterName(model)
	}

	// If the MongoDB version is provided, set it as the current model's MongoDB version
	if model.MongoDBVersion != nil {
		currentModel.MongoDBVersion = model.MongoDBVersion
	}

	return currentModel, nil
}

// extractClusterKey This method generates the key using which the config is fetched
func extractClusterKey(model InputModel) string {
	// Create a buffer to store the key
	var configKey bytes.Buffer

	// Append the T-shirt size to the key
	configKey.WriteString(strings.ToLower(*model.TshirtSize))

	// Append a hyphen to the key
	configKey.WriteString("-")

	// Append the cloud provider to the key
	configKey.WriteString(strings.ToLower(*model.CloudProvider))

	// Convert the buffer to a string and return it as the key
	key := configKey.String()
	return key
}

// generateClusterName This method generates the cluster name which is then assigned to the created cluster
func generateClusterName(model InputModel) *string {
	// Get the current time in the format "02-01-06 15:04:05"
	toRet := time.Now().Format("02-01-06 15:04:05")

	// Replace all colons with hyphens
	toRet = strings.ReplaceAll(toRet, ":", "-")

	// Replace all spaces with hyphens
	toRet = strings.ReplaceAll(toRet, " ", "-")

	// Replace all hyphens with hyphens (no-op)
	toRet = strings.ReplaceAll(toRet, "-", "-")

	// Concatenate the cloud provider, project ID, T-shirt size, and current time to generate the cluster name
	toRet = *model.TshirtSize + "-" + *model.CloudProvider + "-" + toRet + "-" + *model.ProjectId

	// Return the cluster name as a pointer to a string
	return &toRet
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	// Set up the logger
	setup()

	// Validate the input fields
	modelValidation := validateModel(ReadRequiredFields, inputModel)
	if modelValidation != nil {
		util.Warnf(ctx, "read cluster is failing with invalid parameters : %+v", modelValidation.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, modelValidation.Error())
		return handleError(constants.InvalidInputParameter, message, modelValidation)
	}

	// Create a MongoDB client using the public key and private key
	client, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))

	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
	}

	// Check if the project already exists
	_, _, projectErr := client.ProjectsApi.GetProject(context.Background(), cast.ToString(inputModel.ProjectId)).Execute()

	if projectErr != nil {
		util.Warnf(ctx, "Get Project error: %v", projectErr.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ResourceDoesNotExist].Message, constants.Project, *inputModel.ProjectId)
		return handleError(constants.ResourceDoesNotExist, message, projectErr)
	}

	// Read the cluster based on the provided params
	model, resp, err := readCluster(ctx, client, &Model{ProjectId: inputModel.ProjectId, Name: inputModel.ClusterName})

	if err != nil {
		util.Warnf(ctx, "error cluster get- err:%+v resp:%+v", err, resp)
		message := fmt.Sprintf(configuration.GetConfig()[constants.ResourceDoesNotExist].Message, constants.Cluster, *inputModel.ClusterName)
		return handleError(constants.ResourceDoesNotExist, message, err)
	}

	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.ClusterReadSuccess].Code,
		Message:        fmt.Sprintf(configuration.GetConfig()[constants.ClusterReadSuccess].Message, *model.StateName),
	}
}

// Delete This method deletes the cluster based on the clusterName
func Delete(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	// Set up the logger
	setup()

	// Validate the input fields
	modelValidation := validateModel(DeleteRequiredFields, inputModel)
	if modelValidation != nil {
		util.Warnf(ctx, "delete cluster is failing with invalid parameters : %+v", modelValidation.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, modelValidation.Error())
		return handleError(constants.InvalidInputParameter, message, modelValidation)

	}

	// Create a MongoDB client using the public key and private key
	client, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))

	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
	}

	// Check if the project already exists
	_, _, projectErr := client.ProjectsApi.GetProject(context.Background(), cast.ToString(inputModel.ProjectId)).Execute()

	if projectErr != nil {
		util.Warnf(ctx, "Get Project error: %v", projectErr.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ResourceDoesNotExist].Message, constants.Project, *inputModel.ProjectId)
		return handleError(constants.ResourceDoesNotExist, message, projectErr)
	}

	// Set retainBackup to false
	retainBackup := false

	// Create args for the API call to delete the cluster
	args := admin.DeleteClusterApiParams{
		GroupId:       *inputModel.ProjectId,
		ClusterName:   *inputModel.ClusterName,
		RetainBackups: &retainBackup,
	}

	// Make API call to delete the cluster
	_, err := client.MultiCloudClustersApi.DeleteClusterWithParams(context.Background(), &args).Execute()

	if err != nil {
		util.Warnf(ctx, "Delete cluster error: %v", err.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ClusterDeleteError].Message, *inputModel.ClusterName)
		return handleError(constants.ClusterDeleteError, message, err)
	}

	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: configuration.GetConfig()[constants.ClusterDeleteSuccess].Code,
		Message:        fmt.Sprintf(configuration.GetConfig()[constants.ClusterDeleteSuccess].Message, *inputModel.ClusterName),
	}
}

// List handles the List event from the Cloudformation service.
func List(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	// Set up the logger
	setup()

	// Validate the input fields
	modelValidation := validateModel(ListRequiredFields, inputModel)
	if modelValidation != nil {
		util.Warnf(ctx, "list clusters is failing with invalid parameters : %+v", modelValidation.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.InvalidInputParameter].Message, modelValidation.Error())
		return handleError(constants.InvalidInputParameter, message, modelValidation)
	}

	// Create a MongoDB client using the public key and private key
	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)

	if peErr != nil {
		util.Warnf(ctx, "CreateMongoDBClient error: %v", peErr.Error())
		return handleError(constants.MongoClientCreationError, constants.EmptyString, peErr)
	}

	// Check if the project already exists
	_, _, projectErr := client.ProjectsApi.GetProject(context.Background(), cast.ToString(inputModel.ProjectId)).Execute()

	if projectErr != nil {
		util.Warnf(ctx, "Get Project error: %v", projectErr.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ResourceDoesNotExist].Message, constants.Project, *inputModel.ProjectId)
		return handleError(constants.ResourceDoesNotExist, message, projectErr)
	}

	// Create args for the API call to list all clusters associated with the project
	args := admin.ListClustersApiParams{
		GroupId: *inputModel.ProjectId,
	}

	// Make API call to list all clusters associated with the project
	clustersResponse, _, err := client.MultiCloudClustersApi.ListClustersWithParams(context.Background(), &args).Execute()

	if err != nil {
		util.Warnf(ctx, "List clusters error: %v", err.Error())
		message := fmt.Sprintf(configuration.GetConfig()[constants.ClusterListError].Message, *inputModel.ProjectId)
		return handleError(constants.ClusterListError, message, err)
	}

	projectIpList, _, ipListError := client.ProjectIPAccessListApi.ListProjectIpAccessLists(context.Background(), *inputModel.ProjectId).Execute()
	if ipListError != nil {
		util.Warnf(ctx, "Error When fetching Project IP AccessList :%v", ipListError.Error())
		return handleError(constants.ProjectIpAccessListError, constants.EmptyString, ipListError)
	}

	hasPublicAccess := checkIfProjectHasPublicAccess(projectIpList.Results)
	// Create an empty slice of models
	models := make([]*Model, 0)

	// Iterate over the clusters in the response and map them to models
	for i := range clustersResponse.Results {
		model := &Model{}
		mapClusterToModel(ctx, model, &clustersResponse.Results[i])

		// Fetch advanced cluster config
		processArgs, _, clusterErr := client.ClustersApi.GetClusterAdvancedConfiguration(context.Background(), *model.ProjectId, *model.Name).Execute()

		if clusterErr != nil {
			message := fmt.Sprintf(configuration.GetConfig()[constants.ClusterListError].Message, *inputModel.ProjectId)
			return handleError(constants.ClusterAdvancedListError, message, err)
		}

		model.AdvancedSettings = flattenProcessArgs(processArgs)
		model.PublicAccessEnabled = &hasPublicAccess
		models = append(models, model)
	}

	// Return an AtlasRespone with the models and the appropriate message and status code
	return atlasresponse.AtlasRespone{
		Response:       models,
		HttpStatusCode: configuration.GetConfig()[constants.ClusterListSuccess].Code,
	}
}

// mapClusterToModel This method is used to map the cluster object returned from the mongo client to our model
func mapClusterToModel(ctx context.Context, model *Model, cluster *admin.AdvancedClusterDescription) {

	// Map the cluster ID to the model
	if cluster.Id != nil {
		model.Id = cluster.Id
	}

	// Map the project ID to the model
	if cluster.GroupId != nil {
		model.ProjectId = cluster.GroupId
	}

	// Map the cluster name to the model
	if cluster.Name != nil {
		model.Name = cluster.Name
	}

	// Map the backup enabled flag to the model
	if cluster.BackupEnabled != nil {
		model.BackupEnabled = cluster.BackupEnabled
	}

	// Map the BI connector config to the model
	if cluster.BiConnector != nil {
		model.BiConnector = flattenBiConnectorConfig(cluster.BiConnector)
	}

	// Map the connection strings to the model
	if cluster.ConnectionStrings != nil {
		model.ConnectionStrings = flattenConnectionStrings(cluster.ConnectionStrings)
	}

	// Map the cluster type to the model
	if cluster.ClusterType != nil {
		model.ClusterType = cluster.ClusterType
	}

	// Map the created date to the model
	if cluster.CreateDate != nil {
		createdDate := cluster.CreateDate.Format("2006-01-02 15:04:05")
		model.CreatedDate = &createdDate
	}

	// Map the disk size to the model
	if cluster.DiskSizeGB != nil {
		model.DiskSizeGB = cluster.DiskSizeGB
	}

	// Map the encryption at rest provider to the model
	if cluster.EncryptionAtRestProvider != nil {
		model.EncryptionAtRestProvider = cluster.EncryptionAtRestProvider
	}

	// Map the labels to the model
	if cluster.Labels != nil {
		model.Labels = flattenLabels(cluster.Labels)
	}

	// Map the MongoDB major version to the model
	if cluster.MongoDBMajorVersion != nil {
		model.MongoDBMajorVersion = cluster.MongoDBMajorVersion
	}

	// Map the MongoDB version to the model
	if cluster.MongoDBVersion != nil {
		model.MongoDBVersion = cluster.MongoDBVersion
	}

	// Map the paused flag to the model
	if cluster.Paused != nil {
		model.Paused = cluster.Paused
	}

	// Map the PIT enabled flag to the model
	if cluster.PitEnabled != nil {
		model.PitEnabled = cluster.PitEnabled
	}

	// Map the replication specs to the model
	if cluster.ReplicationSpecs != nil {
		model.ReplicationSpecs = flattenReplicationSpecs(ctx, cluster.ReplicationSpecs)
	}

	// Map the root cert type to the model
	if cluster.RootCertType != nil {
		model.RootCertType = cluster.RootCertType
	}

	// Map the state name to the model
	if cluster.StateName != nil {
		model.StateName = cluster.StateName
	}

	// Map the version release system to the model
	if cluster.VersionReleaseSystem != nil {
		model.VersionReleaseSystem = cluster.VersionReleaseSystem
	}

	// Map the termination protection enabled flag to the model
	if cluster.TerminationProtectionEnabled != nil {
		model.TerminationProtectionEnabled = cluster.TerminationProtectionEnabled
	}
}

// HasAdvanceSettings This method checks if the model has advanced settings
func (m *Model) HasAdvanceSettings() bool {
	// This logic is because of a bug in CloudFormation, when we return in_progress in the CREATE,
	// the second time the CREATE gets executed it returns the AdvancedSettings is not nil but its fields are nil
	return m.AdvancedSettings != nil && (m.AdvancedSettings.DefaultReadConcern != nil ||
		m.AdvancedSettings.DefaultWriteConcern != nil ||
		m.AdvancedSettings.FailIndexKeyTooLong != nil ||
		m.AdvancedSettings.JavascriptEnabled != nil ||
		m.AdvancedSettings.MinimumEnabledTLSProtocol != nil ||
		m.AdvancedSettings.NoTableScan != nil ||
		m.AdvancedSettings.OplogSizeMB != nil ||
		m.AdvancedSettings.SampleSizeBIConnector != nil ||
		m.AdvancedSettings.SampleRefreshIntervalBIConnector != nil ||
		m.AdvancedSettings.OplogMinRetentionHours != nil)
}

// expandBiConnector This function is used to expand the BiConnector struct to the admin.BiConnector struct
func expandBiConnector(biConnector *BiConnector) *admin.BiConnector {
	// If the biConnector is nil, return nil
	if biConnector == nil {
		return nil
	}

	// Create a new admin.BiConnector struct and map the fields from the biConnector struct to it
	return &admin.BiConnector{
		Enabled:        biConnector.Enabled,
		ReadPreference: biConnector.ReadPreference,
	}
}

// expandReplicationSpecs This function is used to expand the AdvancedReplicationSpec struct to the admin.ReplicationSpec struct
func expandReplicationSpecs(ctx context.Context, replicationSpecs []AdvancedReplicationSpec) []admin.ReplicationSpec {
	var rSpecs []admin.ReplicationSpec

	// Loop through each replication spec in the input slice
	for i := range replicationSpecs {
		var numShards int

		// Create a new admin.ReplicationSpec struct and map the fields from the AdvancedReplicationSpec struct to it
		rSpec := admin.ReplicationSpec{
			Id:            nil,
			NumShards:     &numShards,
			RegionConfigs: expandRegionsConfig(ctx, replicationSpecs[i].AdvancedRegionConfigs),
		}

		// Map the NumShards field to the new struct if it is not nil
		if replicationSpecs[i].NumShards != nil {
			rSpec.NumShards = replicationSpecs[i].NumShards
		}

		// Map the ZoneName field to the new struct if it is not nil
		if replicationSpecs[i].ZoneName != nil {
			zoneName := cast.ToString(replicationSpecs[i].ZoneName)
			rSpec.ZoneName = &zoneName
		}

		// Append the new struct to the output slice
		rSpecs = append(rSpecs, rSpec)
	}

	// Print the length and contents of the output slice for debugging purposes
	fmt.Printf("specs: len %d %+v", len(replicationSpecs), rSpecs)

	// Return the output slice
	return rSpecs
}

// expandAutoScaling This function is used to expand the AdvancedAutoScaling struct to the admin.AdvancedAutoScalingSettings struct
func expandAutoScaling(scaling *AdvancedAutoScaling) *admin.AdvancedAutoScalingSettings {
	// Create a new admin.AdvancedAutoScalingSettings struct
	advAutoScaling := &admin.AdvancedAutoScalingSettings{}

	// If the scaling is nil, return nil
	if scaling == nil {
		return nil
	}

	// If the Compute field is not nil, map its fields to the new struct
	if scaling.Compute != nil {
		var minInstanceSize string
		if scaling.Compute.MinInstanceSize != nil {
			minInstanceSize = *scaling.Compute.MinInstanceSize
		}
		var maxInstanceSize string
		if scaling.Compute.MaxInstanceSize != nil {
			maxInstanceSize = *scaling.Compute.MaxInstanceSize
		}
		advAutoScaling.Compute = &admin.AdvancedComputeAutoScaling{
			Enabled:          scaling.Compute.Enabled,
			MaxInstanceSize:  &maxInstanceSize,
			MinInstanceSize:  &minInstanceSize,
			ScaleDownEnabled: scaling.Compute.ScaleDownEnabled,
		}
	}

	// If the DiskGB field is not nil, map its fields to the new struct
	if scaling.DiskGB != nil {
		advAutoScaling.DiskGB = &admin.DiskGBAutoScaling{
			Enabled: scaling.DiskGB.Enabled,
		}
	}

	// Return the new struct
	return advAutoScaling
}

// expandRegionsConfig This function is used to expand the AdvancedRegionConfig slice to the admin.CloudRegionConfig slice
func expandRegionsConfig(ctx context.Context, regionConfigs []AdvancedRegionConfig) []admin.CloudRegionConfig {
	var regionsConfigs []admin.CloudRegionConfig

	// Loop through each region config in the input slice
	for _, regionCfg := range regionConfigs {
		// Expand the region config to the admin.CloudRegionConfig struct and append it to the output slice
		regionsConfigs = append(regionsConfigs, expandRegionConfig(ctx, regionCfg))
	}

	// Return the output slice
	return regionsConfigs
}

// expandRegionConfig This function is used to expand the AdvancedRegionConfig struct to the admin.CloudRegionConfig struct
func expandRegionConfig(ctx context.Context, regionCfg AdvancedRegionConfig) admin.CloudRegionConfig {
	// Set the region and provider name fields to their default values
	var region string
	if regionCfg.RegionName != nil {
		region = *regionCfg.RegionName
	}
	providerName := constants.AWS
	if regionCfg.ProviderName != nil {
		providerName = *regionCfg.ProviderName
	}

	// Create a new admin.CloudRegionConfig struct and map the fields from the AdvancedRegionConfig struct to it
	advRegionConfig := admin.CloudRegionConfig{
		ProviderName: &providerName,
		RegionName:   &region,
		Priority:     regionCfg.Priority,
	}

	// If the AutoScaling field is not nil, map its fields to the new struct
	if regionCfg.AutoScaling != nil {
		advRegionConfig.AutoScaling = expandAutoScaling(regionCfg.AutoScaling)
	}

	// If the AnalyticsAutoScaling field is not nil, map its fields to the new struct
	if regionCfg.AnalyticsAutoScaling != nil {
		advRegionConfig.AnalyticsAutoScaling = expandAutoScaling(regionCfg.AnalyticsAutoScaling)
	}

	// If the AnalyticsSpecs field is not nil, map its fields to the new struct
	if regionCfg.AnalyticsSpecs != nil {
		advRegionConfig.AnalyticsSpecs = expandRegionConfigSpec(ctx, regionCfg.AnalyticsSpecs)
	}

	// If the ElectableSpecs field is not nil, map its fields to the new struct
	if regionCfg.ElectableSpecs != nil {
		advRegionConfig.ElectableSpecs = expandRegionConfigSpecElectableSpecs(ctx, regionCfg.ElectableSpecs)
	}

	// If the ReadOnlySpecs field is not nil, map its fields to the new struct
	if regionCfg.ReadOnlySpecs != nil {
		advRegionConfig.ReadOnlySpecs = expandRegionConfigSpec(ctx, regionCfg.ReadOnlySpecs)
	}

	// If the BackingProviderName field is not nil, map its fields to the new struct
	if regionCfg.BackingProviderName != nil {
		advRegionConfig.BackingProviderName = regionCfg.BackingProviderName
	}

	// Return the new struct
	return advRegionConfig
}

// expandRegionConfigSpec This function is used to expand the Specs struct to the admin.DedicatedHardwareSpec struct
func expandRegionConfigSpec(ctx context.Context, spec *Specs) *admin.DedicatedHardwareSpec {
	// If the spec is nil, return nil
	if spec == nil {
		return nil
	}

	// Set the ebsVolumeType and instanceSize fields to their default values
	var ebsVolumeType string
	var instanceSize string
	if spec.EbsVolumeType != nil {
		ebsVolumeType = *spec.EbsVolumeType
	}
	if spec.InstanceSize != nil {
		instanceSize = *spec.InstanceSize
	}

	// Set the val variable to the value of the DiskIOPS field if it is not nil
	var val int
	if spec.DiskIOPS != nil {
		v, err := strconv.Atoi(*spec.DiskIOPS)
		if err == nil {
			val = v
		}
		util.Debugf(ctx, "set diskIops %d", val)
	}

	// Create a new admin.DedicatedHardwareSpec struct and map the fields from the Specs struct to it
	return &admin.DedicatedHardwareSpec{
		DiskIOPS:      &val,
		EbsVolumeType: &ebsVolumeType,
		InstanceSize:  &instanceSize,
		NodeCount:     spec.NodeCount,
	}
}

// expandRegionConfigSpecElectableSpecs This function is used to expand the Specs struct to the admin.HardwareSpec struct
func expandRegionConfigSpecElectableSpecs(ctx context.Context, spec *Specs) *admin.HardwareSpec {
	// If the spec is nil, return nil
	if spec == nil {
		return nil
	}

	// Set the ebsVolumeType and instanceSize fields to their default values
	var ebsVolumeType string
	var instanceSize string
	if spec.EbsVolumeType != nil {
		ebsVolumeType = *spec.EbsVolumeType
	}
	if spec.InstanceSize != nil {
		instanceSize = *spec.InstanceSize
	}

	// Set the val variable to the value of the DiskIOPS field if it is not nil
	var val int
	if spec.DiskIOPS != nil {
		v, err := strconv.Atoi(*spec.DiskIOPS)
		if err == nil {
			val = v
		}
		util.Debugf(ctx, "set diskIops %d", val)
	}

	// Create a new admin.HardwareSpec struct and map the fields from the Specs struct to it
	return &admin.HardwareSpec{
		DiskIOPS:      &val,
		EbsVolumeType: &ebsVolumeType,
		InstanceSize:  &instanceSize,
		NodeCount:     spec.NodeCount,
	}
}

// expandLabelSlice This function is used to expand the Labels slice to the admin.ComponentLabel slice
func expandLabelSlice(labels []Labels) []admin.ComponentLabel {
	// Create a new slice of admin.ComponentLabel structs with the same length as the input slice
	res := make([]admin.ComponentLabel, len(labels))

	// Loop through each label in the input slice
	for i := range labels {
		// Set the key and value fields to their default values
		var key string
		if labels[i].Key != nil {
			key = *labels[i].Key
		}
		var value string
		if labels[i].Key != nil {
			value = *labels[i].Value
		}

		// Create a new admin.ComponentLabel struct and map the fields from the Labels struct to it
		res[i] = admin.ComponentLabel{
			Key:   &key,
			Value: &value,
		}
	}

	// Return the output slice
	return res
}

// flattenAutoScaling This function is used to flatten the admin.AdvancedAutoScalingSettings struct to the AdvancedAutoScaling struct
func flattenAutoScaling(scaling *admin.AdvancedAutoScalingSettings) *AdvancedAutoScaling {
	// If the scaling is nil, return nil
	if scaling == nil {
		return nil
	}

	// Create a new AdvancedAutoScaling struct
	advAutoScaling := &AdvancedAutoScaling{}

	// If the DiskGB field is not nil, map its fields to the new struct
	if scaling.DiskGB != nil {
		advAutoScaling.DiskGB = &DiskGB{Enabled: scaling.DiskGB.Enabled}
	}

	// If the Compute field is not nil, map its fields to the new struct
	if scaling.Compute != nil {
		compute := &Compute{}
		if scaling.Compute.Enabled != nil {
			compute.Enabled = scaling.Compute.Enabled
		}
		if scaling.Compute.ScaleDownEnabled != nil {
			compute.ScaleDownEnabled = scaling.Compute.ScaleDownEnabled
		}
		if scaling.Compute.MinInstanceSize != nil {
			compute.MinInstanceSize = scaling.Compute.MinInstanceSize
		}
		if scaling.Compute.MaxInstanceSize != nil {
			compute.MaxInstanceSize = scaling.Compute.MaxInstanceSize
		}

		advAutoScaling.Compute = compute
	}

	// Return the new struct
	return advAutoScaling
}

// flattenReplicationSpecs This function is used to flatten the admin.ReplicationSpec slice to the AdvancedReplicationSpec slice
func flattenReplicationSpecs(ctx context.Context, replicationSpecs []admin.ReplicationSpec) []AdvancedReplicationSpec {
	var rSpecs []AdvancedReplicationSpec

	// Loop through each replication spec in the input slice
	for ind := range replicationSpecs {
		// Get the region configs from the replication spec
		configs := replicationSpecs[ind].RegionConfigs

		// Create a new AdvancedReplicationSpec struct and map the fields from the ReplicationSpec struct to it
		rSpec := AdvancedReplicationSpec{
			ID:                    replicationSpecs[ind].Id,
			NumShards:             replicationSpecs[ind].NumShards,
			ZoneName:              replicationSpecs[ind].ZoneName,
			AdvancedRegionConfigs: flattenRegionsConfig(ctx, configs),
		}

		// Append the new struct to the output slice
		rSpecs = append(rSpecs, rSpec)
	}

	// Print the output slice for debugging purposes
	fmt.Printf("specs: len %d %+v", len(replicationSpecs), rSpecs)

	// Return the output slice
	return rSpecs
}

// flattenRegionsConfig This function is used to flatten the admin.CloudRegionConfig slice to the AdvancedRegionConfig slice
func flattenRegionsConfig(ctx context.Context, regionConfigs []admin.CloudRegionConfig) []AdvancedRegionConfig {
	var regionsConfigs []AdvancedRegionConfig

	// Loop through each region config in the input slice
	for i := range regionConfigs {
		// Call the flattenRegionConfig function to flatten the region config
		regionsConfigs = append(regionsConfigs, flattenRegionConfig(ctx, &regionConfigs[i]))
	}

	// Return the output slice
	return regionsConfigs
}

// flattenRegionConfig This function is used to flatten the admin.CloudRegionConfig struct to the AdvancedRegionConfig struct
func flattenRegionConfig(ctx context.Context, regionCfg *admin.CloudRegionConfig) AdvancedRegionConfig {
	// Create a new AdvancedRegionConfig struct and map the fields from the CloudRegionConfig struct to it
	advRegConfig := AdvancedRegionConfig{
		AutoScaling:          flattenAutoScaling(regionCfg.AutoScaling),
		AnalyticsAutoScaling: flattenAutoScaling(regionCfg.AnalyticsAutoScaling),
		RegionName:           regionCfg.RegionName,
		Priority:             regionCfg.Priority,
	}

	// If the AnalyticsSpecs field is not nil, map its fields to the new struct
	if regionCfg.AnalyticsSpecs != nil {
		advRegConfig.AnalyticsSpecs = flattenRegionConfigSpec(ctx, regionCfg.AnalyticsSpecs)
	}

	// If the ElectableSpecs field is not nil, map its fields to the new struct
	if regionCfg.ElectableSpecs != nil {
		advRegConfig.ElectableSpecs = flattenRegionConfigHardwareSpecSpec(ctx, regionCfg.ElectableSpecs)
	}

	// If the ReadOnlySpecs field is not nil, map its fields to the new struct
	if regionCfg.ReadOnlySpecs != nil {
		advRegConfig.ReadOnlySpecs = flattenRegionConfigSpec(ctx, regionCfg.ReadOnlySpecs)
	}

	// Return the new struct
	return advRegConfig
}

// flattenRegionConfigHardwareSpecSpec This function is used to flatten the admin.HardwareSpec struct to the Specs struct
func flattenRegionConfigHardwareSpecSpec(ctx context.Context, spec *admin.HardwareSpec) *Specs {
	// If the spec is nil, return nil
	if spec == nil {
		return nil
	}

	// Set the diskIops field to its default value
	var diskIops string
	if spec.DiskIOPS != nil {
		diskIops = strconv.FormatInt(*util.Cast64(spec.DiskIOPS), 10)
		util.Debugf(ctx, "get diskIops %s", diskIops)
	}

	// Create a new Specs struct and map the fields from the HardwareSpec struct to it
	return &Specs{
		DiskIOPS:      &diskIops,
		EbsVolumeType: spec.EbsVolumeType,
		InstanceSize:  spec.InstanceSize,
		NodeCount:     spec.NodeCount,
	}
}

// flattenRegionConfigSpec This function is used to flatten the admin.DedicatedHardwareSpec struct to the Specs struct
func flattenRegionConfigSpec(ctx context.Context, spec *admin.DedicatedHardwareSpec) *Specs {
	// If the spec is nil, return nil
	if spec == nil {
		return nil
	}

	// Set the diskIops field to its default value
	var diskIops string
	if spec.DiskIOPS != nil {
		diskIops = strconv.FormatInt(*util.Cast64(spec.DiskIOPS), 10)
		util.Debugf(ctx, "get diskIops %s", diskIops)
	}

	// Create a new Specs struct and map the fields from the DedicatedHardwareSpec struct to it
	return &Specs{
		DiskIOPS:      &diskIops,
		EbsVolumeType: spec.EbsVolumeType,
		InstanceSize:  spec.InstanceSize,
		NodeCount:     spec.NodeCount,
	}
}

// flattenBiConnectorConfig This function is used to flatten the admin.BiConnector struct to the BiConnector struct
func flattenBiConnectorConfig(biConnector *admin.BiConnector) *BiConnector {
	// If the biConnector is nil, return nil
	if biConnector == nil {
		return nil
	}

	// Create a new BiConnector struct and map the fields from the BiConnector struct to it
	return &BiConnector{
		ReadPreference: biConnector.ReadPreference,
		Enabled:        biConnector.Enabled,
	}
}

// flattenConnectionStrings This function is used to flatten the admin.ClusterConnectionStrings struct to the ConnectionStrings struct
func flattenConnectionStrings(clusterConnStrings *admin.ClusterConnectionStrings) (connStrings *ConnectionStrings) {
	// If the clusterConnStrings is nil, return nil
	if clusterConnStrings != nil {
		// Create a new ConnectionStrings struct and map the fields from the ClusterConnectionStrings struct to it
		connStrings = &ConnectionStrings{
			Standard:        clusterConnStrings.Standard,
			StandardSrv:     clusterConnStrings.StandardSrv,
			Private:         clusterConnStrings.Private,
			PrivateSrv:      clusterConnStrings.PrivateSrv,
			PrivateEndpoint: flattenPrivateEndpoint(clusterConnStrings.PrivateEndpoint),
		}
	}

	// Return the new struct
	return
}

// flattenPrivateEndpoint This function is used to flatten the admin.ClusterDescriptionConnectionStringsPrivateEndpoint slice to the PrivateEndpoint slice
func flattenPrivateEndpoint(pes []admin.ClusterDescriptionConnectionStringsPrivateEndpoint) []PrivateEndpoint {
	var prvEndpoints []PrivateEndpoint

	// If the input slice is nil, return an empty slice
	if pes == nil {
		return prvEndpoints
	}

	// Loop through each private endpoint in the input slice
	for ind := range pes {
		// Create a new PrivateEndpoint struct and map the fields from the ClusterDescriptionConnectionStringsPrivateEndpoint struct to it
		pe := PrivateEndpoint{
			ConnectionString:    pes[ind].ConnectionString,
			SRVConnectionString: pes[ind].SrvConnectionString,
			Type:                pes[ind].Type,
			Endpoints:           flattenEndpoints(pes[ind].Endpoints),
		}

		// Append the new struct to the output slice
		prvEndpoints = append(prvEndpoints, pe)
	}

	// Return the output slice
	return prvEndpoints
}

// flattenProcessArgs This function is used to flatten the admin.ClusterDescriptionProcessArgs struct to the ProcessArgs struct
func flattenProcessArgs(p *admin.ClusterDescriptionProcessArgs) *ProcessArgs {
	// Create a new ProcessArgs struct and map the fields from the ClusterDescriptionProcessArgs struct to it
	return &ProcessArgs{
		DefaultReadConcern:               p.DefaultReadConcern,
		DefaultWriteConcern:              p.DefaultWriteConcern,
		FailIndexKeyTooLong:              p.FailIndexKeyTooLong,
		JavascriptEnabled:                p.JavascriptEnabled,
		MinimumEnabledTLSProtocol:        p.MinimumEnabledTlsProtocol,
		NoTableScan:                      p.NoTableScan,
		OplogSizeMB:                      p.OplogSizeMB,
		SampleSizeBIConnector:            p.SampleSizeBIConnector,
		SampleRefreshIntervalBIConnector: p.SampleRefreshIntervalBIConnector,
		OplogMinRetentionHours:           p.OplogMinRetentionHours,
	}
}

// flattenEndpoints This function is used to flatten the admin.ClusterDescriptionConnectionStringsPrivateEndpointEndpoint slice to the Endpoint slice
func flattenEndpoints(eps []admin.ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) []Endpoint {
	var endPoints []Endpoint

	// Loop through each endpoint in the input slice
	for ind := range eps {
		// Create a new Endpoint struct and map the fields from the ClusterDescriptionConnectionStringsPrivateEndpointEndpoint struct to it
		ep := Endpoint{
			EndpointID:   eps[ind].EndpointId,
			ProviderName: eps[ind].ProviderName,
			Region:       eps[ind].Region,
		}

		// Append the new struct to the output slice
		endPoints = append(endPoints, ep)
	}

	// Return the output slice
	return endPoints
}

// flattenLabels This function is used to flatten the admin.ComponentLabel slice to the Labels slice
func flattenLabels(clusterLabels []admin.ComponentLabel) []Labels {
	// Create a new slice of Labels with the same length as the input slice
	labels := make([]Labels, len(clusterLabels))

	// Loop through each label in the input slice
	for i := range clusterLabels {
		// Create a new Labels struct and map the fields from the ComponentLabel struct to it
		labels[i] = Labels{
			Key:   clusterLabels[i].Key,
			Value: clusterLabels[i].Value,
		}
	}

	// Return the new slice
	return labels
}

// formatMongoDBMajorVersion This function is used to format the MongoDB major version string
func formatMongoDBMajorVersion(val *string) *string {
	// If the input string already contains a dot, return it as is
	if strings.Contains(*val, ".") {
		return val
	}

	// Convert the input string to a float32, format it to one decimal place, and return it as a string pointer
	ret := fmt.Sprintf("%.1f", cast.ToFloat32(val))
	return &ret
}

// readCluster This function is used to read the cluster data from the API and update the current model
func readCluster(ctx context.Context, client *admin.APIClient, currentModel *Model) (*Model, *http.Response, error) {
	// Call the GetCluster API to get the cluster data
	cluster, res, err := client.MultiCloudClustersApi.GetCluster(ctx, *currentModel.ProjectId, *currentModel.Name).Execute()

	// If there was an error or the response status code is not 200, return the current model, response, and error
	if err != nil || res.StatusCode != 200 {
		return currentModel, res, err
	}

	// Call the setClusterData function to update the current model with the cluster data
	setClusterData(ctx, currentModel, cluster)

	// Call the GetClusterAdvancedConfiguration API to get the advanced configuration data
	processArgs, resp, errr := client.ClustersApi.GetClusterAdvancedConfiguration(ctx, *currentModel.ProjectId, *currentModel.Name).Execute()

	// If there was an error or the response status code is not 200, return the current model, response, and error
	if errr != nil || resp.StatusCode != 200 {
		return currentModel, resp, errr
	}

	// Call the flattenProcessArgs function to flatten the advanced configuration data and update the current model
	currentModel.AdvancedSettings = flattenProcessArgs(processArgs)

	// Return the updated current model, response, and error
	return currentModel, res, err
}

// setClusterData This method sets the cluster details to Model
func setClusterData(ctx context.Context, currentModel *Model, cluster *admin.AdvancedClusterDescription) {
	// If the cluster is nil, return
	if cluster == nil {
		return
	}

	// Map the fields from the AdvancedClusterDescription struct to the currentModel struct
	currentModel.ProjectId = cluster.GroupId
	currentModel.Name = cluster.Name
	currentModel.Id = cluster.Id

	if cluster.BackupEnabled != nil {
		currentModel.BackupEnabled = cluster.BackupEnabled
	}
	if cluster.BiConnector != nil {
		currentModel.BiConnector = flattenBiConnectorConfig(cluster.BiConnector)
	}
	// Readonly
	currentModel.ConnectionStrings = flattenConnectionStrings(cluster.ConnectionStrings)
	if cluster.ClusterType != nil {
		currentModel.ClusterType = cluster.ClusterType
	}
	// Readonly
	createdDate := cluster.CreateDate.Format("2006-01-02 15:04:05")
	currentModel.CreatedDate = &createdDate
	if cluster.DiskSizeGB != nil {
		currentModel.DiskSizeGB = cluster.DiskSizeGB
	}
	if cluster.EncryptionAtRestProvider != nil {
		currentModel.EncryptionAtRestProvider = cluster.EncryptionAtRestProvider
	}
	if cluster.Labels != nil {
		currentModel.Labels = flattenLabels(cluster.Labels)
	}
	if cluster.MongoDBMajorVersion != nil {
		currentModel.MongoDBMajorVersion = cluster.MongoDBMajorVersion
	}
	// Readonly
	currentModel.MongoDBVersion = cluster.MongoDBVersion

	if cluster.Paused != nil {
		currentModel.Paused = cluster.Paused
	}
	if cluster.PitEnabled != nil {
		currentModel.PitEnabled = cluster.PitEnabled
	}
	if cluster.RootCertType != nil {
		currentModel.RootCertType = cluster.RootCertType
	}
	if cluster.ReplicationSpecs != nil {
		currentModel.ReplicationSpecs = flattenReplicationSpecs(ctx, cluster.ReplicationSpecs)
	}
	// Readonly
	currentModel.StateName = cluster.StateName
	if cluster.VersionReleaseSystem != nil {
		currentModel.VersionReleaseSystem = cluster.VersionReleaseSystem
	}

	currentModel.TerminationProtectionEnabled = cluster.TerminationProtectionEnabled
}

// createClusterRequest creates the ClusterRequest from the Model
func createClusterRequest(ctx context.Context, currentModel *Model) (*admin.AdvancedClusterDescription, error) {
	// Create a new ClusterRequest struct with the Name and ReplicationSpecs fields from the currentModel struct
	clusterRequest := &admin.AdvancedClusterDescription{
		Name:             currentModel.Name,
		ReplicationSpecs: expandReplicationSpecs(ctx, currentModel.ReplicationSpecs),
	}

	// Map the EncryptionAtRestProvider, ClusterType, BackupEnabled, BiConnector, DiskSizeGB, Labels, MongoDBMajorVersion, PitEnabled, VersionReleaseSystem, and RootCertType fields from the currentModel struct to the clusterRequest struct
	if currentModel.EncryptionAtRestProvider != nil {
		clusterRequest.EncryptionAtRestProvider = currentModel.EncryptionAtRestProvider
	}
	if currentModel.ClusterType != nil {
		clusterRequest.ClusterType = currentModel.ClusterType
	}
	if currentModel.BackupEnabled != nil {
		clusterRequest.BackupEnabled = currentModel.BackupEnabled
	}
	if currentModel.BiConnector != nil {
		clusterRequest.BiConnector = expandBiConnector(currentModel.BiConnector)
	}
	if currentModel.DiskSizeGB != nil {
		clusterRequest.DiskSizeGB = currentModel.DiskSizeGB
	}
	if len(currentModel.Labels) > 0 {
		clusterRequest.Labels = expandLabelSlice(currentModel.Labels)
	}
	if currentModel.MongoDBMajorVersion != nil {
		clusterRequest.MongoDBMajorVersion = formatMongoDBMajorVersion(currentModel.MongoDBMajorVersion)
	}
	if currentModel.PitEnabled != nil {
		clusterRequest.PitEnabled = currentModel.PitEnabled
	}
	if currentModel.VersionReleaseSystem != nil {
		clusterRequest.VersionReleaseSystem = currentModel.VersionReleaseSystem
	}
	if currentModel.RootCertType != nil {
		clusterRequest.RootCertType = currentModel.RootCertType
	}

	// Set the TerminationProtectionEnabled field of the clusterRequest struct to the TerminationProtectionEnabled field of the currentModel struct
	clusterRequest.TerminationProtectionEnabled = currentModel.TerminationProtectionEnabled

	// Return the clusterRequest struct and nil error
	return clusterRequest, nil
}

// validateDefaultLabel This method validates if the default label is present in the Labels slice of the Model. If not, it appends the default label to the slice.
func (m *Model) validateDefaultLabel() {
	// If the default label is not present in the Labels slice, append it to the slice
	if !containsLabelOrKey(m.Labels, defaultLabel) {
		m.Labels = append(m.Labels, defaultLabel)
	}
}

// containsLabelOrKey This function checks if the given item is present in the given list of Labels. It returns true if the item is present, false otherwise.
func containsLabelOrKey(list []Labels, item Labels) bool {
	// Iterate over the list of Labels and check if the current item is equal to the given item or has the same key as the given item
	for _, v := range list {
		if reflect.DeepEqual(v, item) || *v.Key == *item.Key {
			return true
		}
	}

	// If the item is not found, return false
	return false
}
