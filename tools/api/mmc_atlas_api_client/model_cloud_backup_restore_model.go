// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupRestoreModel struct for CloudBackupRestoreModel
type CloudBackupRestoreModel struct {
	Cancelled                  *bool                                         `json:"cancelled,omitempty"`
	ClusterName                *string                                       `json:"clusterName,omitempty"`
	CreatedAt                  *string                                       `json:"createdAt,omitempty"`
	DeliveryType               *string                                       `json:"deliveryType,omitempty"`
	DeliveryUrl                []string                                      `json:"deliveryUrl,omitempty"`
	EnableSynchronousCreation  *bool                                         `json:"enableSynchronousCreation,omitempty"`
	Expired                    *bool                                         `json:"expired,omitempty"`
	ExpiresAt                  *string                                       `json:"expiresAt,omitempty"`
	FinishedAt                 *string                                       `json:"finishedAt,omitempty"`
	Id                         *string                                       `json:"id,omitempty"`
	InstanceName               *string                                       `json:"instanceName,omitempty"`
	Links                      []CloudBackupRestoreLinks                     `json:"links,omitempty"`
	OpLogInc                   *string                                       `json:"opLogInc,omitempty"`
	OpLogTs                    *string                                       `json:"opLogTs,omitempty"`
	PointInTimeUtcSeconds      *int                                          `json:"pointInTimeUtcSeconds,omitempty"`
	Profile                    *string                                       `json:"profile,omitempty"`
	ProjectId                  *string                                       `json:"projectId,omitempty"`
	SnapshotId                 *string                                       `json:"snapshotId,omitempty"`
	SynchronousCreationOptions *CloudBackupRestoreSynchronousCreationOptions `json:"synchronousCreationOptions,omitempty"`
	TargetClusterName          *string                                       `json:"targetClusterName,omitempty"`
	TargetProjectId            *string                                       `json:"targetProjectId,omitempty"`
	Timestamp                  *string                                       `json:"timestamp,omitempty"`
}

// NewCloudBackupRestoreModel instantiates a new CloudBackupRestoreModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupRestoreModel() *CloudBackupRestoreModel {
	this := CloudBackupRestoreModel{}
	return &this
}

// NewCloudBackupRestoreModelWithDefaults instantiates a new CloudBackupRestoreModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupRestoreModelWithDefaults() *CloudBackupRestoreModel {
	this := CloudBackupRestoreModel{}
	return &this
}

// GetCancelled returns the Cancelled field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetCancelled() bool {
	if o == nil || IsNil(o.Cancelled) {
		var ret bool
		return ret
	}
	return *o.Cancelled
}

// GetCancelledOk returns a tuple with the Cancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.Cancelled) {
		return nil, false
	}

	return o.Cancelled, true
}

// HasCancelled returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasCancelled() bool {
	if o != nil && !IsNil(o.Cancelled) {
		return true
	}

	return false
}

// SetCancelled gets a reference to the given bool and assigns it to the Cancelled field.
func (o *CloudBackupRestoreModel) SetCancelled(v bool) {
	o.Cancelled = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *CloudBackupRestoreModel) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CloudBackupRestoreModel) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeliveryType returns the DeliveryType field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetDeliveryType() string {
	if o == nil || IsNil(o.DeliveryType) {
		var ret string
		return ret
	}
	return *o.DeliveryType
}

// GetDeliveryTypeOk returns a tuple with the DeliveryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetDeliveryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryType) {
		return nil, false
	}

	return o.DeliveryType, true
}

// HasDeliveryType returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasDeliveryType() bool {
	if o != nil && !IsNil(o.DeliveryType) {
		return true
	}

	return false
}

// SetDeliveryType gets a reference to the given string and assigns it to the DeliveryType field.
func (o *CloudBackupRestoreModel) SetDeliveryType(v string) {
	o.DeliveryType = &v
}

// GetDeliveryUrl returns the DeliveryUrl field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetDeliveryUrl() []string {
	if o == nil || IsNil(o.DeliveryUrl) {
		var ret []string
		return ret
	}
	return o.DeliveryUrl
}

// GetDeliveryUrlOk returns a tuple with the DeliveryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetDeliveryUrlOk() ([]string, bool) {
	if o == nil || IsNil(o.DeliveryUrl) {
		return nil, false
	}

	return o.DeliveryUrl, true
}

// HasDeliveryUrl returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasDeliveryUrl() bool {
	if o != nil && !IsNil(o.DeliveryUrl) {
		return true
	}

	return false
}

// SetDeliveryUrl gets a reference to the given []string and assigns it to the DeliveryUrl field.
func (o *CloudBackupRestoreModel) SetDeliveryUrl(v []string) {
	o.DeliveryUrl = v
}

// GetEnableSynchronousCreation returns the EnableSynchronousCreation field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetEnableSynchronousCreation() bool {
	if o == nil || IsNil(o.EnableSynchronousCreation) {
		var ret bool
		return ret
	}
	return *o.EnableSynchronousCreation
}

// GetEnableSynchronousCreationOk returns a tuple with the EnableSynchronousCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetEnableSynchronousCreationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSynchronousCreation) {
		return nil, false
	}

	return o.EnableSynchronousCreation, true
}

// HasEnableSynchronousCreation returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasEnableSynchronousCreation() bool {
	if o != nil && !IsNil(o.EnableSynchronousCreation) {
		return true
	}

	return false
}

// SetEnableSynchronousCreation gets a reference to the given bool and assigns it to the EnableSynchronousCreation field.
func (o *CloudBackupRestoreModel) SetEnableSynchronousCreation(v bool) {
	o.EnableSynchronousCreation = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetExpired() bool {
	if o == nil || IsNil(o.Expired) {
		var ret bool
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetExpiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Expired) {
		return nil, false
	}

	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasExpired() bool {
	if o != nil && !IsNil(o.Expired) {
		return true
	}

	return false
}

// SetExpired gets a reference to the given bool and assigns it to the Expired field.
func (o *CloudBackupRestoreModel) SetExpired(v bool) {
	o.Expired = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}

	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *CloudBackupRestoreModel) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetFinishedAt() string {
	if o == nil || IsNil(o.FinishedAt) {
		var ret string
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetFinishedAtOk() (*string, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}

	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given string and assigns it to the FinishedAt field.
func (o *CloudBackupRestoreModel) SetFinishedAt(v string) {
	o.FinishedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudBackupRestoreModel) SetId(v string) {
	o.Id = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetInstanceName() string {
	if o == nil || IsNil(o.InstanceName) {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceName) {
		return nil, false
	}

	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasInstanceName() bool {
	if o != nil && !IsNil(o.InstanceName) {
		return true
	}

	return false
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *CloudBackupRestoreModel) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetLinks() []CloudBackupRestoreLinks {
	if o == nil || IsNil(o.Links) {
		var ret []CloudBackupRestoreLinks
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetLinksOk() ([]CloudBackupRestoreLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []CloudBackupRestoreLinks and assigns it to the Links field.
func (o *CloudBackupRestoreModel) SetLinks(v []CloudBackupRestoreLinks) {
	o.Links = v
}

// GetOpLogInc returns the OpLogInc field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetOpLogInc() string {
	if o == nil || IsNil(o.OpLogInc) {
		var ret string
		return ret
	}
	return *o.OpLogInc
}

// GetOpLogIncOk returns a tuple with the OpLogInc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetOpLogIncOk() (*string, bool) {
	if o == nil || IsNil(o.OpLogInc) {
		return nil, false
	}

	return o.OpLogInc, true
}

// HasOpLogInc returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasOpLogInc() bool {
	if o != nil && !IsNil(o.OpLogInc) {
		return true
	}

	return false
}

// SetOpLogInc gets a reference to the given string and assigns it to the OpLogInc field.
func (o *CloudBackupRestoreModel) SetOpLogInc(v string) {
	o.OpLogInc = &v
}

// GetOpLogTs returns the OpLogTs field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetOpLogTs() string {
	if o == nil || IsNil(o.OpLogTs) {
		var ret string
		return ret
	}
	return *o.OpLogTs
}

// GetOpLogTsOk returns a tuple with the OpLogTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetOpLogTsOk() (*string, bool) {
	if o == nil || IsNil(o.OpLogTs) {
		return nil, false
	}

	return o.OpLogTs, true
}

// HasOpLogTs returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasOpLogTs() bool {
	if o != nil && !IsNil(o.OpLogTs) {
		return true
	}

	return false
}

// SetOpLogTs gets a reference to the given string and assigns it to the OpLogTs field.
func (o *CloudBackupRestoreModel) SetOpLogTs(v string) {
	o.OpLogTs = &v
}

// GetPointInTimeUtcSeconds returns the PointInTimeUtcSeconds field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetPointInTimeUtcSeconds() int {
	if o == nil || IsNil(o.PointInTimeUtcSeconds) {
		var ret int
		return ret
	}
	return *o.PointInTimeUtcSeconds
}

// GetPointInTimeUtcSecondsOk returns a tuple with the PointInTimeUtcSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetPointInTimeUtcSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.PointInTimeUtcSeconds) {
		return nil, false
	}

	return o.PointInTimeUtcSeconds, true
}

// HasPointInTimeUtcSeconds returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasPointInTimeUtcSeconds() bool {
	if o != nil && !IsNil(o.PointInTimeUtcSeconds) {
		return true
	}

	return false
}

// SetPointInTimeUtcSeconds gets a reference to the given int and assigns it to the PointInTimeUtcSeconds field.
func (o *CloudBackupRestoreModel) SetPointInTimeUtcSeconds(v int) {
	o.PointInTimeUtcSeconds = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}

	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *CloudBackupRestoreModel) SetProfile(v string) {
	o.Profile = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}

	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *CloudBackupRestoreModel) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}

	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *CloudBackupRestoreModel) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetSynchronousCreationOptions returns the SynchronousCreationOptions field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetSynchronousCreationOptions() CloudBackupRestoreSynchronousCreationOptions {
	if o == nil || IsNil(o.SynchronousCreationOptions) {
		var ret CloudBackupRestoreSynchronousCreationOptions
		return ret
	}
	return *o.SynchronousCreationOptions
}

// GetSynchronousCreationOptionsOk returns a tuple with the SynchronousCreationOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetSynchronousCreationOptionsOk() (*CloudBackupRestoreSynchronousCreationOptions, bool) {
	if o == nil || IsNil(o.SynchronousCreationOptions) {
		return nil, false
	}

	return o.SynchronousCreationOptions, true
}

// HasSynchronousCreationOptions returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasSynchronousCreationOptions() bool {
	if o != nil && !IsNil(o.SynchronousCreationOptions) {
		return true
	}

	return false
}

// SetSynchronousCreationOptions gets a reference to the given CloudBackupRestoreSynchronousCreationOptions and assigns it to the SynchronousCreationOptions field.
func (o *CloudBackupRestoreModel) SetSynchronousCreationOptions(v CloudBackupRestoreSynchronousCreationOptions) {
	o.SynchronousCreationOptions = &v
}

// GetTargetClusterName returns the TargetClusterName field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetTargetClusterName() string {
	if o == nil || IsNil(o.TargetClusterName) {
		var ret string
		return ret
	}
	return *o.TargetClusterName
}

// GetTargetClusterNameOk returns a tuple with the TargetClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetTargetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetClusterName) {
		return nil, false
	}

	return o.TargetClusterName, true
}

// HasTargetClusterName returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasTargetClusterName() bool {
	if o != nil && !IsNil(o.TargetClusterName) {
		return true
	}

	return false
}

// SetTargetClusterName gets a reference to the given string and assigns it to the TargetClusterName field.
func (o *CloudBackupRestoreModel) SetTargetClusterName(v string) {
	o.TargetClusterName = &v
}

// GetTargetProjectId returns the TargetProjectId field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetTargetProjectId() string {
	if o == nil || IsNil(o.TargetProjectId) {
		var ret string
		return ret
	}
	return *o.TargetProjectId
}

// GetTargetProjectIdOk returns a tuple with the TargetProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetTargetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetProjectId) {
		return nil, false
	}

	return o.TargetProjectId, true
}

// HasTargetProjectId returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasTargetProjectId() bool {
	if o != nil && !IsNil(o.TargetProjectId) {
		return true
	}

	return false
}

// SetTargetProjectId gets a reference to the given string and assigns it to the TargetProjectId field.
func (o *CloudBackupRestoreModel) SetTargetProjectId(v string) {
	o.TargetProjectId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise
func (o *CloudBackupRestoreModel) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupRestoreModel) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}

	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CloudBackupRestoreModel) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *CloudBackupRestoreModel) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o CloudBackupRestoreModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupRestoreModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cancelled) {
		toSerialize["cancelled"] = o.Cancelled
	}
	if !IsNil(o.ClusterName) {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.DeliveryType) {
		toSerialize["deliveryType"] = o.DeliveryType
	}
	if !IsNil(o.DeliveryUrl) {
		toSerialize["deliveryUrl"] = o.DeliveryUrl
	}
	if !IsNil(o.EnableSynchronousCreation) {
		toSerialize["enableSynchronousCreation"] = o.EnableSynchronousCreation
	}
	if !IsNil(o.Expired) {
		toSerialize["expired"] = o.Expired
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.FinishedAt) {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InstanceName) {
		toSerialize["instanceName"] = o.InstanceName
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.OpLogInc) {
		toSerialize["opLogInc"] = o.OpLogInc
	}
	if !IsNil(o.OpLogTs) {
		toSerialize["opLogTs"] = o.OpLogTs
	}
	if !IsNil(o.PointInTimeUtcSeconds) {
		toSerialize["pointInTimeUtcSeconds"] = o.PointInTimeUtcSeconds
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if !IsNil(o.SynchronousCreationOptions) {
		toSerialize["synchronousCreationOptions"] = o.SynchronousCreationOptions
	}
	if !IsNil(o.TargetClusterName) {
		toSerialize["targetClusterName"] = o.TargetClusterName
	}
	if !IsNil(o.TargetProjectId) {
		toSerialize["targetProjectId"] = o.TargetProjectId
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}
