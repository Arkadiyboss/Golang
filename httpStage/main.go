package main

import (
	"fmt"
	"http/geo"
	"http/weather"
)

type Location struct {
	City string `json:"city"`
}

func main() {

	fmt.Println("Привет эта программа поможет тебе узнать погоду в выбранном городе")

	for {

		cityData := InputData[string]("Введите ваш город, на английском языке, или 1 если хочешь чтобы мы определили")

		var getCity *geo.Location

		if cityData == "1" {
			getCity, _ = geo.GetLocation(cityData)
		}

		if cityData == "" {
			InputData[string]("Повтори ввод, город не распознан")
		} else {
			check := geo.CityValidate(cityData)
			if check == true {
				InputData[string]("Повтори ввод, город не распознан44444444444")
			}
		}

		formatNumber := InputData[int]("Введите формат отображения, введи от 1 до 4")

		userTemper, err := weather.Temper(geo.Location{City: cityData}, formatNumber)

		if err != nil {
			fmt.Println(err)
			return
		}

		if cityData == "" {
			fmt.Println("Город пользователя ", string(userTemper.Location.City), "\nПогода в городе пользователя ", userTemper.Weather, "Дата и время обновления ", userTemper.UpdateTime)
		} else {
			fmt.Println("Город пользователя ", string(getCity.City), "\nПогода в городе пользователя ", userTemper.Weather, "Дата и время обновления ", userTemper.UpdateTime)

		}

		again := InputData[int]("Хочешь еще раз? 1 - Да, 2 - Нет")

		if again == 2 {
			return
		}
	}
}

func InputData[T string | int](text string) T {
	var info T
	fmt.Println(text)
	fmt.Scan(&info)
	fmt.Scanln()
	return info
}
