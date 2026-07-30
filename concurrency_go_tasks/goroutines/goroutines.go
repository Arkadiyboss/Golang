package goroutines
import (
	"fmt"
	"io"
	"os"
	"sync"
)

var wg sync.WaitGroup

// PingPong должен запускать две горутины "ping" и "pong",
// которые поочередно выводят строки пять раз каждая.
// Реализуйте синхронизацию через каналы и ожидание завершения.
func PingPong(w io.Writer) {
	ch := make(chan string)

	wg.Add(2)
	go ping("ping", ch)
	go pong("pong", ch)

	wg.Wait()
	close(ch)
	// TODO: реализовать обмен сообщениями между горутинами
}

func ping(text string, ch chan string) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- text
		fmt.Println(<-ch)
	}

}

func pong(text string, ch chan string) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
		ch <- text
	}

}

func main() {
	PingPong(os.Stdout)
}