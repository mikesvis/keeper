package domain

import (
	"encoding/json"
)

type UserSecretDataBankCard struct {
	Number string `json:"number"`
	Month  int64  `json:"month"`
	Year   int64  `json:"year"`
	Cvv    int64  `json:"cvv"`
}

var _ UserSecretData = &UserSecretDataBankCard{}

func NewUserSecretBankCard(number string, month, year, cvv int64) *UserSecretDataBankCard {
	return &UserSecretDataBankCard{
		Number: number,
		Month:  month,
		Year:   year,
		Cvv:    cvv,
	}
}

func newUserSecretBankCardFromData(data []byte) (*UserSecretDataBankCard, error) {
	secretData := new(UserSecretDataBankCard)
	err := json.Unmarshal(data, secretData)
	if err != nil {
		return nil, err
	}
	return secretData, nil
}

func (d *UserSecretDataBankCard) GetType() UserSecretType {
	return UserSecretBankCardType
}

func (d *UserSecretDataBankCard) GetData() ([]byte, error) {
	return json.Marshal(d)
}
