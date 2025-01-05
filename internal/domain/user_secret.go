package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type UserSecretType string

const (
	UserSecretPasswordType = UserSecretType("password")
	UserSecretBankCardType = UserSecretType("bank_card")
	UserSecretTextType     = UserSecretType("text")
	UserSecretFileType     = UserSecretType("file")
)

type UserSecret struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Type      UserSecretType
	Name      string
	Data      *UserSecretData
	CreatedAt time.Time
	UpdatedAt time.Time
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

func GetSecretTypeByString(t string) (UserSecretType, error) {
	switch t {
	case "password", "bank_card", "text", "file":
		return UserSecretType(t), nil
	}

	return "", errors.New("unknown secret type")
}
