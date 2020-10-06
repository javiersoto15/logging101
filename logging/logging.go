package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *Logger

//Logger is the wrapper for the logrus package
type Logger struct {
	*logrus.Logger
}

//Event store messages with the defined format. 1 debug, 2 info, 3 error, 4 panic
type Event struct {
	id      int
	message string
}

//NewLogger initializes the new logger
func NewLogger() *Logger {

	var baseLogger = logrus.New()

	var standardLogger = &Logger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	standardLogger.SetOutput(os.Stdout)

	return standardLogger
}

func StartLogger() {
	Log = NewLogger()
}
