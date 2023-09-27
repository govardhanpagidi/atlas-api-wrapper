// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupScheduleApiDeleteCopiedBackupsView struct for CloudBackupScheduleApiDeleteCopiedBackupsView
type CloudBackupScheduleApiDeleteCopiedBackupsView struct {
	CloudProvider     *string `json:"cloudProvider,omitempty"`
	RegionName        *string `json:"regionName,omitempty"`
	ReplicationSpecId *string `json:"replicationSpecId,omitempty"`
}

// NewCloudBackupScheduleApiDeleteCopiedBackupsView instantiates a new CloudBackupScheduleApiDeleteCopiedBackupsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupScheduleApiDeleteCopiedBackupsView() *CloudBackupScheduleApiDeleteCopiedBackupsView {
	this := CloudBackupScheduleApiDeleteCopiedBackupsView{}
	return &this
}

// NewCloudBackupScheduleApiDeleteCopiedBackupsViewWithDefaults instantiates a new CloudBackupScheduleApiDeleteCopiedBackupsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupScheduleApiDeleteCopiedBackupsViewWithDefaults() *CloudBackupScheduleApiDeleteCopiedBackupsView {
	this := CloudBackupScheduleApiDeleteCopiedBackupsView{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) SetRegionName(v string) {
	o.RegionName = &v
}

// GetReplicationSpecId returns the ReplicationSpecId field value if set, zero value otherwise
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) GetReplicationSpecId() string {
	if o == nil || IsNil(o.ReplicationSpecId) {
		var ret string
		return ret
	}
	return *o.ReplicationSpecId
}

// GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) GetReplicationSpecIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationSpecId) {
		return nil, false
	}

	return o.ReplicationSpecId, true
}

// HasReplicationSpecId returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) HasReplicationSpecId() bool {
	if o != nil && !IsNil(o.ReplicationSpecId) {
		return true
	}

	return false
}

// SetReplicationSpecId gets a reference to the given string and assigns it to the ReplicationSpecId field.
func (o *CloudBackupScheduleApiDeleteCopiedBackupsView) SetReplicationSpecId(v string) {
	o.ReplicationSpecId = &v
}

func (o CloudBackupScheduleApiDeleteCopiedBackupsView) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupScheduleApiDeleteCopiedBackupsView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !IsNil(o.ReplicationSpecId) {
		toSerialize["replicationSpecId"] = o.ReplicationSpecId
	}
	return toSerialize, nil
}
