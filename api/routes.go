package api

func (api *API) routes() staxdhttp.Routes {
	return staxdhttp.Routes{
		// Staxd.
		{Method: http.MethodGet, Path: "/live", Handler: api.liveHandler},
		{Method: http.MethodGet, Path: "/ready", Handler: api.readyHandler},
		{Method: http.MethodGet, Path: "/info", Handler: api.infoHandler},

		// Jobs.
		{Method: http.MethodGet, Path: "/jobs/:id", Handler: api.JobGetHandler},

		// Ecostructures.
		{Method: http.MethodPost, Path: "/ecostructures", Handler: api.EcoCreationHandler},
		{Method: http.MethodGet, Path: "/ecostructures", Handler: api.EcoListHandler},
		{Method: http.MethodGet, Path: "/ecostructures/:id", Handler: api.EcoGetHandler},
		{Method: http.MethodPut, Path: "/ecostructures/:id", Handler: api.EcoUpdateHandler},
		{Method: http.MethodPatch, Path: "/ecostructures/:id", Handler: api.EcoPatchHandler},
		{Method: http.MethodPost, Path: "/ecostructures/:id", Handler: api.EcoPatchHandler},
		{Method: http.MethodDelete, Path: "/ecostructures/:id", Handler: api.EcoDeleteHandler},
		{Method: http.MethodGet, Path: "/ecostructures/:id/pools", Handler: api.EcoPoolsHandler},
		{Method: http.MethodGet, Path: "/ecostructures/:id/deployments", Handler: api.DeploymentEcoHandler},

		// Pools.
		{Method: http.MethodPost, Path: "/pools", Handler: api.PoolCreationHandler},
		{Method: http.MethodGet, Path: "/pools", Handler: api.PoolListHandler},
		{Method: http.MethodGet, Path: "/pools/:id", Handler: api.PoolGetHandler},
		{Method: http.MethodPut, Path: "/pools/:id", Handler: api.PoolUpdateHandler},
		{Method: http.MethodPost, Path: "/pools/:id", Handler: api.PoolPatchHandler},
		{Method: http.MethodPatch, Path: "/pools/:id", Handler: api.PoolPatchHandler},
		{Method: http.MethodDelete, Path: "/pools/:id", Handler: api.PoolDeleteHandler},
		{Method: http.MethodGet, Path: "/pools/:id/ecostructure", Handler: api.PoolEcoHandler},
		{Method: http.MethodGet, Path: "/pools/:id/deployments", Handler: api.DeploymentPoolHandler},

		// Deployments.
		{Method: http.MethodPost, Path: "/deployments", Handler: api.DeploymentCreationHandler},
		{Method: http.MethodGet, Path: "/deployments", Handler: api.DeploymentListHandler},
		{Method: http.MethodGet, Path: "/deployments/:id/stats", Handler: api.DeploymentStatsHandler},
		{Method: http.MethodGet, Path: "/deployments/:id", Handler: api.DeploymentGetHandler},
		{Method: http.MethodPut, Path: "/deployments/:id", Handler: api.DeploymentUpdateHandler},
		{Method: http.MethodPatch, Path: "/deployments/:id", Handler: api.DeploymentUpdateHandler},
		{Method: http.MethodGet, Path: "/deployments/:id/info", Handler: api.DeploymentInfoHandler},
		{Method: http.MethodDelete, Path: "/deployments/:id", Handler: api.DeploymentDeleteHandler},
		{Method: http.MethodGet, Path: "/deployments/:id/logs", Handler: api.DeploymentLogsHandler},
		{Method: http.MethodGet, Path: "/deployments/:id/pools", Handler: api.PoolDeploymentHandler},
		{Method: http.MethodGet, Path: "/deployments/:id/plugins", Handler: api.DeploymentPluginHandler},

		// Errors
		{Method: http.MethodGet, Path: "/errors", Handler: api.ErrorListHandler},
		{Method: http.MethodGet, Path: "/errors/:id", Handler: api.ErrorGetHandler},

		// Plugins.
		{Method: http.MethodGet, Path: "/plugins", Handler: api.getDeploymentPluginsListHandler},
		{Method: http.MethodGet, Path: "/drivers", Handler: api.getDriverPluginsListHandler},

		// API.
		{Method: http.MethodGet, Path: "/api-docs", Handler: api.apiDocsHandler},
	}
}
