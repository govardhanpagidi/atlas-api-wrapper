package constants

const (
	DatabaseUserHandler    = "project/{projectId}/databaseUsers"
	DatabaseUserGetHandler = "project/{projectId}/databaseUsers/{username}"

	ClusterWithProjectId      = "project/{projectId}/cluster"
	ClusterWithGroupIdAndName = "project/{projectId}/cluster/{clusterName}"
	ClusterStatus             = "project/{projectId}/cluster/{clusterName}/status"

	DatabaseHandler       = "database"
	DatabaseDeleteHandler = "database/{databaseName}"

	CollectionHandler       = "database/{databaseName}/collections"
	CollectionDeleteHandler = "database/{databaseName}/collection/{collectionName}"
)
