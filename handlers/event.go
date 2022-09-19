package handlers

import (
	"errors"

	"github.com/konrad-amtenbrink/slack-daily-digest/logic/_slack"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func OnMessage(event slackevents.EventsAPIEvent, client *slack.Client) error {
	switch event.Type {
	case slackevents.CallbackEvent:

		innerEvent := event.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			err := onMention(ev, client)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("unsupported event type")
	}
	return nil
}

func onMention(event *slackevents.AppMentionEvent, client *slack.Client) error {
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

	msg, err := _slack.CreateMessage([]_slack.Thread{{Link: link, Title: "Daily Digest"}})
	if err != nil {
		return err
	}

	err = _slack.PostMessage(msg, client)
	return err
}
