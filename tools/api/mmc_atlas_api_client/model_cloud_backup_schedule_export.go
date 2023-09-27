// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupScheduleExport struct for CloudBackupScheduleExport
type CloudBackupScheduleExport struct {
	ExportBucketId *string `json:"exportBucketId,omitempty"`
	FrequencyType  *string `json:"frequencyType,omitempty"`
}

// NewCloudBackupScheduleExport instantiates a new CloudBackupScheduleExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupScheduleExport() *CloudBackupScheduleExport {
	this := CloudBackupScheduleExport{}
	return &this
}

// NewCloudBackupScheduleExportWithDefaults instantiates a new CloudBackupScheduleExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupScheduleExportWithDefaults() *CloudBackupScheduleExport {
	this := CloudBackupScheduleExport{}
	return &this
}

// GetExportBucketId returns the ExportBucketId field value if set, zero value otherwise
func (o *CloudBackupScheduleExport) GetExportBucketId() string {
	if o == nil || IsNil(o.ExportBucketId) {
		var ret string
		return ret
	}
	return *o.ExportBucketId
}

// GetExportBucketIdOk returns a tuple with the ExportBucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleExport) GetExportBucketIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExportBucketId) {
		return nil, false
	}

	return o.ExportBucketId, true
}

// HasExportBucketId returns a boolean if a field has been set.
func (o *CloudBackupScheduleExport) HasExportBucketId() bool {
	if o != nil && !IsNil(o.ExportBucketId) {
		return true
	}

	return false
}

// SetExportBucketId gets a reference to the given string and assigns it to the ExportBucketId field.
func (o *CloudBackupScheduleExport) SetExportBucketId(v string) {
	o.ExportBucketId = &v
}

// GetFrequencyType returns the FrequencyType field value if set, zero value otherwise
func (o *CloudBackupScheduleExport) GetFrequencyType() string {
	if o == nil || IsNil(o.FrequencyType) {
		var ret string
		return ret
	}
	return *o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleExport) GetFrequencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FrequencyType) {
		return nil, false
	}

	return o.FrequencyType, true
}

// HasFrequencyType returns a boolean if a field has been set.
func (o *CloudBackupScheduleExport) HasFrequencyType() bool {
	if o != nil && !IsNil(o.FrequencyType) {
		return true
	}

	return false
}

// SetFrequencyType gets a reference to the given string and assigns it to the FrequencyType field.
func (o *CloudBackupScheduleExport) SetFrequencyType(v string) {
	o.FrequencyType = &v
}

func (o CloudBackupScheduleExport) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupScheduleExport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExportBucketId) {
		toSerialize["exportBucketId"] = o.ExportBucketId
	}
	if !IsNil(o.FrequencyType) {
		toSerialize["frequencyType"] = o.FrequencyType
	}
	return toSerialize, nil
}
