package main

import "fmt"

func main() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")
}

func promtData(promt string) {
	fmt.Println(promt)
	var res string
	fmt.Scan(&res)
	return res
}