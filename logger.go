package backend_api

import (
	"io"
	"log"
	"os"
)

type GoLogger struct {
	Filename string
	Logger   *log.Logger
}

func (gl *GoLogger) CreateGoLogger(_logName string) {
	gl.Filename = _logName
	logFile, _ := os.OpenFile(_logName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	gl.Logger = log.New(multiWriter, "Go Logger: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (log *GoLogger) goLogger() GoLogger {

	return *log
}

// func createGoLogger(fname string) GoLogger {

// 	_logger := log.New(multiWriter, "My app Name ", log.Lshortfile)

// 	return GoLogger{
// 		Filename: fname,
// 		Logger:   *_logger,
// 	}
// }

func (l *GoLogger) Info(_log string) {
	l.Logger.Println(_log)
}

func (l *GoLogger) Warn(_log string) {
	l.Logger.Panicln(_log)
}

func (l *GoLogger) Error(_log string) {
	l.Logger.Fatalln(_log)
}
