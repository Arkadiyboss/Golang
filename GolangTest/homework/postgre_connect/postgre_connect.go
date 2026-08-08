package postgre_connect

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnect() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://postgres:123@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}	
	if err:= conn.Ping(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Подключение в базе данных успешно")
}