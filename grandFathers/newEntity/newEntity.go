package newEntity

import (
	"context"
	"encoding/json"
	"fmt"
	"grandFather/auth"
	duplicateMethod "grandFather/duplicate"
	"grandFather/respondMethod"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Respond struct {
	Id int `json:"id"`
}

func NewEntity(w http.ResponseWriter, r *http.Request, p *pgxpool.Pool, number int) {

	if r.Method != http.MethodPost {
		http.Error(w, "Метод отличается от POST", 400)
		return
	}

	authHeader := r.Header.Get("Authorization")

	IsTokenOk, err := auth.CheckToken(authHeader, p)

	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Токен не найден", http.StatusNotFound)
			return
		}
		http.Error(w, "Ошибка проверки токена: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if IsTokenOk != true {
		http.Error(w, "Ошибка, токен не действителен", 400)
		return
	}

	var request map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Ошибка парсинга входящих данных", 500)
		return
	}

	switch number {
	case 1:
		if name, ok := request["name"].(string); !ok || name == "" {
			http.Error(w, "Body.name не прошел валидацию", 400)
			return
		}
		if secondName, ok := request["secondName"].(string); !ok || secondName == "" {
			http.Error(w, "Body.secondName не прошел валидацию", 400)
			return
		}

		FloatAge, ok := request["age"].(float64)
		if !ok || FloatAge < 1 {
			http.Error(w, "Body.age не прошел валидацию", 400)
			return
		}

		age := int(FloatAge)

		duplicateMethod.Duplicate(w, p, request, number)
		DbDuplicate, DuplicateId := duplicateMethod.Duplicate(w, p, request, number)
		DuplicateText := "Ошибка создания сущности, уже существует с ID: " + fmt.Sprint(DuplicateId)

		if DbDuplicate {
			http.Error(w, DuplicateText, 400)
			return
		}

		query := `INSERT INTO "resident" (name, "secondName", age) VALUES ($1, $2, $3) RETURNING id`

		var Respond Respond
		err = p.QueryRow(
			context.Background(),
			query,
			request["name"].(string),
			request["secondName"].(string),
			age,
		).Scan(&Respond.Id)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Ошибка записи в БД", 500)
			return
		}

		respondMethod.SendRespond(w, r, Respond)

	case 2:
		if name, ok := request["name"].(string); !ok || name == "" {
			http.Error(w, "Body.name не прошел валидацию", 400)
			return
		}
		if secondName, ok := request["secondName"].(string); !ok || secondName == "" {
			http.Error(w, "Body.secondName не прошел валидацию", 400)
			return
		}

		phoneNumber, ok := request["phoneNumber"].(float64)
		if !ok {
			http.Error(w, "Body.phoneNumber не прошел валидацию", 400)
			return
		}

		intPhoneNumber := int(phoneNumber)

		DbDuplicate, DuplicateId := duplicateMethod.Duplicate(w, p, request, number)
		DuplicateText := "Ошибка создания сущности, уже существует с ID: " + fmt.Sprint(DuplicateId)

		if DbDuplicate {
			http.Error(w, DuplicateText, 400)
			return
		}

		query := `INSERT INTO "guardian" (name, "secondName", "phoneNumber") VALUES ($1, $2, $3) RETURNING id`

		var Respond Respond
		err = p.QueryRow(
			context.Background(),
			query,
			request["name"].(string),
			request["secondName"].(string),
			intPhoneNumber,
		).Scan(&Respond.Id)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Ошибка записи в БД", 500)
			return
		}

		respondMethod.SendRespond(w, r, Respond)

	case 3:
		if name, ok := request["name"].(string); !ok || name == "" {
			http.Error(w, "Body.name не прошел валидацию", 400)
			return
		}
		if secondName, ok := request["secondName"].(string); !ok || secondName == "" {
			http.Error(w, "Body.secondName не прошел валидацию", 400)
			return
		}

		phoneNumber, ok := request["phoneNumber"].(float64)
		if !ok {
			http.Error(w, "Body.phoneNumber не прошел валидацию", 400)
			return
		}

		if post, ok := request["post"].(string); !ok || post == "" {
			http.Error(w, "Body.post не прошел валидацию", 400)
			return
		}

		intPhoneNumber := int(phoneNumber)

		DbDuplicate, DuplicateId := duplicateMethod.Duplicate(w, p, request, number)
		DuplicateText := "Ошибка создания сущности, уже существует с ID: " + fmt.Sprint(DuplicateId)

		if DbDuplicate {
			http.Error(w, DuplicateText, 400)
			return
		}

		query := `INSERT INTO "employee" (name, "secondName", "phoneNumber", post) VALUES ($1, $2, $3, $4) RETURNING id`

		var Respond Respond
		err = p.QueryRow(
			context.Background(),
			query,
			request["name"].(string),
			request["secondName"].(string),
			intPhoneNumber,
			request["post"].(string),
		).Scan(&Respond.Id)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Ошибка записи в БД", 500)
			return
		}

		respondMethod.SendRespond(w, r, Respond)

	case 4:

		room, ok := request["room"].(float64)
		if !ok {
			http.Error(w, "Body.room не прошел валидацию", 400)
			return
		}

		intRoom := int(room)

		cost, ok := request["cost"].(float64)
		if !ok {
			http.Error(w, "Body.cost не прошел валидацию", 400)
			return
		}

		intCost := int(cost)

		var dateIn pgtype.Date
		err := dateIn.Scan(request["dateIn"])
		if err != nil {
			http.Error(w, "Body.dateIn не прошел валидацию", http.StatusBadRequest)
			return
		}

		if !dateIn.Valid {
			http.Error(w, "Body.dateIn не прошел валидацию", http.StatusBadRequest)
			return
		}

		today := time.Now().Format("2006-01-02")
		todayDate, _ := time.Parse("2006-01-02", today)

		if dateIn.Time.Before(todayDate) {
			http.Error(w, "Body.dateIn не может быть в прошлом", http.StatusBadRequest)
			return
		}

		var dateOut pgtype.Date
		err = dateOut.Scan(request["dateOut"])
		if err != nil {
			http.Error(w, "Body.dateOut не прошел валидацию", http.StatusBadRequest)
			return
		}

		if !dateOut.Valid {
			http.Error(w, "Body.dateOut не прошел валидацию", http.StatusBadRequest)
			return
		}

		if dateIn.Time.Before(dateIn.Time) {
			fmt.Println(dateIn.Time)
			http.Error(w, "Body.dateOut не прошел валидацию, дата заезда не должна быть больше даты выезда", http.StatusBadRequest)
			return
		}

		GrandFatherId, ok := request["grandFatherId"].(float64)
		if !ok {
			http.Error(w, "Body.grandFatherId не прошел валидацию", 400)
			return
		}

		IntGrandFatherId := int(GrandFatherId)

		NursingHouseId, ok := request["nursingHouseId"].(float64)
		if !ok {
			http.Error(w, "Body.nursingHouseId не прошел валидацию", 400)
			return
		}

		IntNursingHouseId := int(NursingHouseId)

		DbDuplicate, DuplicateId := duplicateMethod.Duplicate(w, p, request, number)
		DuplicateText := "Ошибка создания сущности, уже существует с ID: " + fmt.Sprint(DuplicateId)

		if DbDuplicate {
			http.Error(w, DuplicateText, 400)
			return
		}

		query := `INSERT INTO "slots" (room, cost, "dateIn", "dateOut", "grandFather_id", "nursingHouse_id") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

		var Respond Respond
		err = p.QueryRow(
			context.Background(),
			query,
			intRoom,
			intCost,
			dateIn,
			request["dateOut"].(string),
			IntGrandFatherId,
			IntNursingHouseId,
		).Scan(&Respond.Id)

		if err != nil {
			fmt.Println(err)
			http.Error(w, "Ошибка записи в БД", 500)
			return
		}

		respondMethod.SendRespond(w, r, Respond)

	}

}
