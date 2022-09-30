package db

import (
	"database/sql"
	"fmt"

	"github.com/konrad-amtenbrink/slack-daily-digest/models"
)

func GetThreads(db *sql.DB) ([]models.Thread, error) {
	sqlStatement := `SELECT * FROM "thread"`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		if err := rows.Scan(&thread.ID, &thread.Title, &thread.Url); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	return threads, nil
}

func AddThread(db *sql.DB, thread models.Thread) error {
	sqlStatement := fmt.Sprintf("INSERT INTO thread (title, url) VALUES ('%s', '%s')",
		thread.Title, thread.Url)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
