// Code based on the AtlasAPI V2 OpenAPI file

package mmc_atlas_api_client

import (
	"encoding/json"
)

// CloudBackupScheduleModel struct for CloudBackupScheduleModel
type CloudBackupScheduleModel struct {
	Id                                *string                                                `json:"Id,omitempty"`
	AutoExportEnabled                 *bool                                                  `json:"autoExportEnabled,omitempty"`
	ClusterId                         *string                                                `json:"clusterId,omitempty"`
	CopySettings                      []CloudBackupScheduleApiAtlasDiskBackupCopySettingView `json:"copySettings,omitempty"`
	DeleteCopiedBackups               []CloudBackupScheduleApiDeleteCopiedBackupsView        `json:"deleteCopiedBackups,omitempty"`
	Export                            *CloudBackupScheduleExport                             `json:"export,omitempty"`
	Links                             []CloudBackupScheduleLink                              `json:"links,omitempty"`
	NextSnapshot                      *string                                                `json:"nextSnapshot,omitempty"`
	Policies                          []CloudBackupScheduleApiPolicyView                     `json:"policies,omitempty"`
	Profile                           *string                                                `json:"profile,omitempty"`
	ReferenceHourOfDay                *int                                                   `json:"referenceHourOfDay,omitempty"`
	ReferenceMinuteOfHour             *int                                                   `json:"referenceMinuteOfHour,omitempty"`
	RestoreWindowDays                 *int                                                   `json:"restoreWindowDays,omitempty"`
	UpdateSnapshots                   *bool                                                  `json:"updateSnapshots,omitempty"`
	UseOrgAndGroupNamesInExportPrefix *bool                                                  `json:"useOrgAndGroupNamesInExportPrefix,omitempty"`
}

// NewCloudBackupScheduleModel instantiates a new CloudBackupScheduleModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudBackupScheduleModel() *CloudBackupScheduleModel {
	this := CloudBackupScheduleModel{}
	return &this
}

// NewCloudBackupScheduleModelWithDefaults instantiates a new CloudBackupScheduleModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudBackupScheduleModelWithDefaults() *CloudBackupScheduleModel {
	this := CloudBackupScheduleModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudBackupScheduleModel) SetId(v string) {
	o.Id = &v
}

// GetAutoExportEnabled returns the AutoExportEnabled field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetAutoExportEnabled() bool {
	if o == nil || IsNil(o.AutoExportEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoExportEnabled
}

// GetAutoExportEnabledOk returns a tuple with the AutoExportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetAutoExportEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoExportEnabled) {
		return nil, false
	}

	return o.AutoExportEnabled, true
}

// HasAutoExportEnabled returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasAutoExportEnabled() bool {
	if o != nil && !IsNil(o.AutoExportEnabled) {
		return true
	}

	return false
}

// SetAutoExportEnabled gets a reference to the given bool and assigns it to the AutoExportEnabled field.
func (o *CloudBackupScheduleModel) SetAutoExportEnabled(v bool) {
	o.AutoExportEnabled = &v
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}

	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *CloudBackupScheduleModel) SetClusterId(v string) {
	o.ClusterId = &v
}

// GetCopySettings returns the CopySettings field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetCopySettings() []CloudBackupScheduleApiAtlasDiskBackupCopySettingView {
	if o == nil || IsNil(o.CopySettings) {
		var ret []CloudBackupScheduleApiAtlasDiskBackupCopySettingView
		return ret
	}
	return o.CopySettings
}

// GetCopySettingsOk returns a tuple with the CopySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetCopySettingsOk() ([]CloudBackupScheduleApiAtlasDiskBackupCopySettingView, bool) {
	if o == nil || IsNil(o.CopySettings) {
		return nil, false
	}

	return o.CopySettings, true
}

// HasCopySettings returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasCopySettings() bool {
	if o != nil && !IsNil(o.CopySettings) {
		return true
	}

	return false
}

// SetCopySettings gets a reference to the given []CloudBackupScheduleApiAtlasDiskBackupCopySettingView and assigns it to the CopySettings field.
func (o *CloudBackupScheduleModel) SetCopySettings(v []CloudBackupScheduleApiAtlasDiskBackupCopySettingView) {
	o.CopySettings = v
}

// GetDeleteCopiedBackups returns the DeleteCopiedBackups field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetDeleteCopiedBackups() []CloudBackupScheduleApiDeleteCopiedBackupsView {
	if o == nil || IsNil(o.DeleteCopiedBackups) {
		var ret []CloudBackupScheduleApiDeleteCopiedBackupsView
		return ret
	}
	return o.DeleteCopiedBackups
}

// GetDeleteCopiedBackupsOk returns a tuple with the DeleteCopiedBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetDeleteCopiedBackupsOk() ([]CloudBackupScheduleApiDeleteCopiedBackupsView, bool) {
	if o == nil || IsNil(o.DeleteCopiedBackups) {
		return nil, false
	}

	return o.DeleteCopiedBackups, true
}

// HasDeleteCopiedBackups returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasDeleteCopiedBackups() bool {
	if o != nil && !IsNil(o.DeleteCopiedBackups) {
		return true
	}

	return false
}

// SetDeleteCopiedBackups gets a reference to the given []CloudBackupScheduleApiDeleteCopiedBackupsView and assigns it to the DeleteCopiedBackups field.
func (o *CloudBackupScheduleModel) SetDeleteCopiedBackups(v []CloudBackupScheduleApiDeleteCopiedBackupsView) {
	o.DeleteCopiedBackups = v
}

// GetExport returns the Export field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetExport() CloudBackupScheduleExport {
	if o == nil || IsNil(o.Export) {
		var ret CloudBackupScheduleExport
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetExportOk() (*CloudBackupScheduleExport, bool) {
	if o == nil || IsNil(o.Export) {
		return nil, false
	}

	return o.Export, true
}

// HasExport returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasExport() bool {
	if o != nil && !IsNil(o.Export) {
		return true
	}

	return false
}

// SetExport gets a reference to the given CloudBackupScheduleExport and assigns it to the Export field.
func (o *CloudBackupScheduleModel) SetExport(v CloudBackupScheduleExport) {
	o.Export = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetLinks() []CloudBackupScheduleLink {
	if o == nil || IsNil(o.Links) {
		var ret []CloudBackupScheduleLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetLinksOk() ([]CloudBackupScheduleLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []CloudBackupScheduleLink and assigns it to the Links field.
func (o *CloudBackupScheduleModel) SetLinks(v []CloudBackupScheduleLink) {
	o.Links = v
}

// GetNextSnapshot returns the NextSnapshot field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetNextSnapshot() string {
	if o == nil || IsNil(o.NextSnapshot) {
		var ret string
		return ret
	}
	return *o.NextSnapshot
}

// GetNextSnapshotOk returns a tuple with the NextSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetNextSnapshotOk() (*string, bool) {
	if o == nil || IsNil(o.NextSnapshot) {
		return nil, false
	}

	return o.NextSnapshot, true
}

// HasNextSnapshot returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasNextSnapshot() bool {
	if o != nil && !IsNil(o.NextSnapshot) {
		return true
	}

	return false
}

// SetNextSnapshot gets a reference to the given string and assigns it to the NextSnapshot field.
func (o *CloudBackupScheduleModel) SetNextSnapshot(v string) {
	o.NextSnapshot = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetPolicies() []CloudBackupScheduleApiPolicyView {
	if o == nil || IsNil(o.Policies) {
		var ret []CloudBackupScheduleApiPolicyView
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetPoliciesOk() ([]CloudBackupScheduleApiPolicyView, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}

	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []CloudBackupScheduleApiPolicyView and assigns it to the Policies field.
func (o *CloudBackupScheduleModel) SetPolicies(v []CloudBackupScheduleApiPolicyView) {
	o.Policies = v
}

// GetProfile returns the Profile field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}

	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *CloudBackupScheduleModel) SetProfile(v string) {
	o.Profile = &v
}

// GetReferenceHourOfDay returns the ReferenceHourOfDay field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetReferenceHourOfDay() int {
	if o == nil || IsNil(o.ReferenceHourOfDay) {
		var ret int
		return ret
	}
	return *o.ReferenceHourOfDay
}

// GetReferenceHourOfDayOk returns a tuple with the ReferenceHourOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetReferenceHourOfDayOk() (*int, bool) {
	if o == nil || IsNil(o.ReferenceHourOfDay) {
		return nil, false
	}

	return o.ReferenceHourOfDay, true
}

// HasReferenceHourOfDay returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasReferenceHourOfDay() bool {
	if o != nil && !IsNil(o.ReferenceHourOfDay) {
		return true
	}

	return false
}

// SetReferenceHourOfDay gets a reference to the given int and assigns it to the ReferenceHourOfDay field.
func (o *CloudBackupScheduleModel) SetReferenceHourOfDay(v int) {
	o.ReferenceHourOfDay = &v
}

// GetReferenceMinuteOfHour returns the ReferenceMinuteOfHour field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetReferenceMinuteOfHour() int {
	if o == nil || IsNil(o.ReferenceMinuteOfHour) {
		var ret int
		return ret
	}
	return *o.ReferenceMinuteOfHour
}

// GetReferenceMinuteOfHourOk returns a tuple with the ReferenceMinuteOfHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetReferenceMinuteOfHourOk() (*int, bool) {
	if o == nil || IsNil(o.ReferenceMinuteOfHour) {
		return nil, false
	}

	return o.ReferenceMinuteOfHour, true
}

// HasReferenceMinuteOfHour returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasReferenceMinuteOfHour() bool {
	if o != nil && !IsNil(o.ReferenceMinuteOfHour) {
		return true
	}

	return false
}

// SetReferenceMinuteOfHour gets a reference to the given int and assigns it to the ReferenceMinuteOfHour field.
func (o *CloudBackupScheduleModel) SetReferenceMinuteOfHour(v int) {
	o.ReferenceMinuteOfHour = &v
}

// GetRestoreWindowDays returns the RestoreWindowDays field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetRestoreWindowDays() int {
	if o == nil || IsNil(o.RestoreWindowDays) {
		var ret int
		return ret
	}
	return *o.RestoreWindowDays
}

// GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetRestoreWindowDaysOk() (*int, bool) {
	if o == nil || IsNil(o.RestoreWindowDays) {
		return nil, false
	}

	return o.RestoreWindowDays, true
}

// HasRestoreWindowDays returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasRestoreWindowDays() bool {
	if o != nil && !IsNil(o.RestoreWindowDays) {
		return true
	}

	return false
}

// SetRestoreWindowDays gets a reference to the given int and assigns it to the RestoreWindowDays field.
func (o *CloudBackupScheduleModel) SetRestoreWindowDays(v int) {
	o.RestoreWindowDays = &v
}

// GetUpdateSnapshots returns the UpdateSnapshots field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetUpdateSnapshots() bool {
	if o == nil || IsNil(o.UpdateSnapshots) {
		var ret bool
		return ret
	}
	return *o.UpdateSnapshots
}

// GetUpdateSnapshotsOk returns a tuple with the UpdateSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetUpdateSnapshotsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateSnapshots) {
		return nil, false
	}

	return o.UpdateSnapshots, true
}

// HasUpdateSnapshots returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasUpdateSnapshots() bool {
	if o != nil && !IsNil(o.UpdateSnapshots) {
		return true
	}

	return false
}

// SetUpdateSnapshots gets a reference to the given bool and assigns it to the UpdateSnapshots field.
func (o *CloudBackupScheduleModel) SetUpdateSnapshots(v bool) {
	o.UpdateSnapshots = &v
}

// GetUseOrgAndGroupNamesInExportPrefix returns the UseOrgAndGroupNamesInExportPrefix field value if set, zero value otherwise
func (o *CloudBackupScheduleModel) GetUseOrgAndGroupNamesInExportPrefix() bool {
	if o == nil || IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		var ret bool
		return ret
	}
	return *o.UseOrgAndGroupNamesInExportPrefix
}

// GetUseOrgAndGroupNamesInExportPrefixOk returns a tuple with the UseOrgAndGroupNamesInExportPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudBackupScheduleModel) GetUseOrgAndGroupNamesInExportPrefixOk() (*bool, bool) {
	if o == nil || IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		return nil, false
	}

	return o.UseOrgAndGroupNamesInExportPrefix, true
}

// HasUseOrgAndGroupNamesInExportPrefix returns a boolean if a field has been set.
func (o *CloudBackupScheduleModel) HasUseOrgAndGroupNamesInExportPrefix() bool {
	if o != nil && !IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		return true
	}

	return false
}

// SetUseOrgAndGroupNamesInExportPrefix gets a reference to the given bool and assigns it to the UseOrgAndGroupNamesInExportPrefix field.
func (o *CloudBackupScheduleModel) SetUseOrgAndGroupNamesInExportPrefix(v bool) {
	o.UseOrgAndGroupNamesInExportPrefix = &v
}

func (o CloudBackupScheduleModel) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudBackupScheduleModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.AutoExportEnabled) {
		toSerialize["autoExportEnabled"] = o.AutoExportEnabled
	}
	if !IsNil(o.ClusterId) {
		toSerialize["clusterId"] = o.ClusterId
	}
	if !IsNil(o.CopySettings) {
		toSerialize["copySettings"] = o.CopySettings
	}
	if !IsNil(o.DeleteCopiedBackups) {
		toSerialize["deleteCopiedBackups"] = o.DeleteCopiedBackups
	}
	if !IsNil(o.Export) {
		toSerialize["export"] = o.Export
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.NextSnapshot) {
		toSerialize["nextSnapshot"] = o.NextSnapshot
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.ReferenceHourOfDay) {
		toSerialize["referenceHourOfDay"] = o.ReferenceHourOfDay
	}
	if !IsNil(o.ReferenceMinuteOfHour) {
		toSerialize["referenceMinuteOfHour"] = o.ReferenceMinuteOfHour
	}
	if !IsNil(o.RestoreWindowDays) {
		toSerialize["restoreWindowDays"] = o.RestoreWindowDays
	}
	if !IsNil(o.UpdateSnapshots) {
		toSerialize["updateSnapshots"] = o.UpdateSnapshots
	}
	if !IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		toSerialize["useOrgAndGroupNamesInExportPrefix"] = o.UseOrgAndGroupNamesInExportPrefix
	}
	return toSerialize, nil
}
