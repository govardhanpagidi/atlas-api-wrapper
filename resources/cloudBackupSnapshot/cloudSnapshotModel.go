package cloudBackupSnapshot

import "fmt"

// Model is autogenerated from the json schema
type Model struct {
	Profile          *string                                              `json:",omitempty"`
	CloudProvider    *string                                              `json:",omitempty"`
	ClusterName      *string                                              `json:",omitempty"`
	InstanceName     *string                                              `json:",omitempty"`
	CreatedAt        *string                                              `json:",omitempty"`
	Description      *string                                              `json:",omitempty"`
	ExpiresAt        *string                                              `json:",omitempty"`
	FrequencyType    *string                                              `json:",omitempty"`
	ProjectId        *string                                              `json:",omitempty"`
	Id               *string                                              `json:",omitempty"`
	IncludeCount     *bool                                                `json:",omitempty"`
	ItemsPerPage     *int                                                 `json:",omitempty"`
	MasterKeyUUID    *string                                              `json:",omitempty"`
	Members          []ApiAtlasDiskBackupShardedClusterSnapshotMemberView `json:",omitempty"`
	MongodVersion    *string                                              `json:",omitempty"`
	PageNum          *int                                                 `json:",omitempty"`
	PolicyItems      []string                                             `json:",omitempty"`
	ReplicaSetName   *string                                              `json:",omitempty"`
	Results          []ApiAtlasDiskBackupShardedClusterSnapshotView       `json:",omitempty"`
	RetentionInDays  *int                                                 `json:",omitempty"`
	SnapshotId       *string                                              `json:",omitempty"`
	SnapshotIds      []string                                             `json:",omitempty"`
	SnapshotType     *string                                              `json:",omitempty"`
	Status           *string                                              `json:",omitempty"`
	StorageSizeBytes *string                                              `json:",omitempty"`
	TotalCount       *float64                                             `json:",omitempty"`
	Type             *string                                              `json:",omitempty"`
}

// ApiAtlasDiskBackupShardedClusterSnapshotMemberView is autogenerated from the json schema
type ApiAtlasDiskBackupShardedClusterSnapshotMemberView struct {
	CloudProvider  *string `json:",omitempty"`
	Id             *string `json:",omitempty"`
	ReplicaSetName *string `json:",omitempty"`
}

// ApiAtlasDiskBackupShardedClusterSnapshotView is autogenerated from the json schema
type ApiAtlasDiskBackupShardedClusterSnapshotView struct {
	CreatedAt        *string                                              `json:",omitempty"`
	Description      *string                                              `json:",omitempty"`
	ExpiresAt        *string                                              `json:",omitempty"`
	FrequencyType    *string                                              `json:",omitempty"`
	Id               *string                                              `json:",omitempty"`
	MasterKeyUUID    *string                                              `json:",omitempty"`
	Members          []ApiAtlasDiskBackupShardedClusterSnapshotMemberView `json:",omitempty"`
	MongodVersion    *string                                              `json:",omitempty"`
	PolicyItems      []string                                             `json:",omitempty"`
	SnapshotIds      []string                                             `json:",omitempty"`
	SnapshotType     *string                                              `json:",omitempty"`
	Status           *string                                              `json:",omitempty"`
	StorageSizeBytes *string                                              `json:",omitempty"`
	Type             *string                                              `json:",omitempty"`
}
type InputModel struct {
	// The name of the cluster to create a snapshot for.
	//
	// required: true
	ClusterName *string `json:"-"`

	// The description of the snapshot.
	//
	// required: false
	Description string `json:"description"`

	// The public key to use for authentication.
	//
	// required: false
	PublicKey *string `json:"-"`

	// The private key to use for authentication.
	//
	// required: false
	PrivateKey *string `json:"-"`

	// The ID of the project where the cluster is located.
	//
	// required: true
	ProjectId *string `json:"-"`

	// The ID of the snapshot to create.
	//
	// required: true
	SnapshotId *string `json:"-"`

	// The number of days to retain the snapshot.
	//
	// required: false
	RetentionInDays *string `json:"retentionInDays"`
}

func (im InputModel) ToString() string {

	clusterName := ""
	description := ""
	projectId := ""
	snapshotId := ""
	retentionInDays := ""
	if im.ClusterName != nil {
		clusterName = *im.ClusterName
	}
	if im.Description != "" {
		description = im.Description
	}
	if im.ProjectId != nil {
		projectId = *im.ProjectId
	}
	if im.SnapshotId != nil {
		snapshotId = *im.SnapshotId
	}
	if im.RetentionInDays != nil {
		retentionInDays = *im.RetentionInDays
	}

	return fmt.Sprintf("ClusterName: %s, Description: %s, ProjectId: %v, SnapshotId: %v, RetentionInDays: %s",
		clusterName, description, projectId, snapshotId, retentionInDays)
}