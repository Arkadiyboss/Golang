package files

import (
	"fmt"
	"os"
)


func ReadInfo(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Возникла ошибка")
	}
	return string(data)
}

func WriteInfo(content []byte, name string) string {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Возникла ошибка", err)
	}
	len, err := file.Write(content)
	if err != nil {
		fmt.Println("Возникла ошибка", err)
	}
	fmt.Println(len)
	defer file.Close()
	return "success"
}