package db

import (
	"database/sql"
	"fmt"

	"github.com/konrad-amtenbrink/slack-daily-digest/logic/_slack"
)

func GetThreads() []_slack.Thread {
	return nil
}

func AddThread(db *sql.DB, thread _slack.Thread) error {
	sqlStatement := fmt.Sprintf("INSERT INTO thread (title, url) VALUES ('%s', '%s')",
		thread.Title, thread.Link)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}
