// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterConnectionStrings struct for ClusterConnectionStrings
type ClusterConnectionStrings struct {
	AwsPrivateLink    *string                  `json:"awsPrivateLink,omitempty"`
	AwsPrivateLinkSrv *string                  `json:"awsPrivateLinkSrv,omitempty"`
	Private           *string                  `json:"private,omitempty"`
	PrivateEndpoint   []ClusterPrivateEndpoint `json:"privateEndpoint,omitempty"`
	PrivateSrv        *string                  `json:"privateSrv,omitempty"`
	Standard          *string                  `json:"standard,omitempty"`
	StandardSrv       *string                  `json:"standardSrv,omitempty"`
}

// NewClusterConnectionStrings instantiates a new ClusterConnectionStrings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConnectionStrings() *ClusterConnectionStrings {
	this := ClusterConnectionStrings{}
	return &this
}

// NewClusterConnectionStringsWithDefaults instantiates a new ClusterConnectionStrings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConnectionStringsWithDefaults() *ClusterConnectionStrings {
	this := ClusterConnectionStrings{}
	return &this
}

// GetAwsPrivateLink returns the AwsPrivateLink field value if set, zero value otherwise
func (o *ClusterConnectionStrings) GetAwsPrivateLink() string {
	if o == nil || IsNil(o.AwsPrivateLink) {
		var ret string
		return ret
	}
	return *o.AwsPrivateLink
}

// GetAwsPrivateLinkOk returns a tuple with the AwsPrivateLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionStrings) GetAwsPrivateLinkOk() (*string, bool) {
	if o == nil || IsNil(o.AwsPrivateLink) {
		return nil, false
	}

	return o.AwsPrivateLink, true
}

// HasAwsPrivateLink returns a boolean if a field has been set.
func (o *ClusterConnectionStrings) HasAwsPrivateLink() bool {
	if o != nil && !IsNil(o.AwsPrivateLink) {
		return true
	}

	return false
}

// SetAwsPrivateLink gets a reference to the given string and assigns it to the AwsPrivateLink field.
func (o *ClusterConnectionStrings) SetAwsPrivateLink(v string) {
	o.AwsPrivateLink = &v
}

// GetAwsPrivateLinkSrv returns the AwsPrivateLinkSrv field value if set, zero value otherwise
func (o *ClusterConnectionStrings) GetAwsPrivateLinkSrv() string {
	if o == nil || IsNil(o.AwsPrivateLinkSrv) {
		var ret string
		return ret
	}
	return *o.AwsPrivateLinkSrv
}

// GetAwsPrivateLinkSrvOk returns a tuple with the AwsPrivateLinkSrv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionStrings) GetAwsPrivateLinkSrvOk() (*string, bool) {
	if o == nil || IsNil(o.AwsPrivateLinkSrv) {
		return nil, false
	}

	return o.AwsPrivateLinkSrv, true
}

// HasAwsPrivateLinkSrv returns a boolean if a field has been set.
func (o *ClusterConnectionStrings) HasAwsPrivateLinkSrv() bool {
	if o != nil && !IsNil(o.AwsPrivateLinkSrv) {
		return true
	}

	return false
}

// SetAwsPrivateLinkSrv gets a reference to the given string and assigns it to the AwsPrivateLinkSrv field.
func (o *ClusterConnectionStrings) SetAwsPrivateLinkSrv(v string) {
	o.AwsPrivateLinkSrv = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise
func (o *ClusterConnectionStrings) GetPrivate() string {
	if o == nil || IsNil(o.Private) {
		var ret string
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionStrings) GetPrivateOk() (*string, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}

	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ClusterConnectionStrings) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given string and assigns it to the Private field.
func (o *ClusterConnectionStrings) SetPrivate(v string) {
	o.Private = &v
}

// GetPrivateEndpoint returns the PrivateEndpoint field value if set, zero value otherwise
func (o *ClusterConnectionStrings) GetPrivateEndpoint() []ClusterPrivateEndpoint {
	if o == nil || IsNil(o.PrivateEndpoint) {
		var ret []ClusterPrivateEndpoint
		return ret
	}
	return o.PrivateEndpoint
}

// GetPrivateEndpointOk returns a tuple with the PrivateEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionStrings) GetPrivateEndpointOk() ([]ClusterPrivateEndpoint, bool) {
	if o == nil || IsNil(o.PrivateEndpoint) {
		return nil, false
	}

	return o.PrivateEndpoint, true
}

// HasPrivateEndpoint returns a boolean if a field has been set.
func (o *ClusterConnectionStrings) HasPrivateEndpoint() bool {
	if o != nil && !IsNil(o.PrivateEndpoint) {
		return true
	}

	return false
}

// SetPrivateEndpoint gets a reference to the given []ClusterPrivateEndpoint and assigns it to the PrivateEndpoint field.
func (o *ClusterConnectionStrings) SetPrivateEndpoint(v []ClusterPrivateEndpoint) {
	o.PrivateEndpoint = v
}

// GetPrivateSrv returns the PrivateSrv field value if set, zero value otherwise
func (o *ClusterConnectionStrings) GetPrivateSrv() string {
	if o == nil || IsNil(o.PrivateSrv) {
		var ret string
		return ret
	}
	return *o.PrivateSrv
}

// GetPrivateSrvOk returns a tuple with the PrivateSrv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionStrings) GetPrivateSrvOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateSrv) {
		return nil, false
	}

	return o.PrivateSrv, true
}

// HasPrivateSrv returns a boolean if a field has been set.
func (o *ClusterConnectionStrings) HasPrivateSrv() bool {
	if o != nil && !IsNil(o.PrivateSrv) {
		return true
	}

	return false
}

// SetPrivateSrv gets a reference to the given string and assigns it to the PrivateSrv field.
func (o *ClusterConnectionStrings) SetPrivateSrv(v string) {
	o.PrivateSrv = &v
}

// GetStandard returns the Standard field value if set, zero value otherwise
func (o *ClusterConnectionStrings) GetStandard() string {
	if o == nil || IsNil(o.Standard) {
		var ret string
		return ret
	}
	return *o.Standard
}

// GetStandardOk returns a tuple with the Standard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionStrings) GetStandardOk() (*string, bool) {
	if o == nil || IsNil(o.Standard) {
		return nil, false
	}

	return o.Standard, true
}

// HasStandard returns a boolean if a field has been set.
func (o *ClusterConnectionStrings) HasStandard() bool {
	if o != nil && !IsNil(o.Standard) {
		return true
	}

	return false
}

// SetStandard gets a reference to the given string and assigns it to the Standard field.
func (o *ClusterConnectionStrings) SetStandard(v string) {
	o.Standard = &v
}

// GetStandardSrv returns the StandardSrv field value if set, zero value otherwise
func (o *ClusterConnectionStrings) GetStandardSrv() string {
	if o == nil || IsNil(o.StandardSrv) {
		var ret string
		return ret
	}
	return *o.StandardSrv
}

// GetStandardSrvOk returns a tuple with the StandardSrv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConnectionStrings) GetStandardSrvOk() (*string, bool) {
	if o == nil || IsNil(o.StandardSrv) {
		return nil, false
	}

	return o.StandardSrv, true
}

// HasStandardSrv returns a boolean if a field has been set.
func (o *ClusterConnectionStrings) HasStandardSrv() bool {
	if o != nil && !IsNil(o.StandardSrv) {
		return true
	}

	return false
}

// SetStandardSrv gets a reference to the given string and assigns it to the StandardSrv field.
func (o *ClusterConnectionStrings) SetStandardSrv(v string) {
	o.StandardSrv = &v
}

func (o ClusterConnectionStrings) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterConnectionStrings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsPrivateLink) {
		toSerialize["awsPrivateLink"] = o.AwsPrivateLink
	}
	if !IsNil(o.AwsPrivateLinkSrv) {
		toSerialize["awsPrivateLinkSrv"] = o.AwsPrivateLinkSrv
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.PrivateEndpoint) {
		toSerialize["privateEndpoint"] = o.PrivateEndpoint
	}
	if !IsNil(o.PrivateSrv) {
		toSerialize["privateSrv"] = o.PrivateSrv
	}
	if !IsNil(o.Standard) {
		toSerialize["standard"] = o.Standard
	}
	if !IsNil(o.StandardSrv) {
		toSerialize["standardSrv"] = o.StandardSrv
	}
	return toSerialize, nil
}
