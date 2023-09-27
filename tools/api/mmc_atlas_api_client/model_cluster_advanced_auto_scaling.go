// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterAdvancedAutoScaling struct for ClusterAdvancedAutoScaling
type ClusterAdvancedAutoScaling struct {
	Compute *ClusterCompute `json:"compute,omitempty"`
	DiskGB  *ClusterDiskGB  `json:"diskGB,omitempty"`
}

// NewClusterAdvancedAutoScaling instantiates a new ClusterAdvancedAutoScaling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAdvancedAutoScaling() *ClusterAdvancedAutoScaling {
	this := ClusterAdvancedAutoScaling{}
	return &this
}

// NewClusterAdvancedAutoScalingWithDefaults instantiates a new ClusterAdvancedAutoScaling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAdvancedAutoScalingWithDefaults() *ClusterAdvancedAutoScaling {
	this := ClusterAdvancedAutoScaling{}
	return &this
}

// GetCompute returns the Compute field value if set, zero value otherwise
func (o *ClusterAdvancedAutoScaling) GetCompute() ClusterCompute {
	if o == nil || IsNil(o.Compute) {
		var ret ClusterCompute
		return ret
	}
	return *o.Compute
}

// GetComputeOk returns a tuple with the Compute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedAutoScaling) GetComputeOk() (*ClusterCompute, bool) {
	if o == nil || IsNil(o.Compute) {
		return nil, false
	}

	return o.Compute, true
}

// HasCompute returns a boolean if a field has been set.
func (o *ClusterAdvancedAutoScaling) HasCompute() bool {
	if o != nil && !IsNil(o.Compute) {
		return true
	}

	return false
}

// SetCompute gets a reference to the given ClusterCompute and assigns it to the Compute field.
func (o *ClusterAdvancedAutoScaling) SetCompute(v ClusterCompute) {
	o.Compute = &v
}

// GetDiskGB returns the DiskGB field value if set, zero value otherwise
func (o *ClusterAdvancedAutoScaling) GetDiskGB() ClusterDiskGB {
	if o == nil || IsNil(o.DiskGB) {
		var ret ClusterDiskGB
		return ret
	}
	return *o.DiskGB
}

// GetDiskGBOk returns a tuple with the DiskGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedAutoScaling) GetDiskGBOk() (*ClusterDiskGB, bool) {
	if o == nil || IsNil(o.DiskGB) {
		return nil, false
	}

	return o.DiskGB, true
}

// HasDiskGB returns a boolean if a field has been set.
func (o *ClusterAdvancedAutoScaling) HasDiskGB() bool {
	if o != nil && !IsNil(o.DiskGB) {
		return true
	}

	return false
}

// SetDiskGB gets a reference to the given ClusterDiskGB and assigns it to the DiskGB field.
func (o *ClusterAdvancedAutoScaling) SetDiskGB(v ClusterDiskGB) {
	o.DiskGB = &v
}

func (o ClusterAdvancedAutoScaling) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterAdvancedAutoScaling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Compute) {
		toSerialize["compute"] = o.Compute
	}
	if !IsNil(o.DiskGB) {
		toSerialize["diskGB"] = o.DiskGB
	}
	return toSerialize, nil
}
