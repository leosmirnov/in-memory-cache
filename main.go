package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	"github.com/leosmirnov/in-memory-cache/pkg/api"
	"github.com/leosmirnov/in-memory-cache/pkg/conf"
	"github.com/leosmirnov/in-memory-cache/pkg/storage"
	"github.com/leosmirnov/in-memory-cache/pkg/storage/inmemory"
)

type Service struct {
	// Config fields.
	cfg *conf.Conf

	// Services.
	apiService *api.API
	kvService  storage.Service

	logger logrus.FieldLogger
}

func main() {
	cfgPath := flag.String("config", "./config.yml", "full file path to config")
	flag.Parse()

	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}

	cfg, err := conf.Configure(logger, *cfgPath)
	if err != nil {
		logrus.Fatalf("failed to read config")
		return
	}

	s := Service{
		cfg:    cfg,
		logger: logger,
	}

	s.kvService = inmemory.NewService(logger, &cfg.Cache.CleanupInterval)
	s.apiService = api.New(cfg.API, logger, s.kvService)
	go s.startAPIService()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-interrupt
	logger.Debug("handle SIGINT, SIGTERM, SIGKILL, SIGQUIT")

	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.App.StopTimeout)
	defer cancel()
	go s.finalize(cancel)

	select {
	case <-ctx.Done():
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			logger.Warn("app stopped with problems")
			return
		}

		logger.Info("app gracefully stopped")
	}
}

func (s *Service) startAPIService() {
	if err := s.apiService.Start(); err != nil && err != http.ErrServerClosed {
		s.logger.WithError(err).Fatal("failed start api service")
	}
}

func (s *Service) finalize(cancel context.CancelFunc) {
	err := s.kvService.Close()
	if err != nil {
		s.logger.WithError(err).Error("failed finalize kv service")
	}
	err = s.apiService.Stop()
	if err != nil {
		s.logger.WithError(err).Error("failed finalize api service")
	}

	cancel()
}
