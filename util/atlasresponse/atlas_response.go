package atlasresponse

import (
	"encoding/json"
	"fmt"
)

type AtlasRespone struct {
	Response       interface{} `json:"response,omitempty"`
	HttpStatusCode int         `json:"code,omitempty"`
	Message        string      `json:"message,omitempty"`
}

func (ar AtlasRespone) String() string {
	responseBytes, _ := json.Marshal(ar.Response)
	return fmt.Sprintf("Response: %s, HttpStatusCode: %d, Message: %s", responseBytes, ar.HttpStatusCode, ar.Message)
}
