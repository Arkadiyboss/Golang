package main

import "fmt"

type Person struct {
    Name string
    Age  int
}
type Family struct {
    Name string
    Age  int
}
type human interface {
    nameAge()
}
func (p Person) String() string {
    return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}
func (p Person) nameAge()  {
}
func (f Family) nameAge()  {
}
func main() {
    alice := Person{Name: "Alice", Age: 30}
    boris := Family{Name: "Boris", Age: 20}
    fmt.Println(human(alice))
    fmt.Print(human(boris))
}