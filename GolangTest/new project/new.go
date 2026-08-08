package main

import (
	"fmt"
)

func main() {
	transaction := []float64{}
	fmt.Println("Привет, введи свою транзакцию")
	var userResult float64 = 0
	var userAverageTrans float64 = 0
	var count int
	var maxTrans float64 = 0
	for {
	var again bool
	userTransaction := input()
	transaction = append(transaction, userTransaction)
	count, again = inputChoice(count)
	maxTrans = maxTransaction(transaction)
	userResult = sum(userResult, userTransaction)
	userAverageTrans = userAverage(userResult, count)

		if again {
			continue
		}
		break
	}
	fmt.Print("Вот ваш баланс: ")
	fmt.Println(userResult)
	fmt.Print("Вот ваша средняя сумма транзакций: ")
	fmt.Println(userAverageTrans)
	fmt.Print("Ваша самая большая транзакция: ")
	fmt.Println(maxTrans)
	fmt.Println("Вот все ваши транзакции")
	fmt.Println(transaction)
}
func input() (float64) {
	var trans float64
	fmt.Scan(&trans)
	return trans
}
func inputChoice(count int) (int, bool) {
	var choice string
	fmt.Println("Хочешь записать еще транзакцию?")	
	fmt.Scan(&choice)
	switch choice {
	case "YES", "Y", "y", "yes":
		fmt.Println("Введи следующую транзакцию")
		return count + 1, true
	case "NO", "no", "N", "n":
		return count, false
	default:
		fmt.Println("Повторите попытку")
		return inputChoice(count)
	}
}	
func sum(userResult float64,userTransaction float64) (float64) {
	result := userResult + userTransaction
	return result
}
func userAverage(userSum float64, userCount int) float64 {
	return userSum/float64(userCount)
}
func maxTransaction(transactions []float64) float64 {
    if len(transactions) == 0 {
        return 0
    }
    max := transactions[0]
    for i := 1; i < len(transactions); i++ {
        if transactions[i] > max {
            max = transactions[i]
        }
    }
    return max
}