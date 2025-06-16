package main

import (
	// "encoding/json"
	// "regexp"
	"net/http"
	"testing"
)

type ChatGPTClient struct {
	logger GoLogger
	Field  string "message"
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestResponse(t *testing.T) {
	name := "Gladys"
	t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, name)
}

func (gpt *ChatGPTClient) GetResponse(w *http.ResponseWriter, r *http.Request) {

}
