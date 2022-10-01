package _slack

import (
	"github.com/konrad-amtenbrink/slack-daily-digest/models"
	"github.com/slack-go/slack"
)

func HandleDigest(users []models.User, threads []models.Thread, client *slack.Client) error {
	message, err := CreateDigestMessage(threads)
	if err != nil {
		return err
	}
	for _, user := range users {
		PostDigestMessage(message, user, client)
	}
	return nil
}
