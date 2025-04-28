package backend_api

import "log"

type backEndRouter struct { 
 selectedServce Object
 path String
 logger Logger
} 

func (r IRoutingStrategy) routePath(String _path) void {
	String[] splitPath  = strings.Split(path, '/')

	for _,v := range splitPath{ 
		switch v {
		case "chatGPT":
			path = _path
			return 
		case "llama":
			path=_path
			return
		default:
		 logger.printLn("Bad Path given: " + _path)	
		}
		
	}
}
func (r IRoutingStrategy) isValidPath(String path) bool{
	return true
}

func (r IRoutingStrategy) getPath() String{
	return path
}
