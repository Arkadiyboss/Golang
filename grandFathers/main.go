package main

import (
	"fmt"
	"grandFather/database"
	findgrandfather "grandFather/findGrandFather"
	newgrandfather "grandFather/newGrandFather"
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

	pool, err := database.ConnectDB("postgres://arkadiy:123@localhost:5432/grandFatherBase?sslmode=disable")


	if err != nil {
		fmt.Println("Не получается подключиться к БД", err)
		return
	}

	defer pool.Close()

	go database.PingDB(pool)

	 http.HandleFunc("/newGrandFather", func(w http.ResponseWriter, r *http.Request) {
        newgrandfather.NewGrandFather(w, r, pool) 
    })
	http.HandleFunc("/findGrandFather", func(w http.ResponseWriter, r *http.Request) {
        findgrandfather.FindGrandFather(w, r, pool)
    })
	http.HandleFunc("/allGrandFathers", func(w http.ResponseWriter, r *http.Request) {
        findgrandfather.FindAllGrandFathers(w, r, pool)
    })
	http.HandleFunc("/deleteGrandFather", DeleteGrandFather)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
