package _slack

import (
	"io/ioutil"
	"os"

	"github.com/konrad-amtenbrink/slack-daily-digest/models"
	"github.com/slack-go/slack"
)

func CreateDigestMessage(threads []models.Thread) (slack.MsgOption, error) {
	messageBlocks, err := createMessageBlocks()
	if err != nil {
		return nil, err
	}

	threadBlocks, err := createThreadBlocks(threads)
	if err != nil {
		return nil, err
	}
	msg := append(messageBlocks[:3], threadBlocks...)
	msg = append(msg, messageBlocks[3:]...)
	return slack.MsgOptionBlocks(msg...), nil
}

func CreateErrorMessage() slack.MsgOption {
	attachment := slack.Attachment{}
	attachment.Pretext = "Please mention me in a thread to add to the daily digest."
	return slack.MsgOptionAttachments(attachment)
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
