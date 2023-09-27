// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// DatabaseUserRoleDefinition struct for DatabaseUserRoleDefinition
type DatabaseUserRoleDefinition struct {
	CollectionName *string `json:"collectionName,omitempty"`
	DatabaseName   *string `json:"databaseName,omitempty"`
	RoleName       *string `json:"roleName,omitempty"`
}

// NewDatabaseUserRoleDefinition instantiates a new DatabaseUserRoleDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseUserRoleDefinition() *DatabaseUserRoleDefinition {
	this := DatabaseUserRoleDefinition{}
	return &this
}

// NewDatabaseUserRoleDefinitionWithDefaults instantiates a new DatabaseUserRoleDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseUserRoleDefinitionWithDefaults() *DatabaseUserRoleDefinition {
	this := DatabaseUserRoleDefinition{}
	return &this
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise
func (o *DatabaseUserRoleDefinition) GetCollectionName() string {
	if o == nil || IsNil(o.CollectionName) {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserRoleDefinition) GetCollectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionName) {
		return nil, false
	}

	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *DatabaseUserRoleDefinition) HasCollectionName() bool {
	if o != nil && !IsNil(o.CollectionName) {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *DatabaseUserRoleDefinition) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise
func (o *DatabaseUserRoleDefinition) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserRoleDefinition) GetDatabaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseName) {
		return nil, false
	}

	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *DatabaseUserRoleDefinition) HasDatabaseName() bool {
	if o != nil && !IsNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *DatabaseUserRoleDefinition) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise
func (o *DatabaseUserRoleDefinition) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserRoleDefinition) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}

	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *DatabaseUserRoleDefinition) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *DatabaseUserRoleDefinition) SetRoleName(v string) {
	o.RoleName = &v
}

func (o DatabaseUserRoleDefinition) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DatabaseUserRoleDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionName) {
		toSerialize["collectionName"] = o.CollectionName
	}
	if !IsNil(o.DatabaseName) {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !IsNil(o.RoleName) {
		toSerialize["roleName"] = o.RoleName
	}
	return toSerialize, nil
}
