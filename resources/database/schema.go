package database

import "fmt"

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
		toString(model.CollectionName),
		toString(model.DatabaseName),
		toString(model.HostName),
	)
}

func (model DeleteInputModel) String() string {
	return fmt.Sprintf(
		"DatabaseName: %s, HostName: %s",
		toString(model.DatabaseName),
		toString(model.HostName),
	)
}

func toString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
