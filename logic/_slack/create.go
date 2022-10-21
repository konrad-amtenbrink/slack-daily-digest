package _slack

import (
	"github.com/konrad-amtenbrink/slack-daily-digest/models"
	"github.com/slack-go/slack"
)

func CreateDigestMessage(threads []models.Thread) (slack.MsgOption, error) {
	messageBlocks := createMessageBlocks()

	threadBlocks, err := createThreadBlocks(threads)
	if err != nil {
		return nil, err
	}

	msg := make([]slack.Block, len(messageBlocks)+len(threadBlocks))
	copy(msg, messageBlocks)
	copy(msg[len(messageBlocks):], threadBlocks)

	footer := getFooter()
	msg = append(msg, &footer)

	return slack.MsgOptionBlocks(msg...), nil
}

func CreateErrorMessage() slack.MsgOption {
	attachment := slack.Attachment{}
	attachment.Pretext = "Please mention me in a thread to add to the daily digest."
	return slack.MsgOptionAttachments(attachment)
}

func createMessageBlocks() []slack.Block {
	result := []slack.Block{}

	header := getHeader()
	result = append(result, &header)

	context := getContext()
	result = append(result, &context)

	divider := getDivider()
	result = append(result, &divider)

	mainSection := getMainSection()
	result = append(result, &mainSection)

	return result
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
