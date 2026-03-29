package main

import "fmt"

// FilterEven возвращает новый срез, содержащий
// только чётные числа из nums в том же порядке.
func FilterEven(in []int) []int {
	// TODO: реализуйте
	even := []int{}
	for i := 0; i < len(in); i++ {
		if in[i]%2 == 0 {
			even = append(even, in[i])
		}
	}
	return even
}
func main() {
	in := []int{5, 4, 9, 2, 7, 6}
	fmt.Println(FilterEven(in))
	// Ожидаемый вывод: [4 2 6]
}