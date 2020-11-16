package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/leosmirnov/in-memory-cache/pkg/api/request"
	"github.com/leosmirnov/in-memory-cache/pkg/api/response"
	"github.com/leosmirnov/in-memory-cache/pkg/api/types"
	"github.com/leosmirnov/in-memory-cache/pkg/constants"
)

func (api *API) AddValue(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rqContentType := r.Header.Get(constants.ContentType)
	if !strings.Contains(rqContentType, constants.ApplicationJSON) && !strings.Contains(rqContentType, constants.TextJSON) {
		response.WriteError(api.logger, w, http.StatusBadRequest, errors.New("unsupported media type"))
		return
	}

	if r.Body == nil {
		response.WriteError(api.logger, w, http.StatusBadRequest, errors.New("no body specified"))
		return
	}

	var body types.KeyValue
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		response.WriteError(api.logger, w, http.StatusBadRequest, fmt.Errorf("unable to decode request body: %s", err))
		return
	}
	if !body.Validate() {
		response.WriteError(api.logger, w, http.StatusBadRequest, errors.New("key and value must not be empty"))
		return
	}

	key := body.Key
	bts := []byte(body.Value)
	exp, _ := time.ParseDuration(body.Expiration)
	err := api.storageSvc.Set(key, bts, exp)
	if err != nil {
		response.WriteError(api.logger, w, http.StatusConflict, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (api *API) GetValue(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	key := params.ByName(request.ParamKey)

	value, err := api.storageSvc.GetByKey(key)
	if err != nil {
		response.WriteError(api.logger, w, http.StatusNotFound, fmt.Errorf("unable to get key: %s", err))
		return
	}
	w.WriteHeader(http.StatusOK)
	response.WriteJSONContentType(w)
	if err = json.NewEncoder(w).Encode(&types.DataBody{Data: &types.Value{Value: string(value)}}); err != nil {
		response.WriteError(api.logger, w, http.StatusInternalServerError, fmt.Errorf("unable to write response: %s", err))
		return
	}
}

func (api *API) ListKeys(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	keys := api.storageSvc.GetKeys()

	resp := &types.KeysList{
		Keys: keys,
	}
	if err := json.NewEncoder(w).Encode(&types.DataBody{Data: resp}); err != nil {
		response.WriteError(api.logger, w, http.StatusInternalServerError, fmt.Errorf("unable to write response: %s", err))
		return
	}
	response.WriteJSONContentType(w)
}

func (api *API) DeleteValue(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	key := params.ByName(request.ParamKey)

	err := api.storageSvc.RemoveKey(key)
	if err != nil {
		response.WriteError(api.logger, w, http.StatusNotFound, fmt.Errorf("unable to get key: %s", err))
		return
	}
	w.WriteHeader(http.StatusOK)
}
