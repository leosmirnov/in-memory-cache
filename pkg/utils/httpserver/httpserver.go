package httpserver

import (
	"context"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type (
	Route struct {
		Method  string
		Path    string
		Handler httprouter.Handle
	}

	Routes []*Route

	Server struct {
		srv    *http.Server
		logger logrus.FieldLogger

		available  bool
		reindexing bool
	}

	Opts struct {
		Port   int
		Host   string
		Routes Routes
		Logger logrus.FieldLogger
	}
)

func New(opts *Opts) *Server {
	server := &Server{
		logger: opts.Logger,
	}

	router := httprouter.New()
	for _, route := range opts.Routes {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	server.srv = &http.Server{
		Addr:    opts.Host + ":" + strconv.Itoa(opts.Port),
		Handler: router,
	}

	return server
}

func (s *Server) Start() error {
	s.logger.Infof("listening http requests on %s", s.srv.Addr)

	if err := s.srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			return err
		}

		return errors.Wrap(err, "http server listen and serve err")
	}
	return nil
}

func (s *Server) Stop() error {
	s.logger.Debug("shutting down API server")

	if err := s.srv.Shutdown(context.TODO()); err != nil {
		return errors.Wrap(err, "http server shutdown err")
	}

	if err := s.srv.Close(); err != nil {
		return errors.Wrap(err, "http server close err")
	}

	s.logger.Debug("API server is shut down")
	return nil
}
