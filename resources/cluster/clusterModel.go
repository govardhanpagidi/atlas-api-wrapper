// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package cluster

import "fmt"

// Model is autogenerated from the json schema
type Model struct {
	AdvancedSettings             *ProcessArgs              `json:"AdvancedSettings,omitempty"`
	BackupEnabled                *bool                     `json:"BackupEnabled,omitempty"`
	BiConnector                  *BiConnector              `json:"biConnector,omitempty"`
	ClusterType                  *string                   `json:"clusterType,omitempty"`
	CreatedDate                  *string                   `json:"createdDate,omitempty"`
	ConnectionStrings            *ConnectionStrings        `json:"ConnectionStrings,omitempty"`
	DiskSizeGB                   *float64                  `json:"diskSizeGB,omitempty"`
	EncryptionAtRestProvider     *string                   `json:"encryptionAtRestProvider,omitempty"`
	Profile                      *string                   `json:"profile,omitempty"`
	ProjectId                    *string                   `json:"projectId,omitempty"`
	Id                           *string                   `json:"id,omitempty"`
	Labels                       []Labels                  `json:"labels,omitempty"`
	MongoDBMajorVersion          *string                   `json:"mongoDBMajorVersion,omitempty"`
	MongoDBVersion               *string                   `json:"mongoDBVersion,omitempty"`
	Name                         *string                   `json:"name,omitempty"`
	Paused                       *bool                     `json:"paused,omitempty"`
	PitEnabled                   *bool                     `json:"pitEnabled,omitempty"`
	ReplicationSpecs             []AdvancedReplicationSpec `json:"replicationSpecs,omitempty"`
	RootCertType                 *string                   `json:"rootCertType,omitempty"`
	StateName                    *string                   `json:"stateName,omitempty"`
	VersionReleaseSystem         *string                   `json:"versionReleaseSystem,omitempty"`
	TerminationProtectionEnabled *bool                     `json:"terminationProtectionEnabled,omitempty"`
}

// ProcessArgs is autogenerated from the json schema
type ProcessArgs struct {
	DefaultReadConcern               *string  `json:"defaultReadConcern,omitempty"`
	DefaultWriteConcern              *string  `json:"defaultWriteConcern,omitempty"`
	FailIndexKeyTooLong              *bool    `json:"failIndexKeyTooLong,omitempty"`
	JavascriptEnabled                *bool    `json:"javascriptEnabled,omitempty"`
	MinimumEnabledTLSProtocol        *string  `json:"minimumEnabledTLSProtocol,omitempty"`
	NoTableScan                      *bool    `json:"noTableScan,omitempty"`
	OplogSizeMB                      *int     `json:"oplogSizeMB,omitempty"`
	SampleSizeBIConnector            *int     `json:"sampleSizeBIConnector,omitempty"`
	SampleRefreshIntervalBIConnector *int     `json:"sampleRefreshIntervalBIConnector,omitempty"`
	OplogMinRetentionHours           *float64 `json:"oplogMinRetentionHours,omitempty"`
}

// BiConnector is autogenerated from the json schema
type BiConnector struct {
	ReadPreference *string `json:"readPreference,omitempty"`
	Enabled        *bool   `json:"enabled,omitempty"`
}

// ConnectionStrings is autogenerated from the json schema
type ConnectionStrings struct {
	Standard          *string           `json:"standard,omitempty"`
	StandardSrv       *string           `json:"standardSrv,omitempty"`
	Private           *string           `json:"private,omitempty"`
	PrivateSrv        *string           `json:"privateSrv,omitempty"`
	PrivateEndpoint   []PrivateEndpoint `json:"privateEndpoint,omitempty"`
	AwsPrivateLinkSrv *string           `json:"awsPrivateLinkSrv,omitempty"`
	AwsPrivateLink    *string           `json:"awsPrivateLink,omitempty"`
}

// PrivateEndpoint is autogenerated from the json schema
type PrivateEndpoint struct {
	ConnectionString    *string    `json:"connectionString,omitempty"`
	Endpoints           []Endpoint `json:"endpoints,omitempty"`
	SRVConnectionString *string    `json:"SRVConnectionString,omitempty"`
	Type                *string    `json:"type,omitempty"`
}

// Endpoint is autogenerated from the json schema
type Endpoint struct {
	EndpointID   *string `json:"endpointID,omitempty"`
	ProviderName *string `json:"providerName,omitempty"`
	Region       *string `json:"region,omitempty"`
}

// Labels is autogenerated from the json schema
type Labels struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// AdvancedReplicationSpec is autogenerated from the json schema
type AdvancedReplicationSpec struct {
	ID                    *string                `json:"id,omitempty"`
	NumShards             *int                   `json:"numShards,omitempty"`
	AdvancedRegionConfigs []AdvancedRegionConfig `json:"advancedRegionConfigs,omitempty"`
	ZoneName              *string                `json:"zoneName,omitempty"`
}

// AdvancedRegionConfig is autogenerated from the json schema
type AdvancedRegionConfig struct {
	AnalyticsAutoScaling *AdvancedAutoScaling `json:"analyticsAutoScaling,omitempty"`
	AutoScaling          *AdvancedAutoScaling `json:"autoScaling,omitempty"`
	RegionName           *string              `json:"regionName,omitempty"`
	BackingProviderName  *string              `json:"BackingProviderName,omitempty"`
	ProviderName         *string              `json:"providerName,omitempty"`
	AnalyticsSpecs       *Specs               `json:"analyticsSpecs,omitempty"`
	ElectableSpecs       *Specs               `json:"electableSpecs,omitempty"`
	Priority             *int                 `json:"priority,omitempty"`
	ReadOnlySpecs        *Specs               `json:"readOnlySpecs,omitempty"`
}

// AdvancedAutoScaling is autogenerated from the json schema
type AdvancedAutoScaling struct {
	DiskGB  *DiskGB  `json:"diskGB,omitempty"`
	Compute *Compute `json:"compute,omitempty"`
}

// DiskGB is autogenerated from the json schema
type DiskGB struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// Compute is autogenerated from the json schema
type Compute struct {
	Enabled          *bool   `json:"enabled,omitempty"`
	ScaleDownEnabled *bool   `json:"scaleDownEnabled,omitempty"`
	MinInstanceSize  *string `json:"minInstanceSize,omitempty"`
	MaxInstanceSize  *string `json:"maxInstanceSize,omitempty"`
}

// Specs is autogenerated from the json schema
type Specs struct {
	DiskIOPS      *string `json:"diskIOPS,omitempty"`
	EbsVolumeType *string `json:"ebsVolumeType,omitempty"`
	InstanceSize  *string `json:"InstanceSize,omitempty"`
	NodeCount     *int    `json:"nodeCount,omitempty"`
}

type InputModel struct {
	ProjectId      *string `json:"projectId,omitempty"`
	ClusterName    *string `json:"clusterName,omitempty"`
	PrivateKey     *string `json:"privateKey,omitempty"`
	PublicKey      *string `json:"publicKey,omitempty"`
	TshirtSize     *string `json:"tshirtSize,omitempty"`
	CloudProvider  *string `json:"cloudProvider,omitempty"`
	MongoDBVersion *string `json:"mongoDBVersion,omitempty"`
}

func (model InputModel) String() string {
	var projectId, clusterName, TshirtSize, cloudProvider, mongoDBVersion string

	if model.ProjectId != nil {
		projectId = *model.ProjectId
	}
	if model.ClusterName != nil {
		clusterName = *model.ClusterName
	}
	if model.TshirtSize != nil {
		TshirtSize = *model.TshirtSize
	}
	if model.CloudProvider != nil {
		cloudProvider = *model.CloudProvider
	}
	if model.MongoDBVersion != nil {
		mongoDBVersion = *model.MongoDBVersion
	}

	return fmt.Sprintf("ProjectId: %s, ClusterName: %s, ClusterSize: %s CloudProvider: %s, MongoDBVersion: %s",
		projectId, clusterName, TshirtSize, cloudProvider, mongoDBVersion)
}
