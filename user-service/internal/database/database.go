package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "host=localhost port=5432 dbname=users sslmode=disable user=postgres password=postgres")
	if err != nil {
		log.Println("unable to connect to database")
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		log.Println("unable to connect to database", err.Error())
		return nil, err
	}
	return conn, nil
}
