package duplicateMethod

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Duplicate(w http.ResponseWriter, p *pgxpool.Pool, request map[string]interface{}, number int) (bool, int) {

	switch number {
	case 1:
		age, _ := request["age"].(float64)

		IntAge := int(age)

		query := `SELECT id FROM "resident" WHERE "name" = $1 AND "secondName" = $2 AND age = $3`
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
	case 5:
		query := `SELECT id FROM "users" WHERE login = $1`
		var id int
		err := p.QueryRow(
			context.Background(),
			query,
			request["login"],
		).Scan(&id)
		if err == nil {
			return true, id
		}
		return false, 0
	default:
		return false, 0
	}
}
