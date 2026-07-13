package main

import (
	"fmt"
	"net/http"
)

func FindGrandFather(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status": "ok", "message": "API работает"}`)
}

func AllGrandFathers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Это страница 'О нас'")
}

func DeleteGrandFather(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Это страница 'О нас'")
}

func main() {
	http.HandleFunc("/newGrandFather", )
	http.HandleFunc("/findGrandFather", FindGrandFather)
	http.HandleFunc("/allGrandFathers", AllGrandFathers)
	http.HandleFunc("/deleteGrandFather", DeleteGrandFather)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
