package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectarComBanco() *sql.DB {
	conexao := "user=hug40 dbname=testes host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
