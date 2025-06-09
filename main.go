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
	server := initServer(serverPort)

}

func initServer(serverPort string) http.Server {
	const keyServerAddress = "serverAddr"
	mux := http.NewServeMux()
	mux.HandleFunc("/", getPath)
	logger := GoLogger.createGoLogger("common.log")
	ctx, cancelCtx := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    ":" + serverPort,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	if err := server.ListenAndServe(); err != nil {
		if err == nil {
			fmt.Printf("error running http server: %s\n", err)
			logger.getLogger.Fatalln("error running http server: %s\n", err)
		}
	}
	return server
}
