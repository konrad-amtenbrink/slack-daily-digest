package _slack

import (
	"log"

	"github.com/konrad-amtenbrink/slack-daily-digest/models"
	"github.com/slack-go/slack"
)

func PostDigestMessage(message slack.MsgOption, user models.User, client *slack.Client) error {
	_, _, err := client.PostMessage(user.ID, message)
	if err != nil {
		log.Printf("error sending message to user %s: %s", user.ID, err)
	}

	return err
}

func PostErrorMessage(message slack.MsgOption, channelID string, client *slack.Client) error {
	_, _, err := client.PostMessage(channelID, message)
	if err != nil {
		log.Printf("error sending message to channel %s: %s", channelID, err)
	}

	return err
}
