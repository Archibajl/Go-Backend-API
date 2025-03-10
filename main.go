package main

import (
	// "crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	// "github.com/gin-gonic/gin"
)

func main() {

	logger := log.Default()
	serverPort := 8080

	go func() {
		// client := http.DefaultClient()
		mux := http.NewServeMux()
		mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server pinged at: %s /\n", r.Method)
			logger.Println("server pinged at: %s /\n {}\n reader {}\n writer {} \n end",
				r.Method, r, w)
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	time.Sleep(10000 * time.Millisecond)

	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	/*r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}).POST("/post", func(w *gin.Context) {
		w.JSON(http.StatusOK, gin.H{
			"message": "u wot",
		})
	})

	r.Run()*/ // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
