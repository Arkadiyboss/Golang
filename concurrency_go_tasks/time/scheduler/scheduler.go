package scheduler

import "time"

// Every запускает f каждые d и возвращает функцию для остановки.
func Every(d time.Duration, f func()) (stop func()) {
	// TODO: периодический вызов функции с возможностью остановки
	stopChannel := make(chan struct{})
	ticker := time.NewTicker(d)

	go func() {
		defer ticker.Stop()
        for {
            select {
            case <-ticker.C:
                f()
            case <-stopChannel:
                return
            }
        }
    }()

	stop = func() {
        close(stopChannel)
    }
	return stop
}
