package logger

import (
	"go.uber.org/zap"
	"keeper/internal/config"
)

// NewLogger Инициализация логирования
func NewLogger(config *config.Config) (*zap.SugaredLogger, error) {
	var sugar *zap.Logger
	var err error

	if config.Environment == "production" {
		sugar, err = zap.NewProduction()
	} else {
		sugar, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	defer sugar.Sync()

	return sugar.Sugar(), nil
}
