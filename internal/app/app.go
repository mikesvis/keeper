package app

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"keeper/internal/config"
	"keeper/internal/postgres"
	"keeper/internal/service/user"
	uStorage "keeper/internal/storage/user"
)

type App struct {
	config      *config.Config
	logger      *zap.SugaredLogger
	db          *sqlx.DB
	userService *user.Service
}

func NewApp(c *config.Config, l *zap.SugaredLogger) (*App, error) {

	//	DB init
	db, err := postgres.NewPostgres(c)
	if err != nil {
		return nil, err
	}

	// User storage
	userStorage, err := uStorage.NewStorage(db, l)
	if err != nil {
		return nil, err
	}

	// User service
	userService := user.NewService(userStorage, l)

	return &App{
		config:      c,
		logger:      l,
		db:          db,
		userService: userService,
	}, nil
}
