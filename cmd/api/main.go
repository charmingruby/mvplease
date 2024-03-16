package main

import (
	"fmt"
	"os"

	"github.com/charmingruby/mvplease/config"
	"github.com/charmingruby/mvplease/internal/account"
	"github.com/charmingruby/mvplease/internal/shared/rest"
	"github.com/charmingruby/mvplease/internal/shared/rest/middlewares"
	"github.com/charmingruby/mvplease/pkg/logger"
	database "github.com/charmingruby/mvplease/pkg/postgres"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Logger
	logger := logger.NewLogger()

	// Environment
	if err := godotenv.Load(); err != nil {
		logger.Info("Environment warning: .env file not found")
	}

	// Configuration
	cfg, err := config.New(logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Configuration error: %s", err.Error()))
		os.Exit(1)
	}

	// Database
	db, err := database.New(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Database error: %s", err.Error()))
		os.Exit(1)
	}
	cfg.SetDatabase(db)

	// Services
	logger.Info("Initializing services...")

	accountService, err := account.NewService(cfg.Database.Conn, cfg.Logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Service initialization error: %s", err.Error()))
		os.Exit(1)
	}

	logger.Info("Services initialized.")

	// Server
	logger.Info("Setting up server...")
	router := mux.NewRouter()
	server, err := rest.NewServer(cfg, router)
	if err != nil {
		logger.Error(fmt.Sprintf("Server error: %s", err.Error()))
		os.Exit(1)
	}

	middlewares := middlewares.NewMiddleware(cfg.Logger)

	if err := account.NewHTTPService(router, middlewares, accountService, cfg); err != nil {
		logger.Error(fmt.Sprintf("Account HTTP error: %s", err.Error()))
		os.Exit(1)
	}
	logger.Info("Server setted.")

	if err := server.Start(); err != nil {
		logger.Errorf("Server error: %s", err.Error())
		os.Exit(1)
	}
}
