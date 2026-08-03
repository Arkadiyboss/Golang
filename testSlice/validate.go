// Package validator содержит функции валидации пользовательских данных.
package validate

import (
	"errors"
	"strings"
	"unicode"
)

// Ошибки валидации пароля.
var (
	ErrPasswordEmpty       = errors.New("password is empty")
	ErrPasswordTooShort    = errors.New("password is too short")
	ErrPasswordTooLong     = errors.New("password is too long")
	ErrPasswordNoUpper     = errors.New("password must contain uppercase letter")
	ErrPasswordNoLower     = errors.New("password must contain lowercase letter")
	ErrPasswordNoDigit     = errors.New("password must contain digit")
	ErrPasswordHasSpace    = errors.New("password must not contain spaces")
	ErrPasswordCommonWord  = errors.New("password is too common")
)

const (
	minPasswordLen = 8
	maxPasswordLen = 64
)

// commonPasswords — список запрещённых очевидных паролей (регистр не учитывается).
var commonPasswords = map[string]struct{}{
	"password": {},
	"qwerty":   {},
	"12345678": {},
	"admin":    {},
	"letmein":  {},
}

// ValidatePassword проверяет пароль по набору правил:
//   - не пустой и не состоит только из пробелов;
//   - длина от 8 до 64 символов включительно;
//   - содержит хотя бы одну заглавную букву (любого алфавита);
//   - содержит хотя бы одну строчную букву;
//   - содержит хотя бы одну цифру;
//   - не содержит пробельных символов;
//   - не входит в список очевидных паролей (регистр не учитывается).
//
// Возвращает первую найденную ошибку или nil, если пароль валиден.
// Порядок проверок: пустота → длина → запрещённое слово → состав символов.
func ValidatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return ErrPasswordEmpty
	}

	if len(password) < minPasswordLen {
		return ErrPasswordTooShort
	}
	if len(password) > maxPasswordLen {
		return ErrPasswordTooLong
	}

	if _, ok := commonPasswords[strings.ToLower(password)]; ok {
		return ErrPasswordCommonWord
	}

	var hasUpper, hasLower, hasDigit bool
	for _, r := range password {
		switch {
		case unicode.IsSpace(r):
			return ErrPasswordHasSpace
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}

	if !hasUpper {
		return ErrPasswordNoUpper
	}
	if !hasLower {
		return ErrPasswordNoLower
	}
	if !hasDigit {
		return ErrPasswordNoDigit
	}

	return nil
}