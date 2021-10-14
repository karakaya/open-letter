package main

import (
	"database/sql"
	"fmt"
)

func connect() *sql.DB {
	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "postgres", "password", "open-letter")
	db, dbErr = sql.Open("postgres", psql)
	if dbErr != nil {
		panic(dbErr)
	}
	return db
}
