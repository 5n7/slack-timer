package slacktimer

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/skmatz/slack-timer/controller"
	"github.com/skmatz/slack-timer/middleware"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	s.router = s.Route()
}

func (s *Server) Run(port int) error {
	log.Printf("Listening on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CombinedLoggingHandler(os.Stdout, s.router)); err != nil {
		return err
	}
	return nil
}

func (s *Server) Route() *mux.Router {
	slackMiddleware := middleware.NewSlack()
	basicChain := alice.New()
	slackChain := basicChain.Append(slackMiddleware.Handler)

	r := mux.NewRouter()
	slackController := controller.NewSlack()

	r.Methods(http.MethodGet).Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	r.Methods(http.MethodPost).Path("/slack").Handler(slackChain.Then(AppHandler{slackController.Post}))
	return r
}
