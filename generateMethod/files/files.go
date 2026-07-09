package files

import (
	"fmt"
	"gov1/output"
	"os"

	"github.com/fatih/color"
)

type JsonDb struct {
	FileName string
}

func NewJsonDb (name string) *JsonDb {
	return &JsonDb{
		FileName: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.FileName)
	if err != nil {
		fmt.Println("Возникла ошибка")
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) string {
	file, err := os.Create(db.FileName)
	if err != nil {
		output.PrintError("Возникла ошибка")
	}
	len, err := file.Write(content)
	if err != nil {
		output.PrintError("Возникла ошибка")
	}
	fmt.Println(len)
	defer file.Close()
	color.Red(fmt.Sprint("Файл успешно записан"))
	return "success"
}