package main

import "fmt"

// Rectangle описывает прямоугольник.
type Rectangle struct {
	Width, Height float64
}

// Area должна вернуть площадь прямоугольника.
func (r Rectangle) Area() float64 {
	// TODO: реализуйте
	size := r.Width*r.Height
	return size
}
func main() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println(r.Area())
	// Ожидаемый вывод: 12
}