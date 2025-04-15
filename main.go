package main

import (
	// "crypto/tls"
	"fmt"
	"net/http"
	"os"
	"backend-api"
	// "time"
	// "github.com/gin-gonic/gin"
)


func main() {

	//Todo implement api get and post functions in an asynchronous fasion
	serverPort := 8080

	server = initServer()

	// go func() {
		
	// }
	// time.Sleep(10000 * time.Millisecond)

	// requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	// res, err := http.Get(requestURL)
	// if err != nil {
	// 	fmt.Printf("error making http request: %s\n", err)
	// 	logger.Fatalln("Error running http server: %s\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("client: got response!\n")
	// fmt.Printf("client: status code: %d\n", res.StatusCode)

}


func instServer() http.Server{
// client := http.DefaultClient()
logger := GoLogger()
mux := http.NewServeMux()
mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("server pinged at: %s /\n", r.Method)
	logger.getLogger.Println("server pinged at: %s /\n {}\n reader {}\n writer {} \n end",
		r.Method, r, w)
})
server := http.Server{
	Addr:    fmt.Sprintf("localhost:%d", serverPort),
	Handler: mux,
}
if err := server.ListenAndServe(); err != nil {
	if err == nil {
		fmt.Printf("error running http server: %s\n", err)
		logger.getLogger.Fatalln("error running http server: %s\n", err)
	}
}
return server
}
