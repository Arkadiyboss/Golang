package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(tasks <-chan int) {
	for task := range tasks {

		fmt.Printf("Обрабатываю задачу %d\n", task)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	tasks := make(chan int)

	go worker(tasks)

	for i := 1; i <= 3; i++ {
		tasks <- i
	}
	close(tasks)
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("Горутин в работе: %d\n", runtime.NumGoroutine())
}
