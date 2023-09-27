// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupScheduleApiPolicyItem struct for CloudBackupScheduleApiPolicyItem
type CloudBackupScheduleApiPolicyItem struct {
	FrequencyInterval *int    `json:"frequencyInterval,omitempty"`
	FrequencyType     *string `json:"frequencyType,omitempty"`
	Id                *string `json:"id,omitempty"`
	RetentionUnit     *string `json:"retentionUnit,omitempty"`
	RetentionValue    *int    `json:"retentionValue,omitempty"`
}

// NewCloudBackupScheduleApiPolicyItem instantiates a new CloudBackupScheduleApiPolicyItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupScheduleApiPolicyItem() *CloudBackupScheduleApiPolicyItem {
	this := CloudBackupScheduleApiPolicyItem{}
	return &this
}

// NewCloudBackupScheduleApiPolicyItemWithDefaults instantiates a new CloudBackupScheduleApiPolicyItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupScheduleApiPolicyItemWithDefaults() *CloudBackupScheduleApiPolicyItem {
	this := CloudBackupScheduleApiPolicyItem{}
	return &this
}

// GetFrequencyInterval returns the FrequencyInterval field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItem) GetFrequencyInterval() int {
	if o == nil || IsNil(o.FrequencyInterval) {
		var ret int
		return ret
	}
	return *o.FrequencyInterval
}

// GetFrequencyIntervalOk returns a tuple with the FrequencyInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItem) GetFrequencyIntervalOk() (*int, bool) {
	if o == nil || IsNil(o.FrequencyInterval) {
		return nil, false
	}

	return o.FrequencyInterval, true
}

// HasFrequencyInterval returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItem) HasFrequencyInterval() bool {
	if o != nil && !IsNil(o.FrequencyInterval) {
		return true
	}

	return false
}

// SetFrequencyInterval gets a reference to the given int and assigns it to the FrequencyInterval field.
func (o *CloudBackupScheduleApiPolicyItem) SetFrequencyInterval(v int) {
	o.FrequencyInterval = &v
}

// GetFrequencyType returns the FrequencyType field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItem) GetFrequencyType() string {
	if o == nil || IsNil(o.FrequencyType) {
		var ret string
		return ret
	}
	return *o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItem) GetFrequencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FrequencyType) {
		return nil, false
	}

	return o.FrequencyType, true
}

// HasFrequencyType returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItem) HasFrequencyType() bool {
	if o != nil && !IsNil(o.FrequencyType) {
		return true
	}

	return false
}

// SetFrequencyType gets a reference to the given string and assigns it to the FrequencyType field.
func (o *CloudBackupScheduleApiPolicyItem) SetFrequencyType(v string) {
	o.FrequencyType = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudBackupScheduleApiPolicyItem) SetId(v string) {
	o.Id = &v
}

// GetRetentionUnit returns the RetentionUnit field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItem) GetRetentionUnit() string {
	if o == nil || IsNil(o.RetentionUnit) {
		var ret string
		return ret
	}
	return *o.RetentionUnit
}

// GetRetentionUnitOk returns a tuple with the RetentionUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItem) GetRetentionUnitOk() (*string, bool) {
	if o == nil || IsNil(o.RetentionUnit) {
		return nil, false
	}

	return o.RetentionUnit, true
}

// HasRetentionUnit returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItem) HasRetentionUnit() bool {
	if o != nil && !IsNil(o.RetentionUnit) {
		return true
	}

	return false
}

// SetRetentionUnit gets a reference to the given string and assigns it to the RetentionUnit field.
func (o *CloudBackupScheduleApiPolicyItem) SetRetentionUnit(v string) {
	o.RetentionUnit = &v
}

// GetRetentionValue returns the RetentionValue field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItem) GetRetentionValue() int {
	if o == nil || IsNil(o.RetentionValue) {
		var ret int
		return ret
	}
	return *o.RetentionValue
}

// GetRetentionValueOk returns a tuple with the RetentionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItem) GetRetentionValueOk() (*int, bool) {
	if o == nil || IsNil(o.RetentionValue) {
		return nil, false
	}

	return o.RetentionValue, true
}

// HasRetentionValue returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItem) HasRetentionValue() bool {
	if o != nil && !IsNil(o.RetentionValue) {
		return true
	}

	return false
}

// SetRetentionValue gets a reference to the given int and assigns it to the RetentionValue field.
func (o *CloudBackupScheduleApiPolicyItem) SetRetentionValue(v int) {
	o.RetentionValue = &v
}

func (o CloudBackupScheduleApiPolicyItem) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupScheduleApiPolicyItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FrequencyInterval) {
		toSerialize["frequencyInterval"] = o.FrequencyInterval
	}
	if !IsNil(o.FrequencyType) {
		toSerialize["frequencyType"] = o.FrequencyType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RetentionUnit) {
		toSerialize["retentionUnit"] = o.RetentionUnit
	}
	if !IsNil(o.RetentionValue) {
		toSerialize["retentionValue"] = o.RetentionValue
	}
	return toSerialize, nil
}
