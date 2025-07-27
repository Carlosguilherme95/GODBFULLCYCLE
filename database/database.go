package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// nome da função maiúscula para poder usar em outros pacotes
func DatabaseConection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/meu_banco")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
