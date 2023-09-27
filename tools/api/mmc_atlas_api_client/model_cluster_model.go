// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterModel struct for ClusterModel
type ClusterModel struct {
	AdvancedSettings             *ClusterProcessArgs              `json:"advancedSettings,omitempty"`
	BackupEnabled                *bool                            `json:"backupEnabled,omitempty"`
	BiConnector                  *ClusterBiConnector              `json:"biConnector,omitempty"`
	ClusterType                  *string                          `json:"clusterType,omitempty"`
	ConnectionStrings            *ClusterConnectionStrings        `json:"connectionStrings,omitempty"`
	CreatedDate                  *string                          `json:"createdDate,omitempty"`
	DiskSizeGB                   *float32                         `json:"diskSizeGB,omitempty"`
	EncryptionAtRestProvider     *string                          `json:"encryptionAtRestProvider,omitempty"`
	Id                           *string                          `json:"id,omitempty"`
	Labels                       []ClusterLabels                  `json:"labels,omitempty"`
	MongoDBMajorVersion          *string                          `json:"mongoDBMajorVersion,omitempty"`
	MongoDBVersion               *string                          `json:"mongoDBVersion,omitempty"`
	Name                         *string                          `json:"name,omitempty"`
	Paused                       *bool                            `json:"paused,omitempty"`
	PitEnabled                   *bool                            `json:"pitEnabled,omitempty"`
	Profile                      *string                          `json:"profile,omitempty"`
	ProjectId                    *string                          `json:"projectId,omitempty"`
	PublicAccessEnabled          *bool                            `json:"publicAccessEnabled,omitempty"`
	ReplicationSpecs             []ClusterAdvancedReplicationSpec `json:"replicationSpecs,omitempty"`
	RootCertType                 *string                          `json:"rootCertType,omitempty"`
	StateName                    *string                          `json:"stateName,omitempty"`
	TerminationProtectionEnabled *bool                            `json:"terminationProtectionEnabled,omitempty"`
	VersionReleaseSystem         *string                          `json:"versionReleaseSystem,omitempty"`
}

// NewClusterModel instantiates a new ClusterModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterModel() *ClusterModel {
	this := ClusterModel{}
	return &this
}

// NewClusterModelWithDefaults instantiates a new ClusterModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterModelWithDefaults() *ClusterModel {
	this := ClusterModel{}
	return &this
}

// GetAdvancedSettings returns the AdvancedSettings field value if set, zero value otherwise
func (o *ClusterModel) GetAdvancedSettings() ClusterProcessArgs {
	if o == nil || IsNil(o.AdvancedSettings) {
		var ret ClusterProcessArgs
		return ret
	}
	return *o.AdvancedSettings
}

// GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetAdvancedSettingsOk() (*ClusterProcessArgs, bool) {
	if o == nil || IsNil(o.AdvancedSettings) {
		return nil, false
	}

	return o.AdvancedSettings, true
}

// HasAdvancedSettings returns a boolean if a field has been set.
func (o *ClusterModel) HasAdvancedSettings() bool {
	if o != nil && !IsNil(o.AdvancedSettings) {
		return true
	}

	return false
}

// SetAdvancedSettings gets a reference to the given ClusterProcessArgs and assigns it to the AdvancedSettings field.
func (o *ClusterModel) SetAdvancedSettings(v ClusterProcessArgs) {
	o.AdvancedSettings = &v
}

// GetBackupEnabled returns the BackupEnabled field value if set, zero value otherwise
func (o *ClusterModel) GetBackupEnabled() bool {
	if o == nil || IsNil(o.BackupEnabled) {
		var ret bool
		return ret
	}
	return *o.BackupEnabled
}

// GetBackupEnabledOk returns a tuple with the BackupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetBackupEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BackupEnabled) {
		return nil, false
	}

	return o.BackupEnabled, true
}

// HasBackupEnabled returns a boolean if a field has been set.
func (o *ClusterModel) HasBackupEnabled() bool {
	if o != nil && !IsNil(o.BackupEnabled) {
		return true
	}

	return false
}

// SetBackupEnabled gets a reference to the given bool and assigns it to the BackupEnabled field.
func (o *ClusterModel) SetBackupEnabled(v bool) {
	o.BackupEnabled = &v
}

// GetBiConnector returns the BiConnector field value if set, zero value otherwise
func (o *ClusterModel) GetBiConnector() ClusterBiConnector {
	if o == nil || IsNil(o.BiConnector) {
		var ret ClusterBiConnector
		return ret
	}
	return *o.BiConnector
}

// GetBiConnectorOk returns a tuple with the BiConnector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetBiConnectorOk() (*ClusterBiConnector, bool) {
	if o == nil || IsNil(o.BiConnector) {
		return nil, false
	}

	return o.BiConnector, true
}

// HasBiConnector returns a boolean if a field has been set.
func (o *ClusterModel) HasBiConnector() bool {
	if o != nil && !IsNil(o.BiConnector) {
		return true
	}

	return false
}

// SetBiConnector gets a reference to the given ClusterBiConnector and assigns it to the BiConnector field.
func (o *ClusterModel) SetBiConnector(v ClusterBiConnector) {
	o.BiConnector = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise
func (o *ClusterModel) GetClusterType() string {
	if o == nil || IsNil(o.ClusterType) {
		var ret string
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetClusterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterType) {
		return nil, false
	}

	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *ClusterModel) HasClusterType() bool {
	if o != nil && !IsNil(o.ClusterType) {
		return true
	}

	return false
}

// SetClusterType gets a reference to the given string and assigns it to the ClusterType field.
func (o *ClusterModel) SetClusterType(v string) {
	o.ClusterType = &v
}

// GetConnectionStrings returns the ConnectionStrings field value if set, zero value otherwise
func (o *ClusterModel) GetConnectionStrings() ClusterConnectionStrings {
	if o == nil || IsNil(o.ConnectionStrings) {
		var ret ClusterConnectionStrings
		return ret
	}
	return *o.ConnectionStrings
}

// GetConnectionStringsOk returns a tuple with the ConnectionStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetConnectionStringsOk() (*ClusterConnectionStrings, bool) {
	if o == nil || IsNil(o.ConnectionStrings) {
		return nil, false
	}

	return o.ConnectionStrings, true
}

// HasConnectionStrings returns a boolean if a field has been set.
func (o *ClusterModel) HasConnectionStrings() bool {
	if o != nil && !IsNil(o.ConnectionStrings) {
		return true
	}

	return false
}

// SetConnectionStrings gets a reference to the given ClusterConnectionStrings and assigns it to the ConnectionStrings field.
func (o *ClusterModel) SetConnectionStrings(v ClusterConnectionStrings) {
	o.ConnectionStrings = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise
func (o *ClusterModel) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}

	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ClusterModel) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ClusterModel) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetDiskSizeGB returns the DiskSizeGB field value if set, zero value otherwise
func (o *ClusterModel) GetDiskSizeGB() float32 {
	if o == nil || IsNil(o.DiskSizeGB) {
		var ret float32
		return ret
	}
	return *o.DiskSizeGB
}

// GetDiskSizeGBOk returns a tuple with the DiskSizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetDiskSizeGBOk() (*float32, bool) {
	if o == nil || IsNil(o.DiskSizeGB) {
		return nil, false
	}

	return o.DiskSizeGB, true
}

// HasDiskSizeGB returns a boolean if a field has been set.
func (o *ClusterModel) HasDiskSizeGB() bool {
	if o != nil && !IsNil(o.DiskSizeGB) {
		return true
	}

	return false
}

// SetDiskSizeGB gets a reference to the given float32 and assigns it to the DiskSizeGB field.
func (o *ClusterModel) SetDiskSizeGB(v float32) {
	o.DiskSizeGB = &v
}

// GetEncryptionAtRestProvider returns the EncryptionAtRestProvider field value if set, zero value otherwise
func (o *ClusterModel) GetEncryptionAtRestProvider() string {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		var ret string
		return ret
	}
	return *o.EncryptionAtRestProvider
}

// GetEncryptionAtRestProviderOk returns a tuple with the EncryptionAtRestProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetEncryptionAtRestProviderOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionAtRestProvider) {
		return nil, false
	}

	return o.EncryptionAtRestProvider, true
}

// HasEncryptionAtRestProvider returns a boolean if a field has been set.
func (o *ClusterModel) HasEncryptionAtRestProvider() bool {
	if o != nil && !IsNil(o.EncryptionAtRestProvider) {
		return true
	}

	return false
}

// SetEncryptionAtRestProvider gets a reference to the given string and assigns it to the EncryptionAtRestProvider field.
func (o *ClusterModel) SetEncryptionAtRestProvider(v string) {
	o.EncryptionAtRestProvider = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ClusterModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterModel) SetId(v string) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise
func (o *ClusterModel) GetLabels() []ClusterLabels {
	if o == nil || IsNil(o.Labels) {
		var ret []ClusterLabels
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetLabelsOk() ([]ClusterLabels, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}

	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ClusterModel) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []ClusterLabels and assigns it to the Labels field.
func (o *ClusterModel) SetLabels(v []ClusterLabels) {
	o.Labels = v
}

// GetMongoDBMajorVersion returns the MongoDBMajorVersion field value if set, zero value otherwise
func (o *ClusterModel) GetMongoDBMajorVersion() string {
	if o == nil || IsNil(o.MongoDBMajorVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBMajorVersion
}

// GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetMongoDBMajorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBMajorVersion) {
		return nil, false
	}

	return o.MongoDBMajorVersion, true
}

// HasMongoDBMajorVersion returns a boolean if a field has been set.
func (o *ClusterModel) HasMongoDBMajorVersion() bool {
	if o != nil && !IsNil(o.MongoDBMajorVersion) {
		return true
	}

	return false
}

// SetMongoDBMajorVersion gets a reference to the given string and assigns it to the MongoDBMajorVersion field.
func (o *ClusterModel) SetMongoDBMajorVersion(v string) {
	o.MongoDBMajorVersion = &v
}

// GetMongoDBVersion returns the MongoDBVersion field value if set, zero value otherwise
func (o *ClusterModel) GetMongoDBVersion() string {
	if o == nil || IsNil(o.MongoDBVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBVersion
}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetMongoDBVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBVersion) {
		return nil, false
	}

	return o.MongoDBVersion, true
}

// HasMongoDBVersion returns a boolean if a field has been set.
func (o *ClusterModel) HasMongoDBVersion() bool {
	if o != nil && !IsNil(o.MongoDBVersion) {
		return true
	}

	return false
}

// SetMongoDBVersion gets a reference to the given string and assigns it to the MongoDBVersion field.
func (o *ClusterModel) SetMongoDBVersion(v string) {
	o.MongoDBVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *ClusterModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterModel) SetName(v string) {
	o.Name = &v
}

// GetPaused returns the Paused field value if set, zero value otherwise
func (o *ClusterModel) GetPaused() bool {
	if o == nil || IsNil(o.Paused) {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.Paused) {
		return nil, false
	}

	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *ClusterModel) HasPaused() bool {
	if o != nil && !IsNil(o.Paused) {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *ClusterModel) SetPaused(v bool) {
	o.Paused = &v
}

// GetPitEnabled returns the PitEnabled field value if set, zero value otherwise
func (o *ClusterModel) GetPitEnabled() bool {
	if o == nil || IsNil(o.PitEnabled) {
		var ret bool
		return ret
	}
	return *o.PitEnabled
}

// GetPitEnabledOk returns a tuple with the PitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetPitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PitEnabled) {
		return nil, false
	}

	return o.PitEnabled, true
}

// HasPitEnabled returns a boolean if a field has been set.
func (o *ClusterModel) HasPitEnabled() bool {
	if o != nil && !IsNil(o.PitEnabled) {
		return true
	}

	return false
}

// SetPitEnabled gets a reference to the given bool and assigns it to the PitEnabled field.
func (o *ClusterModel) SetPitEnabled(v bool) {
	o.PitEnabled = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise
func (o *ClusterModel) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}

	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ClusterModel) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *ClusterModel) SetProfile(v string) {
	o.Profile = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise
func (o *ClusterModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}

	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ClusterModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ClusterModel) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetPublicAccessEnabled returns the PublicAccessEnabled field value if set, zero value otherwise
func (o *ClusterModel) GetPublicAccessEnabled() bool {
	if o == nil || IsNil(o.PublicAccessEnabled) {
		var ret bool
		return ret
	}
	return *o.PublicAccessEnabled
}

// GetPublicAccessEnabledOk returns a tuple with the PublicAccessEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetPublicAccessEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicAccessEnabled) {
		return nil, false
	}

	return o.PublicAccessEnabled, true
}

// HasPublicAccessEnabled returns a boolean if a field has been set.
func (o *ClusterModel) HasPublicAccessEnabled() bool {
	if o != nil && !IsNil(o.PublicAccessEnabled) {
		return true
	}

	return false
}

// SetPublicAccessEnabled gets a reference to the given bool and assigns it to the PublicAccessEnabled field.
func (o *ClusterModel) SetPublicAccessEnabled(v bool) {
	o.PublicAccessEnabled = &v
}

// GetReplicationSpecs returns the ReplicationSpecs field value if set, zero value otherwise
func (o *ClusterModel) GetReplicationSpecs() []ClusterAdvancedReplicationSpec {
	if o == nil || IsNil(o.ReplicationSpecs) {
		var ret []ClusterAdvancedReplicationSpec
		return ret
	}
	return o.ReplicationSpecs
}

// GetReplicationSpecsOk returns a tuple with the ReplicationSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetReplicationSpecsOk() ([]ClusterAdvancedReplicationSpec, bool) {
	if o == nil || IsNil(o.ReplicationSpecs) {
		return nil, false
	}

	return o.ReplicationSpecs, true
}

// HasReplicationSpecs returns a boolean if a field has been set.
func (o *ClusterModel) HasReplicationSpecs() bool {
	if o != nil && !IsNil(o.ReplicationSpecs) {
		return true
	}

	return false
}

// SetReplicationSpecs gets a reference to the given []ClusterAdvancedReplicationSpec and assigns it to the ReplicationSpecs field.
func (o *ClusterModel) SetReplicationSpecs(v []ClusterAdvancedReplicationSpec) {
	o.ReplicationSpecs = v
}

// GetRootCertType returns the RootCertType field value if set, zero value otherwise
func (o *ClusterModel) GetRootCertType() string {
	if o == nil || IsNil(o.RootCertType) {
		var ret string
		return ret
	}
	return *o.RootCertType
}

// GetRootCertTypeOk returns a tuple with the RootCertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetRootCertTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RootCertType) {
		return nil, false
	}

	return o.RootCertType, true
}

// HasRootCertType returns a boolean if a field has been set.
func (o *ClusterModel) HasRootCertType() bool {
	if o != nil && !IsNil(o.RootCertType) {
		return true
	}

	return false
}

// SetRootCertType gets a reference to the given string and assigns it to the RootCertType field.
func (o *ClusterModel) SetRootCertType(v string) {
	o.RootCertType = &v
}

// GetStateName returns the StateName field value if set, zero value otherwise
func (o *ClusterModel) GetStateName() string {
	if o == nil || IsNil(o.StateName) {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetStateNameOk() (*string, bool) {
	if o == nil || IsNil(o.StateName) {
		return nil, false
	}

	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *ClusterModel) HasStateName() bool {
	if o != nil && !IsNil(o.StateName) {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *ClusterModel) SetStateName(v string) {
	o.StateName = &v
}

// GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field value if set, zero value otherwise
func (o *ClusterModel) GetTerminationProtectionEnabled() bool {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.TerminationProtectionEnabled
}

// GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetTerminationProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		return nil, false
	}

	return o.TerminationProtectionEnabled, true
}

// HasTerminationProtectionEnabled returns a boolean if a field has been set.
func (o *ClusterModel) HasTerminationProtectionEnabled() bool {
	if o != nil && !IsNil(o.TerminationProtectionEnabled) {
		return true
	}

	return false
}

// SetTerminationProtectionEnabled gets a reference to the given bool and assigns it to the TerminationProtectionEnabled field.
func (o *ClusterModel) SetTerminationProtectionEnabled(v bool) {
	o.TerminationProtectionEnabled = &v
}

// GetVersionReleaseSystem returns the VersionReleaseSystem field value if set, zero value otherwise
func (o *ClusterModel) GetVersionReleaseSystem() string {
	if o == nil || IsNil(o.VersionReleaseSystem) {
		var ret string
		return ret
	}
	return *o.VersionReleaseSystem
}

// GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterModel) GetVersionReleaseSystemOk() (*string, bool) {
	if o == nil || IsNil(o.VersionReleaseSystem) {
		return nil, false
	}

	return o.VersionReleaseSystem, true
}

// HasVersionReleaseSystem returns a boolean if a field has been set.
func (o *ClusterModel) HasVersionReleaseSystem() bool {
	if o != nil && !IsNil(o.VersionReleaseSystem) {
		return true
	}

	return false
}

// SetVersionReleaseSystem gets a reference to the given string and assigns it to the VersionReleaseSystem field.
func (o *ClusterModel) SetVersionReleaseSystem(v string) {
	o.VersionReleaseSystem = &v
}

func (o ClusterModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvancedSettings) {
		toSerialize["advancedSettings"] = o.AdvancedSettings
	}
	if !IsNil(o.BackupEnabled) {
		toSerialize["backupEnabled"] = o.BackupEnabled
	}
	if !IsNil(o.BiConnector) {
		toSerialize["biConnector"] = o.BiConnector
	}
	if !IsNil(o.ClusterType) {
		toSerialize["clusterType"] = o.ClusterType
	}
	if !IsNil(o.ConnectionStrings) {
		toSerialize["connectionStrings"] = o.ConnectionStrings
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.DiskSizeGB) {
		toSerialize["diskSizeGB"] = o.DiskSizeGB
	}
	if !IsNil(o.EncryptionAtRestProvider) {
		toSerialize["encryptionAtRestProvider"] = o.EncryptionAtRestProvider
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.MongoDBMajorVersion) {
		toSerialize["mongoDBMajorVersion"] = o.MongoDBMajorVersion
	}
	if !IsNil(o.MongoDBVersion) {
		toSerialize["mongoDBVersion"] = o.MongoDBVersion
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Paused) {
		toSerialize["paused"] = o.Paused
	}
	if !IsNil(o.PitEnabled) {
		toSerialize["pitEnabled"] = o.PitEnabled
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.PublicAccessEnabled) {
		toSerialize["publicAccessEnabled"] = o.PublicAccessEnabled
	}
	if !IsNil(o.ReplicationSpecs) {
		toSerialize["replicationSpecs"] = o.ReplicationSpecs
	}
	if !IsNil(o.RootCertType) {
		toSerialize["rootCertType"] = o.RootCertType
	}
	if !IsNil(o.StateName) {
		toSerialize["stateName"] = o.StateName
	}
	if !IsNil(o.TerminationProtectionEnabled) {
		toSerialize["terminationProtectionEnabled"] = o.TerminationProtectionEnabled
	}
	if !IsNil(o.VersionReleaseSystem) {
		toSerialize["versionReleaseSystem"] = o.VersionReleaseSystem
	}
	return toSerialize, nil
}
