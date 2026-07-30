package limiter

import (
	"sync"
	_ "time"
)

// Limiter ограничивает количество событий до 5 в секунду.
type Limiter struct {
	mu sync.RWMutex
	tokens chan struct{}
}

// NewLimiter создаёт новый лимитер с ёмкостью 5 токенов.
func NewLimiter() *Limiter {
	// TODO: инициализировать канал токенов и запуск пополнения
	var lim Limiter
	lim.tokens = make(chan struct{}, 5)
	for i:=1; i <= 5; i++ {
		lim.tokens <- struct{}{}
	}
	return &lim
}

// Allow возвращает true, если событие разрешено в текущий момент.
func (l *Limiter) Allow() bool {
	// TODO: реализовать получение токена из канала
	<- l.tokens
	return false
}

// Stop останавливает лимитер.
func (l *Limiter) Stop() {
	// TODO: остановить пополнение токенов
	l.mu.RLock()
}