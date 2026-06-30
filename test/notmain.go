package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomPasswordDlina := promtDataPassword("Пароль на сколько символов необходим? - ")
	randomPassword := generatePassword(randomPasswordDlina)
	fmt.Println(randomPassword)
}

func promtDataPassword(promt string) int {
	fmt.Print(promt)
	var res int
	fmt.Scan(&res)
	return res
}

func generatePassword(n int) string {
	avaliableSymbols := []string{"a", "b", "c", "d", "1", "2", "3", "4", "5", "!", "@", "#", "$", "%"}
	var pass string
	for i := 0; i <= n; i++ {
		symbol := rand.Intn(cap(avaliableSymbols))
		pass = pass + avaliableSymbols[symbol]
	}
	return pass
}