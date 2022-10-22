package _slack

import (
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestGetHeader(t *testing.T) {
	result := getHeader()

	expected := slack.HeaderBlock{
		Type: slack.MBTHeader,
		Text: slack.NewTextBlockObject("plain_text", ":newspaper:  Slack Daily Digest  :newspaper:", true, false),
	}

	assert.Equal(t, result, expected)
}

func TestGetContext(t *testing.T) {
	result := getContext()

	expected := slack.ContextBlock{
		Type: slack.MBTContext,
		ContextElements: slack.ContextElements{
			Elements: []slack.MixedElement{&slack.TextBlockObject{
				Type: "mrkdwn",
				Text: "Development Announcements",
			}},
		},
	}

	assert.Equal(t, result, expected)
}

func TestGetDivider(t *testing.T) {
	result := getDivider()

	expected := slack.DividerBlock{
		Type: slack.MBTDivider,
	}

	assert.Equal(t, result, expected)
}

func TestGetFooter(t *testing.T) {
	result := getFooter()

	expected := slack.ContextBlock{
		Type: slack.MBTContext,
		ContextElements: slack.ContextElements{
			Elements: []slack.MixedElement{&slack.TextBlockObject{
				Type: "mrkdwn",
				Text: ":pushpin: Do you have something to include in the newsletter? Just mention *me* in a thread and it gets included in the next digest.",
			}},
		},
	}

	assert.Equal(t, result, expected)
}
