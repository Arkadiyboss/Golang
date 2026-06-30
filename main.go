package main

import (
	"fmt"
	"math/rand"
)

type userLogin struct {
	login string
	password string
	url string
} 

func main() {

	randomPasswordDlina := promtDataPassword("Пароль на сколько символов необходим? - ")
	randomPassword := generatePassword(randomPasswordDlina)
	fmt.Println(randomPassword)
	login := promtData("Введите ваш логин: ")
	password := promtData("Введите ваш пароль: ")
	url := promtData("Введите URL: ")

	account := userLogin{
		login: login,
		password: password,
		url: url,
	}

	fmt.Println(account)
}

func promtDataPassword(promt string) int {
	fmt.Print(promt)
	var res int
	fmt.Scan(&res)
	return res
} 

func promtData(promt string) string {
	fmt.Print(promt)
	var res string
	fmt.Scan(&res)
	return res
} 

func generatePassword (n int) string{
	var passCode int
	passNumber := []int{}
	var pass string
	for i := 0; i < n; i = i + 1{
		passCode = rand.Intn(20000)
		passNumber = append(passNumber, passCode)
	}
	for g := 0; g <= len(passNumber)-1; g = g + 1{
		pass = pass + string(passNumber[g])
	}
	return pass
}