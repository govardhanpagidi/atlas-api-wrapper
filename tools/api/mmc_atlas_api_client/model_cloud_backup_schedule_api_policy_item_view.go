// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupScheduleApiPolicyItemView struct for CloudBackupScheduleApiPolicyItemView
type CloudBackupScheduleApiPolicyItemView struct {
	FrequencyInterval *int    `json:"frequencyInterval,omitempty"`
	FrequencyType     *string `json:"frequencyType,omitempty"`
	Id                *string `json:"id,omitempty"`
	RetentionUnit     *string `json:"retentionUnit,omitempty"`
	RetentionValue    *int    `json:"retentionValue,omitempty"`
}

// NewCloudBackupScheduleApiPolicyItemView instantiates a new CloudBackupScheduleApiPolicyItemView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupScheduleApiPolicyItemView() *CloudBackupScheduleApiPolicyItemView {
	this := CloudBackupScheduleApiPolicyItemView{}
	return &this
}

// NewCloudBackupScheduleApiPolicyItemViewWithDefaults instantiates a new CloudBackupScheduleApiPolicyItemView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupScheduleApiPolicyItemViewWithDefaults() *CloudBackupScheduleApiPolicyItemView {
	this := CloudBackupScheduleApiPolicyItemView{}
	return &this
}

// GetFrequencyInterval returns the FrequencyInterval field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItemView) GetFrequencyInterval() int {
	if o == nil || IsNil(o.FrequencyInterval) {
		var ret int
		return ret
	}
	return *o.FrequencyInterval
}

// GetFrequencyIntervalOk returns a tuple with the FrequencyInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItemView) GetFrequencyIntervalOk() (*int, bool) {
	if o == nil || IsNil(o.FrequencyInterval) {
		return nil, false
	}

	return o.FrequencyInterval, true
}

// HasFrequencyInterval returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItemView) HasFrequencyInterval() bool {
	if o != nil && !IsNil(o.FrequencyInterval) {
		return true
	}

	return false
}

// SetFrequencyInterval gets a reference to the given int and assigns it to the FrequencyInterval field.
func (o *CloudBackupScheduleApiPolicyItemView) SetFrequencyInterval(v int) {
	o.FrequencyInterval = &v
}

// GetFrequencyType returns the FrequencyType field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItemView) GetFrequencyType() string {
	if o == nil || IsNil(o.FrequencyType) {
		var ret string
		return ret
	}
	return *o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItemView) GetFrequencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FrequencyType) {
		return nil, false
	}

	return o.FrequencyType, true
}

// HasFrequencyType returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItemView) HasFrequencyType() bool {
	if o != nil && !IsNil(o.FrequencyType) {
		return true
	}

	return false
}

// SetFrequencyType gets a reference to the given string and assigns it to the FrequencyType field.
func (o *CloudBackupScheduleApiPolicyItemView) SetFrequencyType(v string) {
	o.FrequencyType = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItemView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItemView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItemView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudBackupScheduleApiPolicyItemView) SetId(v string) {
	o.Id = &v
}

// GetRetentionUnit returns the RetentionUnit field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItemView) GetRetentionUnit() string {
	if o == nil || IsNil(o.RetentionUnit) {
		var ret string
		return ret
	}
	return *o.RetentionUnit
}

// GetRetentionUnitOk returns a tuple with the RetentionUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItemView) GetRetentionUnitOk() (*string, bool) {
	if o == nil || IsNil(o.RetentionUnit) {
		return nil, false
	}

	return o.RetentionUnit, true
}

// HasRetentionUnit returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItemView) HasRetentionUnit() bool {
	if o != nil && !IsNil(o.RetentionUnit) {
		return true
	}

	return false
}

// SetRetentionUnit gets a reference to the given string and assigns it to the RetentionUnit field.
func (o *CloudBackupScheduleApiPolicyItemView) SetRetentionUnit(v string) {
	o.RetentionUnit = &v
}

// GetRetentionValue returns the RetentionValue field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyItemView) GetRetentionValue() int {
	if o == nil || IsNil(o.RetentionValue) {
		var ret int
		return ret
	}
	return *o.RetentionValue
}

// GetRetentionValueOk returns a tuple with the RetentionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyItemView) GetRetentionValueOk() (*int, bool) {
	if o == nil || IsNil(o.RetentionValue) {
		return nil, false
	}

	return o.RetentionValue, true
}

// HasRetentionValue returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyItemView) HasRetentionValue() bool {
	if o != nil && !IsNil(o.RetentionValue) {
		return true
	}

	return false
}

// SetRetentionValue gets a reference to the given int and assigns it to the RetentionValue field.
func (o *CloudBackupScheduleApiPolicyItemView) SetRetentionValue(v int) {
	o.RetentionValue = &v
}

func (o CloudBackupScheduleApiPolicyItemView) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupScheduleApiPolicyItemView) ToMap() (map[string]interface{}, error) {
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
