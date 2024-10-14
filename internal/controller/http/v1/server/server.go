package server

import (
	"fmt"
	"github.com/Justdanru/bhs-test/config"
	"log/slog"
	"net/http"
)

type HTTPServer struct {
	server *http.Server
	cfg    *config.Config
	logger *slog.Logger
}

func NewHTTPServer(
	cfg *config.Config,
	handler http.Handler,
	logger *slog.Logger,
) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
			Handler: handler,
		},
		logger: logger,
		cfg:    cfg,
	}
}

func (s *HTTPServer) Run() error {
	return s.server.ListenAndServe()
}

func (s *HTTPServer) Shutdown() {
	if err := s.server.Close(); err != nil {
		s.logger.Error("closing http server failed", "error", err)
	}
}
