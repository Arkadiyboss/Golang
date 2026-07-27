package fibonacci

// Fib возвращает канал, из которого можно читать первые n чисел Фибоначчи.
func Fib(n int) <-chan int {
	ch := make(chan int, n)
	x,y := 0,1
	ch <- x
	for i:=1; i < n; i++ {
		ch <- y
		x,y= y, x+y
	}

	close(ch)
	// TODO: отправить последовательность Фибоначчи в канал
	return ch
}