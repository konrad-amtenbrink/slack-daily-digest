package handlers

import (
	"errors"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func EventMessage(event slackevents.EventsAPIEvent, client *slack.Client) error {
	switch event.Type {
	case slackevents.CallbackEvent:

		innerEvent := event.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			err := appMentionEvent(ev, client)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("unsupported event type")
	}
	return nil
}

func appMentionEvent(event *slackevents.AppMentionEvent, client *slack.Client) error {
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	_, err := client.GetUserInfo(event.User)
	if err != nil {
		return err
	}

	threadTs := event.ThreadTimeStamp
	if threadTs == "" {
		attachment := slack.Attachment{}
		attachment.Pretext = "Please mention me in a thread to get a daily digest."
		_, _, err = client.PostMessage(event.Channel, slack.MsgOptionAttachments(attachment))
		if err != nil {
			return errors.New("failed to post message")
		}
		return nil
	}
	link, err := client.GetPermalink(&slack.PermalinkParameters{Channel: event.Channel, Ts: event.ThreadTimeStamp})
	if err != nil {
		return err
	}

	attachment := slack.Attachment{}
	attachment.Pretext = link

	_, _, err = client.PostMessage(channelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return errors.New("failed to post message")
	}
	return nil
}
