package main

import (
	"fmt"
	"keeper/internal/config"
	"keeper/internal/logger"
	"keeper/internal/postgres"
	"os"
)

const exitErrorCode = 1

func main() {
	// Config
	appConfig, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error while initializing server config: %s\n", err.Error())
		os.Exit(exitErrorCode)
	}

	// Logger
	appLogger, err := logger.NewLogger(appConfig)
	if err != nil {
		fmt.Printf("Error while initializing server logger: %s\n", err.Error())
		os.Exit(exitErrorCode)
	}

	// Storage
	db, err := postgres.NewPostgres(appConfig)
	if err != nil {
		fmt.Printf("Error while initializing db: %s\n", err.Error())
		os.Exit(exitErrorCode)
	}

	appLogger.Info(appConfig, db)
}
