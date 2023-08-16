package collection

import (
	"fmt"
	"github.com/atlas-api-helper/util"
)

type InputModel struct {
	CollectionName *string `json:",omitempty"`
	DatabaseName   *string `json:",omitempty"`
	HostName       *string `json:",omitempty"`
	Username       *string `json:",omitempty"`
	Password       *string `json:",omitempty"`
}

type DeleteInputModel struct {
	DatabaseName *string `json:",omitempty"`
	HostName     *string `json:",omitempty"`
	Username     *string `json:",omitempty"`
	Password     *string `json:",omitempty"`
}

func (model InputModel) String() string {
	return fmt.Sprintf(
		"CollectionName: %s, DatabaseName: %s, HostName: %s",
		util.ToString(model.CollectionName),
		util.ToString(model.DatabaseName),
		util.ToString(model.HostName),
	)
}

func (model DeleteInputModel) String() string {
	return fmt.Sprintf(
		"DatabaseName: %s, HostName: %s",
		util.ToString(model.DatabaseName),
		util.ToString(model.HostName),
	)
}
