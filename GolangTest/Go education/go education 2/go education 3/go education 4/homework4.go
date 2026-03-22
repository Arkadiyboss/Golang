package main

import "fmt"

// Person содержит имя и возраст человека.
type Person struct {
	Name string
	Age int
}

// TODO: реализуйте Stringer для Person

func main() {
	alice := Person{Name: "Alice", Age: 30}
	fmt.Printf("%s (%d)", alice.Name, alice.Age)
	// Ожидаемый вывод: Alice (30)
}

type Stringer interface {
    String() string
}