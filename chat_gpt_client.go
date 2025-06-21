package main

import (
	// "encoding/json"
	// "regexp"
	"net/http"
	"testing"
	"time"
)

type ChatGPTClient struct {
	Logger GoLogger
	Client http.Client
	Field  string "message"
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestResponse(t *testing.T) {
	name := "Gladys"
	t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, name)
}

func NewChatGPTClient(_logger GoLogger) *ChatGPTClient {
	client := http.Client{Timeout: 15 * time.Second}
	gptClient := ChatGPTClient{Logger: _logger, Client: client, Field: ""}

	return &gptClient
}

func (gpt *ChatGPTClient) GetResponse(w *http.ResponseWriter, r *http.Request) *http.Response {

	req, err := http.NewRequest("GET", "https://", nil)

	if err != nil {
		gpt.Logger.Error("Error : " + err.Error())
	}

	resp, err := gpt.Client.Do(req)

	if err != nil {
		gpt.Logger.Error("Error : " + err.Error())
	}

	return resp
}
