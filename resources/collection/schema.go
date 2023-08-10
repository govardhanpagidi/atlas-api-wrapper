package collection

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
