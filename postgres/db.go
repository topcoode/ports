package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Postgresconnection() (*sql.DB, error) {
	connStr := "postgres://postgres:mahi@localhost/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
	fmt.Println("Successfully connected to the database!")
	return db, err
}
