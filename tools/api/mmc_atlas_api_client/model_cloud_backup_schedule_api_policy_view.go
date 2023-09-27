// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupScheduleApiPolicyView struct for CloudBackupScheduleApiPolicyView
type CloudBackupScheduleApiPolicyView struct {
	Id          *string                                `json:"id,omitempty"`
	PolicyItems []CloudBackupScheduleApiPolicyItemView `json:"policyItems,omitempty"`
}

// NewCloudBackupScheduleApiPolicyView instantiates a new CloudBackupScheduleApiPolicyView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupScheduleApiPolicyView() *CloudBackupScheduleApiPolicyView {
	this := CloudBackupScheduleApiPolicyView{}
	return &this
}

// NewCloudBackupScheduleApiPolicyViewWithDefaults instantiates a new CloudBackupScheduleApiPolicyView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupScheduleApiPolicyViewWithDefaults() *CloudBackupScheduleApiPolicyView {
	this := CloudBackupScheduleApiPolicyView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyView) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyView) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyView) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudBackupScheduleApiPolicyView) SetId(v string) {
	o.Id = &v
}

// GetPolicyItems returns the PolicyItems field value if set, zero value otherwise
func (o *CloudBackupScheduleApiPolicyView) GetPolicyItems() []CloudBackupScheduleApiPolicyItemView {
	if o == nil || IsNil(o.PolicyItems) {
		var ret []CloudBackupScheduleApiPolicyItemView
		return ret
	}
	return o.PolicyItems
}

// GetPolicyItemsOk returns a tuple with the PolicyItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleApiPolicyView) GetPolicyItemsOk() ([]CloudBackupScheduleApiPolicyItemView, bool) {
	if o == nil || IsNil(o.PolicyItems) {
		return nil, false
	}

	return o.PolicyItems, true
}

// HasPolicyItems returns a boolean if a field has been set.
func (o *CloudBackupScheduleApiPolicyView) HasPolicyItems() bool {
	if o != nil && !IsNil(o.PolicyItems) {
		return true
	}

	return false
}

// SetPolicyItems gets a reference to the given []CloudBackupScheduleApiPolicyItemView and assigns it to the PolicyItems field.
func (o *CloudBackupScheduleApiPolicyView) SetPolicyItems(v []CloudBackupScheduleApiPolicyItemView) {
	o.PolicyItems = v
}

func (o CloudBackupScheduleApiPolicyView) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupScheduleApiPolicyView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PolicyItems) {
		toSerialize["policyItems"] = o.PolicyItems
	}
	return toSerialize, nil
}
