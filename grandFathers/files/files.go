package files

import (
	"fmt"
	"os"

)

func WriteLog(content []byte) bool {
	file, err := os.Create("logInfo.json")
	if err != nil {
		fmt.Println("Возникла ошибка")
		return false
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println("Возникла ошибка")
		return false
	}
	return true
}
