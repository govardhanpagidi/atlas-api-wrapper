package constants

const (
	ProjectHandler          = "project"
	ProjectHandlerWithId    = "project/{Id}"
	DatabaseuserHandler     = "users"
	DatabaseuserGetHandler  = "users/{GroupId}/{DatabaseName}/{Username}"
	DatabaseuserWithGroupId = "users/{GroupId}"
	CustomDbRole            = "customdbrole"
	CustomDbRoleGetHandler  = "customdbrole/{GroupId}/{RoleName}"
	CustomDbRoleWithGroupId = "customdbrole/{GroupId}"
)
