package domain

import (
	"encoding/json"
)

type UserSecretDataPassword struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var _ UserSecretData = &UserSecretDataPassword{}

func NewUserSecretPassword(login, password string) *UserSecretDataPassword {
	return &UserSecretDataPassword{
		Login:    login,
		Password: password,
	}
}

func newUserSecretPasswordFromData(data []byte) (*UserSecretDataPassword, error) {
	secretData := new(UserSecretDataPassword)
	err := json.Unmarshal(data, secretData)
	if err != nil {
		return nil, err
	}
	return secretData, nil
}

func (d *UserSecretDataPassword) GetType() UserSecretType {
	return UserSecretPasswordType
}

func (d *UserSecretDataPassword) GetData() ([]byte, error) {
	return json.Marshal(d)
}
