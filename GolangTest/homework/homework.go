package main

import "fmt"

func homework() {
	const EUR = 100
	const USD = 80
	const RUB = 1
	var USDEUR = USD / EUR
	var RUBEUR = RUB / USDEUR
	fmt.Print(RUBEUR)
}
