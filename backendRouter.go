package backend_api

import (
	"net/http"
	// "strings"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type Router struct {
	routes []Route
	path   string
	logger GoLogger
}

func (r *Router) Handle(method, pattern string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{Method: method, Pattern: pattern, Handler: handler})
}

func (m *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range m.routes {
		if r.Method == route.Method && m.routePath(route.Pattern, r.URL.Path) {
			route.Handler(w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func (r Router) routePath(_pattern string, _path string) bool {
	// splitPath := strings.Split(_path, "/")

	switch _path {
	case "chatGPT":
		r.path = _path
		return true
	case "llama":
		r.path = _path
		return true
	default:
		r.logger.Warn("Bad Path given: " + _path)
		return false
	}
}

func (r Router) isValidPath(_path string) bool {
	return true
}

func (r Router) getPath() string {
	return r.path
}
