package domain

import (
	"errors"
)

type UserSecretData interface {
	GetType() UserSecretType
	GetData() ([]byte, error)
}

func MakeUserSecretData(secretType UserSecretType, data []byte) (UserSecretData, error) {
	switch secretType {
	case UserSecretPasswordType:
		return newUserSecretPasswordFromData(data)
	case UserSecretBankCardType:
		return newUserSecretBankCardFromData(data)
	case UserSecretTextType:
		return newUserSecretTextFromData(data)
	case UserSecretFileType:
		return newUserSecretFileFromData(data)
	}

	return nil, errors.New("invalid secret type")
}
