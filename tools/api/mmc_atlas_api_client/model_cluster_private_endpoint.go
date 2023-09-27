// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterPrivateEndpoint struct for ClusterPrivateEndpoint
type ClusterPrivateEndpoint struct {
	ConnectionString    *string           `json:"connectionString,omitempty"`
	Endpoints           []ClusterEndpoint `json:"endpoints,omitempty"`
	SrvConnectionString *string           `json:"srvConnectionString,omitempty"`
	Type                *string           `json:"type,omitempty"`
}

// NewClusterPrivateEndpoint instantiates a new ClusterPrivateEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterPrivateEndpoint() *ClusterPrivateEndpoint {
	this := ClusterPrivateEndpoint{}
	return &this
}

// NewClusterPrivateEndpointWithDefaults instantiates a new ClusterPrivateEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterPrivateEndpointWithDefaults() *ClusterPrivateEndpoint {
	this := ClusterPrivateEndpoint{}
	return &this
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise
func (o *ClusterPrivateEndpoint) GetConnectionString() string {
	if o == nil || IsNil(o.ConnectionString) {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPrivateEndpoint) GetConnectionStringOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionString) {
		return nil, false
	}

	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *ClusterPrivateEndpoint) HasConnectionString() bool {
	if o != nil && !IsNil(o.ConnectionString) {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *ClusterPrivateEndpoint) SetConnectionString(v string) {
	o.ConnectionString = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise
func (o *ClusterPrivateEndpoint) GetEndpoints() []ClusterEndpoint {
	if o == nil || IsNil(o.Endpoints) {
		var ret []ClusterEndpoint
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPrivateEndpoint) GetEndpointsOk() ([]ClusterEndpoint, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}

	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ClusterPrivateEndpoint) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []ClusterEndpoint and assigns it to the Endpoints field.
func (o *ClusterPrivateEndpoint) SetEndpoints(v []ClusterEndpoint) {
	o.Endpoints = v
}

// GetSrvConnectionString returns the SrvConnectionString field value if set, zero value otherwise
func (o *ClusterPrivateEndpoint) GetSrvConnectionString() string {
	if o == nil || IsNil(o.SrvConnectionString) {
		var ret string
		return ret
	}
	return *o.SrvConnectionString
}

// GetSrvConnectionStringOk returns a tuple with the SrvConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPrivateEndpoint) GetSrvConnectionStringOk() (*string, bool) {
	if o == nil || IsNil(o.SrvConnectionString) {
		return nil, false
	}

	return o.SrvConnectionString, true
}

// HasSrvConnectionString returns a boolean if a field has been set.
func (o *ClusterPrivateEndpoint) HasSrvConnectionString() bool {
	if o != nil && !IsNil(o.SrvConnectionString) {
		return true
	}

	return false
}

// SetSrvConnectionString gets a reference to the given string and assigns it to the SrvConnectionString field.
func (o *ClusterPrivateEndpoint) SetSrvConnectionString(v string) {
	o.SrvConnectionString = &v
}

// GetType returns the Type field value if set, zero value otherwise
func (o *ClusterPrivateEndpoint) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterPrivateEndpoint) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ClusterPrivateEndpoint) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ClusterPrivateEndpoint) SetType(v string) {
	o.Type = &v
}

func (o ClusterPrivateEndpoint) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterPrivateEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConnectionString) {
		toSerialize["connectionString"] = o.ConnectionString
	}
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	if !IsNil(o.SrvConnectionString) {
		toSerialize["srvConnectionString"] = o.SrvConnectionString
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}
