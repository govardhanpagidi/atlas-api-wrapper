package collection

import (
	"fmt"
	"github.com/atlas-api-helper/util"
)

type InputModel struct {
	CollectionNames []*string `json:"CollectionName,omitempty"`
	DatabaseName    *string   `json:"databaseName,omitempty"`
	HostName        *string   `json:"hostName,omitempty"`
	Username        *string   `json:"userName,omitempty"`
	Password        *string   `json:"password,omitempty"`
}

type DeleteInputModel struct {
	CollectionName *string `json:"collectionName,omitempty"`
	DatabaseName   *string `json:"databaseName,omitempty"`
	HostName       *string `json:"hostName,omitempty"`
	Username       *string `json:"userName,omitempty"`
	Password       *string `json:"password,omitempty"`
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
