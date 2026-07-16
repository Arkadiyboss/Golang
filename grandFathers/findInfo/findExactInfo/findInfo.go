package findInfo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FindRespond struct {
	Value []any `json:"value"`
}

type FindRespondFathers struct {
	Id          int         `json:"id"`
	MedInfo_id  pgtype.Int4 `json:"medInfo_id"`
	Guardian_id pgtype.Int4 `json:"guardian_id"`
	Name        string      `json:"name"`
	SecondName  string      `json:"secondName"`
	Age         int         `json:"age"`
	PhoneNumber pgtype.Int4 `json:"phoneNumber"`
	Status      int         `json:"status"`
	CreatedAt   time.Time   `json:"createdAt"`
}

type FindRespondGuardianInfo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	SecondName  string `json:"secondName"`
	PhoneNumber int    `json:"phoneNumber"`
}

type FindRespondEmployee struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	SecondName  string `json:"secondName"`
	PhoneNumber int    `json:"phoneNumber"`
	Post        string `json:"post"`
	Change      int    `json:"change"`
}

type FindRespondSlots struct {
	Id             int       `json:"id"`
	Room           int       `json:"room"`
	Cost           int       `json:"cost"`
	DateIn         time.Time `json:"DateIn"`
	DateOut        time.Time `json:"dateOut"`
	GrandFatherId  int       `json:"grandFatherId"`
	EmployeeId     int       `json:"employeeId"`
	NursingHouseId int       `json:"nursingHouseId"`
}

func FindInfo(w http.ResponseWriter, r *http.Request, p *pgxpool.Pool, number int) {

	if r.Method != http.MethodPost {
		http.Error(w, "Метод отличается от POST", 400)
		return
	}

	var request map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Ошибка парсинга входящих данных", 500)
		return
	}

	switch number {
	case 1, 2, 3:
		if name, ok := request["name"].(string); !ok || name == "" {
			http.Error(w, "Body.name не прошел валидацию", http.StatusBadRequest)
			return
		}

		if secondName, ok := request["secondName"].(string); !ok || secondName == "" {
			http.Error(w, "Body.secondName не прошел валидацию", http.StatusBadRequest)
			return
		}
	case 4:
		if _, ok := request["dateIn"].(string); !ok {
			http.Error(w, "Body.dateIn не прошел валидацию", http.StatusBadRequest)
			return
		}
	}

	switch number {
	case 1:
		query := `SELECT * FROM "grandFather" WHERE name = $1 AND "secondName" = $2`
		var respond FindRespondFathers
		err = p.QueryRow(
			context.Background(),
			query,
			request["name"],
			request["secondName"],
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
		http.Error(w, "Пользователь не найден, FindRespondFather", 500)
		fmt.Println(err)
		return
	}
			SendRespond(w, r, respond)
	case 2:
		query := `SELECT * FROM "guardian" WHERE name = $1 AND "secondName" = $2`
		var respond FindRespondGuardianInfo
		err = p.QueryRow(
			context.Background(),
			query,
			request["name"],
			request["secondName"],
		).Scan(&respond.Id,
			&respond.Name,
			&respond.SecondName,
			&respond.PhoneNumber)
			if err != nil {
		http.Error(w, "Пользователь не найден, FindRespondGuardianInfo", 500)
		fmt.Println(err)
		return
	}
			SendRespond(w, r, respond)
	case 3:
		query := `SELECT * FROM "employee" WHERE name = $1 AND "secondName" = $2`
		var respond FindRespondEmployee
		err = p.QueryRow(
			context.Background(),
			query,
			request["name"],
			request["secondName"],
		).Scan(&respond.Id,
			&respond.Name,
			&respond.SecondName,
			&respond.PhoneNumber,
			&respond.Post,
			&respond.Change)
			if err != nil {
		http.Error(w, "Пользователь не найден, FindRespondEmployee", 500)
		fmt.Println(err)
		return
	}
			SendRespond(w, r, respond)
	case 4:

		dateInStr, ok := request["dateIn"].(string)
    if !ok || dateInStr == "" {
        http.Error(w, "Body.dateIn не прошел валидацию", http.StatusBadRequest)
        return
    }

		dateIn, err := time.Parse("2006-01-02", dateInStr)

    if err != nil {
        http.Error(w, "Body.dateIn должен быть в формате YYYY-MM-DD", http.StatusBadRequest)
        return
    }
		query := `SELECT * FROM "slots" WHERE "dateIn" <= $1`
		var respond FindRespondSlots
		err = p.QueryRow(
			context.Background(),
			query,
			dateIn,
		).Scan(&respond.Id,
			&respond.Room,
			&respond.Cost,
			&respond.DateIn,
			&respond.DateOut,
			&respond.GrandFatherId,
			&respond.EmployeeId,
			&respond.NursingHouseId)
			if err != nil {
		http.Error(w, "Слоты не найден, укажите другую дату заезда, FindRespondSlots", 500)
		fmt.Println(err)
		return
	}
			SendRespond(w, r, respond)
	}

}


func FindAllInfo(w http.ResponseWriter, r *http.Request, p *pgxpool.Pool, number int) {

	if r.Method != http.MethodGet {
		http.Error(w, "Метод отличается от GET", 400)
		return
	}

	switch number {
	case 1:
		query := `SELECT * FROM "grandFather"`

		rows, err := p.Query(
			context.Background(),
			query,
		)
		if err != nil {
			http.Error(w, "Ошибка при попытке поиска", 500)
			return
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
			results = append(results, respond)
			if err != nil {
		http.Error(w, "Пользователи не найдены, FindRespondFathers", 500)
		fmt.Println(err)
		return
	}
		}
		SendRespond(w, r, results)
	case 2:
		query := `SELECT * FROM "guardian"`

		rows, err := p.Query(
			context.Background(),
			query,
		)
		if err != nil {
			http.Error(w, "Ошибка при попытке поиска", 500)
			return
		}
		defer rows.Close()
		var results []FindRespondGuardianInfo

		for rows.Next() {
			var respond FindRespondGuardianInfo
			err := rows.Scan(
				&respond.Id,
				&respond.Name,
				&respond.SecondName,
				&respond.PhoneNumber,
			)
			results = append(results, respond)
			if err != nil {
		http.Error(w, "Пользователи не найдены, FindRespondGuardiansInfo", 500)
		fmt.Println(err)
		return
	}
		}
		SendRespond(w, r, results)
		case 3:
		query := `SELECT * FROM "employee"`

		rows, err := p.Query(
			context.Background(),
			query,
		)
		if err != nil {
			http.Error(w, "Ошибка при попытке поиска", 500)
			return
		}
		defer rows.Close()
		var results []FindRespondEmployee

		for rows.Next() {
			var respond FindRespondEmployee
			err := rows.Scan(
				&respond.Id,
				&respond.Name,
				&respond.SecondName,
				&respond.PhoneNumber,
				&respond.Post,
				&respond.Change,
			)
			results = append(results, respond)
			if err != nil {
		http.Error(w, "Пользователи не найдены, FindRespondEmployee", 500)
		fmt.Println(err)
		return
	}
			
		}
		SendRespond(w, r, results)
		case 4:
		query := `SELECT * FROM "slots"`

		rows, err := p.Query(
			context.Background(),
			query,
		)
		if err != nil {
			http.Error(w, "Ошибка при попытке поиска", 500)
			return
		}
		defer rows.Close()
		var results []FindRespondSlots

		for rows.Next() {
			var respond FindRespondSlots
			err := rows.Scan(
				&respond.Id,
				&respond.Room,
				&respond.Cost,
				&respond.DateIn,
				&respond.DateOut,
				&respond.GrandFatherId,
				&respond.EmployeeId,
				&respond.NursingHouseId,
			)
			results = append(results, respond)
			if err != nil {
		http.Error(w, "Пользователи не найдены, FindRespondSlots", 500)
		fmt.Println(err)
		return
	}
		}
		SendRespond(w, r, results)
	}
}

func SendRespond(w http.ResponseWriter, r *http.Request, respond any) {


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respond)
}