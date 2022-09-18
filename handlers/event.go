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
	_, err := client.GetUserInfo(event.User)
	if err != nil {
		return err
	}

	attachment := slack.Attachment{}
	attachment.Pretext = "Digest added"

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	_, _, err = client.PostMessage(channelID, slack.MsgOptionAttachments(attachment))
	if err != nil {
		return errors.New("failed to post message")
	}
	return nil
}
