package pool

import _ "sync"

// RunPool обрабатывает задачи параллельно в заданном количестве воркеров
// и возвращает сумму результатов.
func RunPool(jobs []int, workers int) int {
	// TODO: реализовать пул воркеров и сбор результатов
	var r int
	ch := make(chan int)
	for i:=1; i <= workers; i++{
		go Run(jobs, ch)
	}
	r = <- ch
	return r
}

func Run(jobs []int, ch chan int) {
	var value int
	for _, valueJobs := range jobs {
		value = value + valueJobs
	}
	ch <- value
}