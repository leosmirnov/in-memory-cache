package api

import (
	"net/http"

	"github.com/leosmirnov/in-memory-cache/pkg/utils/httpserver"
)

func (api *API) routes() httpserver.Routes {
	return httpserver.Routes{
		// Cache operations.
		{Method: http.MethodGet, Path: "/values/:key", Handler: api.GetValue},
		{Method: http.MethodPost, Path: "/values", Handler: api.AddValue},
		{Method: http.MethodDelete, Path: "/values/:key", Handler: api.DeleteValue},
		{Method: http.MethodGet, Path: "/keys", Handler: api.ListKeys},

		// API.
		{Method: http.MethodGet, Path: "/api-docs", Handler: api.apiDocsHandler},

		// health
		{Method: http.MethodGet, Path: "/live", Handler: api.liveHandler},
	}
}
