// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterAdvancedRegionConfig struct for ClusterAdvancedRegionConfig
type ClusterAdvancedRegionConfig struct {
	BackingProviderName  *string                     `json:"BackingProviderName,omitempty"`
	AnalyticsAutoScaling *ClusterAdvancedAutoScaling `json:"analyticsAutoScaling,omitempty"`
	AnalyticsSpecs       *ClusterSpecs               `json:"analyticsSpecs,omitempty"`
	AutoScaling          *ClusterAdvancedAutoScaling `json:"autoScaling,omitempty"`
	ElectableSpecs       *ClusterSpecs               `json:"electableSpecs,omitempty"`
	Priority             *int                        `json:"priority,omitempty"`
	ProviderName         *string                     `json:"providerName,omitempty"`
	ReadOnlySpecs        *ClusterSpecs               `json:"readOnlySpecs,omitempty"`
	RegionName           *string                     `json:"regionName,omitempty"`
}

// NewClusterAdvancedRegionConfig instantiates a new ClusterAdvancedRegionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAdvancedRegionConfig() *ClusterAdvancedRegionConfig {
	this := ClusterAdvancedRegionConfig{}
	return &this
}

// NewClusterAdvancedRegionConfigWithDefaults instantiates a new ClusterAdvancedRegionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAdvancedRegionConfigWithDefaults() *ClusterAdvancedRegionConfig {
	this := ClusterAdvancedRegionConfig{}
	return &this
}

// GetBackingProviderName returns the BackingProviderName field value if set, zero value otherwise
func (o *ClusterAdvancedRegionConfig) GetBackingProviderName() string {
	if o == nil || IsNil(o.BackingProviderName) {
		var ret string
		return ret
	}
	return *o.BackingProviderName
}

// GetBackingProviderNameOk returns a tuple with the BackingProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedRegionConfig) GetBackingProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.BackingProviderName) {
		return nil, false
	}

	return o.BackingProviderName, true
}

// HasBackingProviderName returns a boolean if a field has been set.
func (o *ClusterAdvancedRegionConfig) HasBackingProviderName() bool {
	if o != nil && !IsNil(o.BackingProviderName) {
		return true
	}

	return false
}

// SetBackingProviderName gets a reference to the given string and assigns it to the BackingProviderName field.
func (o *ClusterAdvancedRegionConfig) SetBackingProviderName(v string) {
	o.BackingProviderName = &v
}

// GetAnalyticsAutoScaling returns the AnalyticsAutoScaling field value if set, zero value otherwise
func (o *ClusterAdvancedRegionConfig) GetAnalyticsAutoScaling() ClusterAdvancedAutoScaling {
	if o == nil || IsNil(o.AnalyticsAutoScaling) {
		var ret ClusterAdvancedAutoScaling
		return ret
	}
	return *o.AnalyticsAutoScaling
}

// GetAnalyticsAutoScalingOk returns a tuple with the AnalyticsAutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedRegionConfig) GetAnalyticsAutoScalingOk() (*ClusterAdvancedAutoScaling, bool) {
	if o == nil || IsNil(o.AnalyticsAutoScaling) {
		return nil, false
	}

	return o.AnalyticsAutoScaling, true
}

// HasAnalyticsAutoScaling returns a boolean if a field has been set.
func (o *ClusterAdvancedRegionConfig) HasAnalyticsAutoScaling() bool {
	if o != nil && !IsNil(o.AnalyticsAutoScaling) {
		return true
	}

	return false
}

// SetAnalyticsAutoScaling gets a reference to the given ClusterAdvancedAutoScaling and assigns it to the AnalyticsAutoScaling field.
func (o *ClusterAdvancedRegionConfig) SetAnalyticsAutoScaling(v ClusterAdvancedAutoScaling) {
	o.AnalyticsAutoScaling = &v
}

// GetAnalyticsSpecs returns the AnalyticsSpecs field value if set, zero value otherwise
func (o *ClusterAdvancedRegionConfig) GetAnalyticsSpecs() ClusterSpecs {
	if o == nil || IsNil(o.AnalyticsSpecs) {
		var ret ClusterSpecs
		return ret
	}
	return *o.AnalyticsSpecs
}

// GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedRegionConfig) GetAnalyticsSpecsOk() (*ClusterSpecs, bool) {
	if o == nil || IsNil(o.AnalyticsSpecs) {
		return nil, false
	}

	return o.AnalyticsSpecs, true
}

// HasAnalyticsSpecs returns a boolean if a field has been set.
func (o *ClusterAdvancedRegionConfig) HasAnalyticsSpecs() bool {
	if o != nil && !IsNil(o.AnalyticsSpecs) {
		return true
	}

	return false
}

// SetAnalyticsSpecs gets a reference to the given ClusterSpecs and assigns it to the AnalyticsSpecs field.
func (o *ClusterAdvancedRegionConfig) SetAnalyticsSpecs(v ClusterSpecs) {
	o.AnalyticsSpecs = &v
}

// GetAutoScaling returns the AutoScaling field value if set, zero value otherwise
func (o *ClusterAdvancedRegionConfig) GetAutoScaling() ClusterAdvancedAutoScaling {
	if o == nil || IsNil(o.AutoScaling) {
		var ret ClusterAdvancedAutoScaling
		return ret
	}
	return *o.AutoScaling
}

// GetAutoScalingOk returns a tuple with the AutoScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedRegionConfig) GetAutoScalingOk() (*ClusterAdvancedAutoScaling, bool) {
	if o == nil || IsNil(o.AutoScaling) {
		return nil, false
	}

	return o.AutoScaling, true
}

// HasAutoScaling returns a boolean if a field has been set.
func (o *ClusterAdvancedRegionConfig) HasAutoScaling() bool {
	if o != nil && !IsNil(o.AutoScaling) {
		return true
	}

	return false
}

// SetAutoScaling gets a reference to the given ClusterAdvancedAutoScaling and assigns it to the AutoScaling field.
func (o *ClusterAdvancedRegionConfig) SetAutoScaling(v ClusterAdvancedAutoScaling) {
	o.AutoScaling = &v
}

// GetElectableSpecs returns the ElectableSpecs field value if set, zero value otherwise
func (o *ClusterAdvancedRegionConfig) GetElectableSpecs() ClusterSpecs {
	if o == nil || IsNil(o.ElectableSpecs) {
		var ret ClusterSpecs
		return ret
	}
	return *o.ElectableSpecs
}

// GetElectableSpecsOk returns a tuple with the ElectableSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedRegionConfig) GetElectableSpecsOk() (*ClusterSpecs, bool) {
	if o == nil || IsNil(o.ElectableSpecs) {
		return nil, false
	}

	return o.ElectableSpecs, true
}

// HasElectableSpecs returns a boolean if a field has been set.
func (o *ClusterAdvancedRegionConfig) HasElectableSpecs() bool {
	if o != nil && !IsNil(o.ElectableSpecs) {
		return true
	}

	return false
}

// SetElectableSpecs gets a reference to the given ClusterSpecs and assigns it to the ElectableSpecs field.
func (o *ClusterAdvancedRegionConfig) SetElectableSpecs(v ClusterSpecs) {
	o.ElectableSpecs = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise
func (o *ClusterAdvancedRegionConfig) GetPriority() int {
	if o == nil || IsNil(o.Priority) {
		var ret int
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedRegionConfig) GetPriorityOk() (*int, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}

	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ClusterAdvancedRegionConfig) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int and assigns it to the Priority field.
func (o *ClusterAdvancedRegionConfig) SetPriority(v int) {
	o.Priority = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise
func (o *ClusterAdvancedRegionConfig) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedRegionConfig) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}

	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *ClusterAdvancedRegionConfig) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *ClusterAdvancedRegionConfig) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetReadOnlySpecs returns the ReadOnlySpecs field value if set, zero value otherwise
func (o *ClusterAdvancedRegionConfig) GetReadOnlySpecs() ClusterSpecs {
	if o == nil || IsNil(o.ReadOnlySpecs) {
		var ret ClusterSpecs
		return ret
	}
	return *o.ReadOnlySpecs
}

// GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedRegionConfig) GetReadOnlySpecsOk() (*ClusterSpecs, bool) {
	if o == nil || IsNil(o.ReadOnlySpecs) {
		return nil, false
	}

	return o.ReadOnlySpecs, true
}

// HasReadOnlySpecs returns a boolean if a field has been set.
func (o *ClusterAdvancedRegionConfig) HasReadOnlySpecs() bool {
	if o != nil && !IsNil(o.ReadOnlySpecs) {
		return true
	}

	return false
}

// SetReadOnlySpecs gets a reference to the given ClusterSpecs and assigns it to the ReadOnlySpecs field.
func (o *ClusterAdvancedRegionConfig) SetReadOnlySpecs(v ClusterSpecs) {
	o.ReadOnlySpecs = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *ClusterAdvancedRegionConfig) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedRegionConfig) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *ClusterAdvancedRegionConfig) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *ClusterAdvancedRegionConfig) SetRegionName(v string) {
	o.RegionName = &v
}

func (o ClusterAdvancedRegionConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterAdvancedRegionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackingProviderName) {
		toSerialize["BackingProviderName"] = o.BackingProviderName
	}
	if !IsNil(o.AnalyticsAutoScaling) {
		toSerialize["analyticsAutoScaling"] = o.AnalyticsAutoScaling
	}
	if !IsNil(o.AnalyticsSpecs) {
		toSerialize["analyticsSpecs"] = o.AnalyticsSpecs
	}
	if !IsNil(o.AutoScaling) {
		toSerialize["autoScaling"] = o.AutoScaling
	}
	if !IsNil(o.ElectableSpecs) {
		toSerialize["electableSpecs"] = o.ElectableSpecs
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !IsNil(o.ReadOnlySpecs) {
		toSerialize["readOnlySpecs"] = o.ReadOnlySpecs
	}
	if !IsNil(o.RegionName) {
		toSerialize["regionName"] = o.RegionName
	}
	return toSerialize, nil
}
