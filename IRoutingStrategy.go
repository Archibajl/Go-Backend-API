package backend_api

type IRoutingStrategy interface {
	getPath() string
	route(path string)
	isValidPath(path string) bool
}
