package main

import "net/http"

type IRouter interface {
	ServeHTTP(http.ResponseWriter, http.Request)
	Handle(any, string, http.HandlerFunc)
	routePath(w http.ResponseWriter, r *http.Request)
	isValidPath(path string) bool
}
