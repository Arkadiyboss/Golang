package main

import (
	"fmt"
	"net/http"
	"testServer/mainPage"
	userweather "testServer/userWeather"
)

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Это страница 'О нас'")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"status": "ok", "message": "API работает"}`)
}

func main() {
	http.HandleFunc("/city", userweather.UserWeather)
	http.HandleFunc("/", mainPage.Hello)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/api", apiHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
