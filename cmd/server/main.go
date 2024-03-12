package main

import (
	"fmt"
	"os"

	"github.com/charmingruby/mvplease/pkg/config"
	"github.com/charmingruby/mvplease/pkg/logger"
	"github.com/joho/godotenv"
)

func main() {
	logger := logger.NewLogger()

	if err := godotenv.Load(); err != nil {
		logger.Info("Environment warning: .env file not found")
	}

	_, err := config.New(logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Configuration error: %s", err.Error()))
		os.Exit(1)
	}
}
