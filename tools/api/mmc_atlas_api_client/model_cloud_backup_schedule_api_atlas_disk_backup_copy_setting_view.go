// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupScheduleApiAtlasDiskBackupCopySettingView struct for CloudBackupScheduleApiAtlasDiskBackupCopySettingView
type CloudBackupScheduleApiAtlasDiskBackupCopySettingView struct {
	CloudProvider     *string  `json:"cloudProvider,omitempty"`
	Frequencies       []string `json:"frequencies,omitempty"`
	RegionName        *string  `json:"regionName,omitempty"`
	ReplicationSpecId *string  `json:"replicationSpecId,omitempty"`
	ShouldCopyOplogs  *bool    `json:"shouldCopyOplogs,omitempty"`
}

// NewCloudBackupScheduleApiAtlasDiskBackupCopySettingView instantiates a new CloudBackupScheduleApiAtlasDiskBackupCopySettingView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupScheduleApiAtlasDiskBackupCopySettingView() *CloudBackupScheduleApiAtlasDiskBackupCopySettingView {
	this := CloudBackupScheduleApiAtlasDiskBackupCopySettingView{}
	return &this
}

// NewCloudBackupScheduleApiAtlasDiskBackupCopySettingViewWithDefaults instantiates a new CloudBackupScheduleApiAtlasDiskBackupCopySettingView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupScheduleApiAtlasDiskBackupCopySettingViewWithDefaults() *CloudBackupScheduleApiAtlasDiskBackupCopySettingView {
	this := CloudBackupScheduleApiAtlasDiskBackupCopySettingView{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetFrequencies returns the Frequencies field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetFrequencies() []string {
	if o == nil || IsNil(o.Frequencies) {
		var ret []string
		return ret
	}
	return o.Frequencies
}

// GetFrequenciesOk returns a tuple with the Frequencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetFrequenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.Frequencies) {
		return nil, false
	}

	return o.Frequencies, true
}

// HasFrequencies returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) HasFrequencies() bool {
	if o != nil && !IsNil(o.Frequencies) {
		return true
	}

	return false
}

// SetFrequencies gets a reference to the given []string and assigns it to the Frequencies field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) SetFrequencies(v []string) {
	o.Frequencies = v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) SetRegionName(v string) {
	o.RegionName = &v
}

// GetReplicationSpecId returns the ReplicationSpecId field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetReplicationSpecId() string {
	if o == nil || IsNil(o.ReplicationSpecId) {
		var ret string
		return ret
	}
	return *o.ReplicationSpecId
}

// GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetReplicationSpecIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationSpecId) {
		return nil, false
	}

	return o.ReplicationSpecId, true
}

// HasReplicationSpecId returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) HasReplicationSpecId() bool {
	if o != nil && !IsNil(o.ReplicationSpecId) {
		return true
	}

	return false
}

// SetReplicationSpecId gets a reference to the given string and assigns it to the ReplicationSpecId field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) SetReplicationSpecId(v string) {
	o.ReplicationSpecId = &v
}

// GetShouldCopyOplogs returns the ShouldCopyOplogs field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetShouldCopyOplogs() bool {
	if o == nil || IsNil(o.ShouldCopyOplogs) {
		var ret bool
		return ret
	}
	return *o.ShouldCopyOplogs
}

// GetShouldCopyOplogsOk returns a tuple with the ShouldCopyOplogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) GetShouldCopyOplogsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldCopyOplogs) {
		return nil, false
	}

	return o.ShouldCopyOplogs, true
}

// HasShouldCopyOplogs returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) HasShouldCopyOplogs() bool {
	if o != nil && !IsNil(o.ShouldCopyOplogs) {
		return true
	}

	return false
}

// SetShouldCopyOplogs gets a reference to the given bool and assigns it to the ShouldCopyOplogs field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySettingView) SetShouldCopyOplogs(v bool) {
	o.ShouldCopyOplogs = &v
}

func (o CloudBackupScheduleApiAtlasDiskBackupCopySettingView) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupScheduleApiAtlasDiskBackupCopySettingView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.Frequencies) {
		toSerialize["frequencies"] = o.Frequencies
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	if !IsNil(o.ReplicationSpecId) {
		toSerialize["replicationSpecId"] = o.ReplicationSpecId
	}
	if !IsNil(o.ShouldCopyOplogs) {
		toSerialize["shouldCopyOplogs"] = o.ShouldCopyOplogs
	}
	return toSerialize, nil
}
