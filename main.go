package main

import (
	"io"
	"net/http"
	// "time"
)

var logger = CreateGoLogger("common.log")

func main() {

	serverPort := "8080"
	router := NewRouter(logger)
	router.Handle("GET", "/chat-gpt", nil)
	router.Handle("GET", "/hello", router.hello)
	router.Handle("GET", "/", printOutput)
	server := NewServer(serverPort, router, logger)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("error running http server: %s\n" + err.Error())
	}
}

func printOutput(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	logger.Info(string(body))
}
