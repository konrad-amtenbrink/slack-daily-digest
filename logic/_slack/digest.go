package _slack

import (
	"log"

	"github.com/konrad-amtenbrink/slack-daily-digest/models"
	"github.com/slack-go/slack"
)

func PrepareDigest(users []models.User, threads []models.Thread, client *slack.Client) error {
	msg, err := CreateMessage(threads)
	if err != nil {
		return err
	}

	for _, user := range users {
		_, _, err := client.PostMessage(user.ID, msg)
		if err != nil {
			log.Printf("error sending message to user %s: %s", user.ID, err)
		}
	}
	return nil
}
