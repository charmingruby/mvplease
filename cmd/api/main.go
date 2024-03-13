package main

import (
	"fmt"
	"os"

	"github.com/charmingruby/mvplease/internal/accounts/database/postgres"
	"github.com/charmingruby/mvplease/internal/accounts/domain"
	"github.com/charmingruby/mvplease/internal/config"
	"github.com/charmingruby/mvplease/internal/shared/http"
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

	// Repositories
	logger.Info("Initializing repositories...")

	accountRepository, err := postgres.NewAccountRepository(cfg.Database.Conn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Repositories initialized.")

	// Services
	logger.Info("Initializing services...")

	domain.NewService(&accountRepository)

	logger.Info("Services initialized.")

	// Server
	router := mux.NewRouter()
	server, err := http.NewServer(cfg, router)
	if err != nil {
		logger.Error(fmt.Sprintf("Server error: %s", err.Error()))
		os.Exit(1)
	}

	if err := server.Start(); err != nil {
		logger.Errorf("Server error: %s", err.Error())
		os.Exit(1)
	}
}
