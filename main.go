package main

import (
	// "crypto/tls"
	// "fmt"
	"net/http"
	// "time"
)

func main() {

	logger := CreateGoLogger("common.log")

	serverPort := "8080"
	server := initServer(serverPort)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("error running http server: %s\n" + err.Error())
	}
}

// // Handler function with proper signature
// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "hello world")
// }
