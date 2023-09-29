package collection

import (
	"fmt"
	"github.com/atlas-api-helper/util"
)

// InputModel represents the input for creating a collection.
// swagger:parameters InputModel
type InputModel struct {
	// CollectionNames is an array of collection names to be queried.
	//
	// required: true
	// example: ["default", "users"]
	CollectionNames []*string `json:"collectionNames,omitempty" example:"default"`
	// DatabaseName is the name of the database to be queried.
	//
	// required: true
	// example: "testDatabase"
	DatabaseName *string `json:"-" example:"testDatabase" swagg:"omit"`
	// ProjectId is the ID of the project.
	//
	//required: true
	// example: ""
	ProjectId *string `json:"-"`
	// ClusterName is the name of the cluster.
	//
	// required: true
	// example: ""
	ClusterName *string `json:"-"`
	// HostName is the hostname of the database server.
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

// DeleteInputModel represents the input for deleting a collection.
// swagger:parameters DeleteInputModel
type DeleteInputModel struct {
	// CollectionName is the name of the collection to be deleted.
	//
	// required: true
	CollectionName *string `json:"collectionName,omitempty" example:"testCollection"`
	// DatabaseName is the name of the database to be queried.
	//
	// required: true
	// example: "testDatabase"
	DatabaseName *string `json:"-" example:"testDatabase"`
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
	// HostName is the hostname of the database server.
	//
	// required: false
	// example: "localhost"
	HostName *string `json:"-"`
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
		util.ToStringSlice(model.CollectionNames),
		util.ToString(model.DatabaseName),
		util.ToString(model.HostName),
	)
}

func (model DeleteInputModel) String() string {
	return fmt.Sprintf(
		"CollectionNames: %s, DatabaseName: %s, HostName: %s",
		util.ToString(model.CollectionName),
		util.ToString(model.DatabaseName),
		util.ToString(model.HostName),
	)
}
