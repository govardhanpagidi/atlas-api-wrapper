package constants

const (
	DatabaseUserReqURI             = "project/{projectId}/databaseUsers"
	DatabaseUserWithUsernameReqURI = "project/{projectId}/databaseUsers/{username}"

	ClusterReqURI                = "project/{projectId}/cluster"
	ClusterWithClusterNameReqURI = "project/{projectId}/cluster/{clusterName}"
	ClusterStatusReqURI          = "project/{projectId}/cluster/{clusterName}/status"

	DatabaseReqURI       = "database"
	DatabaseDeleteReqURI = "database/{databaseName}"

	CollectionReqURI       = "database/{databaseName}/collections"
	CollectionDeleteReqURI = "database/{databaseName}/collection/{collectionName}"
)
