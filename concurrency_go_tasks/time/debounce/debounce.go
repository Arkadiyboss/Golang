package debounce

import "time"

// Debounce принимает значения и отдаёт только последнее после паузы d.
func Debounce(d time.Duration, in <-chan int) <-chan int {
    out := make(chan int)

    go func() {
        defer close(out)

        var lastValue int
        var hasValue bool
        timer := time.NewTimer(d)
        timer.Stop() 

        for {
            select {
            case v, ok := <-in:
                if !ok {
                    if hasValue {
                        out <- lastValue
                    }
                    return
                }
                lastValue = v
                hasValue = true
                if !timer.Stop() {
                    <-timer.C
                }
                timer.Reset(d)

            case <-timer.C:
                if hasValue {
                    out <- lastValue
                    hasValue = false
                }
            }
        }
    }()

    return out
}