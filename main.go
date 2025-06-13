package backend_api

import (
	// "crypto/tls"

	"context"
	"fmt"
	"net"
	"net/http"
	// "time"
)

func main() {

	serverPort := "8080"
	initServer(serverPort)

}

func initServer(serverPort string) *http.Server {
	const keyServerAddr = "serverAddr"
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// Initialize logger
	logger := &GoLogger{}
	logger.CreateGoLogger("common.log")

	// Set up context
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx() // ensures cleanup

	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return context.WithValue(ctx, keyServerAddr, l.Addr().String())
		},
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Printf("error running http server: %s\n", err)
		}
	}()

	return server
}

// Handler function with proper signature
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
