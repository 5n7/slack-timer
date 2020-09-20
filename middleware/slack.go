package middleware

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/skmatz/slack-timer/httputil"

	"github.com/slack-go/slack"
)

var (
	signingSecret = os.Getenv("SLACK_SIGNING_SECRET")
)

type Slack struct{}

func NewSlack() *Slack {
	return &Slack{}
}

func (s *Slack) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		verifier, err := slack.NewSecretsVerifier(r.Header, signingSecret)
		if err != nil {
			log.Print(err)
			httputil.RespondJSONError(w, http.StatusInternalServerError, err)
			return
		}

		bodyReader := io.TeeReader(r.Body, &verifier)
		_, err = ioutil.ReadAll(bodyReader)
		if err != nil {
			log.Print(err)
			httputil.RespondJSONError(w, http.StatusInternalServerError, err)
			return
		}

		if err := verifier.Ensure(); err != nil {
			log.Print(err)
			httputil.RespondJSONError(w, http.StatusBadRequest, err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
