package handlers

import (
	"database/sql"
	"errors"
	"os"

	dbClient "github.com/konrad-amtenbrink/slack-daily-digest/db"
	"github.com/konrad-amtenbrink/slack-daily-digest/logic/_slack"
	"github.com/konrad-amtenbrink/slack-daily-digest/models"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func OnMessage(event slackevents.EventsAPIEvent, client *slack.Client, db *sql.DB) error {
	switch event.Type {
	case slackevents.CallbackEvent:

		innerEvent := event.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			err := onMention(ev, client, db)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("unsupported event type")
	}
	return nil
}

func onMention(event *slackevents.AppMentionEvent, client *slack.Client, db *sql.DB) error {
	_, err := client.GetUserInfo(event.User)
	if err != nil {
		return err
	}

	threadTs := event.ThreadTimeStamp
	if threadTs == "" {
		attachment := slack.Attachment{}
		attachment.Pretext = "Please mention me in a thread to get a daily digest."
		err := _slack.PostMessage(slack.MsgOptionAttachments(attachment), client)
		return err
	}

	link, err := client.GetPermalink(&slack.PermalinkParameters{Channel: event.Channel, Ts: event.ThreadTimeStamp})
	if err != nil {
		return err
	}

	thread := models.Thread{Url: link, Title: "Daily Digest"}
	err = dbClient.AddThread(db, thread)
	if err != nil {
		return err
	}

	environment := os.Getenv("ENVIRONMENT")
	if environment == "development" {
		msg, err := _slack.CreateMessage([]models.Thread{thread})
		if err != nil {
			return err
		}

		err = _slack.PostMessage(msg, client)
		if err != nil {
			return err
		}
	}

	return nil
}
