package httputil

import "fmt"

type HTTPError struct {
	Message string `json:"message"`
}

func (h *HTTPError) Error() string {
	return fmt.Sprintf("message=%v", h.Message)
}
