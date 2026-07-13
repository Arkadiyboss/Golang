package newgrandfather

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Request struct {
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
	Age        int    `json:"age"`
	Issue      string `json:"issue"`
}

type Respond struct {
	Id int `json:"id"`
}

func NewGrandFather(w http.ResponseWriter, r *http.Request, p *pgxpool.Pool) {

	if r.Method != http.MethodPost {
		http.Error(w, "Метод отличается от POST", 400)
	}

	var request Request

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Ошибка парсинга входящих данных", 500)
	}

	if request.Name == "" {
		http.Error(w, "Body.name не прошел валидацию", 400)
	}
	if request.SecondName == "" {
		http.Error(w, "Body не прошел валидацию", 400)
	}
	if request.Age < 1 {
		http.Error(w, "Body.age не прошел валидацию", 400)
	}

	query := `INSERT INTO "grandFather" (name, "secondName", age) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err = p.QueryRow(
		context.Background(),
		query,
		request.Name,
		request.SecondName,
		request.Age,
	).Scan(&id)

	if err != nil {
		http.Error(w, "Ошибка записи в БД: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := Respond{Id: id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
