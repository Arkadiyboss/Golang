package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recover ", r)
		}
	}()
	fmt.Print("Рассчет индекса массы тела \n")
	for {
		Height, Weight := getUserInput()
		IMT, err := calculateIMT(Height, Weight)
		 if err != nil {
			//fmt.Println("Неккоректный ввод")
			//continue
			panic("Неккоректный ввод данных")
		}
		switch {
		case IMT < 16:
			fmt.Println("У вас сильный недостаток веса")
		case IMT < 18.5:
			fmt.Println("У вас недостаток веса")
		case IMT < 25:
			fmt.Println("У вас нормальный вес")
		case IMT < 30:
			fmt.Println("У вас избыточный вес")
		default:
			fmt.Println("У вас степень ожирения")
		}
		outputResult(IMT)
		if !retry() {
			break
		}
	}
}
func retry() bool {
	var retryCommand string
	fmt.Println("Хочешь рассчитать еще раз?")
	fmt.Println("Y - да, N - нет: ")
	fmt.Scan(&retryCommand)
	switch retryCommand {
	case "Y", "y", "YES", "yes":
		return true
	case "N", "n", "NO", "no":
		fmt.Println("Давай удачи тогда")
		return false
	default:
		fmt.Println("Давай удачи тогда")
		return false
	}
}
func outputResult(IMT float64) {
	result := fmt.Sprintf("Индекс массы тела: %.2f", IMT)
	fmt.Print(result)
}
func calculateIMT(Height float64, Weight float64) (float64, error) {
	if Height <= 0 || Weight <= 0 {
		return 0, errors.New("NO_PARAMS")
	}
	IMT := Weight / math.Pow(Height, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var Height float64
	var Weight float64
	fmt.Print("ВВедите ваш рост: ")
	fmt.Scan(&Height)
	fmt.Print("ВВедите ваш вес: ")
	fmt.Scan(&Weight)
	return Height, Weight
}
