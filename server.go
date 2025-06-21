package main

import "net/http"

func NewServer(serverPort string, router Router, logger GoLogger) *http.Server {

	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: &router,
	}

	return server
}
