package backend-api
import "log"

type IRoutingStrategy interface
{
	logger Logger
	getPath() String
	route(path String) void
	isValidPath(path String) bool

}


