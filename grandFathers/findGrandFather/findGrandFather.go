package findgrandfather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FindRequest struct {
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
}

type FindRespond struct {
	Value []FindRespondFathers `json:"value"`
}

type FindRespondFathers struct {
	Id          int           `json:"id"`
	MedInfo_id  pgtype.Int4 `json:"medInfo_id"`
	Guardian_id pgtype.Int4 `json:"guardian_id"`
	Name        string        `json:"name"`
	SecondName  string        `json:"secondName"`
	Age         int           `json:"age"`
	PhoneNumber pgtype.Int4 `json:"phoneNumber"`
	Status      int        `json:"status"`
	CreatedAt   time.Time     `json:"createdAt"`
}

func FindGrandFather(w http.ResponseWriter, r *http.Request, p *pgxpool.Pool) {

	if r.Method != http.MethodPost {
		http.Error(w, "Метод отличается от POST", 400)
	}

	var request FindRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Ошибка парсинга входящих данных", 500)
	}

	if request.Name == "" {
		http.Error(w, "Body.name не прошел валидацию", 400)
	}

	if request.SecondName == "" {
		http.Error(w, "Body.secondName не прошел валидацию", 400)
	}

	query := `SELECT * FROM "grandFather" WHERE name = $1 AND "secondName" = $2`

	var respond FindRespondFathers

	err = p.QueryRow(
		context.Background(),
		query,
		request.Name,
		request.SecondName,
	).Scan(&respond.Id,
		&respond.MedInfo_id,
		&respond.Guardian_id,
		&respond.Name,
		&respond.SecondName,
		&respond.Age,
		&respond.PhoneNumber,
		&respond.Status,
		&respond.CreatedAt)

	if err != nil {
		http.Error(w, "Ошибка при попытке поиска", 500)
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respond)
}


func FindAllGrandFathers(w http.ResponseWriter, r *http.Request, p *pgxpool.Pool) {

	if r.Method != http.MethodGet {
		http.Error(w, "Метод отличается от GET", 400)
	}

	query := `SELECT * FROM "grandFather"`


rows, err := p.Query(
    context.Background(),
    query,
)
if err != nil {
    http.Error(w, "Ошибка при попытке поиска", 500)
}
defer rows.Close()

var results []FindRespondFathers

for rows.Next() {
    var respond FindRespondFathers
    err := rows.Scan(
        &respond.Id,
        &respond.MedInfo_id,
        &respond.Guardian_id,
        &respond.Name,
        &respond.SecondName,
        &respond.Age,
        &respond.PhoneNumber,
        &respond.Status,
        &respond.CreatedAt,
    )
    if err != nil {
        http.Error(w, "Ошибка при попытке поиска", 500)
    }
    results = append(results, respond)
}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
		
}