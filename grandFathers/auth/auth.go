package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"grandFather/respondMethod"
	duplicateMethod "grandFather/duplicate"
	"math/rand"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Respond struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type RespondUpdatedToken struct {
	UpdatedToken string `json:"updatedToken"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ1234567890-*!")

func NewUser(w http.ResponseWriter, r *http.Request, p *pgxpool.Pool) {

	var request map[string]interface{}

	if r.Method != http.MethodPost {
		http.Error(w, "Метод отличается от POST", 400)
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Ошибка парсинга входящих данных", 500)
		return
	}

	if login, ok := request["login"].(string); !ok || login == "" {
		http.Error(w, "Body.login не прошел валидацию", 400)
		return
	}

	if password, ok := request["password"].(string); !ok || password == "" {
		http.Error(w, "Body.password не прошел валидацию", 400)
		return
	}

	if role, ok := request["role"].(string); !ok || role == "" {
		http.Error(w, "Body.role не прошел валидацию", 400)
		return
	}

	DbDuplicate, DuplicateId := duplicateMethod.Duplicate(w, p, request, 5)
	DuplicateText := "Ошибка создания сущности, уже существует с ID: " + fmt.Sprint(DuplicateId)

	if DbDuplicate {
		http.Error(w, DuplicateText, 400)
		return
	}

	query := `INSERT INTO "users" (login, password, role, token) VALUES ($1, $2, $3, $4) RETURNING id`

	token := GenerateToken(24)

	var Respond Respond
	err = p.QueryRow(
		context.Background(),
		query,
		request["login"].(string),
		request["password"].(string),
		request["role"].(string),
		token,
	).Scan(&Respond.Id, token)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Ошибка записи в БД", 500)
		return
	}

	respondMethod.SendRespond(w, r, Respond)

}

func GenerateToken(n int) string {

	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(res)
}

func NewUpdatedToken(w http.ResponseWriter, r *http.Request, p *pgxpool.Pool) {
	var request map[string]interface{}

	if r.Method != http.MethodPost {
		http.Error(w, "Метод отличается от POST", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Ошибка парсинга входящих данных", 500)
		return
	}

	if token, ok := request["token"].(string); !ok || token == "" {
		http.Error(w, "Body.token не прошел валидацию", 400)
		return
	}

	query := `SELECT token FROM "users" where token = $1`

	var respond Respond

	err = p.QueryRow(context.Background(), query, request["token"]).Scan(&respond.Token)
	if err != nil {
		if err == pgx.ErrNoRows {
			http.Error(w, "Токен не найден", http.StatusNotFound)
			return
		}
		http.Error(w, "Ошибка при поиске токена: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if respond.Token == "" {
		http.Error(w, "Токен не найден", 500)
		fmt.Println(err)
		return
	}

	var UpdatedToken RespondUpdatedToken

	updated := time.Now()

	UpdatedToken.UpdatedToken = GenerateToken(16)

	query = `UPDATE "users" SET "updatedToken" = $1, "updatedAt" = $2 WHERE "token" = $3`

	_, err = p.Exec(context.Background(), query, UpdatedToken.UpdatedToken, updated, request["token"])

	if err != nil {
		fmt.Println(err, "Не удалось обновить БД")
		http.Error(w, "Ошибка доступа к БД", 500)
	}

	respondMethod.SendRespond(w, r, UpdatedToken)
}

func CheckToken(token string, p *pgxpool.Pool) (bool, error) {
	var respond pgtype.Timestamptz

	query := `SELECT "updatedAt" FROM "users" where "updatedToken" = $1`

	err := p.QueryRow(context.Background(), query, token).Scan(&respond)

	if err != nil {
		return false, err
	}

	fifteenMinutesAgo := time.Now().Add(-15 * time.Minute)
	return respond.Time.After(fifteenMinutesAgo), nil
}
