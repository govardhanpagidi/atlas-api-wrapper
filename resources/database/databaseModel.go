package database

import (
	"fmt"
	"github.com/atlas-api-helper/util"
)

// InputModel represents the input for creating a collection.
// swagger:parameters InputModel
type InputModel struct {
	// The name of the collection to be queried.
	//
	// required: true
	// example: "testCollection"
	CollectionName *string `json:"collectionName,omitempty" example:"testCollection"`
	// The name of the database to be queried.
	//
	// required: false
	// example: "testDatabase"
	DatabaseName *string `json:"databaseName,omitempty" example:"testDatabase"`
	// ProjectId is the ID of the project.
	//
	// required: true
	// example: ""
	ProjectId *string `json:"-"`
	// ClusterName is the name of the cluster.
	//
	// required: true
	// example: ""
	ClusterName *string `json:"-"`
	// The hostname of the database server.
	//
	// required: false
	// example: "localhost"
	HostName *string `json:"-,omitempty"`
	// Username is the username for the database server.
	//
	// required: true
	// example: ""
	Username *string `json:"-"`
	// Password is the password for the database server.
	//
	// required: true
	// example: ""
	Password *string `json:"-"`
	// PublicKey is the public key.
	//
	// required: true
	// example: ""
	PublicKey *string `json:"-"`
	// PrivateKey is the private key.
	//
	// required: true
	// example: ""
	PrivateKey *string `json:"-"`
}

func (model InputModel) String() string {
	return fmt.Sprintf(
		"CollectionNames: %s, DatabaseName: %s, HostName: %s",
		util.ToString(model.CollectionName),
		util.ToString(model.DatabaseName),
		util.ToString(model.HostName),
	)
}
