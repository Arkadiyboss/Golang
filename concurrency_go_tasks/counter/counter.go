package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mu sync.Mutex

var vStruct struct {
	v  int
	mu sync.Mutex
}
var megaNumberStruct struct {
	megaNumber int
	mu         sync.Mutex
}

func main() {
	ch := make(chan int)

	wg.Add(2)
	go worker(ch, 1)
	go worker(ch, 2)

	ch <- 0

	wg.Wait()

	fmt.Println("Дело сделано")
}

func worker(ch chan int, id int) {
	defer wg.Done()

	for {
		now := time.Now()
		_, ok := <-ch
		if !ok {
			return
		}

		if vStruct.v == 1000 {
			close(ch)
			return
		}
		time.Sleep(1 * time.Second)
		megaNumberStruct.mu.Lock()
		megaNumberStruct.megaNumber = megaNumberStruct.megaNumber + (id * 100)

		fmt.Println("Меганомер: ", megaNumberStruct.megaNumber, "Я горутина: ", id, "Время с начала цикла: ", time.Since(now))

		vStruct.v++
		fmt.Printf("Воркер %d: %d\n", id, vStruct.v)
		ch <- vStruct.v
		megaNumberStruct.mu.Unlock()
	}
}
