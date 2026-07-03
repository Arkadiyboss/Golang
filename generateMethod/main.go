package main

import (
	"fmt"
	"gov1/account"
	"gov1/files"
	"math/rand"
	"github.com/fatih/color"
)

func main() {
	randomLogin, _ := promtData("Введи свой логин - ", 1)
	_, randomPasswordDlina := promtData("Пароль на сколько символов необходим? - ", 2)
	randomUrl := promtDataUrl()

	myLogin, err := account.NewLogin(randomLogin, randomPasswordDlina, randomUrl)
	if err != nil {
		recoverThisShit(err.Error())
		return
	}
	content := myLogin.Login + myLogin.Password + myLogin.Url
	fmt.Println("Ваш логин -", myLogin.Login, "Ваш новый пароль -", myLogin.Password, "URL Сайта -", myLogin.Url, "Время создания аккаунта -", myLogin.CreatedAt)
	files.WriteInfo(content, myLogin.Login)
}

func promtData(promt string, number int) (string, int) {
	color.Blue(promt)
	var login string
	var passNumber int
	if number == 1 {
		fmt.Scanln(&login)
		return login, passNumber
	}
	if number == 2 {
		fmt.Scanln(&passNumber)
		return login, passNumber
	}
	return login, passNumber
}

func promtDataUrl() string {
	avaliableSites := []string{"https://google.com", "https://yandex.ru", "https://hltv.org"}
	siteChoice := rand.Intn(cap(avaliableSites))
	site := avaliableSites[siteChoice]
	return site
}

func recoverThisShit(error string) {
	fmt.Println(error)
}
