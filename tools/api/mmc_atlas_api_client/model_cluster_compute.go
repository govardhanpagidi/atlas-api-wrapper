// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterCompute struct for ClusterCompute
type ClusterCompute struct {
	Enabled          *bool   `json:"enabled,omitempty"`
	MaxInstanceSize  *string `json:"maxInstanceSize,omitempty"`
	MinInstanceSize  *string `json:"minInstanceSize,omitempty"`
	ScaleDownEnabled *bool   `json:"scaleDownEnabled,omitempty"`
}

// NewClusterCompute instantiates a new ClusterCompute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterCompute() *ClusterCompute {
	this := ClusterCompute{}
	return &this
}

// NewClusterComputeWithDefaults instantiates a new ClusterCompute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterComputeWithDefaults() *ClusterCompute {
	this := ClusterCompute{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *ClusterCompute) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCompute) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClusterCompute) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClusterCompute) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaxInstanceSize returns the MaxInstanceSize field value if set, zero value otherwise
func (o *ClusterCompute) GetMaxInstanceSize() string {
	if o == nil || IsNil(o.MaxInstanceSize) {
		var ret string
		return ret
	}
	return *o.MaxInstanceSize
}

// GetMaxInstanceSizeOk returns a tuple with the MaxInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCompute) GetMaxInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxInstanceSize) {
		return nil, false
	}

	return o.MaxInstanceSize, true
}

// HasMaxInstanceSize returns a boolean if a field has been set.
func (o *ClusterCompute) HasMaxInstanceSize() bool {
	if o != nil && !IsNil(o.MaxInstanceSize) {
		return true
	}

	return false
}

// SetMaxInstanceSize gets a reference to the given string and assigns it to the MaxInstanceSize field.
func (o *ClusterCompute) SetMaxInstanceSize(v string) {
	o.MaxInstanceSize = &v
}

// GetMinInstanceSize returns the MinInstanceSize field value if set, zero value otherwise
func (o *ClusterCompute) GetMinInstanceSize() string {
	if o == nil || IsNil(o.MinInstanceSize) {
		var ret string
		return ret
	}
	return *o.MinInstanceSize
}

// GetMinInstanceSizeOk returns a tuple with the MinInstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCompute) GetMinInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MinInstanceSize) {
		return nil, false
	}

	return o.MinInstanceSize, true
}

// HasMinInstanceSize returns a boolean if a field has been set.
func (o *ClusterCompute) HasMinInstanceSize() bool {
	if o != nil && !IsNil(o.MinInstanceSize) {
		return true
	}

	return false
}

// SetMinInstanceSize gets a reference to the given string and assigns it to the MinInstanceSize field.
func (o *ClusterCompute) SetMinInstanceSize(v string) {
	o.MinInstanceSize = &v
}

// GetScaleDownEnabled returns the ScaleDownEnabled field value if set, zero value otherwise
func (o *ClusterCompute) GetScaleDownEnabled() bool {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		var ret bool
		return ret
	}
	return *o.ScaleDownEnabled
}

// GetScaleDownEnabledOk returns a tuple with the ScaleDownEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterCompute) GetScaleDownEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ScaleDownEnabled) {
		return nil, false
	}

	return o.ScaleDownEnabled, true
}

// HasScaleDownEnabled returns a boolean if a field has been set.
func (o *ClusterCompute) HasScaleDownEnabled() bool {
	if o != nil && !IsNil(o.ScaleDownEnabled) {
		return true
	}

	return false
}

// SetScaleDownEnabled gets a reference to the given bool and assigns it to the ScaleDownEnabled field.
func (o *ClusterCompute) SetScaleDownEnabled(v bool) {
	o.ScaleDownEnabled = &v
}

func (o ClusterCompute) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterCompute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MaxInstanceSize) {
		toSerialize["maxInstanceSize"] = o.MaxInstanceSize
	}
	if !IsNil(o.MinInstanceSize) {
		toSerialize["minInstanceSize"] = o.MinInstanceSize
	}
	if !IsNil(o.ScaleDownEnabled) {
		toSerialize["scaleDownEnabled"] = o.ScaleDownEnabled
	}
	return toSerialize, nil
}
