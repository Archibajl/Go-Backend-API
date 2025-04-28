package backend_api
import "log"

type IRoutingStrategy interface
{
	logger log.Logger,
	getPath() String,
	route(path String) void,
	isValidPath(path String) bool,

}


