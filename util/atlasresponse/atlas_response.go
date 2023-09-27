package atlasresponse

import (
	"encoding/json"
	"fmt"
)

type AtlasResponse struct {
	Response       interface{} `json:"response,omitempty"`
	HttpStatusCode int         `json:"-"`
	Message        string      `json:"errorMessage,omitempty"`
	ErrorCode      string      `json:"errorCode,omitempty"`
	Status         string      `json:"status,omitempty"`
}

// String returns a string representation of the AtlasResponse object
func (ar AtlasResponse) String() string {
	// Marshal the Response field of the AtlasResponse object to JSON
	responseBytes, _ := json.Marshal(ar.Response)

	// Return a formatted string with the Response, HttpStatusCode, and Message fields of the AtlasResponse object
	return fmt.Sprintf("Response: %s, HttpStatusCode: %d, Message: %s", responseBytes, ar.HttpStatusCode, ar.Message)
}
