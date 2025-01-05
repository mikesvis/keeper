package app

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"keeper/internal/config"
	"keeper/internal/minio"
	"keeper/internal/postgres"
	"keeper/internal/service/secret"
	"keeper/internal/service/user"
	"keeper/internal/storage/files"
	sStorage "keeper/internal/storage/secret"
	uStorage "keeper/internal/storage/user"
)

type App struct {
	config        *config.Config
	logger        *zap.SugaredLogger
	db            *sqlx.DB
	UserService   *user.Service
	SecretService *secret.Service
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

	// Secret storage
	secretStorage, err := sStorage.NewStorage(db, l)
	if err != nil {
		return nil, err
	}

	// Minio client
	minioClient, err := minio.NewMinio(c, l)
	if err != nil {
		return nil, err
	}

	// Minio storage
	fileStorage := files.NewStorage(c, minioClient)

	// Secret service
	secretService := secret.NewService(secretStorage, fileStorage, l)

	return &App{
		config:        c,
		logger:        l,
		db:            db,
		UserService:   userService,
		SecretService: secretService,
	}, nil
}
