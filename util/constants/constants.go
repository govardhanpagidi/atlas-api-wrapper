package constants

const (
	PubKey                  = "ApiKeys.PublicKey"
	PvtKey                  = "ApiKeys.PrivateKey"
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
	CollectionNamePathParam = "collectionName"
	HostNamePathParam       = "hostName"
	PasswordPathParam       = "password"

	DatabaseUserHandlerName = "DatabaseUserhandler"

	ClusterHandler = "ClusterHandler"
	PrivateKey     = "PrivateKey"
	PublicKey      = "PublicKey"
	TshirtSize     = "TshirtSize"
	DBUserName     = "DBUserName"

	PrivateKeyQueryParam = "privateKey"
	PublicKeyQueryParam  = "publicKey"

	InvalidInputParameter    = "INVALID_INPUT_PARAMETER"
	MongoClientCreationError = "MONGO_CLIENT_CREATE_ERROR"
	ClusterDoesNotExist      = "CLUSTER_NOT_EXIST"
	ProjectDoesNotExist      = "PROJECT_NOT_EXIST"
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
)
