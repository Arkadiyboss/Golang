package main

import "fmt"

func main() {
	m := map[string]string{
		"IamtheBoss": "For sure",
	}
	m["Iamthe"] = "123"
	fmt.Print(m)
}