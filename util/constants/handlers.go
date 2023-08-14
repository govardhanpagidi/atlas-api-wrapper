package constants

const (
	DatabaseUserHandler    = "project/{projectId}/databaseUsers"
	DatabaseUserGetHandler = "project/{projectId}/database/{databaseName}/databaseUsers/{username}"

	ClusterWithProjectId      = "project/{projectId}/cluster"
	ClusterWithGroupIdAndName = "project/{projectId}/cluster/{clusterName}"

	DatabaseHandler       = "database"
	DatabaseDeleteHandler = "database/{databaseName}"

	CollectionHandler       = "database/{databaseName}/collections"
	CollectionDeleteHandler = "database/{databaseName}/collection/{collectionName}"
)
