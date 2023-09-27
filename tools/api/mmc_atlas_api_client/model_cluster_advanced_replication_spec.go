// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterAdvancedReplicationSpec struct for ClusterAdvancedReplicationSpec
type ClusterAdvancedReplicationSpec struct {
	AdvancedRegionConfigs []ClusterAdvancedRegionConfig `json:"advancedRegionConfigs,omitempty"`
	Id                    *string                       `json:"id,omitempty"`
	NumShards             *int                          `json:"numShards,omitempty"`
	ZoneName              *string                       `json:"zoneName,omitempty"`
}

// NewClusterAdvancedReplicationSpec instantiates a new ClusterAdvancedReplicationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterAdvancedReplicationSpec() *ClusterAdvancedReplicationSpec {
	this := ClusterAdvancedReplicationSpec{}
	return &this
}

// NewClusterAdvancedReplicationSpecWithDefaults instantiates a new ClusterAdvancedReplicationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterAdvancedReplicationSpecWithDefaults() *ClusterAdvancedReplicationSpec {
	this := ClusterAdvancedReplicationSpec{}
	return &this
}

// GetAdvancedRegionConfigs returns the AdvancedRegionConfigs field value if set, zero value otherwise
func (o *ClusterAdvancedReplicationSpec) GetAdvancedRegionConfigs() []ClusterAdvancedRegionConfig {
	if o == nil || IsNil(o.AdvancedRegionConfigs) {
		var ret []ClusterAdvancedRegionConfig
		return ret
	}
	return o.AdvancedRegionConfigs
}

// GetAdvancedRegionConfigsOk returns a tuple with the AdvancedRegionConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedReplicationSpec) GetAdvancedRegionConfigsOk() ([]ClusterAdvancedRegionConfig, bool) {
	if o == nil || IsNil(o.AdvancedRegionConfigs) {
		return nil, false
	}

	return o.AdvancedRegionConfigs, true
}

// HasAdvancedRegionConfigs returns a boolean if a field has been set.
func (o *ClusterAdvancedReplicationSpec) HasAdvancedRegionConfigs() bool {
	if o != nil && !IsNil(o.AdvancedRegionConfigs) {
		return true
	}

	return false
}

// SetAdvancedRegionConfigs gets a reference to the given []ClusterAdvancedRegionConfig and assigns it to the AdvancedRegionConfigs field.
func (o *ClusterAdvancedReplicationSpec) SetAdvancedRegionConfigs(v []ClusterAdvancedRegionConfig) {
	o.AdvancedRegionConfigs = v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ClusterAdvancedReplicationSpec) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedReplicationSpec) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterAdvancedReplicationSpec) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClusterAdvancedReplicationSpec) SetId(v string) {
	o.Id = &v
}

// GetNumShards returns the NumShards field value if set, zero value otherwise
func (o *ClusterAdvancedReplicationSpec) GetNumShards() int {
	if o == nil || IsNil(o.NumShards) {
		var ret int
		return ret
	}
	return *o.NumShards
}

// GetNumShardsOk returns a tuple with the NumShards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedReplicationSpec) GetNumShardsOk() (*int, bool) {
	if o == nil || IsNil(o.NumShards) {
		return nil, false
	}

	return o.NumShards, true
}

// HasNumShards returns a boolean if a field has been set.
func (o *ClusterAdvancedReplicationSpec) HasNumShards() bool {
	if o != nil && !IsNil(o.NumShards) {
		return true
	}

	return false
}

// SetNumShards gets a reference to the given int and assigns it to the NumShards field.
func (o *ClusterAdvancedReplicationSpec) SetNumShards(v int) {
	o.NumShards = &v
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise
func (o *ClusterAdvancedReplicationSpec) GetZoneName() string {
	if o == nil || IsNil(o.ZoneName) {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterAdvancedReplicationSpec) GetZoneNameOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneName) {
		return nil, false
	}

	return o.ZoneName, true
}

// HasZoneName returns a boolean if a field has been set.
func (o *ClusterAdvancedReplicationSpec) HasZoneName() bool {
	if o != nil && !IsNil(o.ZoneName) {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *ClusterAdvancedReplicationSpec) SetZoneName(v string) {
	o.ZoneName = &v
}

func (o ClusterAdvancedReplicationSpec) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterAdvancedReplicationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdvancedRegionConfigs) {
		toSerialize["advancedRegionConfigs"] = o.AdvancedRegionConfigs
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NumShards) {
		toSerialize["numShards"] = o.NumShards
	}
	if !IsNil(o.ZoneName) {
		toSerialize["zoneName"] = o.ZoneName
	}
	return toSerialize, nil
}
