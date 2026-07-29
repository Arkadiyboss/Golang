package pipeline



// Run строит конвейер из трёх стадий: квадрат, умножение на 2 и суммирование.
func Run(nums []int) int {
	// TODO: реализовать конвейер обработки чисел
	sum := 0
	for _, n := range nums {
		sum += (n * n) * 2
	}
	return sum
}
