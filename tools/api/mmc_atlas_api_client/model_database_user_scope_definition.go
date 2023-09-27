// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// DatabaseUserScopeDefinition struct for DatabaseUserScopeDefinition
type DatabaseUserScopeDefinition struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewDatabaseUserScopeDefinition instantiates a new DatabaseUserScopeDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseUserScopeDefinition() *DatabaseUserScopeDefinition {
	this := DatabaseUserScopeDefinition{}
	return &this
}

// NewDatabaseUserScopeDefinitionWithDefaults instantiates a new DatabaseUserScopeDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseUserScopeDefinitionWithDefaults() *DatabaseUserScopeDefinition {
	this := DatabaseUserScopeDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise
func (o *DatabaseUserScopeDefinition) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserScopeDefinition) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatabaseUserScopeDefinition) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatabaseUserScopeDefinition) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise
func (o *DatabaseUserScopeDefinition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserScopeDefinition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DatabaseUserScopeDefinition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DatabaseUserScopeDefinition) SetType(v string) {
	o.Type = &v
}

func (o DatabaseUserScopeDefinition) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DatabaseUserScopeDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}
