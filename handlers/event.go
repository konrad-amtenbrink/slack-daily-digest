package handlers

import (
	"database/sql"
	"errors"

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
		switch event := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			err := onMention(event, client, db)
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
		message := _slack.CreateErrorMessage()
		return _slack.PostErrorMessage(message, event.Channel, client)
	}

	url, err := client.GetPermalink(&slack.PermalinkParameters{Channel: event.Channel, Ts: event.ThreadTimeStamp})
	if err != nil {
		return err
	}

	replies, _, _, err := client.GetConversationReplies(&slack.GetConversationRepliesParameters{ChannelID: event.Channel, Timestamp: event.ThreadTimeStamp})
	if err != nil {
		return err
	}

	thread := models.Thread{Url: url, Title: replies[0].Msg.Text}

	return dbClient.AddThread(db, thread)
}
