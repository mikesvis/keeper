package user

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Storage struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewStorage(db *sqlx.DB, logger *zap.SugaredLogger) (*Storage, error) {
	err := bootstrap(db)
	if err != nil {
		return nil, err
	}

	return &Storage{db, logger}, nil
}

func bootstrap(db *sqlx.DB) error {
	return nil
}
