package mainPage

import (
	"encoding/json"
	"net/http"
)

type Hi struct {
	Title    string `json:"hello"`
	Subtitle string `json:"whatIsThis"`
}

func Hello(w http.ResponseWriter, r *http.Request) {

	response := &Hi{}

	response.Title = "Привет, это мой сервер по получению погоды"
	response.Subtitle = "Доступные ендпоинты /city?gorod="

	json.NewEncoder(w).Encode(response)
}
