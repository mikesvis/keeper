package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserSecretType string

const (
	UserSecretPasswordType = UserSecretType("password")
	UserSecretBankCardType = UserSecretType("bank_card")
	UserSecretTextType     = UserSecretType("text")
)

type UserSecret struct {
	ID        uuid.UUID
	UserID    uuid.UUID `db:"user_id"`
	Type      UserSecretType
	Name      string
	Data      *UserSecretData
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUserSecret(userID uuid.UUID, t UserSecretType, name string, d *UserSecretData) *UserSecret {
	return &UserSecret{
		ID:        uuid.New(),
		UserID:    userID,
		Type:      t,
		Name:      name,
		Data:      d,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
