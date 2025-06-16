package main

import (
	"net/http"
	"strings"
)

/*
	Router component, setup to handle
*/

// Route structure for routing information
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// Router Struct for routing data
type Router struct {
	routes []Route
	logger GoLogger
}

// Instantiates router
func NewRouter(_goLogger GoLogger) *Router {
	router := &Router{
		logger: _goLogger}
	router.Handle("GET", "/chat-gpt", nil)
	return router
}

func hello(w *http.ResponseWriter, r *http.Request) {

}

func (r *Router) Handle(method, path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{Method: method, Path: path, Handler: handler})
}

func (m *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range m.routes {
		if req.Method == route.Method && matchPattern(route.Path, req.URL.Path) {
			route.Handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func (r *Router) routePath(wrt *http.ResponseWriter, req *http.Request) bool {
	// splitPath := strings.Split(_path, "/")
	_path := req.URL.Path

	switch _path {
	case "/chatGPT":

		return true
	case "/llama":

		return true
	default:
		r.logger.Warn("Bad Path given: " + _path)
		return false
	}
}

func matchPattern(pattern, path string) bool {
	if pattern == path {
		return true
	}
	if strings.HasSuffix(pattern, "/*") {
		base := strings.TrimSuffix(pattern, "/*")
		return strings.HasPrefix(path, base)
	}
	return false
}

func (r Router) isValidPath(_path string) bool {
	return true
}
