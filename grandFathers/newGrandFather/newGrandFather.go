package newgrandfather

import (
	"fmt"
	"net/http"
)

type Request struct {
	Name string `json:"name"`
	SecondName string `json:"secondName"`
	Age int `json:"age"`
}

type Respond struct {
	Id int `json:"id"`
}

func NewGrandFather(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Это страница 'О нас'")
}
