// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupSnapshotModel struct for CloudBackupSnapshotModel
type CloudBackupSnapshotModel struct {
	CloudProvider    *string                                                             `json:"cloudProvider,omitempty"`
	ClusterName      *string                                                             `json:"clusterName,omitempty"`
	CreatedAt        *string                                                             `json:"createdAt,omitempty"`
	Description      *string                                                             `json:"description,omitempty"`
	ExpiresAt        *string                                                             `json:"expiresAt,omitempty"`
	FrequencyType    *string                                                             `json:"frequencyType,omitempty"`
	Id               *string                                                             `json:"id,omitempty"`
	IncludeCount     *bool                                                               `json:"includeCount,omitempty"`
	InstanceName     *string                                                             `json:"instanceName,omitempty"`
	ItemsPerPage     *int                                                                `json:"itemsPerPage,omitempty"`
	MasterKeyUUID    *string                                                             `json:"masterKeyUUID,omitempty"`
	Members          []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember `json:"members,omitempty"`
	MongodVersion    *string                                                             `json:"mongodVersion,omitempty"`
	PageNum          *int                                                                `json:"pageNum,omitempty"`
	PolicyItems      []string                                                            `json:"policyItems,omitempty"`
	Profile          *string                                                             `json:"profile,omitempty"`
	ProjectId        *string                                                             `json:"projectId,omitempty"`
	ReplicaSetName   *string                                                             `json:"replicaSetName,omitempty"`
	Results          []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot       `json:"results,omitempty"`
	RetentionInDays  *int                                                                `json:"retentionInDays,omitempty"`
	SnapshotId       *string                                                             `json:"snapshotId,omitempty"`
	SnapshotIds      []string                                                            `json:"snapshotIds,omitempty"`
	SnapshotType     *string                                                             `json:"snapshotType,omitempty"`
	Status           *string                                                             `json:"status,omitempty"`
	StorageSizeBytes *string                                                             `json:"storageSizeBytes,omitempty"`
	TotalCount       *float32                                                            `json:"totalCount,omitempty"`
	Type             *string                                                             `json:"type,omitempty"`
}

// NewCloudBackupSnapshotModel instantiates a new CloudBackupSnapshotModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupSnapshotModel() *CloudBackupSnapshotModel {
	this := CloudBackupSnapshotModel{}
	return &this
}

// NewCloudBackupSnapshotModelWithDefaults instantiates a new CloudBackupSnapshotModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupSnapshotModelWithDefaults() *CloudBackupSnapshotModel {
	this := CloudBackupSnapshotModel{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *CloudBackupSnapshotModel) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *CloudBackupSnapshotModel) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CloudBackupSnapshotModel) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudBackupSnapshotModel) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}

	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CloudBackupSnapshotModel) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetFrequencyType returns the FrequencyType field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetFrequencyType() string {
	if o == nil || IsNil(o.FrequencyType) {
		var ret string
		return ret
	}
	return *o.FrequencyType
}

// GetFrequencyTypeOk returns a tuple with the FrequencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetFrequencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FrequencyType) {
		return nil, false
	}

	return o.FrequencyType, true
}

// HasFrequencyType returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasFrequencyType() bool {
	if o != nil && !IsNil(o.FrequencyType) {
		return true
	}

	return false
}

// SetFrequencyType gets a reference to the given string and assigns it to the FrequencyType field.
func (o *CloudBackupSnapshotModel) SetFrequencyType(v string) {
	o.FrequencyType = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudBackupSnapshotModel) SetId(v string) {
	o.Id = &v
}

// GetIncludeCount returns the IncludeCount field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetIncludeCount() bool {
	if o == nil || IsNil(o.IncludeCount) {
		var ret bool
		return ret
	}
	return *o.IncludeCount
}

// GetIncludeCountOk returns a tuple with the IncludeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetIncludeCountOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeCount) {
		return nil, false
	}

	return o.IncludeCount, true
}

// HasIncludeCount returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasIncludeCount() bool {
	if o != nil && !IsNil(o.IncludeCount) {
		return true
	}

	return false
}

// SetIncludeCount gets a reference to the given bool and assigns it to the IncludeCount field.
func (o *CloudBackupSnapshotModel) SetIncludeCount(v bool) {
	o.IncludeCount = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetInstanceName() string {
	if o == nil || IsNil(o.InstanceName) {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceName) {
		return nil, false
	}

	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasInstanceName() bool {
	if o != nil && !IsNil(o.InstanceName) {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *CloudBackupSnapshotModel) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetItemsPerPage() int {
	if o == nil || IsNil(o.ItemsPerPage) {
		var ret int
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetItemsPerPageOk() (*int, bool) {
	if o == nil || IsNil(o.ItemsPerPage) {
		return nil, false
	}

	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasItemsPerPage() bool {
	if o != nil && !IsNil(o.ItemsPerPage) {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int and assigns it to the ItemsPerPage field.
func (o *CloudBackupSnapshotModel) SetItemsPerPage(v int) {
	o.ItemsPerPage = &v
}

// GetMasterKeyUUID returns the MasterKeyUUID field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetMasterKeyUUID() string {
	if o == nil || IsNil(o.MasterKeyUUID) {
		var ret string
		return ret
	}
	return *o.MasterKeyUUID
}

// GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetMasterKeyUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.MasterKeyUUID) {
		return nil, false
	}

	return o.MasterKeyUUID, true
}

// HasMasterKeyUUID returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasMasterKeyUUID() bool {
	if o != nil && !IsNil(o.MasterKeyUUID) {
		return true
	}

	return false
}

// SetMasterKeyUUID gets a reference to the given string and assigns it to the MasterKeyUUID field.
func (o *CloudBackupSnapshotModel) SetMasterKeyUUID(v string) {
	o.MasterKeyUUID = &v
}

// GetMembers returns the Members field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetMembers() []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember {
	if o == nil || IsNil(o.Members) {
		var ret []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetMembersOk() ([]CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}

	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember and assigns it to the Members field.
func (o *CloudBackupSnapshotModel) SetMembers(v []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshotMember) {
	o.Members = v
}

// GetMongodVersion returns the MongodVersion field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetMongodVersion() string {
	if o == nil || IsNil(o.MongodVersion) {
		var ret string
		return ret
	}
	return *o.MongodVersion
}

// GetMongodVersionOk returns a tuple with the MongodVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetMongodVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongodVersion) {
		return nil, false
	}

	return o.MongodVersion, true
}

// HasMongodVersion returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasMongodVersion() bool {
	if o != nil && !IsNil(o.MongodVersion) {
		return true
	}

	return false
}

// SetMongodVersion gets a reference to the given string and assigns it to the MongodVersion field.
func (o *CloudBackupSnapshotModel) SetMongodVersion(v string) {
	o.MongodVersion = &v
}

// GetPageNum returns the PageNum field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetPageNum() int {
	if o == nil || IsNil(o.PageNum) {
		var ret int
		return ret
	}
	return *o.PageNum
}

// GetPageNumOk returns a tuple with the PageNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetPageNumOk() (*int, bool) {
	if o == nil || IsNil(o.PageNum) {
		return nil, false
	}

	return o.PageNum, true
}

// HasPageNum returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasPageNum() bool {
	if o != nil && !IsNil(o.PageNum) {
		return true
	}

	return false
}

// SetPageNum gets a reference to the given int and assigns it to the PageNum field.
func (o *CloudBackupSnapshotModel) SetPageNum(v int) {
	o.PageNum = &v
}

// GetPolicyItems returns the PolicyItems field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetPolicyItems() []string {
	if o == nil || IsNil(o.PolicyItems) {
		var ret []string
		return ret
	}
	return o.PolicyItems
}

// GetPolicyItemsOk returns a tuple with the PolicyItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetPolicyItemsOk() ([]string, bool) {
	if o == nil || IsNil(o.PolicyItems) {
		return nil, false
	}

	return o.PolicyItems, true
}

// HasPolicyItems returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasPolicyItems() bool {
	if o != nil && !IsNil(o.PolicyItems) {
		return true
	}

	return false
}

// SetPolicyItems gets a reference to the given []string and assigns it to the PolicyItems field.
func (o *CloudBackupSnapshotModel) SetPolicyItems(v []string) {
	o.PolicyItems = v
}

// GetProfile returns the Profile field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}

	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *CloudBackupSnapshotModel) SetProfile(v string) {
	o.Profile = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}

	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *CloudBackupSnapshotModel) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetReplicaSetName returns the ReplicaSetName field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetReplicaSetName() string {
	if o == nil || IsNil(o.ReplicaSetName) {
		var ret string
		return ret
	}
	return *o.ReplicaSetName
}

// GetReplicaSetNameOk returns a tuple with the ReplicaSetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetReplicaSetNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicaSetName) {
		return nil, false
	}

	return o.ReplicaSetName, true
}

// HasReplicaSetName returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasReplicaSetName() bool {
	if o != nil && !IsNil(o.ReplicaSetName) {
		return true
	}

	return false
}

// SetReplicaSetName gets a reference to the given string and assigns it to the ReplicaSetName field.
func (o *CloudBackupSnapshotModel) SetReplicaSetName(v string) {
	o.ReplicaSetName = &v
}

// GetResults returns the Results field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetResults() []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot {
	if o == nil || IsNil(o.Results) {
		var ret []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetResultsOk() ([]CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}

	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot and assigns it to the Results field.
func (o *CloudBackupSnapshotModel) SetResults(v []CloudBackupSnapshotApiAtlasDiskBackupShardedClusterSnapshot) {
	o.Results = v
}

// GetRetentionInDays returns the RetentionInDays field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetRetentionInDays() int {
	if o == nil || IsNil(o.RetentionInDays) {
		var ret int
		return ret
	}
	return *o.RetentionInDays
}

// GetRetentionInDaysOk returns a tuple with the RetentionInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetRetentionInDaysOk() (*int, bool) {
	if o == nil || IsNil(o.RetentionInDays) {
		return nil, false
	}

	return o.RetentionInDays, true
}

// HasRetentionInDays returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasRetentionInDays() bool {
	if o != nil && !IsNil(o.RetentionInDays) {
		return true
	}

	return false
}

// SetRetentionInDays gets a reference to the given int and assigns it to the RetentionInDays field.
func (o *CloudBackupSnapshotModel) SetRetentionInDays(v int) {
	o.RetentionInDays = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}

	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *CloudBackupSnapshotModel) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetSnapshotIds returns the SnapshotIds field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetSnapshotIds() []string {
	if o == nil || IsNil(o.SnapshotIds) {
		var ret []string
		return ret
	}
	return o.SnapshotIds
}

// GetSnapshotIdsOk returns a tuple with the SnapshotIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetSnapshotIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SnapshotIds) {
		return nil, false
	}

	return o.SnapshotIds, true
}

// HasSnapshotIds returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasSnapshotIds() bool {
	if o != nil && !IsNil(o.SnapshotIds) {
		return true
	}

	return false
}

// SetSnapshotIds gets a reference to the given []string and assigns it to the SnapshotIds field.
func (o *CloudBackupSnapshotModel) SetSnapshotIds(v []string) {
	o.SnapshotIds = v
}

// GetSnapshotType returns the SnapshotType field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetSnapshotType() string {
	if o == nil || IsNil(o.SnapshotType) {
		var ret string
		return ret
	}
	return *o.SnapshotType
}

// GetSnapshotTypeOk returns a tuple with the SnapshotType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetSnapshotTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotType) {
		return nil, false
	}

	return o.SnapshotType, true
}

// HasSnapshotType returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasSnapshotType() bool {
	if o != nil && !IsNil(o.SnapshotType) {
		return true
	}

	return false
}

// SetSnapshotType gets a reference to the given string and assigns it to the SnapshotType field.
func (o *CloudBackupSnapshotModel) SetSnapshotType(v string) {
	o.SnapshotType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CloudBackupSnapshotModel) SetStatus(v string) {
	o.Status = &v
}

// GetStorageSizeBytes returns the StorageSizeBytes field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetStorageSizeBytes() string {
	if o == nil || IsNil(o.StorageSizeBytes) {
		var ret string
		return ret
	}
	return *o.StorageSizeBytes
}

// GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetStorageSizeBytesOk() (*string, bool) {
	if o == nil || IsNil(o.StorageSizeBytes) {
		return nil, false
	}

	return o.StorageSizeBytes, true
}

// HasStorageSizeBytes returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasStorageSizeBytes() bool {
	if o != nil && !IsNil(o.StorageSizeBytes) {
		return true
	}

	return false
}

// SetStorageSizeBytes gets a reference to the given string and assigns it to the StorageSizeBytes field.
func (o *CloudBackupSnapshotModel) SetStorageSizeBytes(v string) {
	o.StorageSizeBytes = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetTotalCount() float32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret float32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetTotalCountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}

	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given float32 and assigns it to the TotalCount field.
func (o *CloudBackupSnapshotModel) SetTotalCount(v float32) {
	o.TotalCount = &v
}

// GetType returns the Type field value if set, zero value otherwise
func (o *CloudBackupSnapshotModel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupSnapshotModel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudBackupSnapshotModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CloudBackupSnapshotModel) SetType(v string) {
	o.Type = &v
}

func (o CloudBackupSnapshotModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupSnapshotModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloudProvider"] = o.CloudProvider
	}
	if !IsNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
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
	if !IsNil(o.IncludeCount) {
		toSerialize["includeCount"] = o.IncludeCount
	}
	if !IsNil(o.InstanceName) {
		toSerialize["instanceName"] = o.InstanceName
	}
	if !IsNil(o.ItemsPerPage) {
		toSerialize["itemsPerPage"] = o.ItemsPerPage
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
	if !IsNil(o.PageNum) {
		toSerialize["pageNum"] = o.PageNum
	}
	if !IsNil(o.PolicyItems) {
		toSerialize["policyItems"] = o.PolicyItems
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.ReplicaSetName) {
		toSerialize["replicaSetName"] = o.ReplicaSetName
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.RetentionInDays) {
		toSerialize["retentionInDays"] = o.RetentionInDays
	}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
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
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}
