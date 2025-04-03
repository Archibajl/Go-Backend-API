
package main

type IRoutingStrategy interface
{
	getPath() String
	route(path String) void
	isValidPath(path String) bool

}


