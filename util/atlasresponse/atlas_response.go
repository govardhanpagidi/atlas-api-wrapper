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

// String returns a string representation of the AtlasRespone object
func (ar AtlasRespone) String() string {
	// Marshal the Response field of the AtlasRespone object to JSON
	responseBytes, _ := json.Marshal(ar.Response)

	// Return a formatted string with the Response, HttpStatusCode, and Message fields of the AtlasRespone object
	return fmt.Sprintf("Response: %s, HttpStatusCode: %d, Message: %s", responseBytes, ar.HttpStatusCode, ar.Message)
}
