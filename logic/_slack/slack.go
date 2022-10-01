package _slack

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func NewClient(token string, appToken string) *slack.Client {
	return slack.New(token, slack.OptionDebug(true), slack.OptionAppLevelToken(appToken))
}

func NewSocketClient(client *slack.Client) *socketmode.Client {
	return socketmode.New(client)
}
