package handlers

import (
	"database/sql"
	"errors"

	dbClient "github.com/konrad-amtenbrink/slack-daily-digest/db"
	"github.com/konrad-amtenbrink/slack-daily-digest/logic/_slack"
	"github.com/konrad-amtenbrink/slack-daily-digest/models"
	"github.com/slack-go/slack"
)

func HandleCommand(command slack.SlashCommand, client *slack.Client, db *sql.DB) error {
	switch command.Command {
	case "/digest":
		return handleDigestCommand(command, client, db)

	case "/subscribe":
		return handleSubscribeCommand(command, client, db)
	}
	return nil
}

func handleSubscribeCommand(command slack.SlashCommand, client *slack.Client, db *sql.DB) error {
	id := command.UserID
	if len(id) == 0 {
		return errors.New("user id does not exist")
	}
	return dbClient.AddUser(db, id)
}

func handleDigestCommand(command slack.SlashCommand, client *slack.Client, db *sql.DB) error {
	users := []models.User{{ID: command.UserID}}
	threads, err := dbClient.GetThreads(db)
	if err != nil {
		return err
	}
	return _slack.HandleDigest(users, threads, client)
}
