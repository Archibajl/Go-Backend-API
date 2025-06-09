package backend_api

import "strings"

type backEndRouter struct {
	path   string
	logger GoLogger
}

func (r backEndRouter) routePath(_path string) {
	splitPath := strings.Split(_path, "/")

	for _, v := range splitPath {
		switch v {
		case "chatGPT":
			r.path = _path
			return
		case "llama":
			r.path = _path
			return
		default:
			r.logger.Warn("Bad Path given: " + _path)
		}

	}
}
func (r backEndRouter) isValidPath(_path string) bool {
	return true
}

func (r backEndRouter) getPath() string {
	return r.path
}
