
// Задача: Реализовать полиморфизм через интерфейсы в Go.
// Нужно определить интерфейс Shape с методом Area() float64.
// Создать два типа: Rectangle и Circle, которые реализуют Shape.
// В main пройтись по срезу Shape и вычислить общую площадь.

package main

import (
    "fmt"
)

// Shape – интерфейс для геометрических фигур
type Shape interface {
    Area() float64
}

// TODO: Определите тип Rectangle с полями Width и Height
//       и реализуйте для него метод Area().

// TODO: Определите тип Circle с полем Radius
//       и реализуйте для него метод Area().

func main() {
    shapes := []Shape{
        // Инициализируйте несколько Rectangle и Circle, например:
        // Rectangle{Width: 3, Height: 4},
        // Circle{Radius: 2.5},
    }

    var total float64
    for _, s := range shapes {
        fmt.Println("Площадь:", s.Area())
        total += s.Area()
    }
    fmt.Println("Общая площадь:", total)
}