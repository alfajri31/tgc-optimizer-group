package services

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

// TIP: You can use GoLand's code completion and refactoring features to improve the following code.'

func ApiMessageService(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Hello, World!"}
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		return
	}
}
