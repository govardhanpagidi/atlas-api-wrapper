// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterBiConnector struct for ClusterBiConnector
type ClusterBiConnector struct {
	Enabled        *bool   `json:"enabled,omitempty"`
	ReadPreference *string `json:"readPreference,omitempty"`
}

// NewClusterBiConnector instantiates a new ClusterBiConnector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterBiConnector() *ClusterBiConnector {
	this := ClusterBiConnector{}
	return &this
}

// NewClusterBiConnectorWithDefaults instantiates a new ClusterBiConnector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterBiConnectorWithDefaults() *ClusterBiConnector {
	this := ClusterBiConnector{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *ClusterBiConnector) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBiConnector) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClusterBiConnector) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClusterBiConnector) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetReadPreference returns the ReadPreference field value if set, zero value otherwise
func (o *ClusterBiConnector) GetReadPreference() string {
	if o == nil || IsNil(o.ReadPreference) {
		var ret string
		return ret
	}
	return *o.ReadPreference
}

// GetReadPreferenceOk returns a tuple with the ReadPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBiConnector) GetReadPreferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ReadPreference) {
		return nil, false
	}

	return o.ReadPreference, true
}

// HasReadPreference returns a boolean if a field has been set.
func (o *ClusterBiConnector) HasReadPreference() bool {
	if o != nil && !IsNil(o.ReadPreference) {
		return true
	}

	return false
}

// SetReadPreference gets a reference to the given string and assigns it to the ReadPreference field.
func (o *ClusterBiConnector) SetReadPreference(v string) {
	o.ReadPreference = &v
}

func (o ClusterBiConnector) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterBiConnector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ReadPreference) {
		toSerialize["readPreference"] = o.ReadPreference
	}
	return toSerialize, nil
}
