// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot struct for CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot
type CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot struct {
	CreatedAt        *string                                                             `json:"createdAt,omitempty"`
	Description      *string                                                             `json:"description,omitempty"`
	ExpiresAt        *string                                                             `json:"expiresAt,omitempty"`
	FrequencyType    *string                                                             `json:"frequencyType,omitempty"`
	Id               *string                                                             `json:"id,omitempty"`
	MasterKeyUUID    *string                                                             `json:"masterKeyUUID,omitempty"`
	Members          []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember `json:"members,omitempty"`
	MongodVersion    *string                                                             `json:"mongodVersion,omitempty"`
	PolicyItems      []string                                                            `json:"policyItems,omitempty"`
	SnapshotIds      []string                                                            `json:"snapshotIds,omitempty"`
	SnapshotType     *string                                                             `json:"snapshotType,omitempty"`
	Status           *string                                                             `json:"status,omitempty"`
	StorageSizeBytes *string                                                             `json:"storageSizeBytes,omitempty"`
	Type             *string                                                             `json:"type,omitempty"`
}

// NewCloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot instantiates a new CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot() *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot {
	this := CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot{}
	return &this
}

// NewCloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotWithDefaults instantiates a new CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotWithDefaults() *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot {
	this := CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}

	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetFrequencyType returns the FrequencyType field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetFrequencyType() string {
	if o == nil || IsNil(o.FrequencyType) {
		var ret string
		return ret
	}
	return *o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetFrequencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FrequencyType) {
		return nil, false
	}

	return o.FrequencyType, true
}

// HasFrequencyType returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasFrequencyType() bool {
	if o != nil && !IsNil(o.FrequencyType) {
		return true
	}

	return false
}

// SetFrequencyType gets a reference to the given string and assigns it to the FrequencyType field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetFrequencyType(v string) {
	o.FrequencyType = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetId(v string) {
	o.Id = &v
}

// GetMasterKeyUUID returns the MasterKeyUUID field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetMasterKeyUUID() string {
	if o == nil || IsNil(o.MasterKeyUUID) {
		var ret string
		return ret
	}
	return *o.MasterKeyUUID
}

// GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetMasterKeyUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.MasterKeyUUID) {
		return nil, false
	}

	return o.MasterKeyUUID, true
}

// HasMasterKeyUUID returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasMasterKeyUUID() bool {
	if o != nil && !IsNil(o.MasterKeyUUID) {
		return true
	}

	return false
}

// SetMasterKeyUUID gets a reference to the given string and assigns it to the MasterKeyUUID field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetMasterKeyUUID(v string) {
	o.MasterKeyUUID = &v
}

// GetMembers returns the Members field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetMembers() []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember {
	if o == nil || IsNil(o.Members) {
		var ret []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetMembersOk() ([]CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}

	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember and assigns it to the Members field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetMembers(v []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) {
	o.Members = v
}

// GetMongodVersion returns the MongodVersion field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetMongodVersion() string {
	if o == nil || IsNil(o.MongodVersion) {
		var ret string
		return ret
	}
	return *o.MongodVersion
}

// GetMongodVersionOk returns a tuple with the MongodVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetMongodVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongodVersion) {
		return nil, false
	}

	return o.MongodVersion, true
}

// HasMongodVersion returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasMongodVersion() bool {
	if o != nil && !IsNil(o.MongodVersion) {
		return true
	}

	return false
}

// SetMongodVersion gets a reference to the given string and assigns it to the MongodVersion field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetMongodVersion(v string) {
	o.MongodVersion = &v
}

// GetPolicyItems returns the PolicyItems field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetPolicyItems() []string {
	if o == nil || IsNil(o.PolicyItems) {
		var ret []string
		return ret
	}
	return o.PolicyItems
}

// GetPolicyItemsOk returns a tuple with the PolicyItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetPolicyItemsOk() ([]string, bool) {
	if o == nil || IsNil(o.PolicyItems) {
		return nil, false
	}

	return o.PolicyItems, true
}

// HasPolicyItems returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasPolicyItems() bool {
	if o != nil && !IsNil(o.PolicyItems) {
		return true
	}

	return false
}

// SetPolicyItems gets a reference to the given []string and assigns it to the PolicyItems field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetPolicyItems(v []string) {
	o.PolicyItems = v
}

// GetSnapshotIds returns the SnapshotIds field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetSnapshotIds() []string {
	if o == nil || IsNil(o.SnapshotIds) {
		var ret []string
		return ret
	}
	return o.SnapshotIds
}

// GetSnapshotIdsOk returns a tuple with the SnapshotIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetSnapshotIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SnapshotIds) {
		return nil, false
	}

	return o.SnapshotIds, true
}

// HasSnapshotIds returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasSnapshotIds() bool {
	if o != nil && !IsNil(o.SnapshotIds) {
		return true
	}

	return false
}

// SetSnapshotIds gets a reference to the given []string and assigns it to the SnapshotIds field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetSnapshotIds(v []string) {
	o.SnapshotIds = v
}

// GetSnapshotType returns the SnapshotType field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetSnapshotType() string {
	if o == nil || IsNil(o.SnapshotType) {
		var ret string
		return ret
	}
	return *o.SnapshotType
}

// GetSnapshotTypeOk returns a tuple with the SnapshotType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetSnapshotTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotType) {
		return nil, false
	}

	return o.SnapshotType, true
}

// HasSnapshotType returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasSnapshotType() bool {
	if o != nil && !IsNil(o.SnapshotType) {
		return true
	}

	return false
}

// SetSnapshotType gets a reference to the given string and assigns it to the SnapshotType field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetSnapshotType(v string) {
	o.SnapshotType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetStatus(v string) {
	o.Status = &v
}

// GetStorageSizeBytes returns the StorageSizeBytes field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetStorageSizeBytes() string {
	if o == nil || IsNil(o.StorageSizeBytes) {
		var ret string
		return ret
	}
	return *o.StorageSizeBytes
}

// GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetStorageSizeBytesOk() (*string, bool) {
	if o == nil || IsNil(o.StorageSizeBytes) {
		return nil, false
	}

	return o.StorageSizeBytes, true
}

// HasStorageSizeBytes returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasStorageSizeBytes() bool {
	if o != nil && !IsNil(o.StorageSizeBytes) {
		return true
	}

	return false
}

// SetStorageSizeBytes gets a reference to the given string and assigns it to the StorageSizeBytes field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetStorageSizeBytes(v string) {
	o.StorageSizeBytes = &v
}

// GetType returns the Type field value if set, zero value otherwise
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) SetType(v string) {
	o.Type = &v
}

func (o CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.FrequencyType) {
		toSerialize["frequencyType"] = o.FrequencyType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MasterKeyUUID) {
		toSerialize["masterKeyUUID"] = o.MasterKeyUUID
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.MongodVersion) {
		toSerialize["mongodVersion"] = o.MongodVersion
	}
	if !IsNil(o.PolicyItems) {
		toSerialize["policyItems"] = o.PolicyItems
	}
	if !IsNil(o.SnapshotIds) {
		toSerialize["snapshotIds"] = o.SnapshotIds
	}
	if !IsNil(o.SnapshotType) {
		toSerialize["snapshotType"] = o.SnapshotType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StorageSizeBytes) {
		toSerialize["storageSizeBytes"] = o.StorageSizeBytes
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}
