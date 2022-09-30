package db

import (
	"database/sql"
	"fmt"

	"github.com/konrad-amtenbrink/slack-daily-digest/models"
)

func GetUsers(db *sql.DB) ([]models.User, error) {
	sqlStatement := `SELECT * FROM "user"`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func AddUser(db *sql.DB, id string) error {
	sqlStatement := fmt.Sprintf(`INSERT INTO "user" (id) VALUES ('%s')`, id)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
