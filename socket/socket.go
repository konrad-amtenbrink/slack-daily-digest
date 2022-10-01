package socket

import (
	"context"
	"database/sql"
	"log"

	"github.com/konrad-amtenbrink/slack-daily-digest/handlers"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

func OnSocketEvent(ctx context.Context, db *sql.DB, client *slack.Client, socketClient *socketmode.Client) {
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
				err := handlers.HandleCommand(command, client, db)
				if err != nil {
					log.Print(err)
				}
			}

		}
	}
}
