package main

import (
	"fmt"
	"gov1/account"
	"gov1/files"
	"math/rand"

	"github.com/fatih/color"
)

var storage *account.Storage

func main() {
	fmt.Println("Привет, что ты хочешь сделать?")
	storage, err := account.NewStorage()
	if err != nil {
				RecoverThisShit(err.Error())
				return
			}
	for {
		input := getMenu()
		switch input {
		case 1:
			createAccount()
			Again("Аккаунт успешно создан, хочешь что-то еще?", nil)
		case 2:
			FindAccounts, err := storage.FindAccount(false)
			if err != nil {
				RecoverThisShit(err.Error())
				return
			}
			Again("Держи данные от аккаунта", FindAccounts)
		case 3:
			FindAccounts, err := storage.FindAccount(true)
			if err != nil {
				RecoverThisShit(err.Error())
				return
			}
			Again("Держи данные от аккаунта", FindAccounts)
		case 4:
			return
		default:
			RecoverThisShit("Повтори попытку, значение не распознанно")
		}
	}
}

func getMenu() int {
	var input int
	fmt.Println("1 если создать аккаунт \n2 если найти аккаунт \n3 если удалить аккаунт \n4 выйти")
	fmt.Println("Введи нужную тебе цифру")
	fmt.Scan(&input)
	fmt.Scanln()
	return input
}

func createAccount() {
	randomLogin, _ := promtData("Введи свой логин - ", 1)
	_, randomPasswordDlina := promtData("Пароль на сколько символов необходим? - ", 2)
	randomUrl := promtDataUrl()

	myLogin, err := account.NewLogin(randomLogin, randomPasswordDlina, randomUrl)
	if err != nil {
		RecoverThisShit(err.Error())
		return
	}

	storage, err := account.NewStorage()
	storage.AddAccount(*myLogin)
	data, err := storage.ToBytes()

	if err != nil {
		RecoverThisShit(err.Error())
		return
	}
	fmt.Println("Ваш логин -", myLogin.Login, "Ваш новый пароль -", myLogin.Password, "URL Сайта -", myLogin.Url, "Время создания аккаунта -", myLogin.CreatedAt)
	files.WriteInfo(data, "passwordBase.json")
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

func RecoverThisShit(text string) {
	color.Red(fmt.Sprint(text))
}

func Again(text string, sss *account.Storage) {
	color.Green(fmt.Sprint(text))
	if sss != nil {
		color.Cyan(fmt.Sprint(sss))
	}
}
