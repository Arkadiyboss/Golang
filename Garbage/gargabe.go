package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{Name: "Bob", Age: 30}
	fmt.Println(user)
	fmt.Println(user.Name, user.Age)
}