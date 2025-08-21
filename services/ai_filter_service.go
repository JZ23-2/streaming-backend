package services

import (
	"bytes"
	"encoding/json"
	"main/dtos"
	"net/http"
	"os"
)

var AI_URL = os.Getenv("AI_URL")

func Init() {
	AI_URL = os.Getenv("AI_URL")
}

func ModerateMessage(text string) (string, error) {
	reqBody, _ := json.Marshal(dtos.ModerationRequest{Text: text})

	url := AI_URL + "/moderate"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return text, nil
	}
	defer resp.Body.Close()

	var result dtos.ModerationResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return text, nil
	}

	if result.IsInappropriate {
		return "message was inappropriate word", nil
	}

	return text, nil
}
