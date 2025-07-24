package main

import (
	// "encoding/json"
	// "regexp"
	"bytes"
	"encoding/json"
	"io"
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

	// req, err := http.NewRequest("GET", "https://", nil)

	url := "https://api.openai.com/v1/chat/completions"

	payload := map[string]interface{}{
		"model": "gpt-4",
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": "Tell me a joke about software engineers."},
		},
		"temperature": 0.7,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		gpt.Logger.Error("Error Failed to encode payload: " + err.Error())
		return &http.Response{StatusCode: 783, Status: "Error 783 JSON syntax error: unable to json.Marshal payload from input."}
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		gpt.Logger.Error("Error Failed to create request: " + err.Error())
		return &http.Response{StatusCode: req.Response.StatusCode, Status: "Error 783 JSON syntax error: unable to json.Marshal payload from input." + req.response.Status}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		gpt.Logger.Error("Error Request failed : " + err.Error())
		err.
		return &http.Response{StatusCode: resp.StatusCode, Status: "Error Request failed : " + resp.Request.Response.Status}
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		gpt.Logger.Error("Error Failed to read response : " + err.Error())
		return &http.Response{StatusCode: respBody.StatusCode, Status: "Error Failed to read response : " + respBody.Request.Response.Status}
	}

	resp, err := gpt.Client.Do(req)

	if err != nil {
		gpt.Logger.Error("Error : " + err.Error())
		return &http.Response{StatusCode: resp.StatusCode, Status: "Error Request failed : " + resp.Request.Response.Status}
	}

	return resp
}
