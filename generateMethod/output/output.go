package output

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(name any) {
	switch t := name.(type) {
	case string:
			color.Red(t)
	case int:
			color.Red("Код ошибки: %d", t)
	default: 
		color.Red("Неизвестный тип ошибки")
	}
	fmt.Println(name)
}

func sum[T int| float32| float64](a, b T)  T {

	return a + b
}

