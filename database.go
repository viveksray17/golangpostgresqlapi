package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func sqlExec(query string) error {
	psqlConnection := fmt.Sprintf("host=localhost port=5432 user=vivek password=%v dbname=vivek sslmode=disable", os.Getenv("PG_PASS"))
	db, _ := sql.Open("postgres", psqlConnection)
	defer db.Close()
	_, err := db.Exec(query)
	return err
}

func sqlQuery(query string) *sql.Rows {
	psqlConnection := fmt.Sprintf("host=localhost port=5432 user=vivek password=%v dbname=vivek sslmode=disable", os.Getenv("PG_PASS"))
	db, err := sql.Open("postgres", psqlConnection)
	defer db.Close()
	checkError(err)
	rows, err := db.Query(query)
	checkError(err)
	return rows
}
