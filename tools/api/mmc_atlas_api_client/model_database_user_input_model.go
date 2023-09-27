// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// DatabaseUserInputModel struct for DatabaseUserInputModel
type DatabaseUserInputModel struct {
	Password  *string `json:"password,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
	Username  *string `json:"username,omitempty"`
}

// NewDatabaseUserInputModel instantiates a new DatabaseUserInputModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseUserInputModel() *DatabaseUserInputModel {
	this := DatabaseUserInputModel{}
	return &this
}

// NewDatabaseUserInputModelWithDefaults instantiates a new DatabaseUserInputModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseUserInputModelWithDefaults() *DatabaseUserInputModel {
	this := DatabaseUserInputModel{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise
func (o *DatabaseUserInputModel) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserInputModel) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}

	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *DatabaseUserInputModel) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *DatabaseUserInputModel) SetPassword(v string) {
	o.Password = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise
func (o *DatabaseUserInputModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserInputModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}

	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DatabaseUserInputModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *DatabaseUserInputModel) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise
func (o *DatabaseUserInputModel) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseUserInputModel) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}

	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *DatabaseUserInputModel) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *DatabaseUserInputModel) SetUsername(v string) {
	o.Username = &v
}

func (o DatabaseUserInputModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DatabaseUserInputModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}
