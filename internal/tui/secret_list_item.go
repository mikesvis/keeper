package tui

import (
	"fmt"
	"keeper/internal/domain"
	"strconv"
)

type SecretListItem struct {
	domain.UserSecret
}

func (i SecretListItem) Title() string       { return i.Name }
func (i SecretListItem) FilterValue() string { return i.Name }
func (i SecretListItem) Description() string {
	switch i.Type {
	case domain.UserSecretPasswordType:
		data := (*i.Data).(*domain.UserSecretDataPassword)

		return fmt.Sprintf("Login: %s Password: %s", data.Login, data.Password)

	case domain.UserSecretBankCardType:
		data := (*i.Data).(*domain.UserSecretDataBankCard)

		exp := strconv.FormatInt(data.Month, 10) + "/" + strconv.FormatInt(data.Year-2000, 10)

		return fmt.Sprintf("exp: %s cvv: %d", exp, data.Cvv)

	case domain.UserSecretTextType:
		data := (*i.Data).(*domain.UserSecretDataText)
		runes := []rune(data.Text)
		if len(runes) > 40 {
			runes = runes[:40]
			runes = append(runes, '.', '.', '.')
		}

		return string(runes)
	}

	return ""
}
