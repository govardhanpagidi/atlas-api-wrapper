// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupScheduleApiPolicy struct for CloudBackupScheduleApiPolicy
type CloudBackupScheduleApiPolicy struct {
	Id          *string                            `json:"id,omitempty"`
	PolicyItems []CloudBackupScheduleApiPolicyItem `json:"policyItems,omitempty"`
}

// NewCloudBackupScheduleApiPolicy instantiates a new CloudBackupScheduleApiPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupScheduleApiPolicy() *CloudBackupScheduleApiPolicy {
	this := CloudBackupScheduleApiPolicy{}
	return &this
}

// NewCloudBackupScheduleApiPolicyWithDefaults instantiates a new CloudBackupScheduleApiPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupScheduleApiPolicyWithDefaults() *CloudBackupScheduleApiPolicy {
	this := CloudBackupScheduleApiPolicy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudBackupScheduleApiPolicy) SetId(v string) {
	o.Id = &v
}

// GetPolicyItems returns the PolicyItems field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicy) GetPolicyItems() []CloudBackupScheduleApiPolicyItem {
	if o == nil || IsNil(o.PolicyItems) {
		var ret []CloudBackupScheduleApiPolicyItem
		return ret
	}
	return o.PolicyItems
}

// GetPolicyItemsOk returns a tuple with the PolicyItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicy) GetPolicyItemsOk() ([]CloudBackupScheduleApiPolicyItem, bool) {
	if o == nil || IsNil(o.PolicyItems) {
		return nil, false
	}

	return o.PolicyItems, true
}

// HasPolicyItems returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicy) HasPolicyItems() bool {
	if o != nil && !IsNil(o.PolicyItems) {
		return true
	}

	return false
}

// SetPolicyItems gets a reference to the given []CloudBackupScheduleApiPolicyItem and assigns it to the PolicyItems field.
func (o *CloudBackupScheduleApiPolicy) SetPolicyItems(v []CloudBackupScheduleApiPolicyItem) {
	o.PolicyItems = v
}

func (o CloudBackupScheduleApiPolicy) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupScheduleApiPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PolicyItems) {
		toSerialize["policyItems"] = o.PolicyItems
	}
	return toSerialize, nil
}
