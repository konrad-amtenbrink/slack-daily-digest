package _slack

import (
	"testing"

	"github.com/konrad-amtenbrink/slack-daily-digest/models"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestCreateMessageBlocks(t *testing.T) {
	result := createMessageBlocks()

	expected := []slack.Block{}

	header := getHeader()
	expected = append(expected, &header)

	context := getContext()
	expected = append(expected, &context)

	divider := getDivider()
	expected = append(expected, &divider)

	mainSection := getMainSection()
	expected = append(expected, &mainSection)

	assert.ElementsMatch(t, result, expected)
}

func TestCreateThreadBlocks(t *testing.T) {
	threads := []models.Thread{
		{
			ID:    1,
			Title: "Thread 1",
			Url:   "https://example.com/123",
		},
	}
	result, _ := createThreadBlocks(threads)

	expected := []slack.Block{}

	threadBlock := slack.SectionBlock{
		Type: "section",
		Text: &slack.TextBlockObject{
			Type: "mrkdwn",
			Text: "Thread 1",
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
				URL:      "https://example.com/123",
			},
		},
	}

	expected = append(expected, &threadBlock)

	assert.ElementsMatch(t, result, expected)
}

func TestCreateDigestMessage(t *testing.T) {
	threads := []models.Thread{
		{
			ID:    1,
			Title: "Thread 1",
			Url:   "https://example.com/123",
		},
	}
	result, err := CreateDigestMessage(threads)

	assert.NotNil(t, result)
	assert.Nil(t, err)
}
