package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Msg string
}

func Respond(w http.ResponseWriter, _ *http.Request, data any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err=%v\n", err)
	}
}
