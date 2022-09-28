package cron

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/konrad-amtenbrink/slack-daily-digest/logic/_slack"
	"github.com/slack-go/slack"
)

func Init(client *slack.Client) {
	s := gocron.NewScheduler(time.UTC)

	// 6pm Europe/Berlin
	s.Every(1).Day().Tag("tag").At("16:00").Do(func() {
		handleCron(client)
	})

	s.StartBlocking()
}

func handleCron(client *slack.Client) {
	err := publishUpdate(client)
	if err != nil {
		log.Print(err)
	}
}

func publishUpdate(client *slack.Client) error {
	msg, err := _slack.CreateMessage([]_slack.Thread{{Title: "Daily Digest send from cron job"}})
	if err != nil {
		return err
	}

	err = _slack.PostMessage(msg, client)
	if err != nil {
		return err
	}
	return nil
}
