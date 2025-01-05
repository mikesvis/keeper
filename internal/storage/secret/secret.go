package secret

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"keeper/internal/domain"
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
		CREATE TABLE IF NOT EXISTS secrets (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID NOT NULL,
			type VARCHAR(10) NOT NULL,
			name VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := db.Exec(createTable)
	return err
}

func (s *Storage) CreateSecret(ctx context.Context, secret *domain.UserSecret) (*domain.UserSecret, error) {
	row := s.db.QueryRowContext(
		ctx,
		"insert into secrets (id, user_id, type, name, created_at, updated_at) values ($1, $2, $3, $4, $5, $6)",
		secret.ID,
		secret.UserID,
		secret.Type,
		secret.Name,
		secret.CreatedAt,
		secret.UpdatedAt,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	storedSecret, err := s.FindById(ctx, secret.ID)
	if err != nil {
		return nil, err
	}

	if storedSecret == nil {
		return nil, errors.New("unable to create secret")
	}

	return storedSecret, nil
}

func (s *Storage) Delete(ctx context.Context, ID uuid.UUID) error {
	_, err := s.db.ExecContext(ctx, "delete from secrets where id = $1", ID)

	return err
}

func (s *Storage) FindById(ctx context.Context, ID uuid.UUID) (*domain.UserSecret, error) {
	var secret domain.UserSecret
	row := s.db.QueryRowContext(ctx, "select id, user_id, type, name, created_at, updated_at from secrets where id = $1", ID)

	err := row.Scan(&secret.ID, &secret.UserID, &secret.Type, &secret.Name, &secret.CreatedAt, &secret.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &secret, nil
}

func (s *Storage) GetAllForUser(ctx context.Context, userID uuid.UUID) ([]*domain.UserSecret, error) {
	result := make([]*domain.UserSecret, 0)
	err := s.db.SelectContext(ctx, &result, "select id, user_id, type, name, created_at, updated_at from secrets where user_id = $1 order by created_at desc", userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
