package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/leosmirnov/in-memory-cache/pkg/api/response"
)

func (api *API) apiDocsHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	buff, err := ioutil.ReadFile("./openapi.json")
	if err != nil {
		response.WriteError(api.logger, w, http.StatusInternalServerError, fmt.Errorf("failed to read swagger json file: %s", err))
		return
	}

	_, err = w.Write(buff)
	if err != nil {
		response.WriteError(api.logger, w, http.StatusInternalServerError, fmt.Errorf("failed to write swagger data to response: %s", err))
	}
}
