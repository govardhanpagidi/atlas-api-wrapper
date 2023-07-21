package responseHandler

import (
	"encoding/json"
	"github.com/atlas-api-helper/util/atlasresponse"
	"github.com/atlas-api-helper/util/logger"
	"net/http"
)

func Write(response atlasresponse.AtlasRespone, w http.ResponseWriter, handlerName string) {
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(response)
	if response.HttpError != "" && err != nil {
		_, _ = logger.Debugf(handlerName+" error:%s", response.HttpError)

		_, _ = w.Write(res)
		return
	}
	_, _ = w.Write(res)

}
