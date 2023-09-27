// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterEndpoint struct for ClusterEndpoint
type ClusterEndpoint struct {
	EndpointID   *string `json:"endpointID,omitempty"`
	ProviderName *string `json:"providerName,omitempty"`
	Region       *string `json:"region,omitempty"`
}

// NewClusterEndpoint instantiates a new ClusterEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterEndpoint() *ClusterEndpoint {
	this := ClusterEndpoint{}
	return &this
}

// NewClusterEndpointWithDefaults instantiates a new ClusterEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterEndpointWithDefaults() *ClusterEndpoint {
	this := ClusterEndpoint{}
	return &this
}

// GetEndpointID returns the EndpointID field value if set, zero value otherwise
func (o *ClusterEndpoint) GetEndpointID() string {
	if o == nil || IsNil(o.EndpointID) {
		var ret string
		return ret
	}
	return *o.EndpointID
}

// GetEndpointIDOk returns a tuple with the EndpointID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEndpoint) GetEndpointIDOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointID) {
		return nil, false
	}

	return o.EndpointID, true
}

// HasEndpointID returns a boolean if a field has been set.
func (o *ClusterEndpoint) HasEndpointID() bool {
	if o != nil && !IsNil(o.EndpointID) {
		return true
	}

	return false
}

// SetEndpointID gets a reference to the given string and assigns it to the EndpointID field.
func (o *ClusterEndpoint) SetEndpointID(v string) {
	o.EndpointID = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise
func (o *ClusterEndpoint) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEndpoint) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}

	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *ClusterEndpoint) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *ClusterEndpoint) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *ClusterEndpoint) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEndpoint) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ClusterEndpoint) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ClusterEndpoint) SetRegion(v string) {
	o.Region = &v
}

func (o ClusterEndpoint) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointID) {
		toSerialize["endpointID"] = o.EndpointID
	}
	if !IsNil(o.ProviderName) {
		toSerialize["providerName"] = o.ProviderName
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}
