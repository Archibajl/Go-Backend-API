package main

import (
	// "crypto/tls"
	"fmt"
	"net/http"
	// "time"
)

func main() {

	serverPort := "8080"
	initServer(serverPort)
}

// Handler function with proper signature
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
