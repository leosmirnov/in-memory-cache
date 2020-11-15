package api

import (
	"github.com/sirupsen/logrus"

	"github.com/leosmirnov/in-memory-cache/pkg/conf"
	"github.com/leosmirnov/in-memory-cache/pkg/storage"
	"github.com/leosmirnov/in-memory-cache/pkg/utils/httpserver"
)

type API struct {
	cfg        *conf.API
	logger     logrus.FieldLogger
	srv        *httpserver.Server
	storageSvc storage.Service
}

func New(cfg *conf.API, logger logrus.FieldLogger, storage storage.Service) *API {
	return &API{
		cfg:        cfg,
		logger:     logger,
		storageSvc: storage,
	}
}

func (api *API) Start() error {
	api.srv = httpserver.New(&httpserver.Opts{
		Port:   api.cfg.Port,
		Host:   api.cfg.Host,
		Routes: api.routes(),
		Logger: api.logger,
	})
	return api.srv.Start()
}

func (api *API) Stop() error {
	return api.srv.Stop()
}
