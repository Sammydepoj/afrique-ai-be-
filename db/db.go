package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {

	createNewsLetterTable := `
	CREATE TABLE IF NOT EXISTS newsletter (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE
	)`

	_, err := DB.Exec(createNewsLetterTable)

	fmt.Println(err, "Error")

	if err != nil {
		panic("Could not create news letter table.")
	}

}
