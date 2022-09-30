package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	dbClient "github.com/konrad-amtenbrink/slack-daily-digest/db"
	"github.com/konrad-amtenbrink/slack-daily-digest/handlers"
	"github.com/konrad-amtenbrink/slack-daily-digest/logic/cron"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func main() {
	godotenv.Load(".env")
	token := os.Getenv("SLACK_AUTH_TOKEN")
	appToken := os.Getenv("SLACK_APP_TOKEN")

	client := slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))

	db, err := dbClient.Init()
	if err != nil {
		log.Printf("DB Connection Error: %v", err)
	}
	defer db.Close()

	environment := os.Getenv("ENVIRONMENT")
	var socketClient *socketmode.Client
	if environment == "production" {
		go func() {
			cron.Init(client)
		}()
		socketClient = socketmode.New(
			client,
		)
	} else {
		socketClient = socketmode.New(
			client,
			socketmode.OptionDebug(true),
			socketmode.OptionLog(log.New(os.Stdout, "socketmode: ", log.Lshortfile|log.LstdFlags)),
		)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context, client *slack.Client, socketClient *socketmode.Client) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Shutting down socketmode listener")
				return
			case event := <-socketClient.Events:
				switch event.Type {
				case socketmode.EventTypeEventsAPI:
					eventsAPIEvent, succ := event.Data.(slackevents.EventsAPIEvent)
					if !succ {
						log.Printf("Could not type cast the event to the EventsAPIEvent: %v\n", event)
						continue
					}
					socketClient.Ack(*event.Request)
					err := handlers.OnMessage(eventsAPIEvent, client, db)
					if err != nil {
						log.Printf("Error: %v\n", err)
					}
				case socketmode.EventTypeSlashCommand:
					command, succ := event.Data.(slack.SlashCommand)
					if !succ {
						log.Printf("Could not type cast the message to a SlashCommand: %v\n", command)
						continue
					}
					socketClient.Ack(*event.Request)
					err := handlers.HandleSlashCommand(command, client)
					if err != nil {
						log.Print(err)
					}
				}

			}
		}
	}(ctx, client, socketClient)

	socketClient.Run()
}
