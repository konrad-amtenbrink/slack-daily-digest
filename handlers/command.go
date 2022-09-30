package handlers

import (
	"github.com/konrad-amtenbrink/slack-daily-digest/logic/_slack"
	"github.com/slack-go/slack"
)

func HandleSlashCommand(command slack.SlashCommand, client *slack.Client) error {
	switch command.Command {
	case "/digest":
		return handleDigestCommand(command, client)

	case "/subscribe":
		return handleDigestCommand(command, client)
	}
	return nil
}

func handleSubscribeCommand(command slack.SlashCommand, client *slack.Client) error {
	return nil // dbClient.
}

func handleDigestCommand(command slack.SlashCommand, client *slack.Client) error {
	msg, err := _slack.CreateMessage([]_slack.Thread{{Title: "Daily Digest send from command"}})
	if err != nil {
		return err
	}

	id := command.ChannelID
	if len(id) == 0 {
		id = command.UserID
	}

	_, _, err = client.PostMessage(id, msg)
	return err
}
