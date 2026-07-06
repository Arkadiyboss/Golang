package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type UserLogin struct {
	Login     string `json:"jsonLogin" xml:"xmlLogin"`
	Password  string `json:"jsonPassword" xml:"xmlPassword"`
	Url       string `json:"jsonUrl" xml:"xmlUrl"`
	CreatedAt time.Time `json:"jsonCreatedAt" xml:"xmlCreatedAt"`
	Yanelox Yanelox
	}

	type Yanelox struct {
		Ya string `json:"jsonYa" xml:"xmlYa"`
		Lox string `json:"jsonLox" xml:"xmlLox"`
	}

func (acc *UserLogin) generateLoginPassword(n int) string {
	avaliableSymbols := []string{"a", "b", "c", "d", "1", "2", "3", "4", "5", "!", "@", "#", "$", "%"}
	var pass string
	for i := 0; i <= n-1; i++ {
		symbol := rand.Intn(cap(avaliableSymbols))
		pass = pass + avaliableSymbols[symbol]
	}
	return pass
}


func NewLogin(randomLogin string, randomPasswordDlina int, randomUrl string) (*UserLogin, error) {
	var accountPassword string
	tempUser := UserLogin{}
	if randomLogin == "" {
		return nil, errors.New("Отсутствует логин, повторите операцию еще раз")
	}
	if randomPasswordDlina == 0 {
		fmt.Println("Не задана длина пароля, пароль будет сгенерирован")
		accountPassword = tempUser.generateLoginPassword(10)
	} else {
		accountPassword = tempUser.generateLoginPassword(randomPasswordDlina)
	}
	_, err := url.ParseRequestURI(randomUrl)
	if err != nil {
		return nil, errors.New("Ошибка валидации данных, запустите программу еще раз")
	}
	return &UserLogin{
		Login:     randomLogin,
		Password:  accountPassword,
		Url:       randomUrl,
		CreatedAt: time.Now(),
		Yanelox: Yanelox{
			Ya: "yanelox100procentov",
			Lox: "etoNeYa",
		},
	}, nil

}