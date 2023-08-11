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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var defaultLabel = Labels{Key: aws.String("Infrastructure Tool"), Value: aws.String("MongoDB Atlas CloudFormation Provider")}

var CreateRequiredFields = []string{constants.ProjectID, constants.PrivateKey, constants.PublicKey, constants.ClusterSize, constants.DBUserName}
var ReadRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.PublicKey, constants.PrivateKey}
var DeleteRequiredFields = []string{constants.ProjectID, constants.ClusterName, constants.PublicKey, constants.PrivateKey}
var ListRequiredFields = []string{constants.ProjectID, constants.PublicKey, constants.PrivateKey}

// setup initializes logger
func setup() {
	util.SetupLogger("mongodb-atlas-cluster")
}

func cast64(i *int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}

// validateModel inputs based on the method
func validateModel(fields []string, model *InputModel) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()

	_, _ = logger.Debugf("Create cluster model : %+v", inputModel)
	// Validate required fields in the request

	modelValidation := validateModel(CreateRequiredFields, inputModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}
	client, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}
	_, res, projectErr := client.ProjectsApi.GetProject(context.Background(), cast.ToString(inputModel.ProjectId)).Execute()
	if projectErr != nil {
		_, _ = logger.Warnf("Get Project error: %v", projectErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      projectErr.Error(),
		}
	}

	/*	endPoints, res, endpointerr := client.PrivateEndpointServicesApi.ListPrivateEndpointServices(ctx, inputModel.ProjectId, "AWS").Execute()
				if endpointerr != nil {
			_, _ = logger.Warnf("Get PrivateEndpoint error: %v", endpointerr.Error())
					return atlasresponse.AtlasRespone{
						Response:       nil,
						HttpStatusCode: res.StatusCode,
						HttpError:      endpointerr.Error(),
					}
				}

				count := len(endPoints)
				if count == 0 {
		_, _ = logger.Warnf("Get PrivateEndpoint Not Configured error: %v", endpointerr.Error())
					return atlasresponse.AtlasRespone{
						Response:       nil,
						HttpStatusCode: http.StatusInternalServerError,
						HttpError:      "No Entpoints configured for this project",
					}
				}
				_, _ = logger.Debugf("Cluster create projectId: %s, clusterName: %s", inputModel.ProjectId, inputModel.ClusterName)
	*/
	currentModel, err := loadCurrentModel(*inputModel)
	if err != nil {
		_, _ = logger.Warnf("Create Current Model error: %v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusInternalServerError,
			HttpError:      err.Error(),
		}
	}
	currentModel.validateDefaultLabel()

	// Prepare cluster request
	clusterRequest, err := setClusterRequest(&currentModel)
	if err != nil {
		_, _ = logger.Warnf("Create Cluster Request error: %v", err.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}

	// Create Cluster
	cluster, res, err := client.MultiCloudClustersApi.CreateCluster(ctx, cast.ToString(currentModel.ProjectId), clusterRequest).Execute()

	if err != nil {
		_, _ = logger.Warnf("Create - Cluster.Create() - error: %+v", err)
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}
	currentModel.StateName = cluster.StateName
	return atlasresponse.AtlasRespone{
		Response:       currentModel,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

// loadCurrentModel This method loads the config.json file from project path
func loadCurrentModel(model InputModel) (Model, error) {
	var currentModel Model
	var ClusterConfig map[string]Model

	content, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return currentModel, err
	}

	// Now let's unmarshall the data into `payload`
	err = json.Unmarshal(content, &ClusterConfig)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return currentModel, err
	}
	key := extractClusterKey(model)
	clusterConfig, ok := ClusterConfig[key]
	if ok {
		currentModel = clusterConfig
	} else {
		return currentModel, errors.New("provided Cluster Size is Invalid: " + *model.ClusterSize)
	}
	if model.ClusterName != nil {
		currentModel.Name = model.ClusterName
	} else {
		currentModel.Name = generateClusterName(model)
	}
	if model.MongoDBVersion != nil {
		currentModel.MongoDBVersion = model.MongoDBVersion
	}
	return currentModel, nil
}

// extractClusterKey This method generates the key using which the config is fetched
func extractClusterKey(model InputModel) string {
	var configKey bytes.Buffer
	configKey.WriteString(strings.ToLower(*model.ClusterSize))
	configKey.WriteString("-")
	configKey.WriteString(strings.ToLower(*model.CloudProvider))
	key := configKey.String()
	return key
}

// generateClusterName This method generates the cluster name which is then assigned to the created cluster
func generateClusterName(model InputModel) *string {
	clusterNamePrefix := extractClusterKey(model)
	toRet := clusterNamePrefix + "-" + uuid.NewString()
	return &toRet
}

// Read handles the Read event from the Cloudformation service.
func Read(inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf("Read() currentModel:%+v", inputModel)

	modelValidation := validateModel(ReadRequiredFields, inputModel)
	if modelValidation != nil {
		_, _ = logger.Warnf("Input Validation error: %v", modelValidation.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}

	client, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	// Read call
	model, resp, err := readCluster(context.Background(), client, &Model{ProjectId: inputModel.ProjectId, Name: inputModel.ClusterName})
	if err != nil {
		_, _ = logger.Warnf("error cluster get- err:%+v resp:%+v", err, resp)
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: resp.StatusCode,
			HttpError:      err.Error(),
		}
	}
	return atlasresponse.AtlasRespone{
		Response:       "clusterStatus: " + *model.StateName,
		HttpStatusCode: resp.StatusCode,
		HttpError:      "",
	}
}

// Delete This method deletes the cluster based on the clusterName
func Delete(inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf("Delete() currentModel:%+v", inputModel)

	modelValidation := validateModel(DeleteRequiredFields, inputModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "Validation error",
		}
	}

	client, peErr := util.NewMongoDBSDKClient(cast.ToString(inputModel.PublicKey), cast.ToString(inputModel.PrivateKey))
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	retainBackup := false
	args := admin.DeleteClusterApiParams{
		GroupId:       *inputModel.ProjectId,
		ClusterName:   *inputModel.ClusterName,
		RetainBackups: &retainBackup,
	}

	response, err := client.MultiCloudClustersApi.DeleteClusterWithParams(context.Background(), &args).Execute()
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: response.StatusCode,
			HttpError:      err.Error(),
		}
	}

	return atlasresponse.AtlasRespone{
		Response:       nil,
		HttpStatusCode: response.StatusCode,
		HttpError:      "",
	}
}

// List handles the List event from the Cloudformation service.

func List(inputModel *InputModel) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf("List() currentModel:%+v", inputModel)

	modelValidation := validateModel(ListRequiredFields, inputModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "Validation error",
		}
	}

	client, peErr := util.NewMongoDBSDKClient(*inputModel.PublicKey, *inputModel.PrivateKey)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}
	// List call
	itemsPerPage := 100
	pageNum := 1
	args := admin.ListClustersApiParams{
		GroupId:      *inputModel.ProjectId,
		ItemsPerPage: &itemsPerPage,
		PageNum:      &pageNum,
	}
	clustersResponse, res, err := client.MultiCloudClustersApi.ListClustersWithParams(context.Background(), &args).Execute()
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: res.StatusCode,
			HttpError:      err.Error(),
		}
	}

	models := make([]*Model, *clustersResponse.TotalCount)
	for i := range clustersResponse.Results {
		model := &Model{}
		mapClusterToModel(model, &clustersResponse.Results[i])
		// Call AdvancedSettings
		processArgs, resp, err2 := client.ClustersApi.GetClusterAdvancedConfiguration(context.Background(), *model.ProjectId, *model.Name).Execute()
		if err2 != nil {
			return atlasresponse.AtlasRespone{
				Response:       nil,
				HttpStatusCode: resp.StatusCode,
				HttpError:      err2.Error(),
			}
		}

		model.AdvancedSettings = flattenProcessArgs(processArgs)
		models = append(models, model)
	}
	return atlasresponse.AtlasRespone{
		Response:       models,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

// mapClusterToModel This method is used to map the cluster object returned from the mongo client to our model
func mapClusterToModel(model *Model, cluster *admin.AdvancedClusterDescription) {
	model.Id = cluster.Id
	model.ProjectId = cluster.GroupId
	model.Name = cluster.Name
	model.BackupEnabled = cluster.BackupEnabled
	model.BiConnector = flattenBiConnectorConfig(cluster.BiConnector)
	model.ConnectionStrings = flattenConnectionStrings(cluster.ConnectionStrings)
	model.ClusterType = cluster.ClusterType
	createdDate := cluster.CreateDate.Format("2006-01-02 15:04:05")
	model.CreatedDate = &createdDate
	model.DiskSizeGB = cluster.DiskSizeGB
	model.EncryptionAtRestProvider = cluster.EncryptionAtRestProvider
	model.Labels = flattenLabels(cluster.Labels)
	model.MongoDBMajorVersion = cluster.MongoDBMajorVersion
	model.MongoDBVersion = cluster.MongoDBVersion
	model.Paused = cluster.Paused
	model.PitEnabled = cluster.PitEnabled
	model.RootCertType = cluster.RootCertType
	model.ReplicationSpecs = flattenReplicationSpecs(cluster.ReplicationSpecs)
	model.StateName = cluster.StateName
	model.VersionReleaseSystem = cluster.VersionReleaseSystem
}

func (m *Model) HasAdvanceSettings() bool {
	/*This logic is because of a bug un Cloud Formation, when we return in_progress in the CREATE
	,the second time the CREATE gets executed
	it returns the AdvancedSettings is not nil but its fields are nil*/
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

func containsLabelOrKey(list []Labels, item Labels) bool {
	for _, v := range list {
		if reflect.DeepEqual(v, item) || *v.Key == *item.Key {
			return true
		}
	}

	return false
}

func expandBiConnector(biConnector *BiConnector) *admin.BiConnector {
	if biConnector == nil {
		return nil
	}
	return &admin.BiConnector{
		Enabled:        biConnector.Enabled,
		ReadPreference: biConnector.ReadPreference,
	}
}

func expandReplicationSpecs(replicationSpecs []AdvancedReplicationSpec) []admin.ReplicationSpec {
	var rSpecs []admin.ReplicationSpec

	for i := range replicationSpecs {
		var numShards int

		//repId := cast.ToString(replicationSpecs[i].ID)
		rSpec := admin.ReplicationSpec{
			Id:            nil,
			NumShards:     &numShards,
			RegionConfigs: expandRegionsConfig(replicationSpecs[i].AdvancedRegionConfigs),
		}

		if replicationSpecs[i].NumShards != nil {
			rSpec.NumShards = replicationSpecs[i].NumShards
		}
		if replicationSpecs[i].ZoneName != nil {
			zoneName := cast.ToString(replicationSpecs[i].ZoneName)
			rSpec.ZoneName = &zoneName
		}
		rSpecs = append(rSpecs, rSpec)
	}

	fmt.Printf("specs: len %d %+v", len(replicationSpecs), rSpecs)
	return rSpecs
}

func expandAutoScaling(scaling *AdvancedAutoScaling) *admin.AdvancedAutoScalingSettings {
	advAutoScaling := &admin.AdvancedAutoScalingSettings{}
	if scaling == nil {
		return nil
	}
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
	if scaling.DiskGB != nil {
		advAutoScaling.DiskGB = &admin.DiskGBAutoScaling{
			Enabled: scaling.DiskGB.Enabled,
		}
	}
	return advAutoScaling
}

func expandRegionsConfig(regionConfigs []AdvancedRegionConfig) []admin.CloudRegionConfig {
	var regionsConfigs []admin.CloudRegionConfig
	for _, regionCfg := range regionConfigs {
		regionsConfigs = append(regionsConfigs, expandRegionConfig(regionCfg))
	}
	return regionsConfigs
}

func expandRegionConfig(regionCfg AdvancedRegionConfig) admin.CloudRegionConfig {
	var region string
	if regionCfg.RegionName != nil {
		region = *regionCfg.RegionName
	}

	providerName := constants.AWS
	if regionCfg.ProviderName != nil {
		providerName = *regionCfg.ProviderName
	}

	advRegionConfig := admin.CloudRegionConfig{
		ProviderName: &providerName,
		RegionName:   &region,
		Priority:     regionCfg.Priority,
	}

	if regionCfg.AutoScaling != nil {
		advRegionConfig.AutoScaling = expandAutoScaling(regionCfg.AutoScaling)
	}
	if regionCfg.AnalyticsAutoScaling != nil {
		advRegionConfig.AnalyticsAutoScaling = expandAutoScaling(regionCfg.AnalyticsAutoScaling)
	}
	if regionCfg.AnalyticsSpecs != nil {
		advRegionConfig.AnalyticsSpecs = expandRegionConfigSpec(regionCfg.AnalyticsSpecs)
	}
	if regionCfg.ElectableSpecs != nil {
		advRegionConfig.ElectableSpecs = expandRegionConfigSpecElectableSpecs(regionCfg.ElectableSpecs)
	}
	if regionCfg.ReadOnlySpecs != nil {
		advRegionConfig.ReadOnlySpecs = expandRegionConfigSpec(regionCfg.ReadOnlySpecs)
	}
	if regionCfg.BackingProviderName != nil {
		advRegionConfig.BackingProviderName = regionCfg.BackingProviderName
	}
	return advRegionConfig
}

func expandRegionConfigSpec(spec *Specs) *admin.DedicatedHardwareSpec {
	if spec == nil {
		return nil
	}
	var ebsVolumeType string
	var instanceSize string
	if spec.EbsVolumeType != nil {
		ebsVolumeType = *spec.EbsVolumeType
	}
	if spec.InstanceSize != nil {
		instanceSize = *spec.InstanceSize
	}
	var val int
	if spec.DiskIOPS != nil {
		v, err := strconv.Atoi(*spec.DiskIOPS)
		if err == nil {
			val = v
		}
		_, _ = logger.Debugf("set diskIops %d", val)
	}
	return &admin.DedicatedHardwareSpec{
		DiskIOPS:      &val,
		EbsVolumeType: &ebsVolumeType,
		InstanceSize:  &instanceSize,
		NodeCount:     spec.NodeCount,
	}
}

func expandRegionConfigSpecElectableSpecs(spec *Specs) *admin.HardwareSpec {
	if spec == nil {
		return nil
	}
	var ebsVolumeType string
	var instanceSize string
	if spec.EbsVolumeType != nil {
		ebsVolumeType = *spec.EbsVolumeType
	}
	if spec.InstanceSize != nil {
		instanceSize = *spec.InstanceSize
	}
	var val int
	if spec.DiskIOPS != nil {
		v, err := strconv.Atoi(*spec.DiskIOPS)
		if err == nil {
			val = v
		}
		_, _ = logger.Debugf("set diskIops %d", val)
	}
	return &admin.HardwareSpec{
		DiskIOPS:      &val,
		EbsVolumeType: &ebsVolumeType,
		InstanceSize:  &instanceSize,
		NodeCount:     spec.NodeCount,
	}
}

func expandLabelSlice(labels []Labels) []admin.ComponentLabel {
	res := make([]admin.ComponentLabel, len(labels))

	for i := range labels {
		var key string
		if labels[i].Key != nil {
			key = *labels[i].Key
		}
		var value string
		if labels[i].Key != nil {
			value = *labels[i].Value
		}
		res[i] = admin.ComponentLabel{
			Key:   &key,
			Value: &value,
		}
	}
	return res
}

func flattenAutoScaling(scaling *admin.AdvancedAutoScalingSettings) *AdvancedAutoScaling {
	if scaling == nil {
		return nil
	}
	advAutoScaling := &AdvancedAutoScaling{}

	if scaling.DiskGB != nil {
		advAutoScaling.DiskGB = &DiskGB{Enabled: scaling.DiskGB.Enabled}
	}
	if scaling.Compute != nil {
		compute := &Compute{}
		if scaling.Compute.Enabled != nil {
			compute.Enabled = scaling.Compute.Enabled
		}
		if scaling.Compute.ScaleDownEnabled != nil {
			compute.ScaleDownEnabled = scaling.Compute.ScaleDownEnabled
		}
		if *scaling.Compute.MinInstanceSize != "" {
			compute.MinInstanceSize = scaling.Compute.MinInstanceSize
		}
		if *scaling.Compute.MaxInstanceSize != "" {
			compute.MaxInstanceSize = scaling.Compute.MaxInstanceSize
		}

		advAutoScaling.Compute = compute
	}
	return advAutoScaling
}

func flattenReplicationSpecs(replicationSpecs []admin.ReplicationSpec) []AdvancedReplicationSpec {
	var rSpecs []AdvancedReplicationSpec
	for ind := range replicationSpecs {
		configs := replicationSpecs[ind].RegionConfigs
		rSpec := AdvancedReplicationSpec{
			ID:                    replicationSpecs[ind].Id,
			NumShards:             replicationSpecs[ind].NumShards,
			ZoneName:              replicationSpecs[ind].ZoneName,
			AdvancedRegionConfigs: flattenRegionsConfig(configs),
		}
		rSpecs = append(rSpecs, rSpec)
	}
	fmt.Printf("specs: len %d %+v", len(replicationSpecs), rSpecs)
	return rSpecs
}

func flattenRegionsConfig(regionConfigs []admin.CloudRegionConfig) []AdvancedRegionConfig {
	var regionsConfigs []AdvancedRegionConfig
	for i := range regionConfigs {
		regionsConfigs = append(regionsConfigs, flattenRegionConfig(&regionConfigs[i]))
	}
	return regionsConfigs
}

func flattenRegionConfig(regionCfg *admin.CloudRegionConfig) AdvancedRegionConfig {
	advRegConfig := AdvancedRegionConfig{
		AutoScaling:          flattenAutoScaling(regionCfg.AutoScaling),
		AnalyticsAutoScaling: flattenAutoScaling(regionCfg.AnalyticsAutoScaling),
		RegionName:           regionCfg.RegionName,
		Priority:             regionCfg.Priority,
	}
	if regionCfg.AnalyticsSpecs != nil {
		advRegConfig.AnalyticsSpecs = flattenRegionConfigSpec(regionCfg.AnalyticsSpecs)
	}
	if regionCfg.ElectableSpecs != nil {
		advRegConfig.ElectableSpecs = flattenRegionConfigHardwareSpecSpec(regionCfg.ElectableSpecs)
	}

	if regionCfg.ReadOnlySpecs != nil {
		advRegConfig.ReadOnlySpecs = flattenRegionConfigSpec(regionCfg.ReadOnlySpecs)
	}

	return advRegConfig
}

func flattenRegionConfigHardwareSpecSpec(spec *admin.HardwareSpec) *Specs {
	if spec == nil {
		return nil
	}
	var diskIops string
	if spec.DiskIOPS != nil {
		diskIops = strconv.FormatInt(*cast64(spec.DiskIOPS), 10)
		_, _ = logger.Debugf("get diskIops %s", diskIops)
	}

	return &Specs{
		DiskIOPS:      &diskIops,
		EbsVolumeType: spec.EbsVolumeType,
		InstanceSize:  spec.InstanceSize,
		NodeCount:     spec.NodeCount,
	}
}

func flattenRegionConfigSpec(spec *admin.DedicatedHardwareSpec) *Specs {
	if spec == nil {
		return nil
	}
	var diskIops string
	if spec.DiskIOPS != nil {
		diskIops = strconv.FormatInt(*cast64(spec.DiskIOPS), 10)
		_, _ = logger.Debugf("get diskIops %s", diskIops)
	}

	return &Specs{
		DiskIOPS:      &diskIops,
		EbsVolumeType: spec.EbsVolumeType,
		InstanceSize:  spec.InstanceSize,
		NodeCount:     spec.NodeCount,
	}
}

func flattenBiConnectorConfig(biConnector *admin.BiConnector) *BiConnector {
	if biConnector == nil {
		return nil
	}

	return &BiConnector{
		ReadPreference: biConnector.ReadPreference,
		Enabled:        biConnector.Enabled,
	}
}

func flattenConnectionStrings(clusterConnStrings *admin.ClusterConnectionStrings) (connStrings *ConnectionStrings) {
	if clusterConnStrings != nil {
		connStrings = &ConnectionStrings{
			Standard:        clusterConnStrings.Standard,
			StandardSrv:     clusterConnStrings.StandardSrv,
			Private:         clusterConnStrings.Private,
			PrivateSrv:      clusterConnStrings.PrivateSrv,
			PrivateEndpoint: flattenPrivateEndpoint(clusterConnStrings.PrivateEndpoint),
		}
	}
	return
}

func flattenPrivateEndpoint(pes []admin.ClusterDescriptionConnectionStringsPrivateEndpoint) []PrivateEndpoint {
	var prvEndpoints []PrivateEndpoint
	if pes == nil {
		return prvEndpoints
	}
	for ind := range pes {
		pe := PrivateEndpoint{
			ConnectionString:    pes[ind].ConnectionString,
			SRVConnectionString: pes[ind].SrvConnectionString,
			Type:                pes[ind].Type,
			Endpoints:           flattenEndpoints(pes[ind].Endpoints),
		}
		prvEndpoints = append(prvEndpoints, pe)
	}
	return prvEndpoints
}

func flattenProcessArgs(p *admin.ClusterDescriptionProcessArgs) *ProcessArgs {
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

func flattenEndpoints(eps []admin.ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) []Endpoint {
	var endPoints []Endpoint
	for ind := range eps {
		ep := Endpoint{
			EndpointID:   eps[ind].EndpointId,
			ProviderName: eps[ind].ProviderName,
			Region:       eps[ind].Region,
		}
		endPoints = append(endPoints, ep)
	}
	return endPoints
}

func flattenLabels(clusterLabels []admin.ComponentLabel) []Labels {
	labels := make([]Labels, len(clusterLabels))
	for i := range clusterLabels {
		labels[i] = Labels{
			Key:   clusterLabels[i].Key,
			Value: clusterLabels[i].Value,
		}
	}
	return labels
}

func formatMongoDBMajorVersion(val *string) *string {
	if strings.Contains(*val, ".") {
		return val
	}
	ret := fmt.Sprintf("%.1f", cast.ToFloat32(val))
	return &ret
}

func readCluster(ctx context.Context, client *admin.APIClient, currentModel *Model) (*Model, *http.Response, error) {
	cluster, res, err := client.MultiCloudClustersApi.GetCluster(ctx, *currentModel.ProjectId, *currentModel.Name).Execute()
	if err != nil || res.StatusCode != 200 {
		return currentModel, res, err
	}
	setClusterData(currentModel, cluster)

	processArgs, resp, errr := client.ClustersApi.GetClusterAdvancedConfiguration(ctx, *currentModel.ProjectId, *currentModel.Name).Execute()

	if errr != nil || resp.StatusCode != 200 {
		return currentModel, resp, errr
	}
	currentModel.AdvancedSettings = flattenProcessArgs(processArgs)
	return currentModel, res, err
}

// setClusterData This method sets the cluster details to Model
func setClusterData(currentModel *Model, cluster *admin.AdvancedClusterDescription) {
	if cluster == nil {
		return
	}

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
		currentModel.ReplicationSpecs = flattenReplicationSpecs(cluster.ReplicationSpecs)
	}
	// Readonly
	currentModel.StateName = cluster.StateName
	if cluster.VersionReleaseSystem != nil {
		currentModel.VersionReleaseSystem = cluster.VersionReleaseSystem
	}

	currentModel.TerminationProtectionEnabled = cluster.TerminationProtectionEnabled
}

// setClusterRequest creates the ClusterRequest from the Model
func setClusterRequest(currentModel *Model) (*admin.AdvancedClusterDescription, error) {
	// Atlas client
	clusterRequest := &admin.AdvancedClusterDescription{
		Name:             currentModel.Name,
		ReplicationSpecs: expandReplicationSpecs(currentModel.ReplicationSpecs),
	}

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

	clusterRequest.TerminationProtectionEnabled = currentModel.TerminationProtectionEnabled
	return clusterRequest, nil
}

func (m *Model) validateDefaultLabel() {
	if !containsLabelOrKey(m.Labels, defaultLabel) {
		m.Labels = append(m.Labels, defaultLabel)
	}
}
