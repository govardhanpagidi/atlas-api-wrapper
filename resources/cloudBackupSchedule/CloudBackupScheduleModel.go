package cloudBackupSchedule

import (
	"encoding/json"
	"fmt"
)

// Model represents the input for updating a cluster.
// swagger:parameters Model
type Model struct {
	// The ID of the project.
	//
	// required: true
	ProjectId *string `json:"-"`

	// The name of the cluster.
	//
	// required: true
	ClusterName *string `json:"-"`

	// The ID of the backup schedule.
	//
	// required: false
	Id *string `json:"Id,omitempty"`

	// Whether auto-export is enabled.
	//
	// required: false
	AutoExportEnabled *bool `json:"autoExportEnabled,omitempty"`

	// Whether to use org and group names in the export prefix.
	//
	// required: false
	UseOrgAndGroupNamesInExportPrefix *bool `json:"useOrgAndGroupNamesInExportPrefix,omitempty"`

	// The export settings.
	//
	// required: false
	Export *Export `json:"export,omitempty"`

	// The copy settings.
	//
	// required: false
	CopySettings []ApiAtlasDiskBackupCopySettingView `json:"copySettings,omitempty"`

	// The delete copied backups settings.
	//
	// required: false
	DeleteCopiedBackups []ApiDeleteCopiedBackupsView `json:"deleteCopiedBackups,omitempty"`

	// The policies.
	//
	// required: false
	Policies []ApiPolicyView `json:"policies,omitempty"`

	// The reference hour of day.
	//
	// required: false
	ReferenceHourOfDay *int `json:"referenceHourOfDay,omitempty"`

	// The reference minute of hour.
	//
	// required: false
	ReferenceMinuteOfHour *int `json:"referenceMinuteOfHour,omitempty"`

	// The number of days in the restore window.
	//
	// required: false
	RestoreWindowDays *int `json:"restoreWindowDays,omitempty"`

	// Whether to update snapshots.
	//
	// required: false
	UpdateSnapshots *bool `json:"updateSnapshots,omitempty"`

	// The ID of the cluster.
	//
	// required: false
	ClusterId *string `json:"clusterId,omitempty"`

	// The next snapshot.
	//
	// required: false
	NextSnapshot *string `json:"nextSnapshot,omitempty"`

	// The profile.
	//
	// required: false
	Profile *string `json:"profile,omitempty"`

	// The links.
	//
	// required: false
	Links []Link `json:"links,omitempty"`

	// The public key to use for authentication.
	//
	// required: false
	PublicKey *string `json:"-"`

	// The private key to use for authentication.
	//
	// required: false
	PrivateKey *string `json:"-"`
}

// Export is autogenerated from the json schema
type Export struct {
	// The ID of the export bucket.
	//
	// required: false
	ExportBucketId *string `json:",omitempty"`

	// The frequency type of the export.
	//
	// required: false
	FrequencyType *string `json:",omitempty"`
}

// ApiAtlasDiskBackupCopySettingView is autogenerated from the json schema
type ApiAtlasDiskBackupCopySettingView struct {
	// The cloud provider where the backup is copied.
	//
	// required: false
	CloudProvider *string `json:",omitempty"`

	// The name of the region where the backup is copied.
	//
	// required: false
	RegionName *string `json:",omitempty"`

	// The ID of the replication specification.
	//
	// required: false
	ReplicationSpecId *string `json:",omitempty"`

	// Whether to copy oplogs.
	//
	// required: false
	ShouldCopyOplogs *bool `json:",omitempty"`

	// The frequencies of the backup copy.
	//
	// required: false
	Frequencies []string `json:",omitempty"`
}

// ApiDeleteCopiedBackupsView is autogenerated from the json schema
type ApiDeleteCopiedBackupsView struct {
	// The cloud provider where the backups are copied.
	//
	// required: false
	CloudProvider *string `json:",omitempty"`

	// The name of the region where the backups are copied.
	//
	// required: false
	RegionName *string `json:",omitempty"`

	// The ID of the replication specification.
	//
	// required: false
	ReplicationSpecId *string `json:",omitempty"`
}

// ApiPolicyView is autogenerated from the json schema
type ApiPolicyView struct {
	// The ID of the policy.
	//
	// required: false
	ID *string `json:",omitempty"`

	// The policy items of the policy.
	//
	// required: false
	PolicyItems []ApiPolicyItemView `json:",omitempty"`
}

// ApiPolicyItemView is autogenerated from the json schema
type ApiPolicyItemView struct {
	// The ID of the policy item.
	//
	// required: false
	ID *string `json:",omitempty"`

	// The frequency type of the policy item.
	//
	// required: false
	FrequencyType *string `json:",omitempty"`

	// The frequency interval of the policy item.
	//
	// required: false
	FrequencyInterval *int `json:",omitempty"`

	// The retention value of the policy item.
	//
	// required: false
	RetentionValue *int `json:",omitempty"`

	// The retention unit of the policy item.
	//
	// required: false
	RetentionUnit *string `json:",omitempty"`
}

// Link is autogenerated from the json schema
type Link struct {
	// The URL of the link.
	//
	// required: false
	Href *string `json:",omitempty"`

	// The relationship of the link to the current resource.
	//
	// required: false
	Rel *string `json:",omitempty"`
}

func (m Model) ToString() string {
	// Create a copy of the Model struct without publicKey and privateKey.
	modelWithoutKeys := Model{
		ProjectId:                         m.ProjectId,
		ClusterName:                       m.ClusterName,
		Id:                                m.Id,
		AutoExportEnabled:                 m.AutoExportEnabled,
		UseOrgAndGroupNamesInExportPrefix: m.UseOrgAndGroupNamesInExportPrefix,
		Export:                            m.Export,
		CopySettings:                      m.CopySettings,
		DeleteCopiedBackups:               m.DeleteCopiedBackups,
		Policies:                          m.Policies,
		ReferenceHourOfDay:                m.ReferenceHourOfDay,
		ReferenceMinuteOfHour:             m.ReferenceMinuteOfHour,
		RestoreWindowDays:                 m.RestoreWindowDays,
		UpdateSnapshots:                   m.UpdateSnapshots,
		ClusterId:                         m.ClusterId,
		NextSnapshot:                      m.NextSnapshot,
		Profile:                           m.Profile,
		Links:                             m.Links,
	}

	// Convert the copy of the struct to a JSON string.
	jsonString, err := json.Marshal(modelWithoutKeys)
	if err != nil {
		return fmt.Sprintf("Error marshaling struct to JSON: %v", err)
	}

	// Return the JSON string as the string representation of the struct.
	return string(jsonString)
}