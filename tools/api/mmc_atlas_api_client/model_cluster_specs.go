// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterSpecs struct for ClusterSpecs
type ClusterSpecs struct {
	InstanceSize  *string `json:"InstanceSize,omitempty"`
	DiskIOPS      *string `json:"diskIOPS,omitempty"`
	EbsVolumeType *string `json:"ebsVolumeType,omitempty"`
	NodeCount     *int    `json:"nodeCount,omitempty"`
}

// NewClusterSpecs instantiates a new ClusterSpecs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterSpecs() *ClusterSpecs {
	this := ClusterSpecs{}
	return &this
}

// NewClusterSpecsWithDefaults instantiates a new ClusterSpecs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterSpecsWithDefaults() *ClusterSpecs {
	this := ClusterSpecs{}
	return &this
}

// GetInstanceSize returns the InstanceSize field value if set, zero value otherwise
func (o *ClusterSpecs) GetInstanceSize() string {
	if o == nil || IsNil(o.InstanceSize) {
		var ret string
		return ret
	}
	return *o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSpecs) GetInstanceSizeOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceSize) {
		return nil, false
	}

	return o.InstanceSize, true
}

// HasInstanceSize returns a boolean if a field has been set.
func (o *ClusterSpecs) HasInstanceSize() bool {
	if o != nil && !IsNil(o.InstanceSize) {
		return true
	}

	return false
}

// SetInstanceSize gets a reference to the given string and assigns it to the InstanceSize field.
func (o *ClusterSpecs) SetInstanceSize(v string) {
	o.InstanceSize = &v
}

// GetDiskIOPS returns the DiskIOPS field value if set, zero value otherwise
func (o *ClusterSpecs) GetDiskIOPS() string {
	if o == nil || IsNil(o.DiskIOPS) {
		var ret string
		return ret
	}
	return *o.DiskIOPS
}

// GetDiskIOPSOk returns a tuple with the DiskIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSpecs) GetDiskIOPSOk() (*string, bool) {
	if o == nil || IsNil(o.DiskIOPS) {
		return nil, false
	}

	return o.DiskIOPS, true
}

// HasDiskIOPS returns a boolean if a field has been set.
func (o *ClusterSpecs) HasDiskIOPS() bool {
	if o != nil && !IsNil(o.DiskIOPS) {
		return true
	}

	return false
}

// SetDiskIOPS gets a reference to the given string and assigns it to the DiskIOPS field.
func (o *ClusterSpecs) SetDiskIOPS(v string) {
	o.DiskIOPS = &v
}

// GetEbsVolumeType returns the EbsVolumeType field value if set, zero value otherwise
func (o *ClusterSpecs) GetEbsVolumeType() string {
	if o == nil || IsNil(o.EbsVolumeType) {
		var ret string
		return ret
	}
	return *o.EbsVolumeType
}

// GetEbsVolumeTypeOk returns a tuple with the EbsVolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSpecs) GetEbsVolumeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EbsVolumeType) {
		return nil, false
	}

	return o.EbsVolumeType, true
}

// HasEbsVolumeType returns a boolean if a field has been set.
func (o *ClusterSpecs) HasEbsVolumeType() bool {
	if o != nil && !IsNil(o.EbsVolumeType) {
		return true
	}

	return false
}

// SetEbsVolumeType gets a reference to the given string and assigns it to the EbsVolumeType field.
func (o *ClusterSpecs) SetEbsVolumeType(v string) {
	o.EbsVolumeType = &v
}

// GetNodeCount returns the NodeCount field value if set, zero value otherwise
func (o *ClusterSpecs) GetNodeCount() int {
	if o == nil || IsNil(o.NodeCount) {
		var ret int
		return ret
	}
	return *o.NodeCount
}

// GetNodeCountOk returns a tuple with the NodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterSpecs) GetNodeCountOk() (*int, bool) {
	if o == nil || IsNil(o.NodeCount) {
		return nil, false
	}

	return o.NodeCount, true
}

// HasNodeCount returns a boolean if a field has been set.
func (o *ClusterSpecs) HasNodeCount() bool {
	if o != nil && !IsNil(o.NodeCount) {
		return true
	}

	return false
}

// SetNodeCount gets a reference to the given int and assigns it to the NodeCount field.
func (o *ClusterSpecs) SetNodeCount(v int) {
	o.NodeCount = &v
}

func (o ClusterSpecs) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterSpecs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InstanceSize) {
		toSerialize["InstanceSize"] = o.InstanceSize
	}
	if !IsNil(o.DiskIOPS) {
		toSerialize["diskIOPS"] = o.DiskIOPS
	}
	if !IsNil(o.EbsVolumeType) {
		toSerialize["ebsVolumeType"] = o.EbsVolumeType
	}
	if !IsNil(o.NodeCount) {
		toSerialize["nodeCount"] = o.NodeCount
	}
	return toSerialize, nil
}
