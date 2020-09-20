package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/skmatz/slack-timer/httputil"
)

type Recover struct{}

func NewRecover() *Recover {
	return &Recover{}
}

func (r *Recover) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				log.Printf("panic: %s", err)
				httputil.RespondJSONError(w, http.StatusInternalServerError, nil)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
