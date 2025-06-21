package main

import "net/http"

func initServer(serverPort string) *http.Server {

	// Initialize logger
	logger := CreateGoLogger("common.log")

	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: NewRouter(logger),
	}

	return server
}
