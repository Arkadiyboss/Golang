package pipelinectx

import "context"

// Run строит конвейер из двух стадий: удвоение и суммирование.
// Конвейер должен останавливаться, если ctx отменён.
// Возвращает итоговую сумму и ошибку контекста при отмене.
func Run(ctx context.Context, nums []int) (int, error) {
    // Стадия 1: отправка чисел в канал
    in := make(chan int)
	out := make(chan int)
	var sum int
    go func() {
        defer close(in)
        for _, n := range nums {
            select {
            case <-ctx.Done():
                return 
            case in <- n:
            }
        }
    }()

    go func() {
        defer close(out)

        for n := range in {
            select {
            case <-ctx.Done():
                return 
            default:
                sum += n * 2 
            }
        }
        out <- sum 
    }()

    result := 0
    for n := range out {
        result += n
    }

    select {
    case <-ctx.Done():
        return result, ctx.Err()
    default:
        return result, nil
    }
}