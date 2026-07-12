package files

import (
	"fmt"
	"os"
)

func Read(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Возникла ошибка")
		return nil, err
	}
	return data, nil
}

func Write(name string, content []byte) (bool, error) {
	file, err := os.Create(name)
	
	if err != nil {
		return false, err
	}
	defer file.Close()
	len, err := file.Write(content)
	if err != nil {
		return false, err
	}
	fmt.Println(len)
	return true, nil
}
