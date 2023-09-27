// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupScheduleApiAtlasDiskBackupCopySetting struct for CloudBackupScheduleApiAtlasDiskBackupCopySetting
type CloudBackupScheduleApiAtlasDiskBackupCopySetting struct {
	CloudProvider     *string  `json:"cloudProvider,omitempty"`
	Frequencies       []string `json:"frequencies,omitempty"`
	RegionName        *string  `json:"regionName,omitempty"`
	ReplicationSpecId *string  `json:"replicationSpecId,omitempty"`
	ShouldCopyOplogs  *bool    `json:"shouldCopyOplogs,omitempty"`
}

// NewCloudBackupScheduleApiAtlasDiskBackupCopySetting instantiates a new CloudBackupScheduleApiAtlasDiskBackupCopySetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupScheduleApiAtlasDiskBackupCopySetting() *CloudBackupScheduleApiAtlasDiskBackupCopySetting {
	this := CloudBackupScheduleApiAtlasDiskBackupCopySetting{}
	return &this
}

// NewCloudBackupScheduleApiAtlasDiskBackupCopySettingWithDefaults instantiates a new CloudBackupScheduleApiAtlasDiskBackupCopySetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupScheduleApiAtlasDiskBackupCopySettingWithDefaults() *CloudBackupScheduleApiAtlasDiskBackupCopySetting {
	this := CloudBackupScheduleApiAtlasDiskBackupCopySetting{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetFrequencies returns the Frequencies field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetFrequencies() []string {
	if o == nil || IsNil(o.Frequencies) {
		var ret []string
		return ret
	}
	return o.Frequencies
}

// GetFrequenciesOk returns a tuple with the Frequencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetFrequenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.Frequencies) {
		return nil, false
	}

	return o.Frequencies, true
}

// HasFrequencies returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) HasFrequencies() bool {
	if o != nil && !IsNil(o.Frequencies) {
		return true
	}

	return false
}

// SetFrequencies gets a reference to the given []string and assigns it to the Frequencies field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) SetFrequencies(v []string) {
	o.Frequencies = v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) SetRegionName(v string) {
	o.RegionName = &v
}

// GetReplicationSpecId returns the ReplicationSpecId field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetReplicationSpecId() string {
	if o == nil || IsNil(o.ReplicationSpecId) {
		var ret string
		return ret
	}
	return *o.ReplicationSpecId
}

// GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetReplicationSpecIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationSpecId) {
		return nil, false
	}

	return o.ReplicationSpecId, true
}

// HasReplicationSpecId returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) HasReplicationSpecId() bool {
	if o != nil && !IsNil(o.ReplicationSpecId) {
		return true
	}

	return false
}

// SetReplicationSpecId gets a reference to the given string and assigns it to the ReplicationSpecId field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) SetReplicationSpecId(v string) {
	o.ReplicationSpecId = &v
}

// GetShouldCopyOplogs returns the ShouldCopyOplogs field value if set, zero value otherwise
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetShouldCopyOplogs() bool {
	if o == nil || IsNil(o.ShouldCopyOplogs) {
		var ret bool
		return ret
	}
	return *o.ShouldCopyOplogs
}

// GetShouldCopyOplogsOk returns a tuple with the ShouldCopyOplogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) GetShouldCopyOplogsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldCopyOplogs) {
		return nil, false
	}

	return o.ShouldCopyOplogs, true
}

// HasShouldCopyOplogs returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) HasShouldCopyOplogs() bool {
	if o != nil && !IsNil(o.ShouldCopyOplogs) {
		return true
	}

	return false
}

// SetShouldCopyOplogs gets a reference to the given bool and assigns it to the ShouldCopyOplogs field.
func (o *CloudBackupScheduleApiAtlasDiskBackupCopySetting) SetShouldCopyOplogs(v bool) {
	o.ShouldCopyOplogs = &v
}

func (o CloudBackupScheduleApiAtlasDiskBackupCopySetting) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupScheduleApiAtlasDiskBackupCopySetting) ToMap() (map[string]interface{}, error) {
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
