// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterLabels struct for ClusterLabels
type ClusterLabels struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewClusterLabels instantiates a new ClusterLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterLabels() *ClusterLabels {
	this := ClusterLabels{}
	return &this
}

// NewClusterLabelsWithDefaults instantiates a new ClusterLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterLabelsWithDefaults() *ClusterLabels {
	this := ClusterLabels{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise
func (o *ClusterLabels) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLabels) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}

	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ClusterLabels) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ClusterLabels) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise
func (o *ClusterLabels) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLabels) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}

	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ClusterLabels) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ClusterLabels) SetValue(v string) {
	o.Value = &v
}

func (o ClusterLabels) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterLabels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}
