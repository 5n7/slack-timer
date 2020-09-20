package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var (
	slackToken = os.Getenv("SLACK_TOKEN")
)

type Slack struct{}

func NewSlack() *Slack {
	return &Slack{}
}

func (s *Slack) Callback(event slackevents.EventsAPIEvent) error {
	api := slack.New(slackToken)

	switch e := event.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		messages := strings.Split(e.Text, " ")
		if len(messages) < 2 {
			return fmt.Errorf("invalid message: %s", e.Text)
		}

		command := messages[1] // first element is the BOT ID (mention)
		switch command {
		case "ping":
			if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText("pong", false)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Slack) Verify(b []byte) (string, error) {
	var r *slackevents.ChallengeResponse
	if err := json.Unmarshal(b, &r); err != nil {
		return "", err
	}
	return r.Challenge, nil
}
