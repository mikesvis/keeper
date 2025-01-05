package validators

import (
	"errors"
	"unicode/utf8"
)

var ErrEmptyLogin = errors.New(`login must not be empty`)
var ErrShortLogin = errors.New(`login must be at least 6 characters long`)
var ErrLongLogin = errors.New(`login must be shorter than 255 characters`)
var ErrEmptyPassword = errors.New(`password must not be empty`)
var ErrShortPassword = errors.New(`password must be at least 6 characters long`)
var ErrLongPassword = errors.New(`password must be shorter than 255 characters`)

func LoginValidator(login string) error {
	if login == "" {
		return ErrEmptyLogin
	}

	if utf8.RuneCountInString(login) < 3 {
		return ErrShortLogin
	}

	if utf8.RuneCountInString(login) > 255 {
		return ErrLongLogin
	}

	return nil
}

func PasswordValidator(password string) error {
	if password == "" {
		return ErrEmptyPassword
	}

	if utf8.RuneCountInString(password) < 6 {
		return ErrShortPassword
	}

	if utf8.RuneCountInString(password) > 255 {
		return ErrLongPassword
	}

	return nil
}
