package user

import (
	"context"
	"database/sql"
	"errors"
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
	createTable := `
		CREATE TABLE IF NOT EXISTS users (
			id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL
		)
	`
	_, err := db.Exec(createTable)
	return err
}

func (s *Storage) GetUserIDByKey(ctx context.Context, key string) (uint64, error) {
	var userID uint64

	query := `SELECT id FROM users WHERE public_key = $1`
	err := s.db.QueryRowContext(ctx, query, key).Scan(&userID)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return userID, nil
}
