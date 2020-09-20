package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/skmatz/slack-timer/service"
	"github.com/slack-go/slack/slackevents"
)

type Slack struct{}

func NewSlack() *Slack {
	return &Slack{}
}

func (s *Slack) Post(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	if r.Header.Get("X-Slack-Retry-Num") != "" {
		return http.StatusInternalServerError, nil, fmt.Errorf("incalid request: retry request")
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, nil, err
	}

	slackEvent, err := slackevents.ParseEvent(json.RawMessage(b), slackevents.OptionNoVerifyToken())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	slackService := service.NewSlack()

	switch slackEvent.Type {
	case slackevents.CallbackEvent:
		if err := slackService.Callback(slackEvent); err != nil {
			return http.StatusInternalServerError, nil, err
		}
	case slackevents.URLVerification:
		challenge, err := slackService.Verify(b)
		if err != nil {
			log.Print(err)
			return http.StatusInternalServerError, nil, err
		}
		if _, err := w.Write([]byte(challenge)); err != nil {
			return http.StatusInternalServerError, nil, err
		}
	}
	return http.StatusOK, nil, nil
}
