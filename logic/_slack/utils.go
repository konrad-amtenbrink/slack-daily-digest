package _slack

import "github.com/slack-go/slack"

func getHeader() slack.HeaderBlock {
	header := slack.HeaderBlock{
		Type: slack.MBTHeader,
		Text: slack.NewTextBlockObject("plain_text", ":newspaper:  Slack Daily Digest  :newspaper:", true, false),
	}
	return header
}

func getContext() slack.ContextBlock {
	context := slack.ContextBlock{
		Type: slack.MBTContext,
		ContextElements: slack.ContextElements{
			Elements: []slack.MixedElement{&slack.TextBlockObject{
				Type: "mrkdwn",
				Text: "Development Announcements",
			}},
		},
	}
	return context
}

func getDivider() slack.DividerBlock {
	divider := slack.DividerBlock{
		Type: slack.MBTDivider,
	}
	return divider
}

func getMainSection() slack.SectionBlock {
	section := slack.SectionBlock{
		Type: slack.MBTSection,
		Text: slack.NewTextBlockObject(slack.MarkdownType, ":loud_sound: *IN CASE YOU MISSED IT* :loud_sound:", false, false),
	}
	return section
}

func getFooter() slack.ContextBlock {
	context := slack.ContextBlock{
		Type: slack.MBTContext,
		ContextElements: slack.ContextElements{
			Elements: []slack.MixedElement{&slack.TextBlockObject{
				Type: "mrkdwn",
				Text: ":pushpin: Do you have something to include in the newsletter? Just mention *me* in a thread and it gets included in the next digest.",
			}},
		},
	}
	return context
}
