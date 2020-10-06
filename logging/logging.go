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

//DefaultFieldHook holds the hook interface and service names
type DefaultFieldHook struct {
	serviceName  string
	instanceName string
}

//Levels interface for logrus hooks
func (h *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//Fire triggers the log with this field
func (h *DefaultFieldHook) Fire(e *logrus.Entry) error {
	e.Data["Service"] = h.serviceName
	return nil
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
	Log.AddHook(&DefaultFieldHook{
		serviceName: "logging tutorials",
	})
}
