package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
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
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			login VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := db.Exec(createTable)
	return err
}

func (s *Storage) ExistByLogin(ctx context.Context, login string) (bool, error) {
	var count int

	query := `SELECT COUNT(1) FROM users WHERE login = $1`
	err := s.db.QueryRowContext(ctx, query, login).Scan(&count)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}

func (s *Storage) Create(ctx context.Context, login, password string) (uuid.UUID, error) {
	var userID uuid.UUID

	err := s.db.QueryRowContext(ctx, "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id", login, password).Scan(&userID)

	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

func (s *Storage) GetUserID(ctx context.Context, login, password string) (uuid.UUID, error) {
	var userID uuid.UUID

	query := `SELECT id FROM users WHERE login = $1 AND password = $2`
	err := s.db.QueryRowContext(ctx, query, login, password).Scan(&userID)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return uuid.Nil, nil
	}

	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
