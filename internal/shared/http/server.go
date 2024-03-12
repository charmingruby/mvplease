package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmingruby/mvplease/internal/config"
	"github.com/gorilla/mux"
)

type Server struct {
	*http.Server
	Config *config.Config
}

func NewServer(cfg *config.Config, router *mux.Router) (*Server, error) {
	if router == nil {
		return nil, fmt.Errorf("invalid server router")
	}

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return &Server{
		Server: server,
		Config: cfg,
	}, nil
}

func (s *Server) Start() error {
	s.Config.Logger.Info("HTTP server is running...")

	if err := s.Server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
