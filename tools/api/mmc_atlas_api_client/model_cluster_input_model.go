// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// ClusterInputModel struct for ClusterInputModel
type ClusterInputModel struct {
	CloudProvider  *string `json:"cloudProvider,omitempty"`
	MongoDBVersion *string `json:"mongoDBVersion,omitempty"`
	ProjectId      *string `json:"projectId,omitempty"`
	TshirtSize     *string `json:"tshirtSize,omitempty"`
}

// NewClusterInputModel instantiates a new ClusterInputModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterInputModel() *ClusterInputModel {
	this := ClusterInputModel{}
	return &this
}

// NewClusterInputModelWithDefaults instantiates a new ClusterInputModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterInputModelWithDefaults() *ClusterInputModel {
	this := ClusterInputModel{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *ClusterInputModel) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInputModel) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ClusterInputModel) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ClusterInputModel) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetMongoDBVersion returns the MongoDBVersion field value if set, zero value otherwise
func (o *ClusterInputModel) GetMongoDBVersion() string {
	if o == nil || IsNil(o.MongoDBVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBVersion
}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInputModel) GetMongoDBVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBVersion) {
		return nil, false
	}

	return o.MongoDBVersion, true
}

// HasMongoDBVersion returns a boolean if a field has been set.
func (o *ClusterInputModel) HasMongoDBVersion() bool {
	if o != nil && !IsNil(o.MongoDBVersion) {
		return true
	}

	return false
}

// SetMongoDBVersion gets a reference to the given string and assigns it to the MongoDBVersion field.
func (o *ClusterInputModel) SetMongoDBVersion(v string) {
	o.MongoDBVersion = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise
func (o *ClusterInputModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInputModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}

	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ClusterInputModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ClusterInputModel) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetTshirtSize returns the TshirtSize field value if set, zero value otherwise
func (o *ClusterInputModel) GetTshirtSize() string {
	if o == nil || IsNil(o.TshirtSize) {
		var ret string
		return ret
	}
	return *o.TshirtSize
}

// GetTshirtSizeOk returns a tuple with the TshirtSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInputModel) GetTshirtSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TshirtSize) {
		return nil, false
	}

	return o.TshirtSize, true
}

// HasTshirtSize returns a boolean if a field has been set.
func (o *ClusterInputModel) HasTshirtSize() bool {
	if o != nil && !IsNil(o.TshirtSize) {
		return true
	}

	return false
}

// SetTshirtSize gets a reference to the given string and assigns it to the TshirtSize field.
func (o *ClusterInputModel) SetTshirtSize(v string) {
	o.TshirtSize = &v
}

func (o ClusterInputModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ClusterInputModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.MongoDBVersion) {
		toSerialize["mongoDBVersion"] = o.MongoDBVersion
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.TshirtSize) {
		toSerialize["tshirtSize"] = o.TshirtSize
	}
	return toSerialize, nil
}
