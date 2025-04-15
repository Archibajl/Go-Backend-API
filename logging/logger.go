package backend-api

import "log"

type GoLogger struct{
	logger Logger
}

func (l GoLogger) getLogger() GoLogger{
	return l.logger
}

func (l GoLogger) setLogger(_logger Logger) void{
	l.logger := log.New(os.Stdout, "backend-api.log", log.LstdFlags )
	return
}
