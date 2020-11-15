package response

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/leosmirnov/in-memory-cache/pkg/api/types"
	"github.com/leosmirnov/in-memory-cache/pkg/constants"
)

func WriteError(logger logrus.FieldLogger, w http.ResponseWriter, code int, err error) {
	if code >= 500 {
		logger.WithError(err).WithField("code", code).Warn("HTTP error")
	}

	WriteJSONContentType(w)
	w.WriteHeader(code)

	resp := types.ErrorBody{}
	resp.Error.Code = code
	resp.Error.Message = err.Error()
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		logger.WithError(err).Error("unable to encode response")
	}
}

func WriteJSONContentType(w http.ResponseWriter) {
	w.Header().Set(constants.ContentType, constants.ApplicationJSONUTF8)
}
