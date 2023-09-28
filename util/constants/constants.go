//go:build !coverage

package constants

type TraceIDKeyType string

const (
	ProjectID       = "ProjectId"
	AWS             = "AWS"
	HostName        = "HostName"
	DatabaseName    = "DatabaseName"
	Username        = "Username"
	Password        = "Password"
	ClusterName     = "ClusterName"
	CollectionName  = "CollectionName"
	CollectionNames = "CollectionNames"

	DatabaseUserHandlerName = "DatabaseUserhandler"

	ClusterHandler               = "ClusterHandler"
	Collectionhandler            = "CollectionHandler"
	ClusterBackupRestoreHandler  = "ClusterBackupRestoreHandler"
	ClusterBackupSnapshotHandler = "ClusterBackupSnapshotHandler"

	PrivateKey = "PrivateKey"
	PublicKey  = "PublicKey"

	PublicKeyHeader     = "x-mongo-publickey"
	PrivateKeyHeader    = "x-mongo-privatekey"
	TshirtSize          = "TshirtSize"
	CloudProvider       = "CloudProvider"
	MongoDBVersion      = "MongoDBVersion"
	MongoDBMajorVersion = "MongoDBMajorVersion"
	RetainBackup        = "RetainBackup"
	CloudBackupHandler  = "CloudBackupHandler"

	InvalidInputParameter          = "INVALID_INPUT_PARAMETER"
	MongoClientCreationError       = "MONGO_CLIENT_CREATE_ERROR"
	ResourceDoesNotExist           = "RESOURCE_NOT_EXIST"
	ClusterModelError              = "CLUSTER_MODEL_ERROR"
	ClusterRequestError            = "CLUSTER_REQUEST_ERROR"
	ClusterCreateError             = "CLUSTER_CREATE_ERROR"
	ClusterUpdateError             = "CLUSTER_UPDATE_ERROR"
	ClusterCreateSuccess           = "CLUSTER_CREATE_SUCCESS"
	UpdateTagsError                = "UPDATE_TAGS_ERROR"
	ClusterReadSuccess             = "CLUSTER_READ_SUCCESS"
	ClusterDeleteError             = "CLUSTER_DELETE_ERROR"
	ClusterDeleteSuccess           = "CLUSTER_DELETE_SUCCESS"
	ClusterListError               = "CLUSTER_LIST_ERROR"
	ClusterAdvancedListError       = "CLUSTER_ADVANCED_LIST_ERROR"
	ClusterListSuccess             = "CLUSTER_LIST_SUCCESS"
	ListEndpointError              = "LIST_ENDPOINT_ERROR"
	NoEndpointConfigured           = "NO_ENDPOINT_ERROR"
	NoEndpointConfiguredForRegion  = "NO_ENDPOINT_REGION_ERROR"
	NoAdvancedClusterConfiguration = "NO_ADVANCED_CLUSTER_CONFIG"
	ProjectIpAccessListError       = "PROJECT_IP_ACCESS_LIST_ERROR"
	ClusterAlreadyDeleteError      = "CLUSTER_ALREADY_REQUESTED_DELETION"
	ClusterNameNotCreated          = "CLUSTER_NAME_NOT_GENERATED"
	ClusterNameNotSet              = "CLUSTER_HOSTNAME_NOT_SET"

	DatabaseListSuccess   = "DATABASE_LIST_SUCCESS"
	CollectionListError   = "COLLECTION_LIST_ERROR"
	CollectionListSuccess = "COLLECTION_LIST_SUCCESS"

	GetRestoreJobError                   = "RESTORE_JOB_READ_ERROR"
	CreateRestoreJobError                = "RESTORE_JOB_CREATE_ERROR"
	GetSnapshotListError                 = "SNAPSHOT_LIST_ERROR"
	CreateSnapshotError                  = "SNAPSHOT_CREATE_ERROR"
	InvalidTargerClusterNameAndProjectId = "INVALID_TARGET_CLUSTER_NAME_AND_PROJECT_ID"
	InvalidPointInTimeError              = "INVALID_POINT_IN_TIME_ERROR"
	ListReplicaSetBackupError            = "LIST_REPLICASET_BACKUP_ERROR"
	ListShardedClusterBackupError        = "LIST_SHARDED_CLUSTER_BACKUP_ERROR"

	GetBackupScheduleError     = "GET_BACKUP_SCHEDULE_ERROR"
	GetPolicyScheduleError     = "READ_POLICY_SCHEDULE_ERROR"
	ValidateExportDetails      = "VALIDATE_EXPORT_DETAILS_ERROR"
	DeleteBackupSchedulesError = "DELETE_BACKUP_SCHEDULE_ERROR"
	UpdateBackupScheduleError  = "UPDATE_BACKUP_SCHEDULE_ERROR"

	ClusterStatusIDLE       = "The cluster is in an idle state, which usually means it's not actively processing requests"
	ClusterStatusUnknown    = "The cluster status is unknown"
	ClusterStatusCREATING   = "The cluster is in the process of being created. It's not yet fully available"
	ClusterStatusDELETING   = "The cluster is being deleted. It's in the process of being removed"
	ClusterStatusUPDATING   = "The cluster is being updated, which can include scaling, configuration changes, or other updates"
	ClusterStatusREPAIRING  = "The cluster is undergoing repairs due to issues"
	ClusterStatusRESTARTING = "The cluster is in the process of being restarted"
	ClusterStatusPAUSED     = "The cluster is paused, which typically means that certain operations are temporarily stopped"
	ClusterStatusACTIVE     = " The cluster is active and fully operational"

	Unknown                 = "UNKNOWN"
	Idle                    = "IDLE"
	Creating                = "CREATING"
	Deleting                = "DELETING"
	Updating                = "UPDATING"
	Repairing               = "REPAIRING"
	Restarting              = "RESTARTING"
	Paused                  = "PAUSED"
	Active                  = "ACTIVE"
	CollectionError         = "COLLECTION_ERROR"
	CollectionSuccess       = "COLLECTION_SUCCESS"
	CollectionDeleteError   = "COLLECTION_DELETE_ERROR"
	CollectionDeleteSuccess = "COLLECTION_DELETE_SUCCESS"

	DatabaseError         = "DATABASE_ERROR"
	DatabaseSuccess       = "DATABASE_SUCCESS"
	DatabaseDeleteError   = "DATABASE_DELETE_ERROR"
	DatabaseDeleteSuccess = "DATABASE_DELETE_SUCCESS"

	UserCreateError           = "USER_CREATE_ERROR"
	UserCreateSuccess         = "USER_CREATE_SUCCESS"
	UserNotFound              = "USER_NOT_FOUND"
	FetchUser                 = "FETCH_USER"
	FetchuserError            = "FETCH_USER_ERROR"
	DeleteDatabaseUserError   = "DELETE_USER_ERROR"
	DeleteDatabaseUserSuccess = "DELETE_USER_SUCCESS"
	UserListError             = "USER_LIST_ERROR"
	UserListSuccess           = "USER_LIST_SUCCESS"

	SwaggerDocsPath     = "/docs/"
	SwaggerDocsJsonPath = "/docs/doc.json"
	DefaultHostMessage  = "is being created"
	ClusterHostName     = "HostName:"
	HostNameSpearator   = "//"

	DbuserDbName    = "admin"
	DbAdminRoleName = "dbAdmin"
	AtlasAdminRole  = "atlasAdmin"

	PublicIp = "0.0.0.0/0"

	ClusterConfigLocation                = "config.json"
	TraceID               TraceIDKeyType = "TraceID"
	Project                              = "Project"
	Cluster                              = "Cluster"

	MongoBaseUrl = "https://cloud.mongodb.com/"
	MongoDbURI   = "mongodb+srv://%s:%s@%s"

	MessageConfigLocation   = "message_Config.json"
	EmptyString             = ""
	DefaultCollectionString = "default"

	CreateOrGetDatabaseUserReqURI = "project/{ProjectId}/databaseUsers"
	DeleteOrGetDatabaseUserReqURI = "project/{ProjectId}/databaseUsers/{Username}"

	CreateOrGetClusterReqURI    = "project/{ProjectId}/cluster"
	DeleteOrUpdateClusterReqURI = "project/{ProjectId}/cluster/{ClusterName}"
	GetClusterReqURI            = "project/{ProjectId}/cluster/{ClusterName}/status"

	CreateDatabaseReqURI = "project/{ProjectId}/cluster/{ClusterName}/database"
	GetDatabaseReqURI    = "project/{ProjectId}/cluster/{ClusterName}/database"
	DeleteDatabaseReqURI = "project/{ProjectId}/cluster/{ClusterName}/database/{DatabaseName}"

	CreateCollectionReqURI = "project/{ProjectId}/cluster/{ClusterName}/database/{DatabaseName}/collections"
	GetCollectionReqURI    = "project/{ProjectId}/cluster/{ClusterName}/database/{DatabaseName}/collections"
	DeleteCollectionReqURI = "project/{ProjectId}/cluster/{ClusterName}/database/{DatabaseName}/collection/{CollectionName}"

	CloudBackupScheduleReqURI = "project/{ProjectId}/clusters/{ClusterName}/backup/schedule"

	GetAllBackupSnapshotReqURI = "project/{ProjectId}/cluster/{ClusterName}/snapshot"
	CreateBackupSnapshotReqURI = "project/{ProjectId}/cluster/{ClusterName}/snapshot"
	// CreateRestoreClusterReqURI is the URI for restoring a cluster.
	CreateRestoreClusterReqURI = "project/{ProjectId}/cluster/{ClusterName}/restore"
	// GetRestoreJobReqURI is the URI for getting a restore job.
	GetRestoreJobReqURI = "project/{ProjectId}/cluster/{ClusterName}/restore"

	InProgress             = "IN_PROGRESS"
	Success                = "SUCCESS"
	Failed                 = "FAILED"
	Expired                = "EXPIRED"
	Cancelled              = "CANCELLED"
	RestoreJobInProgress   = "Restore job is in progress"
	RestoreJobSuccess      = "Restore job completed successfully"
	RestoreJobFailed       = "Restore job failed"
	RestoreJobCancelled    = "Restore job was cancelled"
	RestoreJobExpired      = "Restore job expired"
	TargetProjectIdDefault = "5e8de3e1042f5b33ab81f33b"

	TargetClusterDefault = "u-AWS-22-09-23-10-48-28-5e8de3e1042f5b33ab81f33a"

	Automated   = "Automated"
	PointInTime = "PointInTime"

	DescriptionParam     = "Description"
	RetentionInDaysParam = "RetentionInDays"

	DeliveryTypeQueryParam = "DeliveryType"

	JobId = "JobId"

	SnapShotId            = "SnapshotId"
	OpLogTs               = "OpLogTs"
	OpLogInc              = "OpLogInc"
	PointInTimeUtcSeconds = "PointInTimeUtcSeconds"
	TargetClusterName     = "TargetClusterNameQueryParam"
	TargetProjectId       = "TargetProjectId"
)

func (c TraceIDKeyType) String() string {
	return string(c)
}
