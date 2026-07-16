package newEntity

import (
	"context"
	"encoding/json"
	"fmt"
	findInfo "grandFather/findInfo/findExactInfo"
	"net/http"
	"time"

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

	var request map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&request)

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

		DbDuplicate, DuplicateId := Duplicate(w, p, request, number)
		DuplicateText := "Ошибка создания сущности, уже существует с ID: " + fmt.Sprint(DuplicateId)

		if DbDuplicate {
			http.Error(w, DuplicateText, 400)
			return
		}

		query := `INSERT INTO "grandFather" (name, "secondName", age) VALUES ($1, $2, $3) RETURNING id`

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

		findInfo.SendRespond(w, r, Respond)

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

		DbDuplicate, DuplicateId := Duplicate(w, p, request, number)
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

		findInfo.SendRespond(w, r, Respond)

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

		DbDuplicate, DuplicateId := Duplicate(w, p, request, number)
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

		findInfo.SendRespond(w, r, Respond)

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

		DbDuplicate, DuplicateId := Duplicate(w, p, request, number)
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

		findInfo.SendRespond(w, r, Respond)

	}

}

func Duplicate(w http.ResponseWriter, p *pgxpool.Pool, request map[string]interface{}, number int) (bool, int) {

	switch number {
	case 1:
		age, _ := request["age"].(float64)

		IntAge := int(age)

		query := `SELECT id FROM "grandFather" WHERE "name" = $1 AND "secondName" = $2 AND age = $3`
		var id int
		err := p.QueryRow(
			context.Background(),
			query,
			request["name"],
			request["secondName"],
			IntAge,
		).Scan(&id)

		if err == nil {
			return true, id
		}
		return false, 0
	case 2:
		query := `SELECT id FROM "guardian" WHERE name = $1 AND "secondName" = $2 AND "phoneNumber" = $3`
		var id int
		err := p.QueryRow(
			context.Background(),
			query,
			request["name"],
			request["secondName"],
			request["phoneNumber"],
		).Scan(&id)
		if err == nil {
			return true, id
		}
		return false, 0
	case 3:
		query := `SELECT id FROM "employee" WHERE name = $1 AND "secondName" = $2 AND "phoneNumber" = $3 AND "post" = $4`
		var id int
		err := p.QueryRow(
			context.Background(),
			query,
			request["name"],
			request["secondName"],
			request["phoneNumber"],
			request["post"],
		).Scan(&id)

		if err == nil {
			return true, id
		}
		return false, 0
	case 4:
		query := `SELECT id FROM "slots" WHERE room = $1 AND "cost" = $2 AND "dateIn" = $3 AND "dateOut" = $4 AND "grandFather_id" = $5 AND "nursingHouse_id" = $6`
		var id int
		err := p.QueryRow(
			context.Background(),
			query,
			request["room"],
			request["cost"],
			request["dateIn"],
			request["dateOut"],
			request["grandFatherId"],
			request["nursingHouseId"],
		).Scan(&id)
		if err == nil {
			return true, id
		}
		return false, 0
	default:
		return false, 0
	}
}
