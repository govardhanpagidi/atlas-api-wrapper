// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CollectionInputModel struct for CollectionInputModel
type CollectionInputModel struct {
	CollectionNames []string `json:"collectionNames,omitempty"`
	DatabaseName    *string  `json:"databaseName,omitempty"`
	HostName        *string  `json:"hostName,omitempty"`
}

// NewCollectionInputModel instantiates a new CollectionInputModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionInputModel() *CollectionInputModel {
	this := CollectionInputModel{}
	return &this
}

// NewCollectionInputModelWithDefaults instantiates a new CollectionInputModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionInputModelWithDefaults() *CollectionInputModel {
	this := CollectionInputModel{}
	return &this
}

// GetCollectionNames returns the CollectionNames field value if set, zero value otherwise
func (o *CollectionInputModel) GetCollectionNames() []string {
	if o == nil || IsNil(o.CollectionNames) {
		var ret []string
		return ret
	}
	return o.CollectionNames
}

// GetCollectionNamesOk returns a tuple with the CollectionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionInputModel) GetCollectionNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.CollectionNames) {
		return nil, false
	}

	return o.CollectionNames, true
}

// HasCollectionNames returns a boolean if a field has been set.
func (o *CollectionInputModel) HasCollectionNames() bool {
	if o != nil && !IsNil(o.CollectionNames) {
		return true
	}

	return false
}

// SetCollectionNames gets a reference to the given []string and assigns it to the CollectionNames field.
func (o *CollectionInputModel) SetCollectionNames(v []string) {
	o.CollectionNames = v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise
func (o *CollectionInputModel) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionInputModel) GetDatabaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseName) {
		return nil, false
	}

	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *CollectionInputModel) HasDatabaseName() bool {
	if o != nil && !IsNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *CollectionInputModel) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise
func (o *CollectionInputModel) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionInputModel) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}

	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *CollectionInputModel) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *CollectionInputModel) SetHostName(v string) {
	o.HostName = &v
}

func (o CollectionInputModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CollectionInputModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionNames) {
		toSerialize["collectionNames"] = o.CollectionNames
	}
	if !IsNil(o.DatabaseName) {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !IsNil(o.HostName) {
		toSerialize["hostName"] = o.HostName
	}
	return toSerialize, nil
}
