package httputil

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

func isNil(p interface{}) bool {
	if p == nil {
		return true
	}
	switch reflect.TypeOf(p).Kind() {
	case reflect.Array, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(p).IsNil()
	}
	return false
}

func marshal(payload interface{}) ([]byte, error) {
	if isNil(payload) {
		return []byte("{}"), nil
	}
	return json.Marshal(payload)
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	r, err := marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, werr := w.Write([]byte(err.Error()))
		if werr != nil {
			log.Print(werr)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, werr := w.Write(r)
	if werr != nil {
		log.Print(werr)
	}
}

func RespondJSONError(w http.ResponseWriter, status int, err error) {
	log.Printf("status=%d, error=%s", status, err)
	if status >= 500 {
		RespondJSON(w, status, HTTPError{
			Message: http.StatusText(status),
		})
	} else if e, ok := err.(*HTTPError); ok {
		RespondJSON(w, status, e)
	} else if err != nil {
		h := HTTPError{Message: err.Error()}
		RespondJSON(w, status, h)
	}
}
