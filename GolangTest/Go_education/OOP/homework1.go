// Задача: Реализовать наследование (встраивание) в Go.
// Здесь есть базовый тип Animal с методом Speak.
// Необходимо создать тип Dog, который встраивает Animal и добавляет метод Bark.
// Затем в main вызвать оба метода.

package main

import "fmt"

// Animal – базовый тип
type Animal struct {
    Name string
}

// Speak выводит звук, который издает животное.
func (a Animal) Speak() {
    fmt.Printf("%s издает какой-то звук\n", a.Name)
}
func (a Animal) Bark() {
	fmt.Printf("%s гавкает\n", a.Name)
}
// TODO: Создайте тип Dog, который встраивает Animal.
//       Добавьте метод Bark(), который выводит "<Name> гавкает".
//       В функции main создайте экземпляр Dog с именем "Rex"
//       и вызовите методы Speak() и Bark() соответственно.

func main() {
    Dog := Animal{Name: "Rex"}
	Dog.Speak()
	Dog.Bark()
}