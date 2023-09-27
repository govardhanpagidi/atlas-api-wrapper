// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember struct for CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember
type CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember struct {
	CloudProvider  *string `json:"cloudProvider,omitempty"`
	Id             *string `json:"id,omitempty"`
	ReplicaSetName *string `json:"replicaSetName,omitempty"`
}

// NewCloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember instantiates a new CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember() *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember {
	this := CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember{}
	return &this
}

// NewCloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMemberWithDefaults instantiates a new CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMemberWithDefaults() *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember {
	this := CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) SetId(v string) {
	o.Id = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}

	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

func (o CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ReplicaSetName) {
		toSerialize["replicaSetName"] = o.ReplicaSetName
	}
	return toSerialize, nil
}
