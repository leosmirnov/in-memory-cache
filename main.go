package main

import (
	"flag"
	"fmt"
	"github.com/leosmirnov/in-memory-cache/pkg/api"
	"github.com/leosmirnov/in-memory-cache/pkg/conf"
	"github.com/leosmirnov/in-memory-cache/pkg/storage/inmemory"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type Service struct {
	// Config fields.
	cfg *conf.Conf

	// Services.
	apiService *api.API

	// System fields.
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
		// TODO: error handling
		fmt.Println(err)
	}

	s := Service{
		cfg:    cfg,
		logger: logger,
	}

	kvService := inmemory.NewService()
	s.apiService = api.New(cfg.API, logger, kvService)
	go s.startAPIService()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	<-interrupt
	logger.Debug("handle SIGINT, SIGTERM, SIGKILL, SIGQUIT")

	//ctx, cancel := context.WithTimeout(context.Background(), s.StopTimeout())
	//defer cancel()
	//go s.Stop(cancel)
	//
	//select {
	//case <-ctx.Done():
	//	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
	//		logger.Warn("app stopped with problems")
	//		return
	//	}
	//
	//	logger.Info("app gracefully stopped")
	//}

}

func (s *Service) startAPIService() {
	if err := s.apiService.Start(); err != nil && err != http.ErrServerClosed {
		s.logger.WithError(err).Fatal("failed start api service")
	}
}
