package main

import "net/http"

func initServer(serverPort string) *http.Server {

	// Initialize logger
	logger := GoLogger{}
	logger.CreateGoLogger("common.log")

	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: NewRouter(logger),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("error running http server: %s\n" + err.Error())
		}
	}()

	return server
}
