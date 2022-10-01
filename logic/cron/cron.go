package cron

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	dbClient "github.com/konrad-amtenbrink/slack-daily-digest/db"
	"github.com/konrad-amtenbrink/slack-daily-digest/logic/_slack"
	"github.com/slack-go/slack"
)

func Init(client *slack.Client, db *sql.DB) {
	s := gocron.NewScheduler(time.UTC)

	// 6pm Europe/Berlin
	s.Every(1).Day().Tag("tag").At("16:00").Do(func() {
		handleCron(client, db)
	})

	s.StartBlocking()
}

func handleCron(client *slack.Client, db *sql.DB) {
	err := publishUpdate(client, db)
	if err != nil {
		log.Print(err)
	}
}

func publishUpdate(client *slack.Client, db *sql.DB) error {
	users, err := dbClient.GetUsers(db)
	if err != nil {
		return err
	}
	threads, err := dbClient.GetThreads(db)
	if err != nil {
		return err
	}
	return _slack.HandleDigest(users, threads, client)
}
