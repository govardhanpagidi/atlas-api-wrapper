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
	"context"
	"encoding/json"
	"fmt"
	"github.com/atlas-api-helper/resources/profile"
	"github.com/atlas-api-helper/util"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/constants"
	"github.com/atlas-api-helper/util/logger"
	"github.com/atlas-api-helper/util/validator"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cast"
	"go.mongodb.org/atlas-sdk/v20230201002/admin"
	"go.mongodb.org/atlas/mongodbatlas"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const (
	LabelError      = "you should not set `Infrastructure Tool` label, it is used for internal purposes"
	CallBackSeconds = 40
)

var defaultLabel = Labels{Key: aws.String("Infrastructure Tool"), Value: aws.String("MongoDB Atlas CloudFormation Provider")}

var CreateRequiredFields = []string{constants.ProjectID, constants.Name}
var ReadRequiredFields = []string{constants.ProjectID, constants.Name}
var UpdateRequiredFields = []string{constants.ProjectID, constants.Name}
var DeleteRequiredFields = []string{constants.ProjectID, constants.Name}
var ListRequiredFields = []string{constants.ProjectID}

func setup() {
	util.SetupLogger("mongodb-atlas-cluster")
}

func castNO64(i *int64) *int {
	x := cast.ToInt(&i)
	return &x
}

func cast64(i *int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}

// validateModel inputs based on the method
func validateModel(fields []string, model *Model) error {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()

	_, _ = logger.Debugf("Create cluster model : %+v", currentModel)

	modelValidation := validateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	_, _ = logger.Debugf("Cluster create projectId: %s, clusterName: %s", *currentModel.ProjectId, *currentModel.Name)

	// Callback

	var err error

	currentModel.validateDefaultLabel()

	// Prepare cluster request
	clusterRequest, err := setClusterRequest(currentModel, err)
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      err.Error(),
		}
	}

	// Create Cluster
	cluster, res, err := client.MultiCloudClustersApi.CreateCluster(ctx, *currentModel.ProjectId, clusterRequest).Execute()

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
		Response:       cluster,
		HttpStatusCode: res.StatusCode,
		HttpError:      "",
	}
}

// Read handles the Read event from the Cloudformation service.
func Read(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf("Read() currentModel:%+v", currentModel)

	modelValidation := validateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}
	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	// Read call
	model, resp, err := readCluster(context.Background(), client, currentModel)
	if err != nil {
		_, _ = logger.Warnf("error cluster get- err:%+v resp:%+v", err, resp)
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: resp.StatusCode,
			HttpError:      err.Error(),
		}
	}
	return atlasresponse.AtlasRespone{
		Response:       model,
		HttpStatusCode: resp.StatusCode,
		HttpError:      "",
	}
}

// Update handles the Update event from the Cloudformation service.
func Update(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf("Update() currentModel:%+v", currentModel)

	modelValidation := validateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      modelValidation.Error(),
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	currentModel.validateDefaultLabel()

	// Update Cluster
	model, resp, err := updateCluster(ctx, client, currentModel)
	if err != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: resp.StatusCode,
			HttpError:      err.Error(),
		}
	}

	var state string
	if model.StateName != nil {
		state = *model.StateName
	}
	_, _ = logger.Debugf("state: %+v", state)
	return atlasresponse.AtlasRespone{
		Response:       model,
		HttpStatusCode: resp.StatusCode,
		HttpError:      "",
	}
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf("Delete() currentModel:%+v", currentModel)

	modelValidation := validateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "Validation error",
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
	if peErr != nil {
		_, _ = logger.Warnf("CreateMongoDBClient error: %v", peErr.Error())
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      peErr.Error(),
		}
	}

	//options := &mongodbatlas.DeleteAdvanceClusterOptions{RetainBackups: util.Pointer(false)}
	retainBackup := false
	args := admin.DeleteClusterApiParams{
		GroupId:       *currentModel.ProjectId,
		ClusterName:   *currentModel.Name,
		RetainBackups: &retainBackup,
	}

	response, err := client.MultiCloudClustersApi.DeleteClusterWithParams(ctx, &args).Execute()
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
func List(ctx context.Context, currentModel *Model) atlasresponse.AtlasRespone {
	setup()
	_, _ = logger.Debugf("List() currentModel:%+v", currentModel)

	modelValidation := validateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return atlasresponse.AtlasRespone{
			Response:       nil,
			HttpStatusCode: http.StatusBadRequest,
			HttpError:      "Validation error",
		}
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	client, peErr := util.NewMongoDBSDKClient(ctx)
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
		GroupId:      *currentModel.ProjectId,
		ItemsPerPage: &itemsPerPage,
		PageNum:      &pageNum,
	}
	clustersResponse, res, err := client.MultiCloudClustersApi.ListClustersWithParams(ctx, &args).Execute()
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
		processArgs, resp, err2 := client.ClustersApi.GetClusterAdvancedConfiguration(ctx, *model.ProjectId, *model.Name).Execute()
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

func isClusterInTargetState(client *mongodbatlas.Client, projectID, clusterName, targetState string) (isReady bool, stateName string, mongoCluster *mongodbatlas.AdvancedCluster, err error) {
	cluster, resp, err := client.AdvancedClusters.Get(context.Background(), projectID, clusterName)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return constants.DeletedState == targetState, constants.DeletedState, nil, nil
		}
		return false, constants.Error, nil, fmt.Errorf("error fetching cluster info (%s): %s", clusterName, err)
	}
	_, _ = logger.Debugf("Cluster state: %s, targetState : %s", cluster.StateName, targetState)
	return cluster.StateName == targetState, cluster.StateName, cluster, nil
}

func expandAdvancedSettings(processArgs ProcessArgs) *admin.ClusterDescriptionProcessArgs {
	var args admin.ClusterDescriptionProcessArgs

	if processArgs.DefaultReadConcern != nil {
		args.DefaultReadConcern = processArgs.DefaultReadConcern
	}
	if processArgs.FailIndexKeyTooLong != nil {
		args.FailIndexKeyTooLong = processArgs.FailIndexKeyTooLong
	}
	if processArgs.DefaultWriteConcern != nil {
		args.DefaultWriteConcern = processArgs.DefaultWriteConcern
	}
	if processArgs.JavascriptEnabled != nil {
		args.JavascriptEnabled = processArgs.JavascriptEnabled
	}
	if processArgs.MinimumEnabledTLSProtocol != nil {
		args.MinimumEnabledTlsProtocol = processArgs.MinimumEnabledTLSProtocol
	}
	if processArgs.NoTableScan != nil {
		args.NoTableScan = processArgs.NoTableScan
	}
	if processArgs.OplogSizeMB != nil {
		args.OplogSizeMB = processArgs.OplogSizeMB
	}
	if processArgs.SampleSizeBIConnector != nil {
		args.SampleSizeBIConnector = processArgs.SampleSizeBIConnector
	}
	if processArgs.SampleRefreshIntervalBIConnector != nil {
		args.SampleRefreshIntervalBIConnector = processArgs.SampleRefreshIntervalBIConnector
	}

	if processArgs.OplogMinRetentionHours != nil {
		args.OplogMinRetentionHours = processArgs.OplogMinRetentionHours
	}

	return &args
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

func updateCluster(ctx context.Context, client *admin.APIClient, currentModel *Model) (*Model, *http.Response, error) {
	clusterRequest := admin.AdvancedClusterDescription{}

	if currentModel.BackupEnabled != nil {
		clusterRequest.BackupEnabled = currentModel.BackupEnabled
	}

	if currentModel.BiConnector != nil {
		clusterRequest.BiConnector = expandBiConnector(currentModel.BiConnector)
	}

	if currentModel.ClusterType != nil {
		clusterRequest.ClusterType = currentModel.ClusterType
	}

	if currentModel.DiskSizeGB != nil {
		clusterRequest.DiskSizeGB = currentModel.DiskSizeGB
	}

	if currentModel.EncryptionAtRestProvider != nil {
		clusterRequest.EncryptionAtRestProvider = currentModel.EncryptionAtRestProvider
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

	if currentModel.ReplicationSpecs != nil {
		clusterRequest.ReplicationSpecs = expandReplicationSpecs(currentModel.ReplicationSpecs)
	}

	if currentModel.RootCertType != nil {
		clusterRequest.RootCertType = currentModel.RootCertType
	}

	if currentModel.VersionReleaseSystem != nil {
		clusterRequest.VersionReleaseSystem = currentModel.VersionReleaseSystem
	}
	clusterRequest.TerminationProtectionEnabled = currentModel.TerminationProtectionEnabled

	_, res, err := updateClusterSettings(currentModel, client, *currentModel.ProjectId, ctx)
	if err != nil {
		return nil, res, err
	}
	_, _ = logger.Debugf("params : %+v %+v %+v", ctx, client, clusterRequest)
	cluster, resp, err := client.MultiCloudClustersApi.UpdateCluster(ctx, *currentModel.ProjectId, *currentModel.Name, &clusterRequest).Execute()
	if cluster != nil {
		currentModel.StateName = cluster.StateName
	}

	return currentModel, resp, err
}

func updateAdvancedCluster(ctx context.Context, conn *admin.APIClient,
	request *admin.AdvancedClusterDescription, projectID, name string) (*admin.AdvancedClusterDescription, *http.Response, error) {
	cluster, response, err := conn.MultiCloudClustersApi.UpdateCluster(ctx, projectID, name, request).Execute()
	if err != nil {
		return nil, response, err
	}
	return cluster, response, err
}

func updateClusterSettings(currentModel *Model, client *admin.APIClient,
	projectID string, ctx context.Context) (*Model, *http.Response, error) {

	cluster, res, err := client.MultiCloudClustersApi.GetCluster(ctx, projectID, *currentModel.Name).Execute()
	if err != nil {
		return nil, res, err
	}
	// Update advanced configuration
	if currentModel.AdvancedSettings != nil {
		_, _ = logger.Debugf("AdvancedSettings: %+v", *currentModel.AdvancedSettings)

		advancedConfig := expandAdvancedSettings(*currentModel.AdvancedSettings)
		args, res, err := client.ClustersApi.UpdateClusterAdvancedConfiguration(ctx, projectID, *cluster.Name, advancedConfig).Execute()
		if err != nil {
			return currentModel, res, err
		}
		currentModel.AdvancedSettings = flattenProcessArgs(args)
	}
	// Update pause
	if (currentModel.Paused != nil) && (*currentModel.Paused != *cluster.Paused) {

		cluster, res, err := updateAdvancedCluster(ctx, client, &admin.AdvancedClusterDescription{Paused: currentModel.Paused}, projectID, *currentModel.Name)
		if err != nil {
			_, _ = logger.Warnf("Cluster Pause - error: %+v", err)
			return currentModel, res, err
		}
		setClusterData(currentModel, cluster)
	}

	jsonStr, _ := json.Marshal(currentModel)
	_, _ = logger.Debugf("Cluster Response --- value: %s ", jsonStr)
	return currentModel, res, nil
}

func setClusterRequest(currentModel *Model, err error) (*admin.AdvancedClusterDescription, error) {
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
