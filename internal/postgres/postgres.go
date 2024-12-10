package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"keeper/internal/config"
)

// NewPostgres Инициализация базы данных
func NewPostgres(c *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", c.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	return db, err
}
