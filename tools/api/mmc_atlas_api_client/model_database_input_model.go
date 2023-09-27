// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// DatabaseInputModel struct for DatabaseInputModel
type DatabaseInputModel struct {
	CollectionName *string `json:"collectionName,omitempty"`
	DatabaseName   *string `json:"databaseName,omitempty"`
	HostName       *string `json:"hostName,omitempty"`
}

// NewDatabaseInputModel instantiates a new DatabaseInputModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseInputModel() *DatabaseInputModel {
	this := DatabaseInputModel{}
	return &this
}

// NewDatabaseInputModelWithDefaults instantiates a new DatabaseInputModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseInputModelWithDefaults() *DatabaseInputModel {
	this := DatabaseInputModel{}
	return &this
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise
func (o *DatabaseInputModel) GetCollectionName() string {
	if o == nil || IsNil(o.CollectionName) {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInputModel) GetCollectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionName) {
		return nil, false
	}

	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *DatabaseInputModel) HasCollectionName() bool {
	if o != nil && !IsNil(o.CollectionName) {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *DatabaseInputModel) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise
func (o *DatabaseInputModel) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInputModel) GetDatabaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseName) {
		return nil, false
	}

	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *DatabaseInputModel) HasDatabaseName() bool {
	if o != nil && !IsNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *DatabaseInputModel) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise
func (o *DatabaseInputModel) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseInputModel) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}

	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *DatabaseInputModel) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *DatabaseInputModel) SetHostName(v string) {
	o.HostName = &v
}

func (o DatabaseInputModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DatabaseInputModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionName) {
		toSerialize["collectionName"] = o.CollectionName
	}
	if !IsNil(o.DatabaseName) {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !IsNil(o.HostName) {
		toSerialize["hostName"] = o.HostName
	}
	return toSerialize, nil
}
