package main

import (
	"fmt"
	"os"

	st "github.com/skmatz/slack-timer"
)

var (
	port = 8080
)

func main() {
	s := st.NewServer()
	s.Init()
	if err := s.Run(port); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
