package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *API) liveHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
