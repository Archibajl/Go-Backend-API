package backend_api

import (
	"io"
	"log"
	"os"
)

type GoLogger struct {
	filename string
	logger   log.Logger
}

func (l *GoLogger) goLogger() GoLogger {

	return *l
}

func createGoLogger(fname string) GoLogger {
	logFile, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	_logger := log.New(multiWriter, "My app Name ", log.Lshortfile)

	return GoLogger{
		filename: fname,
		logger:   *_logger,
	}
}

func (l *GoLogger) Info(_log string) {
	l.logger.Println(_log)
}

func (l *GoLogger) Warn(_log string) {
	l.logger.Panicln(_log)
}

func (l *GoLogger) Error(_log string) {
	l.logger.Fatalln(_log)
}
