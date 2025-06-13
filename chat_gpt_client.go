package backend_api

import (
	// "encoding/json"
	// "regexp"
	"testing"
)

type chat_gpt_client struct {
	logger GoLogger
	Field  string "message"
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestResponse(t *testing.T) {
	name := "Gladys"
	t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, name)
}
