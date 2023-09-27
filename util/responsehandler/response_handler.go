package responseHandler

import (
	"encoding/json"
	"net/http"

	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/logger"
)

// Write writes the given AtlasResponse object to the http.ResponseWriter object
func Write(response atlasresponse.AtlasResponse, w http.ResponseWriter, handlerName string) {
	// Set the Content-Type header of the http.ResponseWriter object to "application/json"
	w.Header().Set("Content-Type", "application/json")
	if response.HttpStatusCode != 0 {
		w.WriteHeader(response.HttpStatusCode)
	}
	var res []byte
	var err error
	if response.Response != nil {
		res, err = json.Marshal(response.Response)
	} else {
		// Marshal the given AtlasResponse object to a JSON string
		res, err = json.Marshal(response)
	}

	// If there is an error and the message field of the AtlasResponse object is not empty, log the error and the message
	if response.Message != "" && err != nil {
		_, _ = logger.Debugf(handlerName+" error:%s", response.Message)
	}

	// Write the JSON string to the http.ResponseWriter object
	_, _ = w.Write(res)
}
