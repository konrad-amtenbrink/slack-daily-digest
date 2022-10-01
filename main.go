package main

import (
	"context"
	"log"

	"github.com/konrad-amtenbrink/slack-daily-digest/config"
	dbClient "github.com/konrad-amtenbrink/slack-daily-digest/db"
	"github.com/konrad-amtenbrink/slack-daily-digest/logic/_slack"
	"github.com/konrad-amtenbrink/slack-daily-digest/logic/cron"
	"github.com/konrad-amtenbrink/slack-daily-digest/socket"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	config := config.LoadConfig()

	db, err := dbClient.Connect(config.DatabaseUrl)
	if err != nil {
		log.Printf("DB Connection Error: %v", err)
	}
	defer db.Close()

	client := _slack.NewClient(config.Slack.Token, config.Slack.AppToken)
	socketClient := _slack.NewSocketClient(client)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		cron.Init(client, db)
	}()

	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
		socket.OnSocketEvent(ctx, db, client, socketClient)
	}(ctx, client, socketClient)

	socketClient.Run()
}
