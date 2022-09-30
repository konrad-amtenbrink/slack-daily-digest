package _slack

import (
	"io/ioutil"
	"os"

	"github.com/konrad-amtenbrink/slack-daily-digest/models"
	"github.com/slack-go/slack"
)

func CreateMessage(threads []models.Thread) (slack.MsgOption, error) {
	messageBlocks, err := createMessageBlocks()
	if err != nil {
		return nil, err
	}

	threadBlocks, err := createThreadBlocks(threads)
	if err != nil {
		return nil, err
	}
	msg := append(messageBlocks[:3], threadBlocks...)
	msg = append(msg, messageBlocks[3+len(threadBlocks):]...)
	return slack.MsgOptionBlocks(msg...), nil
}

func PostMessage(msg slack.MsgOption, client *slack.Client) error {
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	_, _, err := client.PostMessage(channelID, msg)
	if err != nil {
		return err
	}
	return nil
}

func createMessageBlocks() ([]slack.Block, error) {
	jsonFile, err := os.Open("./templates/message.json")
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	blocks := slack.Blocks{}
	blocks.UnmarshalJSON(byteValue)
	return blocks.BlockSet, nil
}

func createThreadBlocks(threads []models.Thread) ([]slack.Block, error) {
	result := []slack.Block{}
	for _, thread := range threads {
		threadBlock := slack.SectionBlock{
			Type: "section",
			Text: &slack.TextBlockObject{
				Type: "mrkdwn",
				Text: thread.Title,
			},
			Accessory: &slack.Accessory{
				ButtonElement: &slack.ButtonBlockElement{
					Type: "button",
					Text: &slack.TextBlockObject{
						Type:  "plain_text",
						Text:  "View thread",
						Emoji: true,
					},
					Value:    "view_thread",
					ActionID: "button-action",
					URL:      thread.Url,
				},
			},
		}
		result = append(result, &threadBlock)
	}

	return result, nil
}
