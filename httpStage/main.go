package main

import (
	"fmt"
	"http/geo"
	"http/weather"
	"time"
)

type Location struct {
	City string `json:"city"`
}

func main() {

	fmt.Println("Привет эта программа поможет тебе узнать погоду в выбранном городе")

	cityData := InputData[string]("Введите ваш город, на английском языке")

	formatNumber := InputData[int]("Введите формат отображения, введи от 1 до 4")

	userTemper, err := weather.Temper(geo.Location{City: cityData}, formatNumber)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Город пользователя ", userTemper.Location, "\nПогода в городе пользователя ", userTemper.Weather, "Дата и время обновления ", userTemper.UpdateTime)

	time.Sleep(5 * time.Second)
}

func InputData[T string | int](text string) T {
	var info T
	fmt.Println(text)
	fmt.Scan(&info)
	fmt.Scanln()
	return info
}
