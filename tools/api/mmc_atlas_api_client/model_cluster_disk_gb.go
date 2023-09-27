// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterDiskGB struct for ClusterDiskGB
type ClusterDiskGB struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewClusterDiskGB instantiates a new ClusterDiskGB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDiskGB() *ClusterDiskGB {
	this := ClusterDiskGB{}
	return &this
}

// NewClusterDiskGBWithDefaults instantiates a new ClusterDiskGB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDiskGBWithDefaults() *ClusterDiskGB {
	this := ClusterDiskGB{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *ClusterDiskGB) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDiskGB) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClusterDiskGB) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClusterDiskGB) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o ClusterDiskGB) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterDiskGB) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}
