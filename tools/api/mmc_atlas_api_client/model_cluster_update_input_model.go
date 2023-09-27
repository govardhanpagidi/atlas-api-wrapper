// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterUpdateInputModel struct for ClusterUpdateInputModel
type ClusterUpdateInputModel struct {
	ClusterName         *string `json:"clusterName,omitempty"`
	MongoDBMajorVersion *string `json:"mongoDBMajorVersion,omitempty"`
	ProjectId           *string `json:"projectId,omitempty"`
}

// NewClusterUpdateInputModel instantiates a new ClusterUpdateInputModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterUpdateInputModel() *ClusterUpdateInputModel {
	this := ClusterUpdateInputModel{}
	return &this
}

// NewClusterUpdateInputModelWithDefaults instantiates a new ClusterUpdateInputModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterUpdateInputModelWithDefaults() *ClusterUpdateInputModel {
	this := ClusterUpdateInputModel{}
	return &this
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *ClusterUpdateInputModel) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdateInputModel) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *ClusterUpdateInputModel) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *ClusterUpdateInputModel) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetMongoDBMajorVersion returns the MongoDBMajorVersion field value if set, zero value otherwise
func (o *ClusterUpdateInputModel) GetMongoDBMajorVersion() string {
	if o == nil || IsNil(o.MongoDBMajorVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBMajorVersion
}

// GetMongoDBMajorVersionOk returns a tuple with the MongoDBMajorVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdateInputModel) GetMongoDBMajorVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBMajorVersion) {
		return nil, false
	}

	return o.MongoDBMajorVersion, true
}

// HasMongoDBMajorVersion returns a boolean if a field has been set.
func (o *ClusterUpdateInputModel) HasMongoDBMajorVersion() bool {
	if o != nil && !IsNil(o.MongoDBMajorVersion) {
		return true
	}

	return false
}

// SetMongoDBMajorVersion gets a reference to the given string and assigns it to the MongoDBMajorVersion field.
func (o *ClusterUpdateInputModel) SetMongoDBMajorVersion(v string) {
	o.MongoDBMajorVersion = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise
func (o *ClusterUpdateInputModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdateInputModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}

	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ClusterUpdateInputModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ClusterUpdateInputModel) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o ClusterUpdateInputModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterUpdateInputModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !IsNil(o.MongoDBMajorVersion) {
		toSerialize["mongoDBMajorVersion"] = o.MongoDBMajorVersion
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	return toSerialize, nil
}
