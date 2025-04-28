package backend_api

import (
    "log"
    "testing"
    "regexp"
    "encoding/json"
)

type chat_gpt_client struct{
    logger Logger
    Message string json:"message"
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestResponse(t *testing.T) {
    name := "Gladys"
        t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
}
