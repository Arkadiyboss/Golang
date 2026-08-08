package main

import (
	"fmt"
)

func main() {
	fmt.Print("Привет, это обменик курса валют,\nвведите валюту которую хотите обменять\n")
	for {
	RUB, USD, EUR, operation, currency := userInput()
	end := calculate(RUB, USD, EUR, operation)
	fmt.Printf("И у нас получилось:%.2f %s",end ,currency )
	fmt.Println("\nХочешь рассчитать еще раз? (Y,N)")
	again := inputAgain()
	if again {
		continue
	}
	break
	}
}
func inputAgain() (bool) {
	for {
	var answer string
	fmt.Scan(&answer)
	switch answer {
	case "YES", "yes", "y":
		return true
	case "NO", "no", "n":
		return false
	default:
		fmt.Print("Повтори ввод (Y-да, N-нет)\n")
		continue
	}
}
}
func calculate(RUB float64, USD float64, EUR float64, operation int) (float64) {
	switch operation {
		case 1:
			operationEnd := RUB / USD
		return operationEnd
		case 2:
			operationEnd := RUB / EUR
		return operationEnd
		case 3:
			operationEnd := USD / RUB
		return operationEnd
		case 4:
			operationEnd := USD / EUR
		return operationEnd
		case 5:
			operationEnd := EUR / RUB
		return operationEnd
		case 6:
			operationEnd := EUR / USD
		return operationEnd
		default:
			fmt.Print("Не повезло")
		return 0
	}
}
func userInput() (float64, float64, float64, int, string) {
	fmt.Println("Какую валюту хотите обменять? (RUB,USD,EUR)")
	for {
	var RUB float64
	var USD float64
	var EUR float64
	var choice1 string
	var choice2 string
	fmt.Scan(&choice1)
	switch choice1 {
	case "RUB", "rub":
		for {
		fmt.Println("Курс валюты (RUB)")
		fmt.Scan(&RUB)
		fmt.Println("Какую валюту хотите обменять? (USD,EUR)")
		fmt.Scan(&choice2)
		switch choice2 {
	case "USD", "usd":
		fmt.Println("Курс валюты (USD)")
		fmt.Scan(&USD)
		return RUB,USD, 0, 1, "USD"
	case "EUR", "eur":
		fmt.Println("Курс валюты (EUR)")
		fmt.Scan(&EUR)
		return RUB, 0, EUR, 2, "EUR"
	default: 
		fmt.Print("Неккоректный ввод суммы, повторите попытку\n")
		continue
	}
		}
	case "USD", "usd":
		for {
		fmt.Println("Курс валюты (USD)")
		fmt.Scan(&USD)
		fmt.Println("Какую валюту хотите обменять? (RUB,EUR)")
		fmt.Scan(&choice2)
		switch choice2 {
		case "RUB", "rub":
			fmt.Println("Курс валюты (RUB)")
			fmt.Scan(&RUB)
			return RUB,USD, 0, 3, "RUB"
		case "EUR", "eur":
			fmt.Println("Курс валюты (EUR)")
			fmt.Scan(&EUR)
			return 0, USD, EUR, 4, "EUR"
		default: 
			fmt.Print("Неккоректный ввод суммы, повторите попытку")
			continue
		}
		}
	case "EUR", "eur":
		for {
		fmt.Println("Курс валюты (EUR)")
		fmt.Scan(&EUR)
		fmt.Println("Какую валюту хотите обменять? (RUB,USD)")
		fmt.Scan(&choice2)
		switch choice2 {
		case "RUB", "rub":
			fmt.Println("Курс валюты (RUB)")
			fmt.Scan(&RUB)
			return RUB, 0, EUR, 5, "RUB"
		case "USD", "usd":
			fmt.Println("Курс валюты (USD)")
			fmt.Scan(&USD)
			return 0, USD, EUR, 6, "USD"
		default: 
			fmt.Print("Неккоректный ввод суммы, повторите попытку")
			continue
		}
		}
	default:
		fmt.Println("Некорректный выбор валюты, повторите ввод\nКакую валюту хотите обменять? (RUB,USD,EUR)")
		continue
	}
	}
}