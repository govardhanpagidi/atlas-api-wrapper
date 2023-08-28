//go:build !coverage

package constants

type TraceIDKeyType string

const (
	ClusterNamePathParam    = "clusterName"
	ProjectID               = "ProjectId"
	ProjectIdPathParam      = "projectId"
	AWS                     = "AWS"
	HostName                = "HostName"
	DatabaseName            = "DatabaseName"
	Username                = "Username"
	Password                = "Password"
	DatabaseNamePathParam   = "databaseName"
	UsernamePathParam       = "username"
	ClusterName             = "ClusterName"
	CollectionName          = "CollectionName"
	CollectionNames         = "CollectionNames"
	CollectionNamePathParam = "collectionName"
	HostNamePathParam       = "hostName"
	PasswordPathParam       = "password"

	DatabaseUserHandlerName = "DatabaseUserhandler"

	ClusterHandler = "ClusterHandler"
	PrivateKey     = "PrivateKey"
	PublicKey      = "PublicKey"
	TshirtSize     = "TshirtSize"
	CloudProvider  = "CloudProvider"

	PrivateKeyQueryParam = "privateKey"
	PublicKeyQueryParam  = "publicKey"

	InvalidInputParameter    = "INVALID_INPUT_PARAMETER"
	MongoClientCreationError = "MONGO_CLIENT_CREATE_ERROR"
	ResourceDoesNotExist     = "RESOURCE_NOT_EXIST"
	ClusterModelError        = "CLUSTER_MODEL_ERROR"
	ClusterRequestError      = "CLUSTER_REQUEST_ERROR"
	ClusterCreateError       = "CLUSTER_CREATE_ERROR"
	ClusterCreateSuccess     = "CLUSTER_CREATE_SUCCESS"
	ClusterReadSuccess       = "CLUSTER_READ_SUCCESS"
	ClusterDeleteError       = "CLUSTER_DELETE_ERROR"
	ClusterDeleteSuccess     = "CLUSTER_DELETE_SUCCESS"
	ClusterListError         = "CLUSTER_LIST_ERROR"
	ClusterAdvancedListError = "CLUSTER_ADVANCED_LIST_ERROR"
	ClusterListSuccess       = "CLUSTER_LIST_SUCCESS"
	ListEndpointError        = "LIST_ENDPOINT_ERROR"
	NoEndpointConfigured     = "NO_ENDPOINT_ERROR"

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
	DeleteDatabaseUserError   = "DELETE_USER_ERROR"
	DeleteDatabaseUserSuccess = "DELETE_USER_SUCCESS"
	UserListError             = "USER_LIST_ERROR"
	UserListSuccess           = "USER_LIST_SUCCESS"

	DbuserDbName    = "admin"
	DbAdminRoleName = "dbAdmin"
	AtlasAdminRole  = "atlasAdmin"

	ClusterConfigLocation                = "config.json"
	TraceID               TraceIDKeyType = "traceID"
	Project                              = "Project"
	Cluster                              = "Cluster"

	MongoBaseUrl = "https://cloud.mongodb.com/"
	MongoDbURI   = "mongodb+srv://%s:%s@%s"

	MessageConfigLocation   = "message_Config.json"
	EmptyString             = ""
	DefaultCollectionString = "default"

	CreateOrGetDatabaseUserReqURI = "project/{projectId}/databaseUsers"
	DeleteOrGetDatabaseUserReqURI = "project/{projectId}/databaseUsers/{username}"

	CreateOrGetClusterReqURI = "project/{projectId}/cluster"
	DeleteClusterReqURI      = "project/{projectId}/cluster/{clusterName}"
	GetClusterReqURI         = "project/{projectId}/cluster/{clusterName}/status"

	CreateDatabaseReqURI = "database"
	DeleteDatabaseReqURI = "database/{databaseName}"

	CreateCollectionReqURI = "database/{databaseName}/collections"
	DeleteCollectionReqURI = "database/{databaseName}/collection/{collectionName}"
)

func (c TraceIDKeyType) String() string {
	return string(c)
}
