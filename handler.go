package slacktimer

import (
	"net/http"

	"github.com/skmatz/slack-timer/httputil"
)

type AppHandler struct {
	h func(http.ResponseWriter, *http.Request) (int, interface{}, error)
}

func (a AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, res, err := a.h(w, r)
	if err != nil {
		httputil.RespondJSONError(w, status, err)
		return
	}
	httputil.RespondJSON(w, status, res)
}
