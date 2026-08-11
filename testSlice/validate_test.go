package validate

import (
	"testing"
)

func TestEmptyPassword(t *testing.T) {
	err := ValidatePassword("")

	if err != ErrPasswordEmpty {
		t.Fatalf("Ошибка не соответсвует ожидаемой")
	}
}

func TestShortPassword(t *testing.T) {

	err := ValidatePassword("1234567")

	if err != ErrPasswordTooShort {
		t.Fatalf("Ошибка не соответсвует ожидаемой")
	}
}

func TestLongPassword(t *testing.T) {

	err := ValidatePassword("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789123")

	if err != ErrPasswordTooLong {
		t.Fatalf("Ошибка не соответсвует ожидаемой")
	}
}

func TestCommonWord(t *testing.T) {

	err := ValidatePassword("12345678")

	if err != ErrPasswordCommonWord {
		t.Fatalf("Ошибка не соответсвует ожидаемой")
	}
}

func TestPasswordHasSpace(t *testing.T) {
	err := ValidatePassword("arkadiy ")

	if err != ErrPasswordHasSpace {
		t.Fatalf("Ошибка не соответсвует ожидаемой")
	}
}

func TestUpperLetter(t *testing.T) {
	err := ValidatePassword("arkadiy123")

	if err != ErrPasswordNoUpper {
		t.Fatalf("Ошибка не соответсвует ожидаемой")
	}
}


func TestLowerLetter(t *testing.T) {
	err := ValidatePassword("ARKADIY123")

	if err != ErrPasswordNoLower {
		t.Fatalf("Ошибка не соответсвует ожидаемой")
	}
}

func TestNoDigit(t *testing.T) {
	err := ValidatePassword("arkadiyYYYyyy")

	if err != ErrPasswordNoDigit {
		t.Fatalf("Ошибка не соответсвует ожидаемой")
	}
}