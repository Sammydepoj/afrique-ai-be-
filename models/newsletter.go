package models

import (
	"errors"
	"github.com/afrique-ai-be/db"
)

type UserEmails struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (e *UserEmails) EmailExists() (bool, error) {
	query := `SELECT id FROM newsletter WHERE email = ?`
	row := db.DB.QueryRow(query, e.Email)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (e *UserEmails) Save() error {
	exists, err := e.EmailExists()
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email has already subscribed")
	}

	query := `INSERT INTO newsletter (email) VALUES (?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Email)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	e.ID = id
	return nil
}
