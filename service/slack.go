package service

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/skmatz/slack-timer/etc"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var (
	slackToken     = os.Getenv("SLACK_TOKEN")
	timeFormat     = "15:04:05"
	timeZoneName   = "Asia/Tokyo"
	timeZoneOffset = 9 * 60 * 60
)

type Slack struct{}

func NewSlack() *Slack {
	return &Slack{}
}

func now() string {
	utc := time.Now().UTC()
	jst := time.FixedZone(timeZoneName, timeZoneOffset)
	return utc.In(jst).Format(timeFormat)
}

func (s *Slack) Callback(event slackevents.EventsAPIEvent) error {
	api := slack.New(slackToken)

	switch e := event.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		e.Text = etc.RemoveDuplicateSpace(e.Text)
		messages := strings.Split(e.Text, " ")
		if len(messages) < 2 {
			return fmt.Errorf("invalid message: %s", e.Text)
		}

		commands := messages[1:] // first element is the BOT ID (mention)
		switch commands[0] {
		case "ping":
			if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText("pong", false)); err != nil {
				return err
			}
		case "timer":
			if len(commands) < 2 {
				return fmt.Errorf("timer command got invalid message: %s", e.Text)
			}

			dur, err := strconv.Atoi(commands[1])
			if err != nil {
				return err
			}

			if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText(fmt.Sprintf("timer started at %s", now()), false)); err != nil {
				return err
			}

			if len(commands) > 2 {
				switch commands[2] {
				case "sec":
				case "min":
					dur *= 60
				}
			}

			timer := time.NewTimer(time.Second * time.Duration(dur))
			defer timer.Stop()
			select {
			case <-timer.C:
				if _, _, err := api.PostMessage(e.Channel, slack.MsgOptionText(fmt.Sprintf("timer finished at %s", now()), false)); err != nil {
					return err
				}
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
