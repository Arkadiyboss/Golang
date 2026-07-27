package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Run запускает продюсера, который отправляет числа от 1 до 10, и консюмера,
// который выводит их в writer. Используйте небуферизованный канал и ожидание
// завершения горутин.
func main() {
	ch := make(chan int)
	wg.Add(2)
	go producer(ch)
	go writer(ch)
	wg.Wait()
	close(ch)

	// TODO: реализовать продюсер и консюмер
}

func producer(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	defer wg.Done()
}

func writer(ch chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
	defer wg.Done()
}
