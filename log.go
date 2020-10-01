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
