package cron

import (
	"log"

	"github.com/konrad-amtenbrink/slack-daily-digest/logic/_slack"
	"github.com/robfig/cron/v3"
	"github.com/slack-go/slack"
)

func Init(client *slack.Client) {
	c := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))
	c.AddFunc("0 0 * * * ", func() { publishUpdate(client) })
	c.Start()
}

func publishUpdate(client *slack.Client) {
	msg, err := _slack.CreateMessage([]_slack.Thread{{Title: "Daily Digest send from cron job"}})
	if err != nil {
		log.Default().Print(err)
	}

	err = _slack.PostMessage(msg, client)
	if err != nil {
		log.Default().Print(err)
	}
}
