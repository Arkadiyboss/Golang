package main

import (
	"fmt"
)

// SumInts возвращает сумму всех чисел nums.
// Если срез пуст, должна вернуться 0.
func SumInts(nums []int) int {
	// TODO: реализуйте
	var amount int
	amount = 0
	for i := 0; i < len(nums); i++ {
		amount += nums[i]
	}
	return amount
}
func main() {
	fmt.Println(SumInts([]int{1, 2, 3}))
	fmt.Println(SumInts([]int{}))
	// Ожидаемый вывод:
	// 6
	// 0
}